FROM golang:1.11.2 AS builder
WORKDIR /go/src/Prem
COPY . .
RUN wget --no-verbose -O /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64
RUN chmod +x /usr/local/bin/dep
RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 go install -v ./...

FROM scratch
COPY --from=builder /go/bin/Prem /Prem
# EXPOSING ports
EXPOSE 8080
ENTRYPOINT ["/Prem"]
