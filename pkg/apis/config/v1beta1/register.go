package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	schedschemev1beta1 "k8s.io/kube-scheduler/config/v1beta1"
)



//GroupName是这个package的group name
const GroupName =  "kubescheduler.config.k8s.io"

//SchemeGroupVersion适用于注册对象的group version
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta1"}

var(
	localSchemeBuilder = &schedschemev1beta1.SchemeBuilder
	AddToScheme = localSchemeBuilder.AddToScheme
)

// addKnownTypes registers known types to the given scheme
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		// &CoschedulingArgs{},
		// &NodeResourcesAllocatableArgs{},
		// &CapacitySchedulingArgs{},
		// &TargetLoadPackingArgs{},
		&NetworkTrafficArgs{},
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