#!/bin/bash

source setenv.sh
docker build . --tag $DOCKER_REGISTRY/$CI_PROJECT_NAMESPACE/$CI_PROJECT_NAME:$DOCKER_IMAGE_TAG
docker push $DOCKER_REGISTRY/$CI_PROJECT_NAMESPACE/$CI_PROJECT_NAME:$DOCKER_IMAGE_TAG
