package v1alpha2

import (
	"k8s.io/api/core/v1"
)


// KFJobSpec contains the base struct which will be extended by TFJobSpec, PyTorchSpec, etc
type KFJobSpec struct {
	// CleanPodPolicy defines the policy to kill pods after Job is
	// succeeded.
	// Default to Running.
	CleanPodPolicy *CleanPodPolicy `json:"cleanPodPolicy,omitempty"`

	// ReplicaSpecs is map of ReplicaType and ReplicaSpec
	// specifies the replicas to run.
	// For example,
	//   {
	//     "PS": ReplicaSpec,
	//     "Worker": ReplicaSpec,
	//   }
	ReplicaSpecs map[ReplicaType]*ReplicaSpec `json:"replicaSpecs"`
}

// KFJobStatus contains the base struct which will be extended by TFJobStatus, PyTorchJobStatus, etc
type KFJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}


// CleanPodPolicy describes how to deal with pods when the Job is finished.
type CleanPodPolicy string

const (
	CleanPodPolicyUndefined CleanPodPolicy = ""
	CleanPodPolicyAll       CleanPodPolicy = "All"
	CleanPodPolicyRunning   CleanPodPolicy = "Running"
	CleanPodPolicyNone      CleanPodPolicy = "None"
)

// Type of the replica. Case insensitive.
type ReplicaType string

const (
	ReplicaTypePS ReplicaType = "PS"

	ReplicaTypeWorker ReplicaType = "WORKER"

	ReplicaTypeChief ReplicaType = "CHIEF"

	ReplicaTypeEval ReplicaType = "EVALUATOR"
)

type ReplicaSpec struct {
	// Replicas is the desired number of replicas of the given template.
	// If unspecified, defaults to 1.
	Replicas *int32 `json:"replicas,omitempty"`

	// Template is the object that describes the pod that
	// will be created for this Replica. RestartPolicy in PodTemplateSpec
	// will be overide by RestartPolicy in ReplicaSpec
	Template v1.PodTemplateSpec `json:"template,omitempty"`

	// Restart policy for all Replicas within the Job.
	// One of Always, OnFailure, Never and ExitCode.
	// Default to Never.
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty"`
}


// RestartPolicy describes how the Replicas should be restarted.
// Only one of the following restart policies may be specified.
// If none of the following policies is specified, the default one
// is RestartPolicyAlways.
type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"

	// `ExitCode` policy means that user should add exit code by themselves,
	// operator will check these exit codes to
	// determine the behavior when an error occurs:
	// - 1-127: permanent error, do not restart.
	// - 128-255: retryable error, will restart the pod.
	RestartPolicyExitCode RestartPolicy = "ExitCode"
)
