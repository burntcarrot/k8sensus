package main

import (
	"flag"
	"fmt"
	"log"
)

// createLease creates a new lease lock object.
func createLease(leaseName, podName, namespace string) {
	fmt.Println("Creating lease using the following metadata:")
	fmt.Println("Lease Name: " + leaseName)
	fmt.Println("Pod Name: " + podName)
	fmt.Println("Namespace: " + namespace)

	// TODO: use client-go to create a lock and pass metadata
}

func main() {
	var leaseName string
	flag.StringVar(&leaseName, "lease-name", "", "Lease Name (Lock Name)")

	flag.Parse()

	// validate lease name
	if leaseName == "" {
		log.Fatal("Lease Name not found. Provide a valid lease name through --lease-name.")
	}

	fmt.Println("k8sensus is running!")

	// TODO: create a valid lease
	// dummy lease
	createLease(leaseName, "mypod", "my-namespace")
}
