# Scaling up a replication controller

* Scale up the number of replicas
* Watch as new pods are added

Resize to 4 replicas.

```
kubectl scale rc simple-cms --replicas=4
```

## Watch the new pods come online

```
kubectl get pods --watch
```

Next up: [expose the pods as a service](expose-service.md) for load-balanced 
cluster access.
