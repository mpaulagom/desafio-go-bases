package filerepository

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

type Repository interface {
	LoadAllData() (interface{}, error)
	ParseData() (interface{}, error)
}

type FileRepository struct {
	FileName string
	FileType string
}

func (fr FileRepository) LoadAllData() (i interface{}, err error) {
	file, err := os.Open(fr.FileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	switch fr.FileType {
	case "CSV":
		fileReader := csv.NewReader(file)
		i, err = fileReader.ReadAll()
	default:
		err = errors.New("unexpected error")
	}
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (fr FileRepository) ParseData() (ticketsList []tickets.Ticket, err error) {
	data, error := fr.LoadAllData()
	if error != nil {
		err = error
		return
	}
	records := data.([][]string)
	for _, record := range records[1:] {

		price, err1 := strconv.ParseFloat(record[4], 64)
		if err1 != nil {
			err = err1
			return
		}
		tck := tickets.Ticket{
			Id:         record[0],
			Name:       record[1],
			Email:      record[2],
			FlightTime: record[3],
			Price:      price,
		}
		ticketsList = append(ticketsList, tck)
	}
	return
}
