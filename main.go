package main

import (
  "fmt"
  "os"
  "flag"
	"encoding/binary"
	"github.com/nstruthers/pihex"
)

func main() {
	index := flag.Int64("i", 0, "start index")
	span := flag.Int64("s", 1000, "span")
	flag.Parse()

  if *index < 0 {
    fmt.Printf("Starting index must be >= 0\n")
    os.Exit(1)
  }

  if *span < 2 || *span % 2 == 1 {
    fmt.Printf("Span must be >= 2, and must be even\n")
    os.Exit(1)
  }

	curByte := byte(0)

	for n := *index; n < *index+*span; n++ {
		if (n % 2) == 0 {
			curByte = pihex.Digit(int64(n)) << 4
		} else {
			curByte = curByte | pihex.Digit(int64(n))
      binary.Write(os.Stdout, binary.LittleEndian, curByte)
		}
	}
}
