package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {

	namespace := "pong"
	deploymentName := "pong-dep"
	scale := int32(1)

	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hbl88/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	s, err := client.AppsV1().Deployments(namespace).GetScale(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	sc := *s
	sc.Spec.Replicas = sc.Spec.Replicas + scale
	_, err = client.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, &sc, metav1.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}
}
