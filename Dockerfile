FROM golang:1.21.6 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/qbtuna

FROM golang:1.21.6 AS release

WORKDIR /app
COPY --from=build /app/bin/qbtuna /app/bin/qbtuna

ENTRYPOINT ["/app/bin/qbtuna"]
