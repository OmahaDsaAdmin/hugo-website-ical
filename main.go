package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var ApiKey string
var ExportPath string

func main() {
	// Setup CLI args
	ExportPath = os.Args[1]

	// Multiple API Keys can be supplied
	for i := 2; i < len(os.Args); i++ {
		ApiKey = os.Args[i]

		if ApiKey == "" || ExportPath == "" {
			fmt.Println("missing required args")
			return
		}

		// Setup http request
		url := "https://actionnetwork.org/api/v2/events/"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("OSDI-API-Token", ApiKey)

		// Send the request
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Unmarshal the response body into an event response
		var CalendarEventResponse CalendarEventResponse
		err = json.Unmarshal([]byte(body), &CalendarEventResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		// For each event, write the data to a markdown template
		events := CalendarEventResponse.Embedded.Events
		createAndFillTemplateWithList(events)
	}

	fmt.Println("passed")
}

func createAndFillTemplateWithList(ces []CalendarEvent) {
	// Fix the data in the array
	newCes := []CalendarEvent{}

	for _, e := range ces {
		id := strings.Split(e.Identifiers[0], "_network:")[1]
		e.Id = id
		e.Title = removeDoubleQuotes(e.Title)
		e.Description = stripHTMLTags(e.Description)

		newCes = append(newCes, e)
	}

	// Create template
	t, err := template.ParseFiles("template.ics")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the file, overwriting any old one
	newFilePath := ExportPath + "/" + "calendar_events.ical"
	fmt.Println(newFilePath)
	f, err := os.Create(newFilePath)
	if err != nil {
		log.Println("create file: "+newFilePath, err)
		return
	}

	err = t.Execute(f, newCes)
	if err != nil {
		log.Print("execute: "+newFilePath, err)
		return
	}

	f.Close()
}

func stripHTMLTags(input string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(input, "")
}

func removeDoubleQuotes(input string) string {
	return strings.ReplaceAll(input, "\"", "")
}
