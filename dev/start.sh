#!/bin/sh
eval $(minikube docker-env)
kubectl config use-context minikube

GOOS=linux go build -o ./conductor .
docker build -t conductor:v1 .
kubectl run conductor --image=conductor:v1 --port 8080
kubectl expose deployment conductor --type=LoadBalancer

echo ""
echo ""
echo "---------------------------------"
echo "Finish deployment"
echo ""
echo "Service exposed in: $(minikube service conductor --url)"
echo "---------------------------------"
