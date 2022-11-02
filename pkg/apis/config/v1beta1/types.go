package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkBandwidthArgs struct {
	metav1.TypeMeta `json:",inline"`

	//Prometheus server的address
	Address            *string `json:"prometheusAddress,omitempty"`
	NetworkInterface   *string `json:"network_interface,,omitempty"`
	//TimeRangeInMinutes为时长，聚合network metrics
	TimeRangeInMinutes *int64 `json:"timeRangeInMinutes,omitempty"`
}






