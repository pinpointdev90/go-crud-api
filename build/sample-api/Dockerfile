FROM alpine:3.18

RUN apk add --no-cache ca-certificates

ENV PATH /usr/local/go/bin:$PATH

ENV GOLANG_VERSION 1.20.6

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/sample-api

FROM alpine
COPY --from=builder /build/sample-api /opt/app/
ENTRYPOINT ["/opt/app/sample-api"]