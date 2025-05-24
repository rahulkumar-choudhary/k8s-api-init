# k8s-api-init

A simple Go CLI tool that connects to a Kubernetes cluster using `client-go` and lists all Pods and Deployments from a specified namespace. (more on the tool comming soon..)

## Usage

```bash
go mod init github.com/rahulkumar-choudhary/k8s-api-init
kubectl version
# According to your kubernetes version and the compatibility matrix of the client-go library install the client-go library.
go get k8s.io/client-go@v0.33.1
go mod tidy
go build
./k8s-api-init --kubeconfig /path/to/kubeconfig --namespace kube-system
```

## Flags

* `--kubeconfig`: Path to your kubeconfig file (default: `/Users/rahul/.kube/config`) 
* `--namespace`: Kubernetes namespace (default: `default`)

Note: provide full path to your kubeconfig. 

## Compatibility

* `client-go`: v0.33.1 (Kubernetes 1.31)
* `go`: 1.21+

## 🧩 Troubleshooting

If you see missing `go.sum` errors, run:

```bash
go get k8s.io/client-go@v0.33.1
go mod tidy
```
