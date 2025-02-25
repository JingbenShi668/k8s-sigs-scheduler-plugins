package networkBandwidth

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"k8s.io/klog/v2"
	"time"
)

const(
	//用于获取node network bandwidth的string template
	nodeMeasureQueryTemplate = "sum_over_time(node_network_receive_bytes_total{kubernetes_node=\"%s\"}[%s])"
	//nodeMeasureQueryTemplate = "sum_over_time(node_network_receive_bytes_total{kubernetes_node=\"%s\",device=\"%s\"}[%s])"
)

//处理与networkplugin plugin的交互
type PrometheusHandle struct {
	timeRange        time.Duration
	address          string  //指向k8s cluster的Prometheus地址
	api              v1.API  //存储Prometheus客户端
}

//创建PrometheusHandle实例
func NewPrometheus(address string, timeRange time.Duration)  *PrometheusHandle{
	client, err := api.NewClient(api.Config{
		Address: address,
	})

	if err != nil{
		klog.Fatalf("[NetworkTraffic] Error creating prometheus client: %s", err.Error());
	}

	return &PrometheusHandle{
		timeRange: timeRange,
		address: address,
		api: v1.NewAPI(client),
	}
}

//查询NodeBandWidth pressure
func (p *PrometheusHandle) GetNodeBandwidthMeasure(node string)  (*model.Sample, error){
	query := getNodeBandwidthQuery(node,p.timeRange);
	res,err := p.query(query);
	if err!=nil {
		return nil, fmt.Errorf("[NetworkTraffic] Error querying prometheus: %w", err)
	}

	nodeMeasure := res.(model.Vector);
	if len(nodeMeasure)!=1{
		return nil, fmt.Errorf("[NetworkTraffic] Invalid response, expected 1 value, got %d", len(nodeMeasure))
	}

	return nodeMeasure[0],nil
}

func getNodeBandwidthQuery(node string,timeRange time.Duration)  string{
	return fmt.Sprintf(nodeMeasureQueryTemplate, node, timeRange);
}

func (p *PrometheusHandle) query(query string) (model.Value, error){
	results, warnings, err := p.api.Query(context.Background(), query, time.Now());

	if len(warnings) > 0 {
		klog.Warningf("[NetworkTraffic] Warnings: %v\n", warnings);
	}

	return results,err;
}



