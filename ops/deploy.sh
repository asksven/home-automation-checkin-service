#!/usr/bin/env bash


# To test this locally
# 1. install jinja2
# 2. set env vars using "source ./setenv.sh"
#   $TESTING is either 0 or 1
#   $NAMESPACE the namespace to deploy to
#   $DEPLOY_ENV is either "production", or "development"
#   $CI_BUILD_REF is the sha of the build, e.g. 1ecfd275763eff1d6b4844ea3168962458c9f27a and is used to tag the image
#   $CI_COMMIT_REF_SLUG is the branch name

echo '---$TESTING:' ${TESTING}
echo '---$DEPLOY_ENV:' ${DEPLOY_ENV}
echo '---$NAMESPACE:' ${NAMESPACE}
echo '---$CI_COMMIT_REF_SLUG:' ${CI_COMMIT_REF_SLUG}

if [ "$DEPLOY_ENV" = "production" ]
then
    HOST1="checkin.asksven.io"
elif [ "$DEPLOY_ENV" = "temp" ] 
then
    # for temp we use the ${CI_BUILD_REF} to create a unique URL    
    HOST1="checkin-${CI_BUILD_REF}.asksven.io"
else
    # we use the branch name in the URL
    HOST1="checkin-${CI_COMMIT_REF_SLUG}.asksven.io"
fi
echo '---$HOST1:' ${HOST1}


# $NAMESPACE is defined in .gitlab-ci.yaml

COMMAND="kubectl create namespace $NAMESPACE"

if [ "$TESTING" = "1" ] 
then
  echo '>>>would execute command:' ${COMMAND}
else
  eval $COMMAND
fi

K8S_DIR=./manifests
TARGET_DIR=${K8S_DIR}/.generated

rm -rf ${TARGET_DIR}/*
mkdir -p ${TARGET_DIR}/values


for f in ${K8S_DIR}/*.yaml
do
  echo processing $f
  jinja2 $f --format=yaml --strict \
    -D NAMESPACE=$NAMESPACE \
    -D PROJECT_NAME=${CI_PROJECT_NAME} \
    -D TAG=${CI_BUILD_REF} \
    -D HOST1=${HOST1} \
    -D DEPLOY_ENV=${DEPLOY_ENV} \
    -D IMAGE_URL=${DOCKER_IMAGE_URL} \
    > "${TARGET_DIR}/$(basename ${f})"
done

# create the namespace first
COMMAND1="kubectl --namespace=$NAMESPACE apply -f ${TARGET_DIR}/ns.yaml"
COMMAND2="kubectl --namespace=$NAMESPACE apply -f ${TARGET_DIR}"

if [ "$TESTING" = "1" ] 
then
  echo '>>>would execute command:' ${COMMAND1}
  echo '>>>would execute command:' ${COMMAND2}

else
  eval $COMMAND1
  eval $COMMAND2
fi
