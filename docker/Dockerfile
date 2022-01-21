FROM golang:1.16-alpine

# Copy and download dependency using go mod
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .
COPY config.yml ./
RUN go build

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dankers-io-backend

EXPOSE 3000
ENTRYPOINT ["./dankers-io-backend"]
