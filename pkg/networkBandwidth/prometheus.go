package networkBandwidth

import (
"github.com/prometheus/client_golang/api"
v1 "github.com/prometheus/client_golang/api/prometheus/v1"
"k8s.io/klog/v2"
"time"
)

const(
	//用于获取node network bandwidth的string template
	nodeMeasureQueryTemplate = "sum_over_time(node_network_receive_bytes_total{kubernetes_node=\"%s\",device=\"%s\"}[%s])"
)

//处理与networkplugin plugin的交互
type PrometheusHandle struct {
	networkInterface string
	timeRange        time.Duration
	address          string  //指向k8s cluster的Prometheus地址
	api              v1.API  //存储Prometheus客户端
}

func NewPrometheus(address string, networkInterface string, timeRange time.Duration)  *PrometheusHandle{
	client, err := api.NewClient(api.Config{
		Address: address,
	})

	if err != nil{
		klog.Fatalf("[NetworkTraffic] Error creating prometheus client: %s", err.Error());
	}

	return &PrometheusHandle{
		networkInterface: networkInterface,
		timeRange: timeRange,
		address: address,
		api: v1.NewAPI(client),
	}
}