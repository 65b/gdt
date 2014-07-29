package main

import (
	"fmt"
	"flag"
	)

var key int
var plaintext string

func main() {
     flag.StringVar(&plaintext,"text","ulysses","text to encrypt/decrypt")
     flag.IntVar(&key,"key",0,"index key")
     flag.Parse()
     line := []byte(plaintext)
     txcrypt(key,line)
     fmt.Printf("%s\n", line)
}

func txcrypt(r int, line []byte) {
     k := len(line)
     var x int
     for i := 1 ; i <= k; i++ {
     	 x = (r & 31) + i
	 line[:][i-1] = byte(int(line[:][i-1]) ^ x)
     }
}