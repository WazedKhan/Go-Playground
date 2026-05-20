# stage 1: need to create build as binary
FROM golang:1.26.3-alpine AS builder

# as normal docker file set the working directory
WORKDIR /app

# copy dependency file and download
COPY go.mod go.sum ./
RUN go mod download

# code source code and build it
COPY . .
# CGO_ENABLED=0 disables CGO to create statically linked binary. This ensures that binary
# run in minimal environments like Alpine or Scratch without needing extra C libraries
# GOOS-linux GOARCH=amd64 ensures that binary is build for linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# stage 2: minimal image
FROM alpine:3.21
# installing certificates(needed for HTTPS requests)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# copy the binary from the builder stage: this time we are not copying source code but binary that we already build in stage 1
COPY --from=builder /app/main .

EXPOSE 8000
CMD [ "./main" ]