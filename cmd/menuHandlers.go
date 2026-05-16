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

func insertEvent(s *sched.Schedule, sc *bufio.Scanner) {
	fmt.Print("insert event name: ")
	sc.Scan()
	name := sc.Text()
	if name == "" {
		return
	}
	fmt.Print("insert event duration: ")
	sc.Scan()
	dur, err := time.ParseDuration(sc.Text() + "m")
	if err != nil {
		log.Println(err)
		return
	}
	newEvent := sched.Event{
		Name:     name,
		Duration: dur,
	}
	var idx int

	for {
		fmt.Print("index to insert at: ")
		sc.Scan()
		idx, err = strconv.Atoi(sc.Text())
		if err != nil {
			log.Println(err)
			continue
		}
		break
	}

	err = s.Insert(idx, newEvent)
	if err != nil {
		log.Println(err)
	}

}

func handleTimeChange(s *sched.Schedule, sc *bufio.Scanner) {
	fmt.Print("hour >> ")
	sc.Scan()
	newHour, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Println("could not parse; using default of '9'")
		newHour = 9
	}
	fmt.Print("minutes >> ")
	sc.Scan()
	newMin, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Println("could not parse; using default of '0'")
		newMin = 0
	}
	n := time.Now()

	s.StartDatetimeFromCommandArgs = time.Date(n.Year(), n.Month(), n.Day(), newHour, newMin, 0, 0, time.Local)
	s.Calc()
}

func handleDelete(s *sched.Schedule, sc *bufio.Scanner) {
	fmt.Print("delete index: ")
	sc.Scan()
	deleteInput, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println("Not a valid event index")
	}
	s.DeleteEvent(deleteInput)
	s.Calc()
}

func handleMarkdown(s *sched.Schedule) {
	fmt.Println(s.ToMarkdown())
}
