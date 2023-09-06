package tickets

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrEmptyTicketsCollection = errors.New("the tickets collection is empty")
)

const (
	MORNING      = "morning"
	AFTERNOON    = "afternoon"
	NIGHT        = "night"
	EARLYMORNING = "latenight"
)

type Ticket struct {
	Id         string  //`csv:"id"`
	Name       string  //`csv:"name"`
	Email      string  //`csv:"email"`
	Country    string  //`csv:"country"`
	FlightTime string  //`csv:"hour"`
	Price      float64 //`csv:"price"`
}

type Airline struct {
	Tickets []Ticket
}

func (a *Airline) AddTicket(t Ticket) {
	a.Tickets = append(a.Tickets, t)
}

// ejemplo 1
func (a Airline) GetTotalTickets(destination string) (int, error) {
	var total int
	if len(a.Tickets) == 0 {
		return total, ErrEmptyTicketsCollection
	}
	for _, ticket := range a.Tickets {
		if ticket.Country == destination {
			total++
		}
	}
	return total, nil
}

func hoursStringToFloat(hour string) (float64, error) {
	stringIime := strings.Split(hour, ":")
	hours, err1 := strconv.ParseFloat(stringIime[0], 64) //18:11
	if err1 != nil {
		return 0, err1
	}
	mins, err := strconv.ParseFloat(stringIime[1], 64)
	if err != nil {
		return 0, err
	}
	return hours + (mins / 100), nil // retorna 18,11

}

// ejemplo 2
func (a Airline) GetMornings(time string) (int, error) {
	var peopleCount int
	if len(a.Tickets) == 0 {
		return peopleCount, ErrEmptyTicketsCollection
	}
	for _, ticket := range a.Tickets {
		timeFloat, err := hoursStringToFloat(ticket.FlightTime)
		if err != nil {
			return peopleCount, err
		}
		switch time {
		case MORNING:
			if timeFloat >= 7 && timeFloat <= 12 {
				peopleCount++
			}
		case AFTERNOON:
			if timeFloat >= 13 && timeFloat <= 19 {
				peopleCount++
			}
		case NIGHT:
			if timeFloat >= 20 && timeFloat <= 23 {
				peopleCount++
			}
		case EARLYMORNING:
			if timeFloat >= 0 && timeFloat <= 6 {
				peopleCount++
			}
		default:
			return peopleCount, errors.New("unexpected error")
		}
	}
	return peopleCount, nil
}

// ejemplo 3
func (a Airline) AverageDestination(destination string) (int, error) {
	if len(a.Tickets) == 0 {
		return 0, ErrEmptyTicketsCollection
	}
	totalOfTickets := len(a.Tickets)
	peopleToDest, err := a.GetTotalTickets(destination)
	return peopleToDest / totalOfTickets, err
}
