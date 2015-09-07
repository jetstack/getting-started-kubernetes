# Rolling update

* Canary a new release for live testing
* Roll out the new release using a replication controller rolling-update


## Launch a new canary

Use `kubectl` to create a new replication controller, with a single pod that uses version 1.1.0 of the application.

```
kubectl create -f rc-canary.yml
```

## Validate the replication controller

```
kubectl get rc
```

## Apply rolling-update with the new release

Assuming the canary worked satisfactorily, now is the time to release this version (and new features) so it accessible by all end users.
The `rolling-update` will create a new replication controller and scale up the replicas as it scales down the replicas in the 'old' replication controller (with version 1.0.0 pods).
 
```
kubectl rolling-update simple-cms --update-period=5s -f rc-new.yml
```

## Watch the update live

If you like, watch the pods with `kubectl`.

```
kubectl get pods --watch
```

And also with a loop, `curl`ing the web page (remember, $IP refers to the service IP as before.

```
while true; do curl http://$IP; sleep 1; done
```

Next up: if time permits, [deploy an Nginx proxy](nginx-ssl-proxy.md) for SSL termination.
