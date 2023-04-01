package metrics

import (
	"net/netip"

	"github.com/netsampler/goflow2/decoders/netflow"
	"github.com/netsampler/goflow2/decoders/netflowlegacy"
	"github.com/netsampler/goflow2/decoders/sflow"
	flowmessage "github.com/netsampler/goflow2/pb"
	"github.com/netsampler/goflow2/producer"

	"github.com/prometheus/client_golang/prometheus"
)

func CreatePromProducerDefaultWrapper(cfg *producer.ProducerConfig) producer.ProducerInterface {
	return WarpPromProducer(producer.CreateProducerWithConfig(cfg))
}

type PromProducerWrapper struct {
	wrapped producer.ProducerInterface
}

func (p *PromProducerWrapper) Produce(msg interface{}, args *producer.ProduceArgs) ([]producer.ProducerMessage, error) {
	flowMessageSet, err := p.wrapped.Produce(msg, args)
	if err != nil {
		return flowMessageSet, err
	}
	key := args.Src.Addr().String()
	var nfvariant bool
	var versionStr string
	switch packet := msg.(type) {
	case *sflow.Packet:
		agentStr := "unk"
		agentIp, ok := netip.AddrFromSlice(packet.AgentIP)
		if ok {
			agentStr = agentIp.String()
		}

		SFlowStats.With(
			prometheus.Labels{
				"router":  key,
				"agent":   agentStr,
				"version": "5",
			}).
			Inc()

		for _, samples := range packet.Samples {
			typeStr := "unknown"
			countRec := 0
			switch samplesConv := samples.(type) {
			case sflow.FlowSample:
				typeStr = "FlowSample"
				countRec = len(samplesConv.Records)
			case sflow.CounterSample:
				typeStr = "CounterSample"
				if samplesConv.Header.Format == 4 {
					typeStr = "Expanded" + typeStr
				}
				countRec = len(samplesConv.Records)
			case sflow.ExpandedFlowSample:
				typeStr = "ExpandedFlowSample"
				countRec = len(samplesConv.Records)
			}
			SFlowSampleStatsSum.With(
				prometheus.Labels{
					"router":  key,
					"agent":   agentStr,
					"version": "5",
					"type":    typeStr,
				}).
				Inc()

			SFlowSampleRecordsStatsSum.With(
				prometheus.Labels{
					"router":  key,
					"agent":   agentStr,
					"version": "5",
					"type":    typeStr,
				}).
				Add(float64(countRec))
		}

	case *netflowlegacy.PacketNetFlowV5:
		NetFlowStats.With(
			prometheus.Labels{
				"router":  key,
				"version": "5",
			}).
			Inc()
		NetFlowSetStatsSum.With(
			prometheus.Labels{
				"router":  key,
				"version": "5",
				"type":    "DataFlowSet",
			}).
			Add(float64(packet.Count))

	case *netflow.NFv9Packet:
		NetFlowStats.With(
			prometheus.Labels{
				"router":  key,
				"version": "9",
			}).
			Inc()
		recordCommonNetFlowMetrics(9, key, packet.FlowSets)
		nfvariant = true
		versionStr = "9"

	case *netflow.IPFIXPacket:
		NetFlowStats.With(
			prometheus.Labels{
				"router":  key,
				"version": "10",
			}).
			Inc()
		recordCommonNetFlowMetrics(10, key, packet.FlowSets)
		nfvariant = true
		versionStr = "10"
	}

	if nfvariant {
		for _, msg := range flowMessageSet {
			fmsg, ok := msg.(*flowmessage.FlowMessage)
			if !ok {
				continue
			}
			timeDiff := fmsg.TimeReceived - fmsg.TimeFlowEnd
			NetFlowTimeStatsSum.With(
				prometheus.Labels{
					"router":  key,
					"version": versionStr,
				}).
				Observe(float64(timeDiff))
		}
	}

	return flowMessageSet, err
}

func (p *PromProducerWrapper) Close() {
	p.wrapped.Close()
}

// Wraps a producer with metrics
func WarpPromProducer(wrapped producer.ProducerInterface) producer.ProducerInterface {
	return &PromProducerWrapper{
		wrapped: wrapped,
	}
}

// metrics template system