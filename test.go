package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	// "math"
)

func main() {
	buf := new(bytes.Buffer)
	var pi int = 1
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
}