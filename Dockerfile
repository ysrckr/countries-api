#build stage
FROM --platform=${BUILDPLATFORM} golang:alpine AS builder

ARG BUILDPLATFORM
ARG TARGETPLATFORM

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ENV USER=appuser
ENV UID=10001 

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /go/src/app/
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o /go/bin/app -v ./cmd/api

#final stage
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /go/bin/app /app

USER appuser:appuser

ENTRYPOINT ["/app"]

LABEL Name=countries-api Version=0.0.1
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD "curl --fail http://localhost:8080 || exit 1"
EXPOSE 8000
