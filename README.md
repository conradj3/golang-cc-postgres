# Run in K8s

- Build the producer, consumer and postgres images
- Push all three images to a docker registry that your cluster can access
- Update the image tags in `values.yaml` for each helm chart
- Install keda on the cluster - `az aks update --resource-group (your group) --name (your cluster) --enable-keda`
- Create `keda-demo` namespace
- Configure ingress
- `helm install` the postgres chart
- `helm install` the producer chart
- `helm install` the consumer chart
- If ingress is not enable, port forward the producer service
- Update `Send-Message.ps1` to use your ingress or port-forwarded service

# Ingress

I created an ingress resource to the producer, update the hostname if you would like to do the same. Otherwise, disable the ingress entirely and port forward the producer service.

# Security

The postgres setup script creates three users with the password 'pass' which is obviously not secure. These can be updated along with the associated connection strings if needed.