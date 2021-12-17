package day16

import (
	"bytes"
	"encoding/hex"

	"github.com/icza/bitio"
)

func ReadLiteral(r *bitio.CountReader) (uint64, error) {
	var bits []uint64
	for {
		next, err := r.ReadBool()
		if err != nil {
			return 0, err
		}
		liternalBits, err := r.ReadBits(4)
		if err != nil {
			return 0, err
		}
		bits = append(bits, liternalBits)
		if !next {
			break
		}
	}
	l := uint64(len(bits))
	var literal uint64
	for i, j := uint64(0), l-1; i < l && j >= 0; i++ {
		literal |= bits[i] << uint64(4*j)
		j--
	}
	return literal, nil
}

func ReadNoOfPackets(r *bitio.CountReader) ([]Packet, error) {
	var subpackets []Packet
	noOfSubpackets, err := r.ReadBits(11)
	if err != nil {
		return nil, err
	}
	for i := uint64(0); i < noOfSubpackets; i++ {
		subpacket, err := ReadPacket(r)
		if err != nil {
			return nil, err
		}
		subpackets = append(subpackets, subpacket)
	}
	return subpackets, nil
}

func ReadVariablePackets(r *bitio.CountReader) ([]Packet, error) {
	var subpackets []Packet
	length, _ := r.ReadBits(15)
	end := r.BitsCount + int64(length)
	for r.BitsCount < end {
		subpacket, err := ReadPacket(r)
		if err != nil {
			return nil, err
		}
		subpackets = append(subpackets, subpacket)
	}
	return subpackets, nil
}

func ReadHeaders(r *bitio.CountReader) (uint64, uint64) {
	version, _ := r.ReadBits(3)
	packetType, _ := r.ReadBits(3)
	return version, packetType
}

func ParsePackets(hexdecimal string) (Packet, error) {
	data, _ := hex.DecodeString(hexdecimal)
	r := bitio.NewCountReader(bytes.NewBuffer(data))
	return ReadPacket(r)
}

func ReadPacket(r *bitio.CountReader) (Packet, error) {
	version, packetType := ReadHeaders(r)

	switch packetType {
	case Literal:
		literal, err := ReadLiteral(r)
		if err != nil {
			return nil, err
		}
		return &LiteralPacket{
			version: version,
			literal: literal,
		}, nil
	default:
		isSubpackets, err := r.ReadBool()
		if err != nil {
			return nil, err
		}
		var subpackets []Packet
		if isSubpackets {
			subpackets, err = ReadNoOfPackets(r)
		} else {
			subpackets, err = ReadVariablePackets(r)
		}
		if err != nil {
			return nil, err
		}
		return &OperatorPacket{
			version:    version,
			op:         packetType,
			subpackets: subpackets,
		}, nil
	}
}
