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

### Deps

Since we are on a PoC just install everything from master:

```
go get \
k8s.io/client-go/kubernetes \
k8s.io/client-go/rest \
k8s.io/client-go/kubernetes/typed/apps/v1beta2 \
k8s.io/api/apps/v1beta2 \
k8s.io/api/core/v1 \
k8s.io/client-go/kubernetes/typed/core/v1 \
k8s.io/apimachinery/pkg/apis/meta/v1 \
github.com/docker/docker/client \
github.com/docker/docker/api/types \
github.com/gin-gonic/gin \
k8s.io/apimachinery/pkg/util/intstr
```

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

### Testing deployment

The `example` folder is a simple nodejs app. You can use it to test the deployment

```
cd example/ && tar -zcvf ../example.tar.gz . && cd ..
```

Create a deployment

```
curl -X POST $(minikube service conductor --url)/deploy \
  -F "file=@example.tar.gz" \
  -H "Content-Type: multipart/form-data"
```

And test your app

```
curl $(minikube service example --url)/
```

## TODO

- Improve port detection, now you can only bound to :8080
- Expose the service, now you must do it manually `kubectl expose deployment example --type=LoadBalancer`
- Use random project names instead of the file name, avoid clash
- Allow registering custom domain names to a specific deployment
- Maybe K8s Jobs is a better way to coordinate build - deploy
