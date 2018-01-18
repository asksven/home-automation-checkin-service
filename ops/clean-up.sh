#!/usr/bin/env bash

kubectl --namespace=blink-location delete -f mongo-deployment.yaml
kubectl --namespace=blink-location delete -f web-deployment.yaml
kubectl --namespace=blink-location delete -f mongo-service.yaml
kubectl --namespace=blink-location delete -f web-service.yaml
kubectl --namespace=blink-location delete -f ns.yaml
kubectl --namespace=blink-location delete -f ingress.yaml
