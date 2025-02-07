package main

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {

	// 1. read config
	// 2. create client set
	// 3. create factory
	// 4. generate informer
	// 5. run

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		log.Fatal(err)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	// create informer factorys
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("default"))

	informer := factory.Core().V1().ConfigMaps().Informer()
	// write your logic
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("AddFunc")
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Println("UpdateFunc")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("DeleteFunc")
		},
	})
	stopCh := make(chan struct{})
	defer close(stopCh)
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
