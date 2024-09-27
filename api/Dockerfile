FROM ubuntu:24.04 AS base

RUN apt update && apt upgrade -y


FROM base AS dev

ARG USERNAME

WORKDIR /tmp
RUN apt install -y curl wget tar
RUN wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
ENV GOPATH /root/go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH

RUN go install github.com/air-verse/air@latest

WORKDIR /api
COPY ./go.mod ./go.sum ./
RUN go mod download

CMD ["/root/go/bin/air"]


FROM dev AS build

COPY . /api
RUN go clean
RUN go build -o /tmp/main /api/main.go


FROM base AS prod

COPY --from=build /tmp/main /api/server

CMD ["/api/server"]
