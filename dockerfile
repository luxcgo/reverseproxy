# Build the Go Binary.
FROM golang:1.20 as build
ENV CGO_ENABLED 0

# Copy the source code into the container.
COPY . /service

# Build the service binary.
WORKDIR /service
RUN go build

# Run the Go Binary in Alpine.
FROM alpine:3.18
ARG BUILD_DATE
COPY --from=build /service/reverseproxy /service/reverseproxy
WORKDIR /service
CMD ["./reverseproxy"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.title="reverseproxy" \
    org.opencontainers.image.authors="luxcgo <luxcgo@gmail.com>" \
    org.opencontainers.image.source="https://github.com/luxcgo/reverseproxy" \
    org.opencontainers.image.vendor="luxcgo"