FROM golang:1.15

# Copy and download dependency using go mod
WORKDIR /build
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

# Copy the code into the container
COPY src/. .
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dankers-io-backend
RUN cp /build/main .

EXPOSE 3000
ENTRYPOINT ["/dankers-io-backend/main"]
