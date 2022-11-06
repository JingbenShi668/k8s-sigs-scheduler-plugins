// +k8s:defaulter-gen=true

package v1beta1

//为NetworkBandwidth中的networkInterface和TimeRangeInMinutes设置默认值
func SetDefaultNetworkBandwidthArgs(args *NetworkBandwidthArgs)  {
	if args.Address == nil {
		defaultAdress := "http://8.218.69.59:9090"
		args.Address = &defaultAdress
	}
	
	if args.TimeRangeInMinutes == nil{
		defaultTime := int64(5) //默认prometheus range time为5m
		args.TimeRangeInMinutes = &defaultTime
	}
}
