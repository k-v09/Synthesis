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
	fnr := bufio.NewReader(os.Stdin)
	fmt.Println("/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/")
	fmt.Println("Name of the file? (excluding file extension)")
	fmt.Print("->")
	unName, _ := fnr.ReadString('\n')
	var name string = "waves/" + unName + ".wav"
	w, err := os.Create(name)
	var enc = wav.NewEncoder(w, sr, 32, 1, 1) //fmt.Printf("%b\n", math.Float64bits(52.0))
	return enc, w, err
}

func findwLoop(r *bufio.Reader, waveCase int) int {
	for {
		fmt.Println("What type of wave?")
		fmt.Print("->")
		text, _ := r.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hiya cutie ;)")
			findwLoop(r, waveCase)
		} else if strings.Compare("sine", text) == 0 || strings.Compare("sin", text) == 0 {
			waveCase = 0
			return waveCase
		} else if strings.Compare("square", text) == 0 || strings.Compare("sq", text) == 0 {
			waveCase = 1
			return waveCase
		} else if strings.Compare("triangle", text) == 0 || strings.Compare("t", text) == 0 {
			waveCase = 2
			return waveCase
		} else {
			fmt.Println("INVALID WAVE TYPE")
			fmt.Println("Input should be lowercase")
			findwLoop(r, waveCase)
		}
	}
}

func Input() int {
	var tip int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Syne Shell")
	fmt.Println("/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/")
	findwLoop(reader, tip)
	return tip
}

func main() {
	var (
		qs   float64
		time int = 2 // where time is the duration of the sample in seconds
		osc  Wave
	)
	osc.frequency = 440
	osc.amplitude = 0.5
	osc.angle = 0
	osc.offset = (2 * math.Pi * osc.frequency) / float64(sr)
	wtype := Input()
	code, out, err := inclusion()
	if err != nil {
		panic(fmt.Sprintf("couldn't open audio file - %v", err))
	}
	switch wtype {
	case 0:
		{
			for i := 0; i < sr*time; i++ {
				osc.angle, qs = sinWave(osc)
				code.WriteFrame(float32(qs))
			}
		}
	case 1:
		{
			for i := 0; i < sr*time; i++ {
				osc.angle, qs = sqWave(osc)
				code.WriteFrame(float32(qs))
			}
		}
	case 2:
		{
			for i := 0; i < sr*time; i++ {
				osc.angle, qs = tWave(osc, 1.0)
				code.WriteFrame(float32(qs))
			}
		}
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
