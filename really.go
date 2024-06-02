package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

var sr int = 44100.0

type Wave struct {
	frequency float64
	amplitude float64
	angle     float64
	offset    float64
}

func sinWave(wave Wave) (float64, float64) {
	sample := wave.amplitude * math.Sin(wave.angle)
	var nangle float64 = wave.angle + wave.offset
	return nangle, sample
}

func main() {
	var qs float64
	var time int = 2 // where time is the duration of the sample in seconds
	var osc1 Wave
	osc1.frequency = 440
	osc1.amplitude = 0.5
	osc1.angle = 0
	osc1.offset = (2 * math.Pi * osc1.frequency) / float64(sr)
	file, err := os.Create("waves/test.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	for i := 0; i < sr*time; i++ {
		osc1.angle, qs = sinWave(osc1)
		file.WriteString(fmt.Sprintf("%f", qs) + "\n")
	}
}
