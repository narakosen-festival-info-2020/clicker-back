# syntax = docker/dockerfile:experimental
FROM golang:1.15.2 as dev-builder

WORKDIR /go/src/github.com/narakosen-festival-info-2020/clicker-back/
COPY ./* ./
COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/

RUN CGO_ENABLED=0 go build -o /go/bin/clicker-back ./cmd/clicker_app/


FROM golang:1.15.2 as main-builder

RUN --mount=type=secret,id=credential,dst=/root/.gitconfig \
    go get github.com/narakosen-festival-info-2020/clicker-back/cmd/clicker_app && \
    CGO_ENABLED=0 go build -o /go/bin/clicker-back github.com/narakosen-festival-info-2020/clicker-back/cmd/clicker_app


FROM alpine:3.12 as develop
COPY --from=dev-builder /go/bin/clicker-back .
ENTRYPOINT [ "./clicker-back" ]


FROM alpine:3.12 as main
COPY --from=main-builder /go/bin/clicker-back .
ENTRYPOINT [ "./clicker-back" ]