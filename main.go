package main

import (
	"bufio"
	"fmt"
	"github.com/mapreduce/src"
	"log"
	"os"
)

func main() {
	//input data read from "data.csv"
	fd, err := os.Open("data.csv")
	defer fd.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := src.NewScanner(fd)
	mapper := src.NewMapper()
	reducer := src.NewReducer()
	//map used to store rating by user_id
	collection := make(map[string][]float64)

	//mapper
	for {
		if line := fileScanner.Next(); line != nil {
			id, rating := mapper.Map(string(line))
			collection[id] = append(collection[id], rating)
		} else {
			break
		}
	}

	//reducer && write the average data into "output.data"
	fd2, err := os.OpenFile("output.data", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer fd2.Close()
	if err != nil {
		log.Fatal(err)
	}
	if fd2 == nil {
		writer := bufio.NewWriter(fd2)
		for key, value := range collection {
			avg := reducer.Reduce(value)
			line := fmt.Sprintf("%s:%g\n", key, avg)
			writer.WriteString(line)
		}
	}

	//find the average rating based on the input of user
	fileScanner = src.NewScanner(fd2)
	var user_id string
	for {
		//cleanup process
		fd2.Seek(0, os.SEEK_SET)
		user_id = ""

		fmt.Print("Enter the user id that you want to konw (hit q to quit the program): ")
		fmt.Scanf("%s", &user_id)
		switch user_id {
		case "":
			fmt.Println("FATAL: Please at least enter something!!")
		case "q":
			fmt.Println("Quit the program.\n")
			return
		default:
			//the type of avg here is string
			if avg, ok := fileScanner.Find(user_id); ok {
				fmt.Printf("%s:%s\n", user_id, avg)
			} else {
				fmt.Println("I cannot find this number!!")
			}
		}
	}
}
