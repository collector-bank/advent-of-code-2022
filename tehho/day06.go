package main

func CheckForDuplicate(b []byte) bool {
	for i := 0; i < len(b); i++ {
		count := 0

		for j := 0; j < len(b); j++ {
			if b[i] == b[j] {
				count++
			}
		}

		if 2 <= count {
			return true
		}
	}
	return false
}

func Day06part01(input string) int {

	size := 4
	temp := make([]byte, size)
	for i := 0; i < size; i++ {
		temp[i] = input[i]
	}

	for i := size; i < len(input); i++ {
		if !CheckForDuplicate(temp) {
			return i
		}
		for j := 0; j < size-1; j++ {
			temp[j] = temp[j+1]
		}
		temp[size-1] = input[i]
	}

	return len(input)
}

func Day06part02(input string) int {

	size := 14
	temp := make([]byte, size)
	for i := 0; i < size; i++ {
		temp[i] = input[i]
	}

	for i := size; i < len(input); i++ {
		if !CheckForDuplicate(temp) {
			return i
		}
		for j := 0; j < size-1; j++ {
			temp[j] = temp[j+1]
		}
		temp[size-1] = input[i]
	}

	return len(input)
}
