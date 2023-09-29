# TSA aka Tanzu Sample Application

[![CI](https://github.com/bzhtux/tanzu-sample-app/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/bzhtux/tanzu-sample-app/actions/workflows/ci.yml)

## Registry credentials

See this kubernetes doc [section](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/).

As an example you can create a registry secret by providing credentials on the command line as below:

```shell
kubectl create secret docker-registry regcred --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>
```

where:

* `<your-registry-server` is your Private Docker Registry FQDN. Use "https://index.docker.io/v1/" for DockerHub.
* `your-name` is your Docker username.
* `your-pword` is your Docker password.
* `your-email` is your Docker email.

## Storage

TSA use PVC to store SQL data (SQLite file). So `tsa` will be deploy as a `StatefulSet` including PVC object.

## Deployment

For the impatient you can run the following command within your terminal:

```shell
kubectl apply -f https://raw.githubusercontent.com/bzhtux/tsa/main/manifest.yml
```

## Access

Point your browser to the ogress url mentioned in the manifest file:

```text
https://sub.fqdn
```
