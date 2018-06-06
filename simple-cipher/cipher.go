package cipher

import (
	"regexp"
	"strings"
)

// Cipher doc here
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// Impl doc here
type Impl struct {
	distanceGenerator func(chan int, int)
}

func round(r rune, offset int) rune {
	r += rune(offset)
	if r < 'a' {
		r += 26
	} else if r > 'z' {
		r -= 26
	}
	return r
}

// Encode doc here
func (cipher *Impl) Encode(str string) string {
	str = regexp.MustCompile("[^a-zA-Z]*").ReplaceAllString(str, "")
	str = strings.ToLower(str)
	distance := make(chan int)
	go cipher.distanceGenerator(distance, len(str))
	return strings.Map(func(r rune) rune { return round(r, <-distance) }, str)
}

// Decode doc here
func (cipher *Impl) Decode(str string) string {
	str = regexp.MustCompile("[^a-zA-Z]*").ReplaceAllString(str, "")
	distance := make(chan int)
	go cipher.distanceGenerator(distance, len(str))
	return strings.Map(func(r rune) rune { return round(r, 0-<-distance) }, str)
}

// NewCaesar doc here
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift doc here
func NewShift(distance int) Cipher {
	if distance <= -26 || distance >= 26 || distance == 0 {
		return nil
	}
	return &Impl{distanceGenerator: func(ch chan int, size int) {
		for i := 0; i < size; i++ {
			ch <- distance
		}
		close(ch)
	}}
}

// NewVigenere doc here
func NewVigenere(key string) Cipher {
	if len(regexp.MustCompile("[^a]+").ReplaceAllString(key, "")) == len(key) || regexp.MustCompile("[^a-z]+").MatchString(key) {
		return nil
	}
	return &Impl{distanceGenerator: func(ch chan int, size int) {
		for i := 0; i < size; i++ {
			ch <- int(key[i%len(key)] - 'a')
		}
		close(ch)
	}}
}
