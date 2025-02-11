package pkg

import (
	v3 "k8s.io/client-go/informers/core/v1"
	v4 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	v2 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	client        kubernetes.Interface
	serviceLister v1.ServiceLister
	ingressLister v2.IngressLister
	queue         workqueue.RateLimitingInterface
}

func (c *Controller) enqueue(obj interface{}) {

}

func (c Controller) updateServiceFunc(obj interface{}, obj2 interface{}) {

}

func (c Controller) addServiceFunc(obj interface{}) {
	c.enqueue(obj)
}

func (c Controller) delIngressFunc(obj interface{}) {

}

func (c *Controller) Run(stopCh <-chan struct{}) {

	<-stopCh
}

// NewController
func NewController(client kubernetes.Interface, serviceInformer v3.ServiceInformer, ingressInformer v4.IngressInformer) *Controller {

	c := &Controller{
		client:        client,
		serviceLister: serviceInformer.Lister(),
		ingressLister: ingressInformer.Lister(),
	}

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addServiceFunc,
		UpdateFunc: c.updateServiceFunc,
	})

	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.delIngressFunc,
	})
	return c
}
