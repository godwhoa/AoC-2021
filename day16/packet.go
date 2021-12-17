package day16

import "math"

const (
	Sum         = 0
	Product     = 1
	Min         = 2
	Max         = 3
	Literal     = 4
	GreaterThan = 5
	LessThan    = 6
	EqualTo     = 7
)

type Packet interface {
	Version() uint64
	Type() uint64
	Sub() []Packet
	Evalute() uint64
}

type OperatorPacket struct {
	op         uint64
	version    uint64
	subpackets []Packet
}

func (o *OperatorPacket) Type() uint64 {
	return o.op
}

func (o *OperatorPacket) Version() uint64 {
	return o.version
}

func (o *OperatorPacket) Sub() []Packet {
	return o.subpackets
}

func (o *OperatorPacket) Evalute() uint64 {
	switch o.op {
	case Sum:
		sum := uint64(0)
		for _, sub := range o.subpackets {
			sum += sub.Evalute()
		}
		return sum
	case Product:
		product := uint64(1)
		for _, sub := range o.subpackets {
			product *= sub.Evalute()
		}
		return product
	case Min:
		min := uint64(math.MaxUint64)
		for _, sub := range o.subpackets {
			if v := sub.Evalute(); v < min {
				min = v
			}
		}
		return min
	case Max:
		max := uint64(0)
		for _, sub := range o.subpackets {
			if v := sub.Evalute(); v > max {
				max = v
			}
		}
		return max
	case GreaterThan:
		if o.subpackets[0].Evalute() > o.subpackets[1].Evalute() {
			return 1
		}
		return 0
	case LessThan:
		if o.subpackets[0].Evalute() < o.subpackets[1].Evalute() {
			return 1
		}
		return 0
	case EqualTo:
		if o.subpackets[0].Evalute() == o.subpackets[1].Evalute() {
			return 1
		}
		return 0
	default:
		return 0
	}
}

type LiteralPacket struct {
	version uint64
	literal uint64
}

func (p *LiteralPacket) Type() uint64 {
	return Literal
}

func (p *LiteralPacket) Version() uint64 {
	return p.version
}

func (p *LiteralPacket) Sub() []Packet {
	return []Packet{}
}

func (p *LiteralPacket) Evalute() uint64 {
	return p.literal
}
