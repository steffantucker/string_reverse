/*
 * Author: Steffan Tucker
 *
 * string_reverse will reverse N number of strings
 * N is the first expected input
 * following N inputs will be strings to be reversed
 * strings are reversed and output as "Case #x: [reversed string]"
 * input is from command line or a file can be passed in by flag
 * output is to command line or a file passed in by flag
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "input", "", "Input file")
	flag.StringVar(&inputFile, "i", "", "Input file")

	var outputFile string
	flag.StringVar(&outputFile, "output", "", "Output file")
	flag.StringVar(&outputFile, "o", "", "Output file")

	flag.Parse()

	var reader *bufio.Reader
	// read from command line
	if inputFile == "" {
		reader = bufio.NewReader(os.Stdin)
	} else { // read from file
		f, err := os.Open(inputFile)
		if err != nil {
			fmt.Println("Error opening file:", err)
		}
		defer f.Close()
		reader = bufio.NewReader(f)
	}
	// first line input should be count of lines to read
	count := cleanInput(reader)

	// first input should be a number, throw error
	// if it can't be converted to an int
	numcount, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println("Must enter a number", err)
		return
	}

	// make a slice, we know how many strings
	// should be entered be entered
	wordsReverse := make([]string, numcount)

	for i := 0; i < numcount; i++ {
		words := cleanInput(reader)

		// reverse string, save to array
		wordsReverse[i] = reverseString(words)
	}

	// write to console
	if outputFile == "" {
		for i, value := range wordsReverse {
			fmt.Printf("Case #%d: %s\n", i+1, value)
		}
	} else { // write to file
		out, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, 0755)
		defer out.Close()
		if err != nil {
			fmt.Println("Error opening output file:", err)
			for i, value := range wordsReverse {
				fmt.Printf("Case #%d: %s\n", i+1, value)
			}
		}

		for i, value := range wordsReverse {
			_, err := io.WriteString(out, fmt.Sprintf("Case #%d: %s\n", i+1, value))
			if err != nil {
				fmt.Println("Error writting to file:", err)
			}
		}
		out.Sync()
	}
	return
}

func reverseString(input string) string {
	// make an array from the words in the string
	words := strings.Split(input, " ")
	// don't need to do anything if there's
	// no words or only 1 word
	if len(words) <= 1 {
		return input
	}

	// loop to reverse the words
	// switch the first word with the last word
	// then the second word with the second to last word
	// and so on
	for i, j := 0, len(words)-1; i < j && j >= len(words)/2; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// ReadString doesn't clean the input of newline char
func cleanInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error reading line", err)
	}
	// windows does carriage return and new line,
	// other OSes don't (I believe)
	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)

	return input
}
