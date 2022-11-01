package config

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	schedscheme "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

//GroupName是在package中所要使用的group name
const GroupName = "kubescheduler.config.k8s.io"

//SchemeGroupVersion是用于注册对象的group version
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName,Version: runtime.APIVersionInternal}

var(
	localSchemeBuilder = &schedscheme.SchemeBuilder
	//AddToScheme是把API group & version注册给scheme的global function
	AddToScheme = localSchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme)  {
	scheme.AddKnownTypes(SchemeGroupVersion, &NetworkBandwidthArgs{},)
}


