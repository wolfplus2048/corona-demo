FROM golang:latest AS build
RUN mkdir -p /src
COPY . /src
WORKDIR /src
RUN export GO111MODULE=on \
     && export GOPROXY=https://goproxy.io/ \
     && go build -trimpath -o server

FROM wolfplus2048/corona:latest
RUN mkdir -p /work
COPY --from=build /src/server /work
COPY config-docker.toml /work/config.toml
WORKDIR /work
CMD ./server
