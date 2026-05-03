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

func addEvents(s *sched.Schedule) {
	sc := bufio.NewScanner(os.Stdin)

	for {
		err := s.AddEvent()
		if err != nil {
			log.Println(err)
		}
		fmt.Print("add another? (n) >> ")
		sc.Scan()
		if sc.Text() == "n" {
			break
		}
	}
}

func editEvent(idx int, s *sched.Schedule) {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("new name? (Enter to keep): ")
	sc.Scan()
	name := sc.Text()
	fmt.Print("new duration? (Enter to keep): ")
	sc.Scan()
	dur := sc.Text()
	if name != "" {
		s.Events[idx].Name = name
	}

	if dur != "" {
		parsedDuration, err := time.ParseDuration(dur + "m")
		if err != nil {
			log.Println(err)
			return
		}
		s.Events[idx].Duration = parsedDuration
	}

}

func insertEvent() {
	return
}

func Run(args []string) int {
	isRunning := true
	dateFromArgs, err := sched.ParseISO(args[0], args[1], args[3], args[4], args[5])
	if err != nil {
		log.Println(err)
	}

	s := sched.Schedule{
		StartDatetimeFromCommandArgs: dateFromArgs,
	}

	sc := bufio.NewScanner(os.Stdin)

	for isRunning {
		fmt.Println(s.ToString())
		fmt.Println("(a) add events | (t) change start time | (d) delete event | (s) save schedule to text | (b) save to binary| (l) load | (q) quit")
		sc.Scan()
		input := sc.Text()

		switch input {
		case "a":
			addEvents(&s)
			s.Calc()
			continue
		case "t":
			fmt.Print("hour >> ")
			sc.Scan()
			newHour := sc.Text()
			fmt.Print("minutes >> ")
			sc.Scan()
			newMin := sc.Text()
			newDate, err := sched.ParseISO(args[0], args[1], args[3], newHour, newMin)
			if err != nil {
				log.Println(err)
				continue
			}
			s.StartDatetimeFromCommandArgs = newDate
			s.Calc()

		case "d":
			fmt.Print("delete index: ")
			sc.Scan()
			deleteInput, err := strconv.Atoi(sc.Text())
			if err != nil {
				fmt.Println("Not a valid event index")
			}
			s.DeleteEvent(deleteInput)
			s.Calc()
			continue
		case "s":
			s.SaveToString(sched.MakeTimestamp() + ".txt")
			continue
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
