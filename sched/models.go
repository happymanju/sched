package sched

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

type Event struct {
	Name      string
	Duration  time.Duration
	StartTime time.Time
	EndTime   time.Time
}

func (e Event) ToString() string {
	return fmt.Sprintf("%s\t%s\t%s\t%d", e.Name, e.StartTime.Format("15:04"), e.EndTime.Format("15:04"), e.Duration)
}

type Schedule struct {
	Events                       []Event
	StartDatetimeFromCommandArgs time.Time
}

func (s *Schedule) AddEvent() error {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Event name: ")
	sc.Scan()
	name := sc.Text()
	fmt.Print("Duration: ")
	sc.Scan()
	durString := sc.Text() + "m"

	dur, err := time.ParseDuration(durString)
	if err != nil || dur < 0 {
		return errors.New("invalid duration; could not parse")
	}

	newEvent := Event{
		Name:     name,
		Duration: dur,
	}
	s.Events = append(s.Events, newEvent)
	return nil
}

func (s *Schedule) DeleteEvent(idx int) {
	if idx < 0 || idx > len(s.Events)-1 {
		return
	}
	old := s.Events
	s.Events = []Event{}
	for k, v := range old {
		if k != idx {
			s.Events = append(s.Events, v)
		}
	}
}

func (s *Schedule) InsertBefore(idx int, e Event) error {
	if idx-1 < 0 || idx-1 > len(s.Events)-1 {
		return errors.New("index not valid to insert event")
	}
	newEvents := []Event{e}
	if idx-1 == 0 {
		s.Events = append(newEvents, s.Events...)
	} else {
		s.Events = slices.Insert(s.Events, idx-1, e)
	}
	return nil
}

func (s *Schedule) ToString() string {
	printedSched := ""
	for k, v := range s.Events {
		printedSched += fmt.Sprintf("%d\t%s\n", k, v.ToString())
	}
	return printedSched
}

func (s *Schedule) Calc() {
	if len(s.Events) == 0 {
		return
	}
	s.Events[0].StartTime = s.StartDatetimeFromCommandArgs
	for k, v := range s.Events {
		v.EndTime = v.StartTime.Add(v.Duration)
		s.Events[k] = v
	}
}

func (s *Schedule) SaveToString(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	outputString := s.ToString()
	bw := bufio.NewWriter(f)

	bw.Write([]byte(outputString))
	err = bw.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Save() error {
	f, err := os.Create("sched.bin")
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)

	err = enc.Encode(s)
	if err != nil {
		return err
	}
	return nil

}

func (s *Schedule) Load() error {
	f, err := os.Open("sched.bin")
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)

	dec := gob.NewDecoder(br)

	err = dec.Decode(s)
	if err != nil {
		s = &Schedule{}
		return err
	}
	return nil
}
