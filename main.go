package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/rahul/.kube/config", "kubeconfig file location")
	namespace := flag.String("namespace", "default", "namespace")
	flag.Parse() 

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(fmt.Errorf("failed to load kubeconfig: %w", err))
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Errorf("failed to create clientset: %w", err))
	}

	ctx := context.Background()
	
	pods, err := clientset.CoreV1().Pods(*namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to list pods: %w", err))
	}

	fmt.Printf("Pods from %s namespace are: \n", *namespace)
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
	
	deployments, err := clientset.AppsV1().Deployments(*namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to list deployments: %w", err))
	}

	fmt.Printf("\nDeployments from %s namespace are: \n", *namespace)
	for _, d := range deployments.Items {
		fmt.Printf("- %s\n", d.Name)
	}
}