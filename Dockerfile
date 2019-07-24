FROM golang

ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor
COPY . $GOPATH/src/reverse
WORKDIR $GOPATH/src/reverse
RUN go mod tidy
RUN go build -o app ./server/server.go
RUN go build -o cli ./client/client.go
EXPOSE 5300
CMD ./app