package main

import (
	"context"
	"fmt"
	"path/filepath"

	apiV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	checkErr(err)

	client, err := kubernetes.NewForConfig(config)
	checkErr(err)

	podClient := client.CoreV1().Pods(apiV1.NamespaceDefault)
	deploymentClient := client.AppsV1().Deployments(apiV1.NamespaceDefault)
	rsClient := client.AppsV1().ReplicaSets(apiV1.NamespaceDefault)

	pods, err := podClient.List(context.TODO(), metav1.ListOptions{})
	checkErr(err)

	deployments, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	checkErr(err)

	rsList, err := rsClient.List(context.TODO(), metav1.ListOptions{})
	checkErr(err)

	fmt.Println(`NAME                                    NODE       STATUS      POD-IP      START
---------------------------------------------------------------------------`)
	for _, pod := range pods.Items {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n", pod.Name, pod.Spec.NodeName, string(pod.Status.Phase),
			pod.Status.PodIP, pod.Status.StartTime.Format("2006-01-02-03 16:05:06"))
	}

	fmt.Println("\n" + `NAME                   READY   UP-TO-DATE   AVAILABLE 
---------------------------------------------------------------------------`)
	for _, d := range deployments.Items {
		fmt.Printf("%s\t%d\t%d\t%d\n", d.Name, d.Status.ReadyReplicas, d.Status.UpdatedReplicas, d.Status.AvailableReplicas)
	}

	fmt.Println("\n" + `NAME                              DESIRED      CURRENT
---------------------------------------------------------------------------`)
	for _, rs := range rsList.Items {
		fmt.Printf("%s\t\t%d\t\t%d\n", rs.Name, *rs.Spec.Replicas, rs.Status.ReadyReplicas)
	}
}
