package pytorchjob

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	kubeflowv1alpha2 "github.com/ankushagarwal/kf-training-controller/pkg/apis/kubeflow/v1alpha2"
	kubeflowv1alpha2client "github.com/ankushagarwal/kf-training-controller/pkg/client/clientset/versioned/typed/kubeflow/v1alpha2"
	kubeflowv1alpha2informer "github.com/ankushagarwal/kf-training-controller/pkg/client/informers/externalversions/kubeflow/v1alpha2"
	kubeflowv1alpha2lister "github.com/ankushagarwal/kf-training-controller/pkg/client/listers/kubeflow/v1alpha2"

	"github.com/ankushagarwal/kf-training-controller/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for PyTorchJob resources goes here.

func (bc *PyTorchJobController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on pytorchjob.PyTorchJobController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=kubeflow,version=v1alpha2,kind=PyTorchJob,resource=pytorchjobs
type PyTorchJobController struct {
	// INSERT ADDITIONAL FIELDS HERE
	pytorchjobLister kubeflowv1alpha2lister.PyTorchJobLister
	pytorchjobclient kubeflowv1alpha2client.KubeflowV1alpha2Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	pytorchjobrecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &PyTorchJobController{
		pytorchjobLister: arguments.ControllerManager.GetInformerProvider(&kubeflowv1alpha2.PyTorchJob{}).(kubeflowv1alpha2informer.PyTorchJobInformer).Lister(),

		pytorchjobclient:   arguments.Clientset.KubeflowV1alpha2(),
		pytorchjobrecorder: arguments.CreateRecorder("PyTorchJobController"),
	}

	// Create a new controller that will call PyTorchJobController.Reconcile on changes to PyTorchJobs
	gc := &controller.GenericController{
		Name:             "PyTorchJobController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&kubeflowv1alpha2.PyTorchJob{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a PyTorchJob Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the PyTorchJobController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
