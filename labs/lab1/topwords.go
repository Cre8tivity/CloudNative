// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace

// Note: You should use `checkError` to handle potential errors.

package textproc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func topWords(path string, K int) []WordCount {
	// Your code here.....

	// Step 1: scan strings from passage file individually, and place into a string vector

	f, err := os.Open(path)

	checkError(err)

	defer f.Close() // close file

	words := []string{} // initialize empty string vector

	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		// append each word in the file to map
		words = append(words, scan.Text())
	}

	// Step 2: populate a map of string values mapped to integers to count the occurences of each word

	m := make(map[string]int)

	for _, word := range words {
		_, match := m[word]
		if match {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	// Step 3: turn map into WordCount array to use sort function

	countarr := []WordCount{}
	for k, v := range m {
		countarr = append(countarr, WordCount{Word: k, Count: v})
	}

	sortWordCounts(countarr) // this is a given function to make sure highest values are at the front of the vector

	// Step 4: return the highest K occurences of words
	result_array := []WordCount{}
	for i := 0; i < K; i++ {
		result_array = append(result_array, countarr[i])
	}

	return result_array

}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
