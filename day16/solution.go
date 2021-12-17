package day16

import (
	"bytes"
	"encoding/hex"

	"github.com/icza/bitio"
)

func Evalute(hexstring string) (evaled uint64) {
	data, _ := hex.DecodeString(hexstring)
	r := bitio.NewCountReader(bytes.NewBuffer(data))
	packet, _ := ReadPacket(r)
	return packet.Evalute()
}

func SumAllVersions(hexstring string) (summed uint64) {
	data, _ := hex.DecodeString(hexstring)
	r := bitio.NewCountReader(bytes.NewBuffer(data))
	packet, _ := ReadPacket(r)

	queue := []Packet{packet}
	for len(queue) > 0 {
		packet = queue[0]
		queue = queue[1:]
		summed += packet.Version()
		queue = append(queue, packet.Sub()...)
	}
	return summed
}
