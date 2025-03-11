package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func writeToFile(content string, filePath string) {
	keyword := strings.Trim(fmt.Sprintf(content), "\x00")
	err := os.WriteFile(filePath, bytes.Trim([]byte(keyword), "\x00"), 0200)
	if err != nil {
		log.Fatal(err)
	}
}

func writeCountToFile(count int, filePath string) {
	//fmt.Println("W.t.F")

	// Write count to the file using the passed file handle (f)
	keyword := strings.Trim(fmt.Sprintf("count=%d", count), "\x00")

	err := os.WriteFile(filePath, bytes.Trim([]byte(keyword), "\x00"), 0200)
	if err != nil {
		log.Fatal(err)
	}
	//DEBUGGING THE count=x function
	// fmt.Println("Wrote to the file:", keyword)
	// for index, runeValue := range keyword {
	// 	fmt.Printf("Index: %d, Rune: %c\n", index, runeValue)
	// }

}

func main() {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Errorf("reading data, error: %w", err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Read the file content
	content := make([]byte, fileInfo.Size())
	_, err = file.Read(content)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the content to string
	strContent := string(content)

	// Split the string by '='
	parts := strings.Split(strContent, "=")
	if len(parts) != 2 {
		log.Fatal("Invalid content format")
	}

	// Parse the numerical value (count)
	//fmt.Println("parts")
	//fmt.Println(parts)
	countStr := parts[1]
	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatal("Error converting string to int:", err)
	}

	// read the "signal.txt" file and print its text

	signal, err := os.OpenFile("signal.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Errorf("reading signal.txt, error: %w", err)
	}
	defer signal.Close()

	// // Get file size
	sig_fileInfo, err := signal.Stat()
	strSignal := ""

	if err != nil {
		fmt.Errorf("FILE SIZE, error: %w", err)
		log.Fatal("File Size Error?")
	}
	fmt.Println("length of signal.txt: ", sig_fileInfo.Size())
	if sig_fileInfo.Size() > 0 { //signal.txt is NOT empty

		sig_content := make([]byte, sig_fileInfo.Size())
		for {
			n, err := signal.Read(sig_content)
			if err != nil {
				if err == io.EOF {
					// End of file
					//fmt.Println("EOF -- going to break")
					strSignal = string(sig_content)
					//fmt.Println("here is strSignal: ", strSignal)
					break
				}
				fmt.Println(err)
				return

			} else { // "no error on this byte"
				sig_content = sig_content[:n]
			}
		}
	} else { //the signal.txt was empty. tell the user to put something
		strSignal = "hey! welcome to my app. go the signal.txt file and add something!"
	}

	fmt.Println(strSignal) // Prints with a newline

	//increment count
	count++
	writeCountToFile(count, "data.txt")
	fmt.Printf("Invocation Number: %d\n", count) // Formatted printing
	//close file

}

func TestStrain(t *testing.T) {
	n := 50
	threshold := 20 //keep threshold LESS than n.

	signal := "Welcome to My App"
	altSignal := "We are testing in CI/CD"
	for x := range n {
		main()
		if x == 2 {
			writeToFile(signal, "signal.txt")
		}
		if x == threshold {
			writeToFile(altSignal, "signal.txt")
		}
	}
}
