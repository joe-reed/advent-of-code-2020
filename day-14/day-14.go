package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	. "utils"
)

func puzzle1(input []string) (r int) {
	mask := ""
	mem := make(map[int]int)
	for _, i := range input {
		if newMask, err := getMask(i); err == nil {
			mask = newMask
			continue
		}

		address, val := getAssignment(i)
		val = applyMaskV1(val, mask)
		mem[address] = val
	}

	for _, v := range mem {
		r += v
	}
	return
}

func puzzle2(input []string) (r int) {
	mask := ""
	mem := make(map[int]int)
	for _, i := range input {
		if newMask, err := getMask(i); err == nil {
			mask = newMask
			continue
		}

		address, val := getAssignment(i)
		addresses := applyMaskV2(address, mask)
		for _, a := range addresses {
			mem[a] = val
		}
	}

	for _, v := range mem {
		r += v
	}
	return
}

func getMask(i string) (string, error) {
	if strings.Contains(i, "mask") {
		return strings.Split(i, " = ")[1], nil
	}
	return "", errors.New("No mask found")
}

func getAssignment(i string) (address, value int) {
	r := regexp.MustCompile(`\[(\d+)\]`)
	m := r.FindStringSubmatch(i)
	address = ConvertToInt(m[1])
	value = ConvertToInt(strings.Split(i, " = ")[1])
	return
}

func applyMaskV1(val int, mask string) int {
	binaryVal := convertIntToBinary(val, len(mask))

	for i, v := range mask {
		if string(v) == "X" {
			continue
		}
		binaryVal = replaceChar(binaryVal, i, string(v))
	}
	return convertBinaryToInt(binaryVal)
}

func applyMaskV2(address int, mask string) (r []int) {
	binaryAddress := convertIntToBinary(address, len(mask))

	for i, v := range mask {
		switch string(v) {
		case "0":
			continue
		case "1":
			binaryAddress = replaceChar(binaryAddress, i, string(v))
		case "X":
			binaryAddress = replaceChar(binaryAddress, i, "X")
		}
	}

	xCount := strings.Count(binaryAddress, "X")
	for i := 0; i < int(math.Pow(2, float64(xCount))); i++ {
		b := convertIntToBinary(i, xCount)
		k := 0
		address := binaryAddress
		for j, c := range binaryAddress {
			if string(c) != "X" {
				continue
			}
			address = replaceChar(address, j, string(b[k]))
			k++
		}
		r = append(r, convertBinaryToInt(address))
	}
	return
}

func convertIntToBinary(val, length int) string {
	return fmt.Sprintf("%0"+fmt.Sprint(length)+"s", strconv.FormatInt(int64(val), 2))
}

func convertBinaryToInt(binary string) int {
	r, _ := strconv.ParseInt(binary, 2, 64)
	return int(r)
}

func replaceChar(s string, i int, c string) string {
	return s[:i] + c + s[i+1:]
}
