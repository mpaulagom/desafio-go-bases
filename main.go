package main

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/desafio-go-bases/internal/filerepository"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/gocarina/gocsv"
)

func main() {

	repo := filerepository.FileRepository{
		FileName: "tickets.csv",
		FileType: "CSV",
	}
	airline := tickets.Airline{}
	in, err := os.Open("tickets.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &airline.Tickets); err != nil {
		panic(err)
	}

	tickets, err := repo.ParseData()
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, rec := range tickets {
		fmt.Println(rec.Email)

	}
}
