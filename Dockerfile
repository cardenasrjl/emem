FROM golang:1.13

RUN apt-get -y update \
    && apt-get -y upgrade \
    && apt-get install -y git bash gnupg2 ca-certificates vim curl tzdata make default-mysql-client

WORKDIR /app

# Enables go modules inside GOPATH
ENV GO111MODULE=on

# Copy go modules
COPY go.mod go.sum ./

# Install dependecies.
RUN go mod download

COPY . ./

# RUN make build
RUN go build -v -o emem cmd/server/main.go

CMD ["/app/emem"]
