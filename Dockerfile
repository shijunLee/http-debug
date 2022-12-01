FROM golang:1.19 AS build
COPY . /workdir/
 
RUN cd /workdir/ && go mod tidy && go mod download && go build -o http-debug main.go  && chmod a+x /workdir/http-debug

FROM alpine:3.16
COPY --from=build  /workdir/http-debug /usr/bin/http-debug

CMD [ "/usr/bin/http-debug" ]