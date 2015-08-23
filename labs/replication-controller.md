# Create a replication controller

* Create a replication controller to maintain multiple pod replicas
* View the details of the controller and the pods

## Create a replication controller

This replication controller will ensure 2 replicas exist.

```
kubectl create -f rc.yml
```

## List all replication controllers

```
kubectl get rc
```

[`rc` is a shorthand for `replicationControllers`]

This should now show the 'simple-cms' replication controller with 2 replicas (pods). 
See the detail of these pods as before using `kubectl`.

``` 
kubectl get pods 
```

Next up: [scale the replication controller](scale-replication-controller.md).
