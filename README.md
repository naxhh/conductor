# Conductor

Conductor is an API responsible of directing the deployments with k8s

Is in charge of:

- Receive a .tar.gz with the project to deploy
- Building & publishing the dockerfile image in the project
- Deploy a Pod exposing the project ports
- Manage DNS of each project


## PoC

This is a PoC for a deployment systeme similar to `now`


### Main goal

The main goal of this project is to beat the deployment time of our current deploy system

### Second class goals

- Able to be used in some simple use cases like demos
- Simple deployment, 0 configuration
- App auto scaling

### Final goal

The final goal of this project, when it's not a PoC anymore, will be to be able to handle the requests of the NodeJS SSR server

## Dev

GoDep produces very slow compile times.
We recommend to execute `GOOS=linux go install` once in the project so vendor get's cached and compile times are faster again.

Faster compile times = happy devs

### Build

We use the in-cluster client. So in order to test this you should have minikube.

to build & deploy in k8s

```
./dev/start.sh
```

Usage:

```
curl $(minikube service conductor --url)/ping
```

When done

```
./dev/stop.sh
```
