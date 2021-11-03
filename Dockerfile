# Currently tested on Go 1.16 only
FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o k8sensus main.go

FROM alpine
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/k8sensus .

ENTRYPOINT ["./k8sensus"]
