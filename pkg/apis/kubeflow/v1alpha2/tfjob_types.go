package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the TFJob resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TFJobSpec defines the desired state of TFJob
type TFJobSpec struct {
	KFJobSpec
}

// TFJobStatus defines the observed state of TFJob
type TFJobStatus struct {
	KFJobStatus
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TFJob
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=tfjobs
type TFJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TFJobSpec   `json:"spec,omitempty"`
	Status TFJobStatus `json:"status,omitempty"`
}
