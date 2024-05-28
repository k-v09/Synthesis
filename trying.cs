using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;


 class Filer {
    static void Main(string[] args) {
        String path = @"D:\Example.txt";
        if (File.Exists(path)) {
            Console.WriteLine("File Exists");
        }
        Console.ReadKey();
    }
}

// where 44100 describes the sample rate
const int sr = 44100;
class SinWave {
    // -1 <= amp <= 1
    float frequency, amplitude, angle = 0.0f, offset = 0.0;
    public SinWave (float freq, float amp) {
        this.frequency = freq;
        this.amplitude = amp;
        offset = 2*System.Math.pi*frequency/sr
    }
    float process() {
        atuto sample =  amplitude * System.Math.sin(angle);
        angle += offset;
        return sample;
    }

}

public static void main(String[] args) {
    // time is the time in seconds of how long the sample audio will be
    int time = 2;
    SinWave sw = new SinWave(440, 0.5);
    for (int i = 0; i < sr * time; i++) {
        sw.process();
    }
    System.out.println("Hello world");
}