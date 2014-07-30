package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
)

type Game struct {
	DataFileLen1   uint32
	DataFileLen2   uint32
	PatchLevelVers uint32
	MajorVers      uint32
	MinorVers      uint32
	MaxScore       uint32
	StrBit         uint32
	Egmxsc         uint32
	Rlnt           uint32
	Rdesc2         uint32
	Rdesc1         [200]uint32
	Rexit          [200]uint32
	Ractio         [200]uint32
	Rval           [200]uint32
	Rflag          [200]uint32
	Xlnt           uint32
	Travel         [1000]uint32
	Olnt           uint32
	Odesc1         [300]uint32
	Odesc2         [300]uint32
	Odesco         [300]uint32
	Oactio         [300]uint32
	Oflag1         [300]uint32
	Oflag2         [300]uint32
	Ofval          [300]uint32
	Otval          [300]uint32
	Osize          [300]uint32
	Ocapac         [300]uint32
	Oroom          [300]uint32
	Oadv           [300]uint32
	Ocan           [300]uint32
	Oread          [300]uint32
	R2lnt          uint32
	O2             [20]uint32
	R2             [20]uint32
	Clnt           uint32
	Ctick          [30]uint32
	Cactio         [30]uint32
	//supposed to be bitmasks
	Cflag  [30]uint32
	Ccncel [30]uint32
	Vlnt   uint32
	Villns [4]uint32
	Vprob  [4]uint32
	Vopps  [4]uint32
	Vbest  [4]uint32
	Vmelee [4]uint32
	Alnt   uint32
	Aroom  [4]uint32
	Ascore [4]uint32
	Avehic [4]uint32
	Aobj   [4]uint32
	Aactio [4]uint32
	Astren [4]uint32
	Aflag  [4]uint32
	Mbase  uint32
	Mlnt   uint32
	Rtext  [1500]uint32
	StringIndex [300]uint32
}

var dbfile string
var dbline int
var stringsOffset int
var stringsIndex []uint32
var gamedata Game

func main() {
	flag.StringVar(&dbfile, "file", "", "dundat file to dump")

	flag.Parse()

	if dbfile != "" {
		filehandle, err := os.Open(dbfile)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n", dbfile, err.Error())
			return
		}
		defer filehandle.Close()
		gamefile_struct(filehandle, &gamedata)

		fmt.Printf("Datafile Section 1 Size: \033[1;33m%d\033[0m\n", gamedata.DataFileLen1)
		fmt.Printf("Datafile Section 2 Size: \033[1;33m%d\033[0m\n", gamedata.DataFileLen2)
		fmt.Printf("Version: \033[1;33m%d.%d\033[0m\n", gamedata.MajorVers, gamedata.MinorVers)
		fmt.Printf("Patchlevel: \033[1;33m%d\033[0m\n", gamedata.PatchLevelVers)
		fmt.Printf("Max Score: \033[1;33m%d\033[0m\n", gamedata.MaxScore)
		stringsIndex = f_get_ints(filehandle, int(gamedata.DataFileLen2))
		stringsOffset = 4 * (2 + int(gamedata.DataFileLen1) + int(gamedata.DataFileLen2))
		fmt.Printf("String data offset: \033[1;33m%d\033[0m\n", stringsOffset)
		for y := 0; y < 300; y++ {
			simple_print(filehandle, y)
		}
	}
}

/*
This is kind of different, the original used crazy pointer arithmetic
We are reading 4 bytes as a uint32 binary encoding, then casting it back to int
And storing it in the array we return, which is cleaner
*/
func f_get_ints(filehandle *os.File, number int) []uint32 {
	result := make([]uint32, number)
	for x := 0; x < number; x++ {
		bits := make([]byte, 4)
		count, err := filehandle.Read(bits)
		if count < 1 {
			if err != nil {
				switch err {
				case io.EOF:
					return result
				default:
					fmt.Printf("Error reading from file: %s\n", err.Error())
					return result
				}
			}
		} else {
			var val uint32
			buf := bytes.NewReader(bits)
			err := binary.Read(buf, binary.BigEndian, &val)
			if err != nil {
				fmt.Printf("Error reading binary encoding: %s\n", err.Error())
			}
			result[x] = val
		}
	}
	return result
}

/* This is better than f_get_ints, we just pass it a struct to populate from the file */
func gamefile_struct(filehandle *os.File, obj *Game) {
	err := binary.Read(filehandle, binary.BigEndian, obj)
	if err != nil {
		fmt.Printf("Could not read bytes as structure: %s\n", err.Error())
	}
}

func simple_print(filehandle *os.File, index int) {
	start := int64(stringsOffset + int(gamedata.StringIndex[index]))
	result := make([]byte,80)
	filehandle.Seek(start,0)
	for x := 0; x < 79; x++ {
		var mychar byte
		err := binary.Read(filehandle, binary.BigEndian, &mychar)
		if err != nil {
		       return
		} else {
			result[:][x] = mychar
			result[:][x] = byte(int(result[x]) ^ ((index + x + 3) & 0xff))
		}
	}
	fmt.Printf("%s\n",result)
}

