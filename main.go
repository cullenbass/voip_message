package main

import (
	"github.com/zaf/g711"
	"go.uber.org/zap"
	// "io/ioutil"
	"os"
	"github.com/youpy/go-wav"
	"io"
	"encoding/binary"
	"bytes"
)

func main() {
	logger, err := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	out, err := os.Create("./out.raw")
	defer out.Close()
	if err != nil {
		sugar.Fatal("Failed to create out.raw")
	}

	file, err := os.Open("./message.wav")
	if err != nil {
		sugar.Fatal("Failed to read message.wav.", err)
	}
	defer file.Close()
	reader := wav.NewReader(file)
	sugar.Info(reader.Format())
	rawBytes := new(bytes.Buffer)
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}

		for _, sample := range samples {
			raw := reader.IntValue(sample, 0)
			err = binary.Write(rawBytes, binary.LittleEndian, int16(raw))
			if err != nil {
				sugar.Fatal(err)
			}
		}
	}
	sugar.Info("Len of rawBytes", rawBytes.Len())
	encoder, err := g711.NewUlawEncoder(out, 0)
	if err != nil {
		sugar.Fatal("Failed to create encoder.", err)
	}
	i, err := encoder.Write(rawBytes.Bytes())
	// i, err := out.Write(rawBytes.Bytes())
	if err != nil {
		sugar.Error("Failed to encode to ULAW.", err)
	}
	sugar.Info("Bytes written ", i)
	out.Sync()
}
