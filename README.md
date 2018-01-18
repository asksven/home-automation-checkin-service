# A simple API for managing check-ins

The API managed objects/persons by name and tracks their location. An object/person can only be in one location at any given time.

## Credits

https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94

## Run

### Locally
1. Install and start a local `mongod`
1. `go get github.com/golang/glog github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux`
1. `go run app.go`
1. access the API under `http://localhost:3000`

#### With docker-compose
1. run `docker-compose up` and access the API under `http://localhost:3000`

#### With k8s

1. `source setenv.sh`
1. `docker build . --tag $DOCKER_REGISTRY/$CI_PROJECT_NAMESPACE/$CI_PROJECT_NAME:latest`
1. `docker login $DOCKER_REGISTRY -u $DOCKER_REGISTRY_USER -p $DOCKER_REGISTRY_PASSWORD`
1. `docker push $DOCKER_REGISTRY/$CI_PROJECT_NAMESPACE/$CI_PROJECT_NAME:latest`

## API

- `GET /checkins`: list all checkins, e.g. `curl http://localhost:3000/checkins`
- `GET /checkins/{location}`: list all checkins for a given location, e.g. `curl http://localhost:3000/checkins/home`
- `GET /checkin/{name}`: retrieves a checkin by name, e.g. `curl http://localhost:3000/checkin/paul`
- `POST /checkin`: add a checkin, e.g. `curl -sSX POST -d '{"name":"peter","location":"home"}' http://localhost:3000/checkin`
- `DELETE /checkin`: deletes a checkin by name, e.g. `curl -sSX DELETE -d '{"name":"paul"}' http://localhost:3000/checkin`
