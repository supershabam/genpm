package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func words(in io.Reader) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(in)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return out
}

func isNPM(words []string) bool {
	if len(words) != 3 {
		return false
	}
	if !strings.HasPrefix(words[0], "n") {
		return false
	}
	if !strings.HasPrefix(words[1], "p") {
		return false
	}
	if !strings.HasPrefix(words[2], "m") {
		return false
	}
	return true
}

func npms(in <-chan string) <-chan string {
	out := make(chan string)
	history := []string{}
	go func() {
		defer close(out)
		for word := range in {
			history = append(history, word)
			if len(history) > 3 {
				history = history[1:]
			}
			if isNPM(history) {
				out <- strings.Join(history, " ")
			}
		}
	}()
	return out
}

func main() {
	for npms := range npms(words(os.Stdin)) {
		fmt.Println(npms)
	}
}
