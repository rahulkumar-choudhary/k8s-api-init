package utils

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/tools/clientcmd"
)

// sortedContexts retrieves and sorts the context names from the kubeconfig file.
func sortedContexts(flagPath string, userContext string) ([]string, error) {
	config, err := clientcmd.LoadFromFile(flagPath)
	if err != nil {
		return nil, err
	}
	// if default context is specified, return only that context
	fmt.Println(config.CurrentContext)
	if userContext == "default" {
		return []string{config.CurrentContext}, nil
	}
	// if userContext is "all", "*", or "a", return all contexts
	contexts := make([]string, 0, len(config.Contexts))
	for _, context := range config.Contexts {
		contexts = append(contexts, context.Cluster)
	}
	sort.Strings(contexts)
	userContext = strings.ToLower(userContext)
	if userContext == "all" || userContext == "*" || userContext == "a" {
		return contexts, nil
	}
	if userContext == "default" {
		return []string{config.CurrentContext}, nil
	}
	for _, context := range config.Contexts {
		if context.Cluster == userContext {
			return []string{context.Cluster}, nil
		}
	}
	return nil, fmt.Errorf("context %s not found in kubeconfig file", userContext)
}

// Work is the main function that initializes the clientset for each context
// and logs the data for the specified resources in the given namespace.
// It iterates through all contexts in the kubeconfig file and performs the logging.
func Work(ctx context.Context, options Flags) {
	contexts, err := sortedContexts(options.Kubeconfig, options.Context)
	if err != nil {
		fmt.Printf("Error loading contexts from kubeconfig file: %v\n", err)
		return
	}

	if len(contexts) == 0 {
		fmt.Println("No contexts found in kubeconfig file.")
		return
	}
	for _, context := range contexts {
		fmt.Println("Cluster:", context)
		clientset, err := contextBasedClientset(options.Kubeconfig, context)
		if err != nil {
			fmt.Printf("Error creating clientset for context %q: %v\n", context, err)
			continue
		}
		logData(ctx, options, clientset)
	}
}

// contextBasedClientset creates a Kubernetes clientset for the specified context
// using the kubeconfig file.
func contextBasedClientset(kubeConfigPath, contextName string) (*kubernetes.Clientset, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = kubeConfigPath
	configOverrides := &clientcmd.ConfigOverrides{
		CurrentContext: contextName,
	}

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides,
	)

	// Build the REST client config
	restConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to build REST config for context %q: %w", contextName, err)
	}

	// Create the Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Clientset for context %q: %w", contextName, err)
	}
	return clientset, nil
}

func logData(ctx context.Context, flag Flags, clientset *kubernetes.Clientset) error {
	namespaces, err := namespaceList(ctx, flag.Namespace, clientset)
	if err != nil {
		return err
	}
	flag.Resource = strings.ToLower(flag.Resource)
	logPattern := "%-40s %-12s %-10s\n"

	for _, namespace := range namespaces {
		fmt.Printf("\nNamespace: %s\n", namespace)
		if flag.Resource == "all" || flag.Resource == "pods" {
			pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("failed to list pods in namespace %s: %w", namespace, err)
			}
			fmt.Printf("Pods\n")
			fmt.Printf(logPattern, "NAME", "PHASE", "POD IP")
			repeatPattern("-", patternLength)
			for _, pod := range pods.Items {
				fmt.Printf(logPattern, pod.Name, pod.Status.Phase, pod.Status.PodIP)
			}
			logEmptyMessage(len(pods.Items))
		}

		if flag.Resource == "all" || flag.Resource == "deployments" {
			deployments, err := clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				panic(fmt.Errorf("failed to list deployments: %w", err))
			}
			fmt.Printf("\nDeployments\n")
			fmt.Printf(logPattern, "NAME", "REPLICAS", "AVAILABLE")
			repeatPattern("-", patternLength)
			for _, d := range deployments.Items {
				fmt.Printf(logPattern, d.Name, strconv.Itoa(int(*d.Spec.Replicas)), strconv.Itoa(int(d.Status.AvailableReplicas)))
			}
			logEmptyMessage(len(deployments.Items))
		}

		if flag.Resource == "all" || flag.Resource == "services" {
			services, err := clientset.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("failed to list services in namespace %s: %w", namespace, err)
			}
			fmt.Printf("\nServices\n")
			fmt.Printf(logPattern, "NAME", "TYPE", "CLUSTER IP")
			repeatPattern("-", patternLength)
			for _, svc := range services.Items {
				fmt.Printf(logPattern, svc.Name, svc.Spec.Type, svc.Spec.ClusterIP)
			}
			logEmptyMessage(len(services.Items))
		}
		repeatPattern("#", patternLength)
	}
	repeatPattern("*", patternLength)
	return nil
}

func namespaceList(ctx context.Context, namespace string, clientset *kubernetes.Clientset) ([]string, error) {
	ns, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces: %w", err)
	}
	var namespaces []string
	for _, n := range ns.Items {
		namespaces = append(namespaces, n.Name)
	}
	if namespace == "" || namespace == "all" || namespace == "*" || namespace == "a" {
		sort.Strings(namespaces)
		return namespaces, nil
	}
	for _, n := range namespaces {
		if n == namespace {
			return []string{namespace}, nil
		}
	}
	return nil, fmt.Errorf("namespace %s not found", namespace)
}
