package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in the format of 'question,answer'")
	flag.Parse()
	
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file name: %s\n", *csvFilename))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV files")
	}


	problems := parseLines(lines) 
	//fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == p.a {
			correct += 1
			fmt.Println("Correct!!")
		} else {
			fmt.Println("Wrong!!")
		}
	}

	fmt.Printf("You answered %d questions correctly out of %d\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}