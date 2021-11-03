package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

// create a new clientset
var client *clientset.Clientset

// createLease creates a new lease lock object.
func createLease(leaseName, podName, namespace string) *resourcelock.LeaseLock {
	fmt.Println("Creating lease using the following metadata:")
	fmt.Println("Lease Name: " + leaseName)
	fmt.Println("Pod Name: " + podName)
	fmt.Println("Namespace: " + namespace)

	return &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      leaseName,
			Namespace: namespace,
		},
		Client: client.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: podName,
		},
	}
}

// elect helps in electing a new leader by using the leaderelection API.
func elect(lock *resourcelock.LeaseLock, ctx context.Context, id string) {

	// TODO: investigate lease duration and renew deadlines
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:          lock,
		LeaseDuration: 15 * time.Second,
		RenewDeadline: 10 * time.Second,
		RetryPeriod:   2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(c context.Context) {
				sampleTask()
			},
			OnStoppedLeading: func() {
				log.Println("Evicted as leader: finding new leaders..")
			},
			OnNewLeader: func(identity string) {
				if identity == id {
					log.Println("Still the leader.")
					return
				}
				log.Println("New leader is: " + identity)
			},
		},
		ReleaseOnCancel: true,
	})
}

// sampleTask is ran when a LeaderElector starts running
func sampleTask() {
	for {
		log.Println("k8sensus is running sample task.")
		time.Sleep(10 * time.Second)
	}
}

func main() {
	var leaseName string
	flag.StringVar(&leaseName, "lease-name", "", "Lease Name (Lock Name)")

	flag.Parse()

	// validate lease name
	if leaseName == "" {
		log.Fatalln("Lease Name not found. Provide a valid lease name through --lease-name.")
	}

	fmt.Println("🚢🏗️ k8sensus is running!")

	config, err := rest.InClusterConfig()
	// dies if no config is given
	client = clientset.NewForConfigOrDie(config)

	if err != nil {
		log.Fatalln("Failed to get kube config.")
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a lease lock
	lock := createLease(leaseName, "mypod", "my-namespace")

	// run leader election
	elect(lock, ctx, "mypod")
}
