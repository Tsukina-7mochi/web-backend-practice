FROM ubuntu:24.04 AS base

RUN apt update && apt upgrade -y


FROM base AS dev

ARG USERNAME

WORKDIR /tmp
RUN apt install -y wget tar
RUN wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz

CMD ["/bin/bash"]


FROM dev AS build

COPY . /api
RUN /usr/local/go/bin/go clean
RUN /usr/local/go/bin/go build -o /tmp/main /api/main.go


FROM base AS prod

COPY --from=build /tmp/main /api/server

CMD ["/api/server"]
