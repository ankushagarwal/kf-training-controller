package tfjob

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	kubeflowv1alpha2 "github.com/ankushagarwal/kf-training-controller/pkg/apis/kubeflow/v1alpha2"
	"github.com/ankushagarwal/kf-training-controller/pkg/inject/args"
	"github.com/ankushagarwal/kf-training-controller/pkg/controller/base"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/eventhandlers"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/predicates"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"

)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for TFJob resources goes here.

func (bc *TFJobController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on tfjob.TFJobController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=kubeflow,version=v1alpha2,kind=TFJob,resource=tfjobs
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=service,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:informers:group=core,version=v1,kind=Pod
// +kubebuilder:informers:group=core,version=v1,kind=Service
type TFJobController struct {
	base.BaseJobController
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &TFJobController{
		base.BaseJobController{
			InjectArgs: arguments,
			EventRecorder: arguments.CreateRecorder(
				"TFJobController"),},
	}

	// Create a new controller that will call TFJobController.Reconcile on changes to TFJobs
	gc := &controller.GenericController{
		Name:             "TFJobController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&kubeflowv1alpha2.TFJob{}); err != nil {
		return gc, err
	}

	// Watch Pods
	tfJobsLookup := func(k types.ReconcileKey) (
		interface{}, error) {
		d, err := bc.Clientset.
			KubeflowV1alpha2().
			TFJobs(k.Namespace).
			Get(k.Name, metav1.GetOptions{})
		return d, err
	}
	if err := gc.WatchControllerOf(
		&corev1.Pod{},
		eventhandlers.Path{tfJobsLookup},
		predicates.ResourceVersionChanged); err != nil {
		return gc, err
	}
	return gc, nil
}
