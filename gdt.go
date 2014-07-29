package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"encoding/binary"
)

var dbfile string
var dbline int
var dboffset int

var datafile_len1 int
var datafile_len2 int
var patchlevel_vers int
var major_vers int
var minor_vers int

func main() {
	flag.StringVar(&dbfile, "file", "", "dundat file to dump")
	flag.IntVar(&dboffset, "offset",45268,"strings offset in datafile")
	
	flag.Parse()
	
	if dbfile != "" {
		filehandle,err := os.Open(dbfile)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n",dbfile,err.Error())
			return
		}
		
		//Read first five int values from database
		//Print the useful info
		gamedata := f_get_ints(filehandle, 5)
		datafile_len1 = gamedata[0]
		datafile_len2 = gamedata[1]
		major_vers = gamedata[2]
		minor_vers = gamedata[3]
		patchlevel_vers = gamedata[4]
		
		fmt.Printf("Datafile Section 1 Size: \033[1;33m%d\033[0m\n",datafile_len1)
		fmt.Printf("Datafile Section 2 Size: \033[1;33m%d\033[0m\n",datafile_len2)
		fmt.Printf("Version: \033[1;33m%d.%d\033[0m\n",major_vers,minor_vers)
		fmt.Printf("Patchlevel: \033[1;33m%d\033[0m\n",patchlevel_vers)
		
	}
}

/*
This is kind of different, the original used crazy pointer arithmetic 
We are reading 4 bytes as a uint32 binary encoding, then casting it back to int
And storing it in the array we return, which is cleaner
*/
func f_get_ints(filehandle *os.File, number int) []int {
	result := make([]int,number)
	for x := 0; x < number; x++ {
		bits := make([]byte,4)
		count, err := filehandle.Read(bits)
		if count < 1 {
			if err != nil {
				fmt.Printf("Error reading from file: %s\n", err.Error())
			}
		} else {
			var val uint32
			fmt.Printf("read %d bytes",count)
			fmt.Printf("bytes %s\n",bits)
			buf := bytes.NewReader(bits)
			err := binary.Read(buf, binary.BigEndian, &val)
			if err !=nil {
				fmt.Printf("Error reading binary encoding: %s\n", err.Error())
			}
			fmt.Printf("ints %d\n",val)
			result[x]=int(val)
		}
	}
	return result
}

/*
Not sure if needed, this is some crypto function from the game engine
*/
func txcrypt(r int, line []byte) {
	k := len(line)
	var x int
	for i := 1; i <= k; i++ {
		x = (r & 31) + i
		line[:][i-1] = byte(int(line[:][i-1]) ^ x)
	}
}

