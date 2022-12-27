package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	fms := []FreqMap{}
	for _, s := range l {
		wg.Add(1)
		m := FreqMap{}
		fms = append(fms, m)

		go frequency(&wg, s, &m)
	}

	wg.Wait()

	r := FreqMap{}
	for _, m := range fms {
		mergeMap(&r, &m)
	}

	return r
}

func frequency(wg *sync.WaitGroup, s string, m *FreqMap) {
	defer wg.Done()

	for _, r := range s {
		(*m)[r]++
	}
}

func mergeMap(to, from *FreqMap) {
	if from == nil {
		return
	}

	for k, v := range *from {
		val, ok := (*to)[k]
		if ok {
			(*to)[k] = val + v
		} else {
			(*to)[k] = v
		}
	}
}
