package main

func puzzle1(keys []int) int {
	cardLoopSize := getLoopSize(keys[0])
	encryptionKey := keys[1]
	for i := 1; i <= cardLoopSize; i++ {
		encryptionKey = transform(encryptionKey, keys[1])
	}
	return encryptionKey
}

func getLoopSize(key int) (result int) {
	subjectNumber := 7
	loopValue := subjectNumber
	for loopValue != key {
		result++
		loopValue = transform(loopValue, subjectNumber)
	}
	return
}

func transform(input int, subjectNumber int) int {
	input *= subjectNumber
	input %= 20201227
	return input
}
