FROM golang

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
COPY . $GOPATH/src/reverse
WORKDIR $GOPATH/src/reverse
RUN go build -o app ./server/server.go
EXPOSE 5300
CMD ./app
