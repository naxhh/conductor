#!/bin/sh
eval $(minikube docker-env)
kubectl config use-context minikube

kubectl delete service conductor
kubectl delete deployment conductor

echo ""
echo ""
echo "---------------------------------"
echo "Deployment deleted"
echo "---------------------------------"
