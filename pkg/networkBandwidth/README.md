# Overview
This folder holds the NetworkBandwidth plugin implemented.

## Example scheduler config

```yaml
apiVersion: kubescheduler.config.k8s.io/v1beta1
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: "/etc/kubernetes/scheduler.conf"
profiles:
- schedulerName: default-scheduler
  plugins:
    score:
      enabled:
      - name: NetworkBandwidth
      disabled:
      - name: "*"
  pluginConfig:
  - name: NetworkBandwidth
    args:
      prometheusAddress: "Your own Prometheus address on k8s cluser"
      timeRangeInMinutes: 3
```

## Demo


