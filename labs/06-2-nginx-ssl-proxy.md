# Deploy a Nginx SSL proxy

* Generate self-signed certificates and create Kubernetes secrets
* Deploy a replication controller
* Expose pods as a service with a NodePort

## Generate certificates

A certificate and key file will need to be generated, as well as a DHE param.

```
openssl req \
-x509 \
-nodes \
-days 365 \
-newkey rsa:2048 \
-keyout simple-cms.key \
-out simple-cms.crt \
-subj "/CN=simple-cms/O=simple-cms"

```

```
openssl dhparam -out dhparam.pem 2048
```

## Create a Kubernetes secret

The certificate, key and DHE param files should be base64 encoded and edited in the `nginx-secret.yml` file (the 
`-w 0` ensures there are no line breaks as this is not permitted in a secret).


```
cat simple-cms.key | base64 -w 0
```
```
cat simple-cms.crt | base64 -w 0
```
```
cat dheparam.pem | base64 -w 0
```

## Create the replication controller

```
kubectl create -f nginx-ssl-proxy-rc.yml
```

## Check the pods created successfully

```
kubectl get pods
```

## Expose as a service

The service will expose the Nginx pods. To ensure it is accessible from outside the cluster (i.e. on the host), a 
`NodePort` is requested. This is a randomly chosen port that is opened on *all* nodes.

```
kubectl create -f nginx-ssl-proxy.yml                          
```

## Find details of the service

We need to find the `NodePort`. Use `kubectl describe` and a port 80 and port 443 `NodePort` should be stated.

```
kubectl describe service nginx-ssl-proxy
```

## Test it out

Using any one of the node IP addresses and the `NodePort`, point a browser and use https.

```
https://$nodeip:$nodeport
```
