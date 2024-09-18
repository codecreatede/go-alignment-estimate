package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 *
 Author Gaurav Sablok
 Universitat Potsdam
 Date 2024-9-19

It takes a whole genome aligned files and then estimates the linear block by using the first genome as a reference genome.
You can pass as many genome as a reference genome and it will take the first genome as a reference. This only compares the
single line block. In the alignmentGO package, you can find the specific for defining your own block.

 */

 func main () {

 alignment := flag.String("alignment", "path to the alignment file", "file")

 flag.Parse()

	type alignmentIDStore struct {
		id string
	}

	type alignmentSeqStore struct {
		seq string
	}

	fOpen, err := os.Open(*alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignmentID := []alignmentIDStore{}
	alignmentSeq := []alignmentSeqStore{}
	sequenceSpec := []string{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentID = append(alignmentID, alignmentIDStore{
				id: strings.Replace(string(line), ">", "", -1),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentSeq = append(alignmentSeq, alignmentSeqStore{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceSpec = append(sequenceSpec, string(line))
		}
	}

	counterA := 0
	counterT := 0
	counterG := 0
	counterC := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "A" && string(sequenceSpec[i+1][j]) == "T",
			 || string(sequenceSpec[i+1][j] == "C" || string(sequenceSpec[i+1][j]) == "G" {
					counterA++
				}
				if string(sequenceSpec[i][j]) == "T" && string(sequenceSpec[i+1][j]) == "C",
					|| string(sequenceSpec[i+1][j]) == "G" || string(sequenceSpec[i+1][j]) == "A" {
						counterT++
					}
					if string(sequenceSpec[i][j]) == "C" && string(sequenceSpec[i+1][j]) == "G",
						|| string(sequenceSpec[i+1][j]) == "A" || string(sequenceSpec[i+1][j]) == "T" {
							counterC++
						}
						if string(sequenceSpec[i][j]) == "G" && string(sequenceSpec[i+1][j]) == "A",
							|| string(SequenceSpec[i+1][j]) == "T" || string(sequenceSpec[i+1][j]) == "C" {
								counterG++
							} else {
								continue
							}
		}
	}
	fmt.Printf("The collinearity block of line width 1 for A as a base pattern is %d\n", counterA)
	fmt.Printf("The collinearity block of line width 1 for T as a base pattern is %d\n", counterT)
	fmt.Printf("The collinearity block of line width 1 for G as a base pattern is %d\n", counterG)
	fmt.Printf("The collinearity block of line width 1 for C as a base pattern is %d\n", counterC)
}
