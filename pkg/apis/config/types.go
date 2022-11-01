package config

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//存放用于解析来自KubeSchedulerConfiguration的信息
type NetworkBandwidthArgs struct {
	metav1.TypeMeta

	Address string //Prometheus server的address
	NetworkInterface string
	TimeRangeInMinutes int64 //TimeRangeInMinutes为时长，聚合network metrics
}

