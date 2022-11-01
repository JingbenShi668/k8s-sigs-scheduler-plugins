package v1beta1

import "k8s.io/apimachinery/pkg/runtime"

//将defaulters functions添加到given scheme
func RegisterDefaults(scheme *runtime.Scheme)  error{
	scheme.AddTypeDefaultingFunc(&NetworkBandwidthArgs{},func(obj interface{}){
		SetObjectDefaultNetworkBandwidthArgs(obj.(*NetworkBandwidthArgs))
	})

	return nil
}

func SetObjectDefaultNetworkBandwidthArgs(in *NetworkBandwidthArgs)  {
	SetDefaultNetworkBandwidthArgs(in)
}


