package base

import (
	"k8s.io/client-go/tools/record"
	"github.com/ankushagarwal/kf-training-controller/pkg/inject/args"
)

type BaseJobController struct {
	args.InjectArgs
	EventRecorder record.EventRecorder
}
