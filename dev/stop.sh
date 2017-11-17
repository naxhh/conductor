#!/bin/sh
eval $(minikube docker-env)
kubectl config use-context minikube

kubectl delete service example
kubectl delete deployment example

kubectl delete service conductor
kubectl delete deployment conductor
docker rmi -f conductor:v1

echo ""
echo ""
echo "---------------------------------"
echo "Deployment deleted"
echo "---------------------------------"
