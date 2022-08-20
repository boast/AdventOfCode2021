package main

import (
	"fmt"
	"math"
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day19Part1() int {
	file, err := os.Open("assets/day19.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	scanners := ParseScanner(lines)

	alignedScanners := make(map[int]struct{})

	alignedScanners[0] = struct{}{}
	scanners[0].p = &position{}

	for len(alignedScanners) < len(scanners) {
		for i, s1 := range scanners {
			for j, s2 := range scanners {
				if i == j {
					continue
				}
				if _, hasS1 := alignedScanners[i]; !hasS1 {
					continue
				}
				if _, hasS2 := alignedScanners[j]; hasS2 {
					continue
				}

				intersection := s1.compare(&s2)

				if intersection == nil {
					continue
				}

				s1.align(&s2, intersection)
				alignedScanners[j] = struct{}{}
			}
		}
	}

	beacons := make(map[position]struct{})

	for _, s := range scanners {
		for _, sig := range s.signals {
			beacons[sig.p] = struct{}{}
		}
	}

	return len(beacons)
}

func Day19Part2() int {
	file, err := os.Open("assets/day19.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	scanners := ParseScanner(lines)

	alignedScanners := make(map[int]*position)

	alignedScanners[0] = &position{}
	scanners[0].p = &position{}

	for len(alignedScanners) < len(scanners) {
		for i, s1 := range scanners {
			for j, s2 := range scanners {
				if i == j {
					continue
				}
				if _, hasS1 := alignedScanners[i]; !hasS1 {
					continue
				}
				if _, hasS2 := alignedScanners[j]; hasS2 {
					continue
				}

				intersection := s1.compare(&s2)

				if intersection == nil {
					continue
				}

				s1.align(&s2, intersection)
				alignedScanners[j] = s2.p
			}
		}
	}

	max := float64(0)
	for _, p1 := range alignedScanners {
		for _, p2 := range alignedScanners {
			max = math.Max(max, math.Abs(float64(p1.x-p2.x))+math.Abs(float64(p1.y-p2.y))+math.Abs(float64(p1.z-p2.z)))
		}
	}

	return int(max)
}

type scanner struct {
	signals []*signal
	p       *position
}

type position struct {
	x, y, z int
}

type signal struct {
	p         position
	id        int
	relatives map[string]int
}

type signalMatch struct {
	fingerprint string
	id1, id2    int
}

type scannerCompare struct {
	s1, s2       *signal
	intersection []signalMatch
}

func (s *scanner) addSignal(x, y, z int) {
	newSignal := signal{p: position{x: x, y: y, z: z}, id: len(s.signals), relatives: make(map[string]int)}

	for _, otherSignal := range s.signals {
		otherSignal.align(&newSignal)
	}

	s.signals = append(s.signals, &newSignal)
}

func (s *scanner) compare(otherScanner *scanner) *scannerCompare {
	for _, otherSignal := range otherScanner.signals {
		for _, thisSignal := range s.signals {
			intersection := otherSignal.compare(thisSignal)

			if len(intersection) >= 11 {
				return &scannerCompare{s1: otherSignal, s2: thisSignal, intersection: intersection}
			}
		}
	}

	return nil
}

func (s *scanner) align(otherScanner *scanner, data *scannerCompare) {
	for _, matches := range data.intersection {
		relativeHere := s.signals[matches.id2]
		dx0 := data.s2.p.x - relativeHere.p.x
		dy0 := data.s2.p.y - relativeHere.p.y
		dz0 := data.s2.p.z - relativeHere.p.z

		relativeThere := otherScanner.signals[matches.id1]
		dx1 := data.s1.p.x - relativeThere.p.x
		dy1 := data.s1.p.y - relativeThere.p.y
		dz1 := data.s1.p.z - relativeThere.p.z

		if math.Abs(float64(dx0)) == math.Abs(float64(dy0)) || math.Abs(float64(dz0)) == math.Abs(float64(dy0)) || math.Abs(float64(dx0)) == math.Abs(float64(dz0)) {
			continue
		}

		rotationMap := make([]int, 9)

		switch dx0 {
		case dx1:
			rotationMap[0] = 1
			break
		case -dx1:
			rotationMap[0] = -1
			break
		case dy1:
			rotationMap[3] = 1
			break
		case -dy1:
			rotationMap[3] = -1
			break
		case dz1:
			rotationMap[6] = 1
			break
		case -dz1:
			rotationMap[6] = -1
			break
		}

		switch dy0 {
		case dx1:
			rotationMap[1] = 1
			break
		case -dx1:
			rotationMap[1] = -1
			break
		case dy1:
			rotationMap[4] = 1
			break
		case -dy1:
			rotationMap[4] = -1
			break
		case dz1:
			rotationMap[7] = 1
			break
		case -dz1:
			rotationMap[7] = -1
			break
		}

		switch dz0 {
		case dx1:
			rotationMap[2] = 1
			break
		case -dx1:
			rotationMap[2] = -1
			break
		case dy1:
			rotationMap[5] = 1
			break
		case -dy1:
			rotationMap[5] = -1
			break
		case dz1:
			rotationMap[8] = 1
			break
		case -dz1:
			rotationMap[8] = -1
			break
		}

		for _, currentSignal := range otherScanner.signals {
			oldPosition := position{x: currentSignal.p.x, y: currentSignal.p.y, z: currentSignal.p.z}

			currentSignal.p.x = oldPosition.x*rotationMap[0] + oldPosition.y*rotationMap[3] + oldPosition.z*rotationMap[6]
			currentSignal.p.y = oldPosition.x*rotationMap[1] + oldPosition.y*rotationMap[4] + oldPosition.z*rotationMap[7]
			currentSignal.p.z = oldPosition.x*rotationMap[2] + oldPosition.y*rotationMap[5] + oldPosition.z*rotationMap[8]
		}

		otherScanner.p = &position{
			x: data.s2.p.x - data.s1.p.x,
			y: data.s2.p.y - data.s1.p.y,
			z: data.s2.p.z - data.s1.p.z,
		}

		for _, currentSignal := range otherScanner.signals {
			currentSignal.p.x += otherScanner.p.x
			currentSignal.p.y += otherScanner.p.y
			currentSignal.p.z += otherScanner.p.z
		}

		break
	}
}

func (s *signal) align(otherSignal *signal) {
	dX := math.Abs(float64(s.p.x - otherSignal.p.x))
	dY := math.Abs(float64(s.p.y - otherSignal.p.y))
	dZ := math.Abs(float64(s.p.z - otherSignal.p.z))

	fingerPrint := strings.Join([]string{
		fmt.Sprintf("%.5f", math.Sqrt(math.Pow(dX, 2)+math.Pow(dY, 2)+math.Pow(dZ, 2))),
		fmt.Sprintf("%g", math.Min(dX, math.Min(dY, dZ))),
		fmt.Sprintf("%g", math.Max(dX, math.Max(dY, dZ))),
	}, ",")

	s.relatives[fingerPrint] = otherSignal.id
	otherSignal.relatives[fingerPrint] = s.id
}

func (s *signal) compare(otherSignal *signal) []signalMatch {
	signalMatches := make([]signalMatch, 0)

	for fingerprint, id := range s.relatives {
		if otherId, exists := otherSignal.relatives[fingerprint]; exists {
			signalMatches = append(signalMatches, signalMatch{fingerprint: fingerprint, id1: id, id2: otherId})
		}
	}

	return signalMatches
}

func ParseScanner(lines []string) []scanner {
	scanners := make([]scanner, 0)
	var currentScanner scanner

	for _, line := range lines {
		if line == "" {
			scanners = append(scanners, currentScanner)
			continue
		}

		if strings.Contains(line, "scanner") {
			currentScanner = scanner{signals: make([]*signal, 0)}
			continue
		}

		beaconCoordinates := strings.Split(line, ",")

		x, err := strconv.Atoi(beaconCoordinates[0])
		panic.Check(err)
		y, err := strconv.Atoi(beaconCoordinates[1])
		panic.Check(err)
		z, err := strconv.Atoi(beaconCoordinates[2])
		panic.Check(err)

		currentScanner.addSignal(x, y, z)
	}

	return append(scanners, currentScanner)
}
