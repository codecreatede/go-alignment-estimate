package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

func main() {
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

	counterAT := 0
	counterAG := 0
	counterAC := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "A" && string(sequenceSpec[i+1][j]) == "T" {
				counterAT++
			}
			if string(sequenceSpec[i][j]) == "A" && string(sequenceSpec[i+1][j]) == "C" {
				counterAG++
			}
			if string(sequenceSpec[i][j]) == "A" && string(sequenceSpec[i+1][j]) == "G" {
				counterAC++
			}
		}
	}

	counterTG := 0
	counterTC := 0
	counterTA := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "T" && string(sequenceSpec[i+1][j]) == "G" {
				counterTA++
			}
			if string(sequenceSpec[i][j]) == "T" && string(sequenceSpec[i+1][j]) == "C" {
				counterTC++
			}
			if string(sequenceSpec[i][j]) == "T" && string(sequenceSpec[i+1][j]) == "A" {
				counterTA++
			}
		}
	}

	counterGC := 0
	counterGA := 0
	counterGT := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "G" && string(sequenceSpec[i+1][j]) == "C" {
				counterGC++
			}
			if string(sequenceSpec[i][j]) == "G" && string(sequenceSpec[i+1][j]) == "A" {
				counterGA++
			}
			if string(sequenceSpec[i][j]) == "G" && string(sequenceSpec[i+1][j]) == "T" {
				counterGT++
			}
		}
	}

	counterCA := 0
	counterCT := 0
	counterCG := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "C" && string(sequenceSpec[i+1][j]) == "A" {
				counterCA++
			}
			if string(sequenceSpec[i][j]) == "C" && string(sequenceSpec[i+1][j]) == "T" {
				counterCT++
			}
			if string(sequenceSpec[i][j]) == "C" && string(sequenceSpec[i+1][j]) == "G" {
				counterCG++
			}
		}
	}

	fmt.Printf(
		"The collinearity block for A as a base pattern and T as a mismatch is %d\n",
		counterAT,
	)
	fmt.Printf("The collinearity block for A as a base pattern G as a mismatch is  %d\n", counterAG)
	fmt.Printf(
		"The collinearity block for A as a base pattern and C as a mismatch is %d\n",
		counterAC,
	)

	fmt.Printf(
		"The collinearity block for T as a base pattern and G as a mismatch is %d\n",
		counterTG,
	)
	fmt.Printf("The collinearity block for T as a base pattern C as a mismatch is  %d\n", counterTC)
	fmt.Printf(
		"The collinearity block for T as a base pattern and A as a mismatch is %d\n",
		counterTA,
	)

	fmt.Printf(
		"The collinearity block for G as a base pattern and C as a mismatch is %d\n",
		counterGC,
	)
	fmt.Printf("The collinearity block for G as a base pattern A as a mismatch is  %d\n", counterGA)
	fmt.Printf(
		"The collinearity block for G as a base pattern and T as a mismatch is %d\n",
		counterGT,
	)

	fmt.Printf(
		"The collinearity block for C as a base pattern and A as a mismatch is %d\n",
		counterCA,
	)
	fmt.Printf("The collinearity block for C as a base pattern T as a mismatch is  %d\n", counterCT)
	fmt.Printf(
		"The collinearity block for C as a base pattern and G as a mismatch is %d\n",
		counterCG,
	)
}
