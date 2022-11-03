package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	schedschemev1beta2 "k8s.io/kube-scheduler/config/v1beta2"
)

// GroupName就是在package所要使用的group name
const GroupName = "kubescheduler.config.k8s.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta1"}

var (
	localSchemeBuilder = &schedschemev1beta2.SchemeBuilder
	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = localSchemeBuilder.AddToScheme
)

// addKnownTypes registers known types to the given scheme
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		// &CoschedulingArgs{},
		// &NodeResourcesAllocatableArgs{},
		// &CapacitySchedulingArgs{},
		// &TargetLoadPackingArgs{},
		&NetworkBandwidthArgs{},
	)
	return nil
}

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
	localSchemeBuilder.Register(RegisterDefaults)
	localSchemeBuilder.Register(RegisterConversions)
}