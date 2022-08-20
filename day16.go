package main

import (
	"fmt"
	"os"
	"panic"
	"readers"
	"strconv"
)

type Packet struct {
	Version      int8
	TypeId       PacketType
	Value        int
	LengthTypeId bool
	Length       int
	Subpackages  []Packet
}

type PacketType int

const (
	Sum PacketType = iota
	Product
	Min
	Max
	Literal
	Greater
	Less
	Equal
)

func Day16Part1() int {
	file, err := os.Open("assets/day16.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	binary := ""

	for _, i := range lines[0] {
		binary += HexToBin(string(i))
	}

	p, _ := ParsePacket(binary)

	return SumVersion(p)
}

func Day16Part2() int {
	file, err := os.Open("assets/day16.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	binary := ""

	for _, i := range lines[0] {
		binary += HexToBin(string(i))
	}

	p, _ := ParsePacket(binary)

	return ExecPacket(p)
}

func ExecPacket(p Packet) int {
	switch p.TypeId {
	case Sum:
		sum := 0

		for _, subpackage := range p.Subpackages {
			sum += ExecPacket(subpackage)
		}

		return sum
	case Product:
		product := ExecPacket(p.Subpackages[0])

		for i := 1; i < len(p.Subpackages); i++ {
			product *= ExecPacket(p.Subpackages[i])
		}

		return product
	case Min:
		min := ExecPacket(p.Subpackages[0])

		for i := 1; i < len(p.Subpackages); i++ {
			if newMin := ExecPacket(p.Subpackages[i]); newMin < min {
				min = newMin
			}
		}

		return min
	case Max:
		max := ExecPacket(p.Subpackages[0])

		for i := 1; i < len(p.Subpackages); i++ {
			if newMax := ExecPacket(p.Subpackages[i]); newMax > max {
				max = newMax
			}
		}

		return max
	case Literal:
		return p.Value
	case Greater:
		if ExecPacket(p.Subpackages[0]) > ExecPacket(p.Subpackages[1]) {
			return 1
		} else {
			return 0
		}
	case Less:
		if ExecPacket(p.Subpackages[0]) < ExecPacket(p.Subpackages[1]) {
			return 1
		} else {
			return 0
		}
	case Equal:
		if ExecPacket(p.Subpackages[0]) == ExecPacket(p.Subpackages[1]) {
			return 1
		} else {
			return 0
		}
	default:
		panic.Panic("Could not parse packet")
		return -1
	}
}

func SumVersion(p Packet) int {
	sum := int(p.Version)

	for _, subpackage := range p.Subpackages {
		sum += SumVersion(subpackage)
	}

	return sum
}

func ParsePacket(binary string) (Packet, string) {
	var p Packet

	version, err := strconv.ParseUint(binary[:3], 2, 3)
	panic.Check(err)
	p.Version = int8(version)

	typeId, err := strconv.ParseUint(binary[3:6], 2, 3)
	panic.Check(err)
	p.TypeId = PacketType(typeId)

	cursor := 6

	switch p.TypeId {
	case Literal:
		size := 0
		value := ""
		for true {
			flag := binary[cursor]
			value += binary[cursor+1 : cursor+5]
			cursor += 5
			size += 4
			if flag == "0"[0] {
				break
			}
		}

		v, err := strconv.ParseUint(value, 2, size)
		panic.Check(err)
		p.Value = int(v)
		break
	default:
		p.LengthTypeId = binary[6] != "0"[0]
		cursor++
		p.Subpackages = make([]Packet, 0)

		if p.LengthTypeId {
			countSubpackets, err := strconv.ParseUint(binary[7:18], 2, 11)
			panic.Check(err)
			cursor += 11

			p.Length = int(countSubpackets)
			subpacketsBinary := binary[cursor:]

			for i := 0; i < p.Length; i++ {
				subpacket, subpacketsBinaryNew := ParsePacket(subpacketsBinary)
				p.Subpackages = append(p.Subpackages, subpacket)
				subpacketsBinary = subpacketsBinaryNew
			}

			return p, subpacketsBinary
		} else {
			lengthSubpackets, err := strconv.ParseUint(binary[7:22], 2, 15)
			panic.Check(err)
			cursor += 15

			p.Length = int(lengthSubpackets)
			subpacketsBinary := binary[cursor : cursor+p.Length]

			for len(subpacketsBinary) > 0 {
				subpacket, subpacketsBinaryNew := ParsePacket(subpacketsBinary)
				p.Subpackages = append(p.Subpackages, subpacket)
				subpacketsBinary = subpacketsBinaryNew
			}
			return p, binary[cursor+p.Length:]
		}
	}

	return p, binary[cursor:]
}

func HexToBin(hex string) string {
	ui, err := strconv.ParseUint(hex, 16, 4)
	panic.Check(err)

	return fmt.Sprintf("%04b", ui)
}
