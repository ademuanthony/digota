FROM golang

RUN mkdir -p /go/src/github.com/digota/digota

ENV GO111MODULE=on

ADD . /go/src/github.com/digota/digota
WORKDIR /go/src/github.com/digota/digota

RUN go mod download

RUN ls /go/src/github.com/digota/digota

RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/digota/digota  -watch github.com/digota/digota
