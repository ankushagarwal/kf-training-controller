# Build and run locally
```
GOBIN=${PWD}/bin go install ${PWD#$GOPATH/src/}/cmd/controller-manager
bin/controller-manager --kubeconfig ~/.kube/config

```

# Build docker image
`docker build . -f Dockerfile.controller`

# Generate go code
`kubebuilder generate`

# Generate CRDs
`kubebuilder create config --crds --output hack/crds.yaml`
