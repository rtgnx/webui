FROM golang:1.15-alpine as build
WORKDIR /webui
COPY . .
ARG TARGET=webui
RUN go mod vendor && \
        CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /webui main.go && \
        apk add upx binutils && \
        strip /webui && \
        upx /webui && \
        ls -alh /n

FROM scratch
LABEL org.opencontainers.image.source https://github.com/netauth/webui
EXPOSE 8080
ENTRYPOINT ["/webui"]
COPY --from=build /webui /webui