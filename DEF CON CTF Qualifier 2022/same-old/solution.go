package main

import (
	"fmt"
	"hash/crc32"
	"sync"
)

var availableChars = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var teamName = []byte("HiCr33k")

var Threads = 10
var MaxLength = 10
var threadWaitGroup sync.WaitGroup

var crcTable = crc32.MakeTable(crc32.IEEE)
var theCrc = crc32.Checksum([]byte("the"), crcTable)

func CheckVariationsOfSpecificLength(byteList []byte, length int, threadNumber int) {
	if length == 0 || len(byteList) == 0 || length > MaxLength {
		threadWaitGroup.Done()
		fmt.Printf("Thread %d: Finished everyting\n", threadNumber)
		return
	}

	fmt.Printf("Thread %d: Checking for %d\n", threadNumber, length)

	var chars = *new([]byte)
	var ids = *new([]int)
	numberOfItems := len(byteList)

	for i := 0; i < length; i++ {
		ids = append(ids, 0)
		chars = append(chars, byteList[0])
	}

	if checkByteList(chars) {
		fmt.Printf("Thread %d: Found for %s\n", threadNumber, JoinBytesToString(append(teamName, chars...)))
	}

	for {
		ids[length-1]++
		lastIncremented := length - 1

		for ; lastIncremented >= 0; lastIncremented-- {
			if ids[lastIncremented] == numberOfItems {
				ids[lastIncremented] = 0
				ids[lastIncremented-1]++
			} else {
				break
			}
		}

		for i := lastIncremented; i < length; i++ {
			chars[i] = byteList[ids[i]]
		}

		if checkByteList(chars) {
			fmt.Printf("Thread %d: Found for %s\n", threadNumber, JoinBytesToString(append(teamName, chars...)))
		}

		if ids[0] == numberOfItems-1 {
			break
		}
	}

	fmt.Printf("Thread %d: Finished checking for %d\n", threadNumber, length)
	CheckVariationsOfSpecificLength(byteList, length+Threads, threadNumber)
}

func checkByteList(lst []byte) bool {
	return crc32.Checksum(append(teamName, lst...), crcTable) == theCrc
}

func JoinBytesToString(lst []byte) string {
	result := ""

	for _, v := range lst {
		result += string(v)
	}

	return result
}

func main() {
	fmt.Println("Main: starting!")
	for i := 0; i < Threads; i++ {
		threadWaitGroup.Add(1)
		go CheckVariationsOfSpecificLength(availableChars, i+1, i+1)
	}

	fmt.Println("Main: Waiting for threads to finish!")
	threadWaitGroup.Wait()
	fmt.Println("Main: Completed!")
}
