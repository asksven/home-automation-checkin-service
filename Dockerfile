# build stage
FROM golang:1.12 AS build-env
WORKDIR /go/src/github.com/asksven/home-automation-checkin-service

#ADD . .
ADD . /go/src/github.com/asksven/home-automation-checkin-service

RUN go get github.com/BurntSushi/toml github.com/golang/glog github.com/gorilla/mux gopkg.in/mgo.v2 gopkg.in/mgo.v2/bson
RUN cd /go/src/github.com/asksven/home-automation-checkin-service && go install -v && CGO_ENABLED=0 GOOS=linux go build -o goapp . 

# final stage
FROM alpine:3.9
WORKDIR /app
COPY --from=build-env /go/src/github.com/asksven/home-automation-checkin-service/goapp /app/
COPY --from=build-env /go/src/github.com/asksven/home-automation-checkin-service/config.toml /app/

ENTRYPOINT ./goapp
EXPOSE 3000