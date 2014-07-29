package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
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
}

var dbfile string
var dbline int
var dboffset int

func main() {
	flag.StringVar(&dbfile, "file", "", "dundat file to dump")
	flag.IntVar(&dboffset, "offset", 45268, "strings offset in datafile")

	flag.Parse()

	if dbfile != "" {
		filehandle, err := os.Open(dbfile)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n", dbfile, err.Error())
			return
		}

		var gamedata Game

		gamefile_struct(filehandle, &gamedata)

		fmt.Printf("Datafile Section 1 Size: \033[1;33m%d\033[0m\n", gamedata.DataFileLen1)
		fmt.Printf("Datafile Section 2 Size: \033[1;33m%d\033[0m\n", gamedata.DataFileLen2)
		fmt.Printf("Version: \033[1;33m%d.%d\033[0m\n", gamedata.MajorVers, gamedata.MinorVers)
		fmt.Printf("Patchlevel: \033[1;33m%d\033[0m\n", gamedata.PatchLevelVers)

	}
}

/* This is better than f_get_ints, we just pass it a struct to populate from the file */
func gamefile_struct(filehandle *os.File, obj *Game) {
	err := binary.Read(filehandle, binary.BigEndian, obj)
	if err != nil {
		fmt.Printf("Could not read bytes as structure: %s\n", err.Error())
	}
}
