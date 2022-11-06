package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=true

//NetworkBandwidthArgs用于配置NetworkBandwidth plugin
type NetworkBandwidthArgs struct {
	metav1.TypeMeta `json:",inline"`

	//Prometheus server的address
	Address            *string `json:"prometheusAddress,omitempty"`
	//TimeRangeInMinutes为时长，聚合network metrics
	TimeRangeInMinutes *int64 `json:"timeRangeInMinutes,omitempty"`
}






