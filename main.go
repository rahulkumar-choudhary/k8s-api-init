package main

import (
	"context"

	"github.com/rahulkumar-choudhary/k8s-api-init/utils"
)

func main() {
	// Initialize flags
	options := utils.ParseFlags()

	// The Work function will handle the main logic of the application
	utils.Work(context.Background(), options)
}
