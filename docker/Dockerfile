FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache curl

ENV PATH="/go/bin:${PATH}"

# Install Go tools
RUN go install \
    github.com/google/wire/cmd/wire@latest && \
    go install github.com/air-verse/air@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/nsf/gocode@latest && \
    go install github.com/acroca/go-symbols@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/lint/golint@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest


COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
EXPOSE 5000

CMD ["air", "-c", ".air.toml"]
