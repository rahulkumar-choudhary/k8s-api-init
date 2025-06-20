# k8s-api-init

`k8s-api-init` is a CLI tool for interacting with Kubernetes clusters. It allows you to list resources such as Pods, Deployments, and Services across multiple contexts and namespaces.

## Usage

Run the tool with the following options:

```
Usage: k8s-api-init [options]
Options:
-c, --context    (default: default)
        Kubernetes context to use
-k, --kubeconfig (default: /Users/{username}/.kube/config)
        Path to kubeconfig file
-n, --namespace  (default: default)
        Kubernetes namespace to use
-r, --resource   (default: all)
        Resource to list: pods, deployments, services, or all
```

### Examples

#### Default Context and Namespace
Run the tool without any options to use the default context and namespace:
```
go run main.go
```

Example Output:
```
cluster2
Cluster: cluster2

Namespace: default
Pods
NAME                                     PHASE        POD IP    
------------------------------------------------------------------
                                 N/A

Deployments
NAME                                     REPLICAS     AVAILABLE 
------------------------------------------------------------------
                                 N/A

Services
NAME                                     TYPE         CLUSTER IP
------------------------------------------------------------------
kubernetes                               ClusterIP    10.96.0.1 
##################################################################
******************************************************************
```

#### All Contexts
Use the `-c all` option to list resources across all contexts:
```
go run main.go -c all
```

Example Output:
```
cluster2
Cluster: cluster1

Namespace: default
Pods
NAME                                     PHASE        POD IP    
------------------------------------------------------------------
                                 N/A

Deployments
NAME                                     REPLICAS     AVAILABLE 
------------------------------------------------------------------
                                 N/A

Services
NAME                                     TYPE         CLUSTER IP
------------------------------------------------------------------
kubernetes                               ClusterIP    10.96.0.1 
##################################################################
******************************************************************
Cluster: cluster2

Namespace: default
Pods
NAME                                     PHASE        POD IP    
------------------------------------------------------------------
                                 N/A

Deployments
NAME                                     REPLICAS     AVAILABLE 
------------------------------------------------------------------
                                 N/A

Services
NAME                                     TYPE         CLUSTER IP
------------------------------------------------------------------
kubernetes                               ClusterIP    10.96.0.1 
##################################################################
******************************************************************
Cluster: cluster3

Namespace: default
Pods
NAME                                     PHASE        POD IP    
------------------------------------------------------------------
                                 N/A

Deployments
NAME                                     REPLICAS     AVAILABLE 
------------------------------------------------------------------
                                 N/A

Services
NAME                                     TYPE         CLUSTER IP
------------------------------------------------------------------
kubernetes                               ClusterIP    10.96.0.1 
##################################################################
******************************************************************
```

#### Specify Namespace
Use the `-n` option to specify a namespace:
```
go run main.go -n my-namespace
```

#### Specify Resource Type
Use the `-r` option to filter resources by type (e.g., `pods`, `deployments`, `services`):
```
go run main.go -r pods
```

## Features
- Supports multiple Kubernetes contexts.
- Lists resources such as Pods, Deployments, and Services.
- Filters by namespace and resource type.
- Reads configuration from a kubeconfig file.

## Requirements
- Go installed on your system.
- A valid kubeconfig file with access to the desired Kubernetes clusters.

## Installation
Clone the repository and build the tool:
```
git clone https://github.com/your-repo/k8s-api-init.git
cd k8s-api-init
go build
```

Run the tool:
```
./k8s-api-init [options]
```

## License
This project is licensed under the MIT License.
