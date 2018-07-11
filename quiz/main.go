package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 10, "time limit for each questions")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	

	problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}


	fmt.Printf("You answered %d questions correctly out of %d\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
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