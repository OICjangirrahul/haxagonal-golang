# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/air-verse/air@latest


COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 5000

CMD ["air", "-c", ".air.toml"]
