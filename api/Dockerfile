FROM golang:latest as builder

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN  GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main ${MAIN_PATH}

FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]