package hackerrank

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func WordCount(s string) int {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	words := strings.Split(s, " ")
	var count int = 0
	for _, w := range words {
		w = strings.Replace(w, "-", "", -1)
		w = strings.Replace(w, ",", "", -1)
		w = strings.Replace(w, "?", "", -1)
		w = strings.Replace(w, "!", "", -1)
		w = strings.Replace(w, "@", "", -1)
		w = strings.Replace(w, "$", "", -1)
		w = strings.Replace(w, "%", "", -1)
		w = strings.Replace(w, "^", "", -1)
		w = strings.Replace(w, "&", "", -1)
		w = strings.Replace(w, "(", "", -1)
		w = strings.Replace(w, ")", "", -1)
		w = strings.Replace(w, ".", "", -1)
		w = strings.Replace(w, ";", "", -1)
		w = strings.Replace(w, ":", "", -1)
		w = strings.Replace(w, "\"", "", -1)
		w = strings.Replace(w, "<", "", -1)
		w = strings.Replace(w, ">", "", -1)
		w = strings.Replace(w, "'", "", -1)

		// fmt.Println(w)
		if isAlpha(w) {
			count++
		}
	}
	return count

}

func IsPalindrom(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func EfficientJanitor(weights []float32, limit float32) int {
	// fmt.Println(weights)
	sort.Slice(weights, func(a, b int) bool {
		return weights[a] < weights[b]
	})
	cnt, left, right := 0, 0, len(weights)-1
	for left <= right {
		cnt++
		if weights[left]+weights[right] <= limit {
			left++
			right--
		} else {
			right--
		}
	}
	return cnt
}

func gcd(den int, num int) int {
	if num == 0 {
		return den
	}
	// fmt.Printf("den: %d | num: %d | dennum: %v\n", den, num, den%num)
	return gcd(num, den%num)
}

func ProperFractions(den int) int {
	cnt := 0
	for i := 1; i < den; i++ {
		if gcd(den, i) == 1 {
			fmt.Printf("callee denom: %d | num: %d\n", den, i)
			cnt++
		}
	}

	return cnt
}

func RemoveKDigits(s string, k int) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > ch {
			stack = stack[:len(stack)-1]
			k--
		}
		if len(stack) == 0 && ch == '0' {
			continue
		}
		stack = append(stack, ch)
	}
	if k >= len(stack) {
		return "0"
	}
	return string(stack[:len(stack)-k])
}

func HCRunFuncs() {

	fmt.Printf("FractionsCount for 15 %d\n", ProperFractions(15))

	tcEfficientJan := []float32{1.01, 1.99, 2.5, 1.5, 1.01}
	fmt.Printf("Efficient Jan %d\n", EfficientJanitor(tcEfficientJan, 3))
	var count int = 0
	for i := 1; i <= 585; i++ {
		bin := strconv.FormatInt(int64(i), 2)
		if IsPalindrom(bin) && IsPalindrom(strconv.Itoa(i)) {
			count++
		}
	}
	fmt.Printf("double base palindrom less than equal to 585: %d\n", count)
	tcIsPalindrom := []string{"a", "aba", "aa", "585", "aabbaaa"}
	for _, s := range tcIsPalindrom {
		fmt.Printf("IsPalindrom(%s), %v\n", s, IsPalindrom(s))
	}

	tcWordCount := []string{"How many eggs are in a half-dozen, 13?", "he is a good programmer, he won 865 competitions, but sometimes he dont. What do you think? All test-cases should pass. Done-done?"}
	for _, s := range tcWordCount {
		fmt.Println(WordCount(s))
	}
}
