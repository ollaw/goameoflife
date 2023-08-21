FROM golang:1.19

LABEL org.opencontainers.image.authors="ollagabriele@pm.me"

WORKDIR /app

# Dep
COPY go.mod go.sum ./
RUN go mod download

# Go files
COPY src/ src/
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /goameoflife

#Run
ENTRYPOINT ["/goameoflife"]