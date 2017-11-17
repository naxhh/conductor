#!/bin/sh
eval $(minikube docker-env)
kubectl config use-context minikube

GOOS=linux go build -o ./conductor .
docker build -t conductor:v1 .
kubectl create -f ./Deployment.yaml
kubectl expose deployment conductor --type=LoadBalancer

echo ""
echo ""
echo "---------------------------------"
echo "Finish deployment"
echo ""
echo "Service exposed in: $(minikube service conductor --interval=1 --url)"
echo "---------------------------------"
