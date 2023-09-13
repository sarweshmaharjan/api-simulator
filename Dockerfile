FROM golang:1.20-alpine
EXPOSE 9999 9990
WORKDIR /app
COPY . ./
COPY ./cmd/.env ./.env
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

RUN go mod download
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o . ./cmd/main.go

