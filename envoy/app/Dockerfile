FROM golang:latest as builder

WORKDIR /app
COPY ./proto ./proto
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/helloworld

# ---
FROM scratch

COPY --from=builder /go/bin/helloworld /go/bin/helloworld

ENTRYPOINT ["/go/bin/helloworld"]
