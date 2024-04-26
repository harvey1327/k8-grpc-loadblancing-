package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hbl88/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	s, err := client.AppsV1().Deployments("pong").GetScale(context.Background(), "pong-dep", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	sc := *s
	sc.Spec.Replicas = sc.Spec.Replicas - 1
	_, err = client.AppsV1().Deployments("pong").UpdateScale(context.TODO(), "pong-dep", &sc, metav1.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}
}
