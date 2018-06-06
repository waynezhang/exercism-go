package clock

import "fmt"

type Clock struct {
	minutes int
}

func cal(minutes int) int {
	var step = 60 * 24
	minutes %= step
	for minutes < 0 {
		minutes += step
	}
	return minutes
}

func New(hour, min int) Clock {
	var clock Clock
	clock.minutes = cal(hour*60 + min)
	return clock
}

func (clock Clock) Add(min int) Clock {
	clock.minutes = cal(clock.minutes + min)
	return clock
}
func (clock Clock) Subtract(min int) Clock {
	clock.minutes = cal(clock.minutes - min)
	return clock
}

func (clock Clock) String() string {
	var hour = clock.minutes / 60
	var min = clock.minutes % 60
	return fmt.Sprintf("%02d:%02d", hour, min)
}
