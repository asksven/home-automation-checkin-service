#!/usr/bin/env bash


kubectl --namespace=blink-location apply -f ns.yaml
kubectl --namespace=blink-location create secret docker-registry k8sdevregistry-azurecr-io-key --docker-server=$DOCKER_REGISTRY --docker-username=$DOCKER_REGISTRY_USER --docker-password=$DOCKER_REGISTRY_PASSWORD --docker-email=foo.bar@pwc.com
kubectl --namespace=blink-location apply -f pvc-azure.yaml
kubectl --namespace=blink-location apply -f mongo-deployment.yaml
kubectl --namespace=blink-location apply -f mongo-service.yaml
kubectl --namespace=blink-location apply -f web-deployment.yaml
kubectl --namespace=blink-location apply -f web-service.yaml
kubectl --namespace=blink-location apply -f ingress.yaml
