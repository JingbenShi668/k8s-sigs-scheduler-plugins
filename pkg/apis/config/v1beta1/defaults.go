package v1beta1

//为NetworkBandwidth中的networkInterface和TimeRangeInMinutes设置默认值
func SetDefaultNetworkBandwidthArgs(args *NetworkBandwidthArgs)  {
	if args.TimeRangeInMinutes == nil{
		defaultTime := int64(5)
		args.TimeRangeInMinutes = &defaultTime
	}

	if args.NetworkInterface == nil || *args.NetworkInterface == "" {
		netInterface := "test"
		args.NetworkInterface = &netInterface
	}

}
