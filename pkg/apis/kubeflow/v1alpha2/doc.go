// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/ankushagarwal/kf-training-controller/pkg/apis/kubeflow
// +k8s:defaulter-gen=TypeMeta
// +groupName=kubeflow.org
package v1alpha2 // import "github.com/ankushagarwal/kf-training-controller/pkg/apis/kubeflow/v1alpha2"
