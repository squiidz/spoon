package api

import (
	"fmt"
	"net/http"
	"strings"
)

const omdb = "http://www.omdbapi.com/?"

func GetMovieData(title string) []byte {
	title = strings.Replace(title, " ", "-", 10)
	fmt.Println("title transform:", title)
	fullURL := fmt.Sprintf("%st=%s&y=&plot=short&r=json", omdb, title)
	fmt.Println("Full URL:", fullURL)
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println(err)
	}

	var data []byte
	resp.Body.Read(data)
	resp.Body.Close()
	fmt.Println(string(data))
	return data
}
