package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readFile(fileName string) {
	fmt.Println("Reading", fileName)
	rand.Seed(time.Now().UTC().UnixNano())
	sleepInt := randInt(1, 5)
	h := time.Duration(sleepInt) * time.Second
	time.Sleep(h)
	fmt.Println("Done reading", fileName)

}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func getFiles(limit int, doneFiles chan string) {
	for i := 0; i <= limit; i++ {
		fileName := fmt.Sprintf("file%d.txt", i)
		readFile(fileName)
		doneFiles <- fileName
	}
	close(doneFiles)
}

func main() {
	fmt.Println("Application Starting")

	doneFiles := make(chan string)
	go getFiles(5, doneFiles)

	for g := range doneFiles {
		// NOTE: This code and the readFile code occur concurrently
		fmt.Println("finished with", g)
	}

}
