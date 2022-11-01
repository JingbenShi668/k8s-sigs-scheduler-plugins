package networkBandwidth

import (
	"fmt"
	"golang.org/x/net/context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	framework "k8s.io/kubernetes/pkg/scheduler/framework"
	"sigs.k8s.io/k8s-sigs-scheduler-plugins/pkg/apis/config"
	"time"
)
type NetworkBandwidth struct {
	handle     framework.Handle
	prometheus *PrometheusHandle
}

//用于Registry和configurations的plugin name
const name = "NetworkBandwidth"

var _ = framework.ScorePlugin(&NetworkBandwidth{})

func (n *NetworkBandwidth)  Name() string{
	return name;
}

//初始化新的NetworkBandwidth plugin, 并返回
func New(obj runtime.Object, h framework.Handle) (framework.Plugin, error)  {
	args,ok := obj.(*config.NetworkBandwidthArgs)
	if !ok {
		return nil, fmt.Errorf("[NetworkTraffic] want args to be of type NetworkTrafficArgs, got %T", obj);
	}

	klog.Infof("[NetworkTraffic] args received. NetworkInterface: %s; TimeRangeInMinutes: %d, Address: %s", args.NetworkInterface, args.TimeRangeInMinutes, args.Address);

	return &NetworkBandwidth{
		handle: h,
		prometheus: NewPrometheus(args.Address, args.NetworkInterface, time.Minute*time.Duration(args.TimeRangeInMinutes)),
	},nil
}

func (n *NetworkBandwidth) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status)  {
	nodeBandwidth,err := n.prometheus.GetNodeBandwidthMeasure(nodeName)
	if err!= nil{
		return 0,framework.NewStatus(framework.Error,fmt.Sprintf("error getting node bandwidth measure: %s", err))
	}

	klog.Infof("[NetworkTraffic] node '%s' bandwidth: %s", nodeName, nodeBandwidth.Value)
	return int64(nodeBandwidth.Value),nil
}

func (n *NetworkBandwidth)  ScoreExtensions() framework.ScoreExtensions{
	return n
}

//归一化node评分
func (n *NetworkBandwidth)  NormalizeScore(ctx context.Context, state *framework.CycleState, p *v1.Pod, scores framework.NodeScoreList) *framework.Status{
	var highscore int64
	for _,node := range scores{
		if highscore<node.Score {
			highscore = node.Score
		}
	}
	for i,node:= range scores{
		scores[i].Score = framework.MaxNodeScore - (node.Score*framework.MaxNodeScore/highscore)
	}

	klog.Infof("[NetworkTraffic] Nodes final score: %v", scores)
	return nil
}