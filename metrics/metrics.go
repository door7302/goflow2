package metrics

import (
	//"strconv"
	//"time"
	"fmt"

	"github.com/netsampler/goflow2/decoders/netflow"
	"github.com/netsampler/goflow2/decoders/netflowlegacy"
	"github.com/netsampler/goflow2/decoders/sflow"
	"github.com/netsampler/goflow2/utils"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	MetricTrafficBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_traffic_bytes",
			Help: "Bytes received by the application.",
		},
		[]string{"remote_ip", "local_ip", "local_port", "type"},
	)
	MetricTrafficPackets = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_traffic_packets",
			Help: "Packets received by the application.",
		},
		[]string{"remote_ip", "local_ip", "local_port", "type"},
	)
	MetricPacketSizeSum = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "flow_traffic_summary_size_bytes",
			Help:       "Summary of packet size.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"remote_ip", "local_ip", "local_port", "type"},
	)
	DecoderStats = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_decoder_count",
			Help: "Decoder processed count.",
		},
		[]string{"worker", "name"},
	)
	DecoderErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_decoder_error_count",
			Help: "Decoder processed error count.",
		},
		[]string{"worker", "name"},
	)
	DecoderTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "flow_summary_decoding_time_us",
			Help:       "Decoding time summary.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"name"},
	)
	DecoderProcessTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "flow_summary_processing_time_us",
			Help:       "Processing time summary.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"name"},
	)
	NetFlowStats = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_nf_count",
			Help: "NetFlows processed.",
		},
		[]string{"router", "version"},
	)
	NetFlowErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_nf_errors_count",
			Help: "NetFlows processed errors.",
		},
		[]string{"router", "error"},
	)
	NetFlowSetRecordsStatsSum = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_nf_flowset_records_sum",
			Help: "NetFlows FlowSets sum of records.",
		},
		[]string{"router", "version", "type"}, // data-template, data, opts...
	)
	NetFlowSetStatsSum = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_nf_flowset_sum",
			Help: "NetFlows FlowSets sum.",
		},
		[]string{"router", "version", "type"}, // data-template, data, opts...
	)
	NetFlowTimeStatsSum = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "flow_process_nf_delay_summary_seconds",
			Help:       "NetFlows time difference between time of flow and processing.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"router", "version"},
	)
	NetFlowTemplatesStats = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_nf_templates_count",
			Help: "NetFlows Template count.",
		},
		[]string{"router", "version", "obs_domain_id", "template_id", "type"}, // options/template
	)
	SFlowStats = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_sf_count",
			Help: "sFlows processed.",
		},
		[]string{"router", "agent", "version"},
	)
	SFlowErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_sf_errors_count",
			Help: "sFlows processed errors.",
		},
		[]string{"router", "error"},
	)
	SFlowSampleStatsSum = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_sf_samples_sum",
			Help: "SFlows samples sum.",
		},
		[]string{"router", "agent", "version", "type"}, // counter, flow, expanded...
	)
	SFlowSampleRecordsStatsSum = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "flow_process_sf_samples_records_sum",
			Help: "SFlows samples sum of records.",
		},
		[]string{"router", "agent", "version", "type"}, // data-template, data, opts...
	)
)

func init() {
	/*prometheus.MustRegister(MetricTrafficBytes)
	prometheus.MustRegister(MetricTrafficPackets)
	prometheus.MustRegister(MetricPacketSizeSum)*/

	//prometheus.MustRegister(DecoderStats)
	prometheus.MustRegister(DecoderErrors)
	//prometheus.MustRegister(DecoderTime)
	//prometheus.MustRegister(DecoderProcessTime)

	//prometheus.MustRegister(NetFlowStats)
	prometheus.MustRegister(NetFlowErrors)
	/*prometheus.MustRegister(NetFlowSetRecordsStatsSum)
	prometheus.MustRegister(NetFlowSetStatsSum)
	prometheus.MustRegister(NetFlowTimeStatsSum)
	prometheus.MustRegister(NetFlowTemplatesStats)

	prometheus.MustRegister(SFlowStats)*/
	prometheus.MustRegister(SFlowErrors)
	//prometheus.MustRegister(SFlowSampleStatsSum)
	//prometheus.MustRegister(SFlowSampleRecordsStatsSum)
}

func PromDecoderWrapper(wrapped utils.DecoderFunc) utils.DecoderFunc {
	return func(msg interface{}) error {
		pkt, ok := msg.(*utils.Message)
		if !ok {
			return fmt.Errorf("flow is not *Message")
		}
		key := pkt.Src.String()

		err := wrapped(msg)
		if err != nil {
			switch err.(type) {
			case *netflowlegacy.ErrorVersion:
				NetFlowErrors.With(
					prometheus.Labels{
						"router": key,
						"error":  "error_version",
					}).
					Inc()
			case *netflow.ErrorTemplateNotFound:
				NetFlowErrors.With(
					prometheus.Labels{
						"router": key,
						"error":  "template_not_found",
					}).
					Inc()
			case *sflow.ErrorVersion:
				SFlowErrors.With(
					prometheus.Labels{
						"router": key,
						"error":  "error_version",
					}).
					Inc()
			case *sflow.ErrorIPVersion:
				SFlowErrors.With(
					prometheus.Labels{
						"router": key,
						"error":  "error_ip_version",
					}).
					Inc()
			case *sflow.ErrorDataFormat:
				SFlowErrors.With(
					prometheus.Labels{
						"router": key,
						"error":  "error_data_format",
					}).
					Inc()
			default:
				// FIXME
				NetFlowErrors.With(
					prometheus.Labels{
						"router": key,
						"error":  "error_decoding",
					}).
					Inc()
				SFlowErrors.With(
					prometheus.Labels{
						"router": key,
						"error":  "error_decoding",
					}).
					Inc()
			}
		}
		return err
	}
}
