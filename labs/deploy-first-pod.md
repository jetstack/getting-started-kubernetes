# Deploying your first pod

* Deploy a pod to a local Kubernetes cluster with kubectl
* View information about a running pod

This lab assumes you have a running Kubernetes cluster. For this first lab, it is
recommended to use a local (Docker) cluster - [see the lab](setup-local-cluster.md).

## Create a pod

```
kubectl run simple-cms \
  --labels="name=simple-cms" \
  --image=jetstack/simple-cms:1.0.0
```

## View running pods

```
kubectl get pods
```

## View details of a running pod

```
kubectl describe pods simple-cms
```

Note the IP address ($IP).

## Test the web service

```
curl http://$IP:8081
```

Next up: [use a replication controller](replication-controller.md) and request Kubernetes automatically maintain multiple pod
replicas.
