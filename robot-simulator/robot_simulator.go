package robot

// TODO Test Case 3 not PASSed!!!

import (
	"regexp"
)

// Test Case 1

// Action doc here
type Action func(robot *Step2Robot, room Rect)

const (
	// N doc here
	N = iota
	// E doc here
	E
	// S doc here
	S
	// W doc here
	W
)

// Right doc here
func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

// Left doc here
func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

// Advance doc here
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case S:
		Step1Robot.Y--
	case E:
		Step1Robot.X++
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	bytes := make([]byte, 0)
	for {
		if d == 0 {
			break
		}
		digit := d % 10
		bytes = append([]byte{'0' + byte(digit)}, bytes...)
		d /= 10
	}
	return string(bytes)
}

// Test Case 2

// StartRobot doc here
func StartRobot(commands chan Command, actions chan Action) {
	for cmd := range commands {
		switch cmd {
		case 'L':
			actions <- func(robot *Step2Robot, _ Rect) { robot.Dir = (robot.Dir + 3) % 4 }
		case 'R':
			actions <- func(robot *Step2Robot, _ Rect) { robot.Dir = (robot.Dir + 1) % 4 }
		case 'A':
			actions <- func(robot *Step2Robot, rect Rect) {
				switch robot.Dir {
				case N:
					robot.Pos.Northing++
					if robot.Pos.Northing > rect.Max.Northing {
						robot.Pos.Northing = rect.Max.Northing
					}
				case S:
					robot.Pos.Northing--
					if robot.Pos.Northing < rect.Min.Northing {
						robot.Pos.Northing = rect.Min.Northing
					}
				case E:
					robot.Pos.Easting++
					if robot.Pos.Easting > rect.Max.Easting {
						robot.Pos.Easting = rect.Max.Easting
					}
				case W:
					robot.Pos.Easting--
					if robot.Pos.Easting < rect.Min.Easting {
						robot.Pos.Easting = rect.Min.Easting
					}
				}
			}
		}
	}
	close(actions)
}

// Room doc here
func Room(extent Rect, robot Step2Robot, actions chan Action, rep chan Step2Robot) {
	for act := range actions {
		act(&robot, extent)
	}
	rep <- robot
	close(rep)
}

// Test Case 1

// Action3 doc here
type Action3 func(robots []Step3Robot, room Rect)

// StartRobot3 doc here
func StartRobot3(name, script string, actions chan Action3, log chan string) {
	if name == "" {
		log <- "A robot without a name"
	}
	if regexp.MustCompile("[^RLA]+").MatchString(script) {
		log <- "An undefined command in a script"
	}
	for _, cmd := range script {
		switch cmd {
		case 'L':
			actions <- func(robots []Step3Robot, _ Rect) {
				for i := range robots {
					robots[i].Dir = (robots[i].Dir + 3) % 4
				}
			}
		case 'R':
			actions <- func(robots []Step3Robot, _ Rect) {
				for i := range robots {
					robots[i].Dir = (robots[i].Dir + 1) % 4
				}
			}
		case 'A':
			actions <- func(robots []Step3Robot, rect Rect) {
				for i := range robots {
					robot := &robots[i]
					switch robot.Dir {
					case N:
						robot.Pos.Northing++
						if robot.Pos.Northing > rect.Max.Northing {
							log <- "A robot attemting to advance into a wall"
							robot.Pos.Northing = rect.Max.Northing
						}
					case S:
						robot.Pos.Northing--
						if robot.Pos.Northing < rect.Min.Northing {
							log <- "A robot attemting to advance into a wall"
							robot.Pos.Northing = rect.Min.Northing
						}
					case E:
						robot.Pos.Easting++
						if robot.Pos.Easting > rect.Max.Easting {
							log <- "A robot attemting to advance into a wall"
							robot.Pos.Easting = rect.Max.Easting
						}
					case W:
						robot.Pos.Easting--
						if robot.Pos.Easting < rect.Min.Easting {
							log <- "A robot attemting to advance into a wall"
							robot.Pos.Easting = rect.Min.Easting
						}
					}
				}
			}
		}
	}
}

var roomMap [][]int

// Room3 doc here
func Room3(extent Rect, robots []Step3Robot, actions chan Action3, report chan []Step3Robot, log chan string) {
	roomMap = make([][]int, extent.Max.Easting-extent.Min.Easting+1)
	for i := range roomMap {
		roomMap[i] = make([]int, extent.Max.Northing-extent.Min.Northing+1)
	}

	roomHash := make(map[string]bool)
	for _, robot := range robots {
		if _, exist := roomHash[robot.Name]; exist {
			log <- "Duplicate robot names"
		}
		roomHash[robot.Name] = true

		if robot.Pos.Easting < extent.Min.Easting || robot.Pos.Northing < extent.Min.Northing || robot.Pos.Easting > extent.Max.Easting || robot.Pos.Northing > extent.Max.Northing {
			log <- "A robot placed outside of the room"
		} else {

			if roomMap[robot.Pos.Easting-extent.Min.Easting][robot.Pos.Northing-extent.Min.Northing] > 0 {
				log <- "Robots placed at the same place"
			}
			roomMap[robot.Pos.Easting-extent.Min.Easting][robot.Pos.Northing-extent.Min.Northing]++
		}
	}

	for act, ok := <-actions; ok; act, ok = <-actions {
		act(robots, extent)
	}
	report <- robots
	close(report)
}
