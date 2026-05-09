package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/happymanju/sched/sched"
)

func clearScreen() {
	fmt.Println("\033[2J\033[H")
}

func Run(args []string) int {
	var newDate time.Time = time.Now()
	var err error
	sc := bufio.NewScanner(os.Stdin)
	newDate = time.Date(newDate.Year(), newDate.Month(), newDate.Day(), 9, 0, 0, 0, time.Local)

	isRunning := true

	s := sched.Schedule{
		StartDatetimeFromCommandArgs: newDate,
	}

	for isRunning {
		clearScreen()
		fmt.Println(s.ToString())
		fmt.Println("(a) add events | (i) insert event | (t) change start time | (d) delete event | (s) save schedule to text | (b) save to binary| (l) load | (q) quit")
		sc.Scan()
		input := sc.Text()

		switch input {
		case "a":
			addEvents(&s)
			s.Calc()
			continue
		case "i":
			insertEvent(&s, sc)
			s.Calc()
			continue
		case "t":
			handleTimeChange(&s, sc)
			continue
		case "d":
			handleDelete(&s, sc)
			continue
		case "s":
			s.SaveToString(sched.MakeTimestamp() + ".txt")
			continue
		case "b":
			err = s.Save()
			if err != nil {
				log.Println(err)
				continue
			}
		case "l":
			err := s.Load()
			if err != nil {
				log.Println(err)
			}
			continue
		case "q":
			isRunning = false
			continue
		default:
			idx, err := strconv.Atoi(input)
			if err != nil {
				log.Println(err)
			}
			if idx < 0 || idx > len(s.Events)-1 {
				fmt.Println("Not a valid event index")
			} else {
				editEvent(idx, &s)
				s.Calc()
			}

		}
	}
	return 0
}
