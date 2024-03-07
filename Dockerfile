FROM golang:1.22.1-alpine3.19 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd 
USER app
CMD ["/app/app"]    