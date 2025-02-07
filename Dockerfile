FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /opt/app -buildvcs=false

FROM scratch
COPY --from=builder /opt/app /opt/app
ENTRYPOINT ["/opt/app"]
