FROM golang:latest@sha256:2b1cbf278ce05a2a310a3d695ebb176420117a8cfcfcc4e5e68a1bef5f6354da AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /opt/app -buildvcs=false

FROM scratch
COPY --from=builder /opt/app /opt/app
ENTRYPOINT ["/opt/app"]
