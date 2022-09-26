/*	~~~~~~~~~
	k8s informers
	Author: Leo
	Usage: go run main.go
	添加删除pod, 查看pods事件
*/
package main

import (
	"log"
	"path/filepath"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
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

	clientset, err := kubernetes.NewForConfig(config)
	checkErr(err)

	sharedInformer := informers.NewSharedInformerFactory(clientset, time.Minute)
	informer := sharedInformer.Core().V1().Pods().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Add pod: %v\n", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			mObj1 := oldObj.(v1.Object)
			mObj2 := newObj.(v1.Object)
			log.Printf("Update pod: %v\t%v\n", mObj1.GetName(), mObj2.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Delete pod: %v\n", mObj.GetName())
		},
	})

	stopCh := make(chan struct{})
	informer.Run(stopCh)
}
