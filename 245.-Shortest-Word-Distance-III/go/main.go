package main

func algoOne(wordsDict []string, word1 string, word2 string) int {

	m := make(map[string][]int)

	for i, v := range wordsDict {
		m[v] = append(m[v], i)
	}

	shortest := 10001
	for _, i := range m[word1] {
		for _, ii := range m[word2] {
			if i != ii && abs(i, ii) < shortest {
				shortest = abs(i, ii)
			}
		}
	}

	return shortest
}

func algoTwo(wordsDict []string, word1 string, word2 string) int {

	var s1, s2 []int

	for i, v := range wordsDict {
		if v == word1 {
			s1 = append(s1, i)
		} else if v == word2 {
			s2 = append(s2, i)
		}
	}

	shortest := 10001
	if len(s2) == 0 {
		s2 = s1
	}
	for _, i := range s1 {
		for _, ii := range s2 {
			if i != ii && abs(i, ii) < shortest {
				shortest = abs(i, ii)
			}
		}
	}

	return shortest
}

func algoThree(wordsDict []string, word1 string, word2 string) int {

	shortest, word1i, word2i := 10001, -1, -1

	for i, v := range wordsDict {

		if word1 == v && word1 == word2 {
			if word1i != -1 {
				dif := abs(word1i, i)
				if dif < shortest {
					shortest = dif
				}
			}
			word1i = i
		} else if word1 == v {
			word1i = i
			if word2i != -1 {
				dif := abs(word1i, word2i)
				if dif < shortest {
					shortest = dif
				}
			}
		} else if word2 == v {
			word2i = i
			if word1i != -1 {
				dif := abs(word1i, word2i)
				if dif < shortest {
					shortest = dif
				}
			}
		}
	}

	return shortest
}

func abs(v, vv int) int {
	if v-vv < 0 {
		return vv - v
	}
	return v - vv
}
