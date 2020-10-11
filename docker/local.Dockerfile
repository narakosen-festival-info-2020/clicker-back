# ==============================================================================
# docker - local.Dockerfile
# ==============================================================================

# golang:1.15.2-buster
FROM golang@sha256:0978cc067eb3f53901c00b70a024f182baa371bdfe7f35f3d64e56cab2471c4d
ENV LANG C.UTF-8

ENV APP_HOME $GOPATH/src/clicker-back
RUN mkdir $APP_HOME
WORKDIR $APP_HOME

# Prepare App
COPY . $APP_HOME

# CMD go run main.go
CMD /bin/bash
