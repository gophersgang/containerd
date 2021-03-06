package prometheus

import (
	"github.com/crosbymichael/cgroups"
	metrics "github.com/docker/go-metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var pidMetrics = []*metric{
	{
		name: "pids",
		help: "The limit to the number of pids allowed",
		unit: metrics.Unit("limit"),
		vt:   prometheus.GaugeValue,
		getValues: func(stats *cgroups.Stats) []value {
			if stats.Pids == nil {
				return nil
			}
			return []value{
				{
					v: float64(stats.Pids.Limit),
				},
			}
		},
	},
	{
		name: "pids",
		help: "The current number of pids",
		unit: metrics.Unit("current"),
		vt:   prometheus.GaugeValue,
		getValues: func(stats *cgroups.Stats) []value {
			if stats.Pids == nil {
				return nil
			}
			return []value{
				{
					v: float64(stats.Pids.Current),
				},
			}
		},
	},
}
