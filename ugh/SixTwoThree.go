package stt

import (
	"fmt"
	"math"
	"os"
)

func sinWave(f float64, an float64, am float64, o float64) (float64, float64) {
	samp := am * math.Sin(2*math.Pi*f*an)
	nan := an + o
	return samp, nan
}

func main() {
	var (
		freq, angle, amp, offset, w float64
	)
	freq = 440
	angle = 0
	amp = 0.5
	offset = (2 * math.Pi * freq) / 100.0
	q, err := os.Create("useless/frown.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		w, angle = sinWave(freq, angle, amp, offset)
		q.WriteString(fmt.Sprintf("%f", w) + " -> " + fmt.Sprintf("%f", float32(w)) + "\n")
	}
}
