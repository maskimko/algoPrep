package main

import "fmt"

//Does not work

// chunksForSpan accepts start packet number and end packet number (excluded)
func chunksForSpan(start, end int64) int32 {
	var left int64 = end - start + 1
	var chunks int32 = 0
	for left > 0 {
		var chunkSize int64 = 1
		for ; chunkSize <= left; chunkSize *= 2 {
		}
		left -= chunkSize / 2
		chunks++
	}
	return chunks
}

func computeSpans(totalPackets int64, uploadedChunks [][]int64) [][]int64 {
	var spans [][]int64
	var spanStart int64 = 1
	for n := range uploadedChunks {
		start := uploadedChunks[n][0]
		end := uploadedChunks[n][1] + 1
		if spanStart < start {
			spans = append(spans, []int64{spanStart, start})
		}
		spanStart = end
		if spanStart >= totalPackets {
			break
		}
	}
	if spanStart < totalPackets {
		spans = append(spans, []int64{spanStart, totalPackets})
	}
	return spans
}

func minimumChunksRequired(totalPackets int64, uploadedChunks [][]int64) int32 {
	// Write your code here
	var chunks int32 = 0
	spans := computeSpans(totalPackets, uploadedChunks)
	for n := range spans {
		chunks += chunksForSpan(spans[n][0], spans[n][1])
	}
	return chunks
}

func main() {

	word := "aaa"
	chars := []rune(word)
	fmt.Printf("%c\n", chars[1]+1)

	fmt.Printf("chunks %d\n", chunksForSpan(2, 15))
}
