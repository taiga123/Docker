FROM golang:1.8.0

ENV GOPATH $GOPATH:/go/src

RUN apt-get update && \
    apt-get upgrade -y
    
RUN go get gopkg.in/olivere/elastic.v5

# install revel and revel-cli
RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel

# change current dir
RUN mkdir /go/src/myapp

EXPOSE 9000
