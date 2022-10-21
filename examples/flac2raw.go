package main

import (
	"fmt"
	"os"

	"github.com/drgolem/go-flac/flac"
)

func main() {
	fmt.Println("example decode flac to wav file")

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: flac2raw <infile.flac> <outfile.raw>")
		fmt.Fprintln(os.Stderr, "play: ffplay -ar 44100 -ac 2 -f s16le <outfile.raw>")
		return
	}

	inFile := os.Args[1]
	outFile := os.Args[2]
	fmt.Printf("infile: %s, outfile: %s\n", inFile, outFile)

	fmt.Printf("libFLAC version: %s\n", flac.GetVersion())

	outBitsPerSample := 16
	outBytesPerSample := outBitsPerSample / 8

	dec, err := flac.NewFlacFrameDecoder(outBitsPerSample)
	if err != nil {
		panic(err)
	}
	defer dec.Delete()

	err = dec.OpenFile(inFile)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	defer dec.Close()

	fmt.Printf("decoder state: %s\n", dec.GetResolvedState())

	fmt.Printf("current sample: %d\n", dec.TellCurrentSample())
	fmt.Printf("total samples: %d\n", dec.TotalSamples())

	rate, channels, bitsPerSample := dec.GetFormat()
	fmt.Printf("Format: [%d:%d:%d]\n", rate, channels, bitsPerSample)

	fOut, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	defer fOut.Close()

	audioSamples := 4 * 1024
	audioBufferBytes := audioSamples * channels * 4 // 4 bytes for 16 or 24 bit samples
	audio := make([]byte, audioBufferBytes)

	for {
		sampleCnt, err := dec.DecodeSamples(audioSamples, audio)
		if err != nil {
			fmt.Printf("ERR: %v\n", err)
			break
		}
		if sampleCnt == 0 {
			break
		}

		bytesToWrite := sampleCnt * channels * outBytesPerSample
		fOut.Write(audio[:bytesToWrite])
	}
	fOut.Sync()
	fmt.Printf("current sample: %d\n", dec.TellCurrentSample())
}
