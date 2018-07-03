package inject

import (
	kubeflowv1alpha2 "github.com/ankushagarwal/kf-training-controller/pkg/apis/kubeflow/v1alpha2"
	rscheme "github.com/ankushagarwal/kf-training-controller/pkg/client/clientset/versioned/scheme"
	"github.com/ankushagarwal/kf-training-controller/pkg/controller/pytorchjob"
	"github.com/ankushagarwal/kf-training-controller/pkg/controller/tfjob"
	"github.com/ankushagarwal/kf-training-controller/pkg/inject/args"
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&kubeflowv1alpha2.PyTorchJob{}, arguments.Informers.Kubeflow().V1alpha2().PyTorchJobs()); err != nil {
			return err
		}
		if err := arguments.ControllerManager.AddInformerProvider(&kubeflowv1alpha2.TFJob{}, arguments.Informers.Kubeflow().V1alpha2().TFJobs()); err != nil {
			return err
		}

		// Add Kubernetes informers

		if c, err := pytorchjob.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		if c, err := tfjob.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &kubeflowv1alpha2.PyTorchJobCRD)
	Injector.CRDs = append(Injector.CRDs, &kubeflowv1alpha2.TFJobCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"kubeflow.org"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "kubeflow.org",
		Version: "v1alpha2",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
