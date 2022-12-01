FROM golang:1.19 AS build
COPY . /workdir/
 
RUN cd /workdir/ && go mod tidy && go mod download && go build -o http-debug main.go  && chmod a+x /workdir/http-debug

FROM alpine:3.17.0
COPY --from=build  /workdir/http-debug /usr/local/bin/http-debug
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
CMD [ "http-debug" ]