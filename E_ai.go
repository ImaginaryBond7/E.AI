package main

import (
	"encoding/binary"
	"fmt"
)

type DecodedData struct {
	Short1      uint16
	Chars12     string
	SingleByte  byte
	Chars8      string
	Short2      uint16
	Chars15     string
	Long4       uint32
}

func DecodePacket(packet []byte) (*DecodedData, error) {
	if len(packet) != 44 {
		return nil, fmt.Errorf("Invalid packet size. Expected 44 bytes, but got %d", len(packet))
	}

	data := &DecodedData{}

	data.Short1 = binary.BigEndian.Uint16(packet[0:2])
	data.Chars12 = string(packet[2:14])
	data.SingleByte = packet[14]
	data.Chars8 = string(packet[15:23])
	data.Short2 = binary.BigEndian.Uint16(packet[23:25])
	data.Chars15 = string(packet[25:40])
	data.Long4 = binary.BigEndian.Uint32(packet[40:44])

	return data, nil
}

func main() {
	packet := []byte{
		'\x04', '\xD2', '\x6B', '\x65', '\x65', '\x70', '\x64', '\x65', '\x63', '\x6F', '\x64', '\x69', '\x6E', '\x67', '\x38', '\x64', '\x6F', '\x6E', '\x74', '\x73', '\x74', '\x6F', '\x70', '\x03', '\x15', '\x63', '\x6F', '\x6E', '\x67', '\x72', '\x61', '\x74', '\x75', '\x6C', '\x61', '\x74', '\x69', '\x6F', '\x6E', '\x73', '\x07', '\x5B', '\xCD', '\x15',
	}

	decodedData, err := DecodePacket(packet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Decoded struct: {%d, %q, %d, %q, %d, %q, , %d}\n",
		decodedData.Short1, decodedData.Chars12,
		decodedData.SingleByte, decodedData.Chars8,
		decodedData.Short2, decodedData.Chars15,
		decodedData.Long4)
}
