package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the formate of 'question,answer")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file %s\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the csv provided csv file")
	}
	problems := parseLines(lines)
	suffle(problems)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//<-timer.C
	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("problem #%d: %s = \n", i+1, p.q)
		answerch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerch <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerch:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
}

//fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
func suffle(slice []problem) {
	for i := range slice {
		//println(i)
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		//rand.Seed(time.Now().UnixNano())
		//j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
