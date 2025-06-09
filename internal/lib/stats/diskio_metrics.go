package statsserver

import (
	"github.com/cri-o/cri-o/internal/lib/sandbox"
	"time"

	"github.com/cri-o/cri-o/internal/config/cgmgr"
	types "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func generateSandboxDiskIoMetrics(sb *sandbox.Sandbox, disk *cgmgr.DiskIoStats) []*types.Metric {

	var diskIoMetrics = []*containerMetric{
		{
			desc: containerBlkIoDeviceUsageTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, serviceByte := range disk.IoServiceBytes {
					total = total + serviceByte.Stats["Total"]
				}
				return metricValues{{
					value:      total,
					labels:     nil,
					metricType: types.MetricType_COUNTER,
				}}
			},
		},
		{
			desc: containerFsIoCurrent,
			valueFunc: func() metricValues {
				var values metricValues
				for _, stats := range disk.IoQueued {
					values = append(values, metricValue{
						value:      stats.Stats["Total"],
						labels:     nil,
						metricType: types.MetricType_GAUGE,
					})
				}
				return values
			},
		},
		{
			desc: containerFsIoTimeSecondsTotal,
			valueFunc: func() metricValues {
				var values metricValues
				for _, stats := range disk.IoServiceTime {
					values = append(values, metricValue{
						value:      stats.Stats["Total"] / uint64(time.Second),
						labels:     nil,
						metricType: types.MetricType_COUNTER,
					})
				}
				return values
			},
		},
		{
			desc: containerFsIoTimeWeightedSecondsTotal,
			valueFunc: func() metricValues {
				var values metricValues
				for _, stats := range disk.IoWaitTime {
					values = append(values, metricValue{
						value:      stats.Stats["Total"] / uint64(time.Second),
						labels:     nil,
						metricType: types.MetricType_COUNTER,
					})
				}
				return values
			},
		},
		{
			desc: containerFsReadsBytesTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoServiceBytes {
					total += stat.Stats["Read"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
		{
			desc: containerFsReadSecondsTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoServiceTime {
					total += stat.Stats["Read"]
				}
				return metricValues{{
					value: total / uint64(time.Second),
				}}
			},
		},
		{
			desc: containerFsReadsMergedTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoMerged {
					total += stat.Stats["Read"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
		{
			desc: containerFsReadsTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoServiced {
					total += stat.Stats["Read"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
		{
			desc: containerFsSectorReadsTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.Sectors {
					total += stat.Stats["Read"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
		{
			desc: containerFsSectorWritesTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.Sectors {
					total += stat.Stats["Write"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
		{
			desc: containerFsWritesBytesTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoServiceBytes {
					total += stat.Stats["Write"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
		{
			desc: containerFsWriteSecondsTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoServiceTime {
					total += stat.Stats["Write"]
				}
				return metricValues{{
					value: total / uint64(time.Second),
				}}
			},
		},
		{
			desc: containerFsWritesMergedTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoMerged {
					total += stat.Stats["Write"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
		{
			desc: containerFsWritesTotal,
			valueFunc: func() metricValues {
				var total uint64
				for _, stat := range disk.IoServiced {
					total += stat.Stats["Write"]
				}
				return metricValues{{
					value: total,
				}}
			},
		},
	}

	return computeSandboxMetrics(sb, diskIoMetrics, "disk")

}
