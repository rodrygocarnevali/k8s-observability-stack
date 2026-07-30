FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o demo-api .


FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/demo-api .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/demo-api"]