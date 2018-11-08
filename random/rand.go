package random

import (
	"math/rand"
	"time"
)

var (
	defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

const (
	Digits       = "0123456789"
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	SpecialChars = "!@#$%&*+-=?"

	Space = " "
)

func Chars(length int) string {
	s := Digits + LowerLetters + UpperLetters
	b := make([]byte, length)
	for i := range b {
		b[i] = s[rand.Int63()%int64(len(s))]
	}
	return string(b)
}

// Password will generate a random password
// Minimum number length of 5 if less than
func CustomChars(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	// Make sure the num minimun is at least 5
	if num < 5 {
		num = 5
	}
	i := 0
	b := make([]byte, num)
	var passString string

	if lower {
		passString += LowerLetters
		b[i] = LowerLetters[rand.Int63()%int64(len(LowerLetters))]
		i++
	}
	if upper {
		passString += UpperLetters
		b[i] = UpperLetters[rand.Int63()%int64(len(UpperLetters))]
		i++
	}
	if numeric {
		passString += Digits
		b[i] = Digits[rand.Int63()%int64(len(Digits))]
		i++
	}
	if special {
		passString += SpecialChars
		b[i] = SpecialChars[rand.Int63()%int64(len(SpecialChars))]
		i++
	}
	if space {
		passString += Space
		b[i] = Space[rand.Int63()%int64(len(Space))]
		i++
	}

	// Set default if empty
	if passString == "" {
		passString = LowerLetters + Digits
	}

	// Loop through and add it up
	for i <= num-1 {
		b[i] = passString[rand.Int63()%int64(len(passString))]
		i++
	}

	// Shuffle bytes
	for i := range b {
		j := defaultRand.Intn(i + 1)
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func Int(max int) int {
	return defaultRand.Intn(max)
}

// 当max>=n:从0到max中随机出n个不重复的数；当max<n,随机出可能n个可能重复的数
func NInt(max, n int, repeat ...bool) []int {
	if max < 0 || n < 0 {
		return []int{}
	}

	arr := []int{}

	if max < n {
		// for i := 0; i < max; i++ {
		// 	arr = append(arr, i)
		// }

		for i := 0; i < n; i++ {
			arr = append(arr, Int(max))
		}
	} else {
		for i := 0; i < n; i++ {
		RAND:
			r := Int(max)
			for _, v := range arr {
				if v == r {
					goto RAND
				}
			}
			arr = append(arr, r)
		}

	}
	return arr
}

// [from, to)
func Range(from, to int) int {
	if from == to {
		return from
	}
	if from > to {
		panic("'to' mush be bigger than 'from'")
	}
	return Int(to-from) + from
}
