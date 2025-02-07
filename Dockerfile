FROM golang:latest@sha256:927112936d6b496ed95f55f362cc09da6e3e624ef868814c56d55bd7323e0959 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /opt/app -buildvcs=false

FROM scratch
COPY --from=builder /opt/app /opt/app
ENTRYPOINT ["/opt/app"]
