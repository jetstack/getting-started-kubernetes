# Set up a local Kubernetes cluster

* Deploy a Dockerised Kubernetes cluster (Hyperkube)

## Run an etcd container

```
docker run \
--net=host \
-d gcr.io/google_containers/etcd:2.0.9 \
/usr/local/bin/etcd --addr=127.0.0.1:4001 --bind-addr=0.0.0.0:4001 --data-dir=/var/etcd/data
```

## Run a kubelet master

The kubelet will start the API server, controller manager and scheduler as containers in a pod (declared in `/etc/kubernetes/manifests`).

```
docker run \
--net=host \
-d \
-v /var/run/docker.sock:/var/run/docker.sock \
jetstack/hyperkube:v1.0.3 \
/hyperkube kubelet --api_servers=http://127.0.0.1:8080 --v=2 --address=0.0.0.0 --enable_server --hostname_override=127.0.0.1 --config=/etc/kubernetes/manifests --cluster_dns=10.0.0.10 --cluster_domain=cluster.local
```

## Run a kube-proxy

```
docker run \
-d \
--net=host \
--privileged \
jetstack/hyperkube:v1.0.3 \
/hyperkube proxy --master=http://127.0.0.1:8080 --v=2
```

Note the IP address ($IP).

## See the Docker containers start

```
docker ps
```

## Download the kubectl CLI tool

```
wget https://storage.googleapis.com/kubernetes-release/release/v1.0.3/bin/linux/amd64/kubectl
chmod u+x kubectl
sudo mv kubectl /usr/local/bin/ 
```

## Verify the Kubernetes cluster is running

```
kubectl cluster-info
```

Now to [deploy a first pod](deploy-first-pod.md).

Read more at the Jetstack blog - [Kubernetes: Getting Started With a Local Deployment](http://www.jetstack.io/new-blog/2015/7/6/getting-started-with-a-local-deployment).
