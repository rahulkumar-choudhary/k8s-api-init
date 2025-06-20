package utils

import (
	"flag"
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
)

func ParseFlags() Flags {
	var flags Flags
	cliFlags(&flags.Kubeconfig, "kubeconfig", clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename(), "Path to kubeconfig file")
	cliFlags(&flags.Namespace, "namespace", "default", "Kubernetes namespace to use")
	cliFlags(&flags.Resource, "resource", "all", "Resource to list: pods, deployments, services, or all")
	cliFlags(&flags.Context, "context", "default", "Kubernetes context to use")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", "k8s-api-init")
		fmt.Println("Options:")
		// Print all flags with their default values and usage without shorthand flags
		flag.VisitAll(func(f *flag.Flag) {
			if len(f.Name) == 1 {
				return // skip shorthand flags in usage output
			}
			fmt.Printf("-%s, --%-10s (default: %s)\n", string(f.Name[0]), f.Name, f.DefValue)
			fmt.Printf("\t%s\n", f.Usage)
		})
	}
	flag.Parse()
	return flags
}

func cliFlags(p *string, name string, value string, usage string) {
	shortHandChar := name[0]
	flag.StringVar(p, name, value, usage)
	flag.StringVar(p, string(shortHandChar), *p, "")
}
