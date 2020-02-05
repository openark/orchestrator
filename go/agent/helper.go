package agent

func enginesSupported(first, second []Engine) bool {
	hashmap := make(map[Engine]int)
	for _, value := range first {
		hashmap[value]++
	}

	for _, value := range second {
		if count, found := hashmap[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			hashmap[value] = count - 1
		}
	}

	return true
}
