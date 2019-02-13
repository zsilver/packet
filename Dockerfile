FROM golang:1.11.5 as builder
COPY ./src/pkg ./src/pkg
COPY ./src/cmd ./src/cmd
COPY ./build_cli.sh ./
RUN ./build_cli.sh

FROM alpine:3.6
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/cli .
ENTRYPOINT ["./cli"]