package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_minDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	got := minDistance(word1, word2)
	fmt.Println(fmt.Sprintf("minDistance result is %v", got))
	time.Sleep(100 * time.Millisecond)

	word1 = "intention"
	word2 = "execution"
	got = minDistance(word1, word2)
	fmt.Println(fmt.Sprintf("minDistance result is %v", got))
	time.Sleep(100 * time.Millisecond)

	word1 = "pneumonoultramicroscopicsilicovolcanoconiosis"
	word2 = "ultramicroscopically"
	got = minDistance(word1, word2)
	fmt.Println(fmt.Sprintf("minDistance result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
