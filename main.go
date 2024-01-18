package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type GluetunResponse struct {
	Port int `json:"port"`
}

func main() {
	// Grab env vars
	gluetunHost := os.Getenv("GLUETUN_HOST")
	qbittorrentHost := os.Getenv("QBITTORRENT_HOST")
	qbittorrentUsername := os.Getenv("QBITTORRENT_USERNAME")
	qbittorrentPassword := os.Getenv("QBITTORRENT_PASSWORD")

	// Get refresh interval from environment variable, or default to 60 seconds
	var interval int
	intervalString, found := os.LookupEnv("INTERVAL_S")
	if !found {
		intervalString = "60"
	}

	// Convert interval into an int
	interval, err := strconv.Atoi(intervalString)
	if err != nil {
		panic(err)
	}

	for {
		err := refresh(RefreshParams{
			GluetunHost:         gluetunHost,
			QbittorrentHost:     qbittorrentHost,
			QbittorrentUsername: qbittorrentUsername,
			QbittorrentPassword: qbittorrentPassword,
		})

		if err != nil {
			fmt.Println("Oh no! Something went wrong:", err)
		} else {
			fmt.Println("Success!")
		}

		fmt.Printf("Running again in %d seconds...\n", interval)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

type RefreshParams struct {
	GluetunHost         string
	QbittorrentHost     string
	QbittorrentUsername string
	QbittorrentPassword string
}

func refresh(params RefreshParams) error {
	// Get port number from gluetun
	fmt.Println("Getting port from gluetun...")
	response, err := http.Get(fmt.Sprintf("%s/v1/openvpn/portforwarded", params.GluetunHost))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	gluetunResponse := GluetunResponse{}
	err = json.Unmarshal(body, &gluetunResponse)

	fmt.Println("Port from gluetun:", gluetunResponse.Port)

	// Log into qbittorrent
	fmt.Println("Logging into qbittorrent...")
	form := url.Values{}
	form.Set("username", params.QbittorrentUsername)
	form.Set("password", params.QbittorrentPassword)
	response, err = http.PostForm(fmt.Sprintf("%s/api/v2/auth/login", params.QbittorrentHost), form)
	if err != nil {
		return err
	}

	fmt.Println("Login response status code:", response.Status)

	body, err = io.ReadAll(response.Body)
	fmt.Println("Login response body:", string(body))

	cookie := response.Header.Get("set-cookie")
	if cookie == "" {
		return fmt.Errorf("no cookie returned by qbittorrent")
	}

	// Update port in qbittorrent
	fmt.Println("Updating listen port in qbittorrent...")
	payload := fmt.Sprintf("json={\"listen_port\": \"%d\"}", gluetunResponse.Port)

	fmt.Println("Payload to be sent:", payload)

	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/v2/app/setPreferences", params.QbittorrentHost),
		strings.NewReader(payload),
	)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", params.QbittorrentHost)
	request.Header.Set("Origin", params.QbittorrentHost)
	request.Header.Set("Cookie", cookie)
	client := http.Client{}

	response, err = client.Do(request)

	fmt.Println("Response status code:", response.Status)

	return nil
}
