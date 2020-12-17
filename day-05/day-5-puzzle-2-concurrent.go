package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
	. "utils"
)

func main() {
	defer PrintTimeSince(time.Now())

	jobs := make(chan string)
	go func() {
		file, err := os.Open("./input.txt")
		Check(err)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	c := make(chan int)
	var wg sync.WaitGroup

	for job := range jobs {
		wg.Add(1)
		go calculateId(job, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	ids := []int{}
	for id := range c {
		ids = append(ids, id)
	}

	fmt.Println(findMissingId(ids))
}

func findMissingId(ids []int) int {
	idsCopy := make([]int, len(ids))
	copy(idsCopy, ids)
	sort.Ints(idsCopy)

	previousId := idsCopy[0] - 1
	for _, id := range idsCopy {
		if id != previousId+1 {
			return previousId + 1
		}
		previousId = id
	}
	return 0
}

func calculateId(boardingPass string, c chan<- int, wg *sync.WaitGroup) {
	row := calculateRow(boardingPass[:7])
	column := calculateColumn(boardingPass[7:])
	c <- row*8 + column
	wg.Done()
}

func calculateRow(rowSpecifier string) int {
	return evaluateLocation(rowSpecifier, 128, "F")
}

func calculateColumn(columnSpecifier string) int {
	return evaluateLocation(columnSpecifier, 8, "L")
}

func evaluateLocation(specifier string, spaceSize int, startHalfCharacter string) int {
	start, end := 0, spaceSize
	for _, char := range specifier {
		average := (end + start) / 2
		if string(char) == startHalfCharacter {
			end = average
		} else {
			start = average
		}
	}
	return end - 1
}
