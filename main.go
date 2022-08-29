package main

import (
	"context"
	"fmt"

	"time"
  
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)

	namespace := "default"
   
	GetPods(clientset, ctx, namespace)
	fmt.Println("pods successfully displayed ")
	 time.Sleep(9999 * time.Second)
		


}

func GetPods(clientset *kubernetes.Clientset, ctx context.Context,
	namespace string) {

	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	fmt.Println(" ---- Pods from namespace ", namespace, " ----")
	if len(pods.Items) > 0 {
		for _, pod := range pods.Items {

			fmt.Printf("%s \n", pod.Name)

		}
	} else {
		fmt.Println("No resources found in namespace", namespace)
	}

	if err != nil {
		fmt.Println(err)
	}
}
