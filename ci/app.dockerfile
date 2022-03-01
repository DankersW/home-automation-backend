FROM golang:1.16 as builder
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o home-automation-backend .

FROM alpine:3.15.0
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/home-automation-backend ./
COPY ci/config.yml .

EXPOSE 3000
ENTRYPOINT ["./home-automation-backend"]