# Exposing as a service

* Create a service using kubectl


## Create a service to expose pods

The service will expose all `name=simple-cms` pods. Use the manifest file with `kubectl`.

```
kubectl create -f service.yml
```

## View details of the service

The service will have been assigned a VIP (Virtual IP) - let's find out what it is. Note the IP and PORT.

```
kubectl get se
```

[`se` is another shorthand - for `services`]

## Check the service works

Use a browser or `curl` to be sure the service correctly exposes the web serving pods. These requests will be load-balanced using a simple round-robin method.

```
curl http://IP:PORT
```

Next up: [deploy a canary replication controller](rolling-update.md).
