// +k8s:defaulter-gen=true

package v1beta1

//为NetworkBandwidth中的networkInterface和TimeRangeInMinutes设置默认值
func SetDefaultNetworkBandwidthArgs(args *NetworkBandwidthArgs)  {
	if args.TimeRangeInMinutes == nil{
		defaultTime := int64(5)
		args.TimeRangeInMinutes = &defaultTime
	}
}
