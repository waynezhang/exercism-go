package twelve

import (
	"bytes"
	"fmt"
)

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var parts = [][]byte{
	[]byte("twelve Drummers Drumming"),
	[]byte("eleven Pipers Piping"),
	[]byte("ten Lords-a-Leaping"),
	[]byte("nine Ladies Dancing"),
	[]byte("eight Maids-a-Milking"),
	[]byte("seven Swans-a-Swimming"),
	[]byte("six Geese-a-Laying"),
	[]byte("five Gold Rings"),
	[]byte("four Calling Birds"),
	[]byte("three French Hens"),
	[]byte("two Turtle Doves"),
	[]byte("and a Partridge in a Pear Tree"),
}

var formatString = "On the %s day of Christmas my true love gave to me, %s."

// Song document here
func Song() string {
	var buffer bytes.Buffer
	for i := range days {
		buffer.WriteString(Verse(i + 1))
		buffer.WriteRune('\n')
	}
	return buffer.String()
}

// Verse document here
func Verse(day int) string {
	var b []byte
	if day == 1 {
		b = bytes.Replace(parts[len(parts)-1], []byte("and "), []byte{}, -1)
	} else {
		b = bytes.Join(parts[len(parts)-day:], []byte(", "))
	}
	return fmt.Sprintf(formatString, days[day-1], string(b))
}
