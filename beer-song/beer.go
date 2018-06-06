package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Song doc here
func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}

// Verse doc here
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", errors.New("")
	}
	switch n {
	case 0:
		return `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
`, nil
	case 1:
		return `1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.
`, nil
	case 2:
		return `2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall.
`, nil
	default:
		return fmt.Sprintf(`%d bottles of beer on the wall, %d bottles of beer.
Take one down and pass it around, %d bottles of beer on the wall.
`, n, n, n-1), nil
	}
}

// Verses doc here
func Verses(to, from int) (string, error) {
	if to < from {
		return "", errors.New("")
	}

	verses := make([]string, to-from+1)
	for i := to; i >= from; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", errors.New("")
		}
		verses[to-i] = v
	}
	return strings.Join(append(verses, ""), "\n"), nil
}
