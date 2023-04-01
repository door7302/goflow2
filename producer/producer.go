package producer

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/netsampler/goflow2/decoders/netflow"
	"github.com/netsampler/goflow2/decoders/netflowlegacy"
	"github.com/netsampler/goflow2/decoders/sflow"
)

// Interface of the messages
// Passed to the Formatter
// Can be protobuf
type ProducerMessage interface {
}

type ProducerInterface interface {
	Produce(msg interface{}, args *ProduceArgs) ([]ProducerMessage, error)
	// add commit where the data can be sent to a sync pool
	// Commit([]ProducerMessage) error
	Close()
}

type ProduceArgs struct {
	//MessageFactory     interface{} // for use with sync pool?
	SamplingRateSystem SamplingRateSystem

	Src            netip.AddrPort
	Dst            netip.AddrPort
	SamplerAddress netip.Addr
	TimeReceived   time.Time
}

type ProtoProducer struct {
	cfgMapped *producerConfigMapped
}

func (p *ProtoProducer) enrich(flowMessageSet []ProducerMessage, cb func(msg *ProtoProducerMessage)) {
	for _, msg := range flowMessageSet {
		fmsg, ok := msg.(*ProtoProducerMessage)
		if !ok {
			continue
		}
		cb(fmsg)
	}
}

func (p *ProtoProducer) Produce(msg interface{}, args *ProduceArgs) (flowMessageSet []ProducerMessage, err error) {
	tr := uint64(args.TimeReceived.Unix())
	sa, _ := args.SamplerAddress.MarshalBinary()
	switch msgConv := msg.(type) {
	case *netflowlegacy.PacketNetFlowV5: //todo: rename PacketNetFlowV5
		flowMessageSet, err = ProcessMessageNetFlowLegacy(msgConv)

		p.enrich(flowMessageSet, func(fmsg *ProtoProducerMessage) {
			fmsg.SamplerAddress = sa
		})
	case *netflow.NFv9Packet:
		flowMessageSet, err = ProcessMessageNetFlowV9Config(msgConv, args.SamplingRateSystem, p.cfgMapped)

		p.enrich(flowMessageSet, func(fmsg *ProtoProducerMessage) {
			fmsg.TimeReceived = tr
			fmsg.SamplerAddress = sa
		})
	case *netflow.IPFIXPacket:
		flowMessageSet, err = ProcessMessageIPFIXConfig(msgConv, args.SamplingRateSystem, p.cfgMapped)

		p.enrich(flowMessageSet, func(fmsg *ProtoProducerMessage) {
			fmsg.TimeReceived = tr
			fmsg.SamplerAddress = sa
		})
	case *sflow.Packet:
		flowMessageSet, err = ProcessMessageSFlowConfig(msgConv, p.cfgMapped)

		p.enrich(flowMessageSet, func(fmsg *ProtoProducerMessage) {
			fmsg.TimeReceived = tr
			fmsg.TimeFlowStart = tr
			fmsg.TimeFlowEnd = tr
		})
	default:
		return flowMessageSet, fmt.Errorf("flow not recognized")
	}

	return flowMessageSet, err
}

func (p *ProtoProducer) Close() {}

func CreateProducerWithConfig(cfg *ProducerConfig) ProducerInterface {
	return &ProtoProducer{
		cfgMapped: mapConfig(cfg),
	}
}

// Producer that keeps the same format
// as the original flow samples.
// This can be used for debugging (eg: getting NetFlow Option Templates)
type RawProducer struct {
}

func (p *RawProducer) Produce(msg interface{}, args *ProduceArgs) ([]ProducerMessage, error) {
	// should return msg wrapped
	// []*interface{msg,}
	return []ProducerMessage{msg}, nil
}

func (p *RawProducer) Close() {}