# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9

WORKDIR /go/src/github.com/asksven/home-automation-checkin-service

#ADD . .
ADD . /go/src/github.com/asksven/home-automation-checkin-service

RUN go get github.com/BurntSushi/toml github.com/golang/glog github.com/gorilla/mux gopkg.in/mgo.v2 gopkg.in/mgo.v2/bson
RUN cd /go/src/github.com/asksven/home-automation-checkin-service && go install -v && go build -o app . 


ENTRYPOINT /go/src/github.com/asksven/home-automation-checkin-service/app
#CMD ["app"]

# Document that the service listens on port 8080.
EXPOSE 3000