FROM golang:alpine AS builder
LABEL file-version="1.0.0" \
      last-update="2023-12-27" \
      created-by="@atsuiest"

WORKDIR /application
# Copy all source files
COPY . .
# Get dependencies
RUN go mod download
RUN go mod verify
# Build binary executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/exec

# Define image for executable
FROM alpine:3.12
# Coping user password
COPY --from=builder /etc/passwd /etc/passw
# Copy executable
COPY --from=builder /go/bin/exec /go/bin/exec
# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Run executable
CMD ["/go/bin/exec"]