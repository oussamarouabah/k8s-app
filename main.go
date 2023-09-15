package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		fmt.Printf("name = %v\ncreated-at = %v\nns = %v\n-----------\n", pod.Name, pod.CreationTimestamp, pod.Namespace)
	}

	deps, err := clientset.AppsV1().Deployments("default").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, dep := range deps.Items {
		fmt.Printf("name = %v\ncreated-at = %v\nns = %v\n-----------\n", dep.Name, dep.CreationTimestamp, dep.Namespace)
	}
}
