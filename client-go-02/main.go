package main

import (
	"github.com/demo/client-go-02/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	// client set
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("default"))
	ingressInformer := factory.Networking().V1().Ingresses()
	serviceInformer := factory.Core().V1().Services()

	controller := pkg.NewController(clientset, serviceInformer, ingressInformer)

	stopCh := make(chan struct{})
	defer close(stopCh)
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	controller.Run(stopCh)
}
