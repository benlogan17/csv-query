package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"demo/concurrency/commands"
	"demo/concurrency/output"
	"demo/concurrency/reader"
)

func main() {
	start := time.Now()
	var input string

	flag.StringVar(&input, "query", "SELECT *", "Query for flag")

	flag.Parse()

	spaceTerminator := []string{" "}

	location := strings.Index(strings.ToLower(input), "from") + 4

	filename := reader.GetNextWord(input, spaceTerminator, &location)
	csv := reader.ReadCSV(filename)

	data := reader.ConverToDataMap(csv)

	firstSpace := strings.Index(input, " ")

	nextCommand := input[0:firstSpace]
	data = *commands.Select(&data, input, &location)

	for location < len(input) {
		switch strings.ToLower(nextCommand) {
		case "where":
			commands.Where(&data, input, &location)
		case "order":
			commands.Order(&data, &location, input)
		default:
			fmt.Printf("Unrecognized statemnt: %s\n", nextCommand)
		}
		nextCommand = reader.GetNextWord(input, spaceTerminator, &location)
	}
	output.DisplayData(&data)

	fmt.Printf("\ntook %v", time.Since(start))
}
