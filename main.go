package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/rahul/.kube/config", "Path to the kubeconfig file")
	namespace := flag.String("namespace", "default", "Kubernetes namespace to query")
	resource := flag.String("resource", "all", "Resource to list: pods, deployments, services, or all")
	flag.Parse()

	if !flagPassed("resource") {
		fmt.Println("Usage: ./k8s-api-init --kubeconfig <path> --namespace <name> --resource <pods|deployments|services|all>")
		fmt.Println("Example: ./k8s-api-init --kubeconfig /path/to/kubeconfig --namespace kube-system --resource pods")
		fmt.Println("Note: we can skip the --kubeconfig if specified in the code already.")
		os.Exit(0)
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(fmt.Errorf("failed to load kubeconfig: %w", err))
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Errorf("failed to create clientset: %w", err))
	}

	ctx := context.Background()
	res := strings.ToLower(*resource)

	if res == "pods" || res == "all" {
		pods, err := clientset.CoreV1().Pods(*namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to list pods: %w", err))
		}
		fmt.Printf("\nPods in namespace '%s':\n", *namespace)
		fmt.Printf("%-40s %-12s %-15s\n", "NAME", "PHASE", "POD IP")
		fmt.Println(strings.Repeat("-", 70))
		for _, pod := range pods.Items {
			fmt.Printf("%-40s %-12s %-15s\n", pod.Name, pod.Status.Phase, pod.Status.PodIP)
		}
	}

	if res == "deployments" || res == "all" {
		deployments, err := clientset.AppsV1().Deployments(*namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to list deployments: %w", err))
		}
		fmt.Printf("\nDeployments in namespace '%s':\n", *namespace)
		fmt.Printf("%-30s %-10s %-10s\n", "NAME", "REPLICAS", "AVAILABLE")
		fmt.Println(strings.Repeat("-", 50))
		for _, d := range deployments.Items {
			fmt.Printf("%-30s %-10d %-10d\n", d.Name, *d.Spec.Replicas, d.Status.AvailableReplicas)
		}
	}

	if res == "services" || res == "all" {
		services, err := clientset.CoreV1().Services(*namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to list services: %w", err))
		}
		fmt.Printf("\nServices in namespace '%s':\n", *namespace)
		fmt.Printf("%-30s %-15s %-15s\n", "NAME", "TYPE", "CLUSTER IP")
		fmt.Println(strings.Repeat("-", 60))
		for _, svc := range services.Items {
			fmt.Printf("%-30s %-15s %-15s\n", svc.Name, svc.Spec.Type, svc.Spec.ClusterIP)
		}
	}
}

// function checks if a specific flag (e.g., --resource) was explicitly set by the user on the command line.
func flagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

