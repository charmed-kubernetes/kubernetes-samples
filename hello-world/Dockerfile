FROM --platform=$TARGETPLATFORM golang:1.18-alpine
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY main.go .
RUN echo "Building hello-world for ${TARGETPLATFORM}"
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o hello-world main.go

FROM alpine:latest
RUN adduser -D hello
USER hello
COPY --from=0 /app/hello-world .
ENV PORT 8080
CMD ["./hello-world"]
