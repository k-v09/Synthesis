package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/go-audio/wav"
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
func sqWave(wave Wave) (float64, float64) {
	ng, sw := sinWave(wave)
	sample := wave.amplitude * (sw / math.Abs(sw))
	return ng, sample
}
func tWave(wave Wave, waveFactor float64) (float64, float64) {
	nn, bw := sinWave(wave)
	sample := wave.amplitude * (2 / math.Pi) * math.Asin(waveFactor*bw)
	return nn, sample
}

func inclusion() (*wav.Encoder, *os.File, error) {
	w, err := os.Create("waves/wave0.wav")
	var enc = wav.NewEncoder(w, sr, 32, 1, 1) //fmt.Printf("%b\n", math.Float64bits(52.0))
	return enc, w, err
}

func Input() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("------------------------")

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hiya cutie ;)")
			break
		}
	}
}

func main() {
	var (
		qs   float64
		time int = 2 // where time is the duration of the sample in seconds
		osc1 Wave
	)
	osc1.frequency = 440
	osc1.amplitude = 0.5
	osc1.angle = 0
	osc1.offset = (2 * math.Pi * osc1.frequency) / float64(sr)
	Input()
	code, out, err := inclusion()
	if err != nil {
		panic(fmt.Sprintf("couldn't open audio file - %v", err))
	}

	for i := 0; i < sr*time; i++ {
		osc1.angle, qs = sinWave(osc1)
		// osc2.angle, q2 = sqWave(osc2)
		// osc3.angle, q3 = tWave(osc3, 1.0)
		code.WriteFrame(float32(qs))
	}
	code.Close()
	out.Close()
	out, err = os.Open("waves/wave0.wav")
	if err != nil {
		panic(err)
	}
	d2 := wav.NewDecoder(out)
	d2.ReadInfo()
	fmt.Println("New file ->", d2)
	out.Close()
}
