package tickets

import (
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	//Arrange
	ticketsList := []Ticket{
		Ticket{"1", "Tait Mc Caughan", "tmc0@scribd.com", "Finland", "17:11", 785},
		Ticket{"2", "Padget McKee", "pmckee1@hexun.com", "China", "20:19", 537},
		Ticket{"3", "Yalonda Jermyn", "yjermyn2@omniture.com", "China", "18:11", 579},
		Ticket{"4", "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238},
	}
	a := Airline{
		Tickets: ticketsList,
	}
	expectedTotal := 2
	//Act
	actualTotal, err := a.GetTotalTickets("China")
	//Assert
	if err != nil || actualTotal != expectedTotal {
		t.Errorf("Err total of tickets from China expected %d, but %d was received", expectedTotal, actualTotal)
	}
}
