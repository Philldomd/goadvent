package taskmaster

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GetPussleInput(year string, day string) string {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	return readBody(getAdvent(url))
}

func PostAnswer(year string, day string, level string, answer string) string {
	form := url.Values{}
	form.Add("answer", answer)
	form.Add("level", level)
	url_address := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", year, day)
  request, err := http.NewRequest("POST", url_address, strings.NewReader(form.Encode()))
	request.Header.Add("Cookie", collectSessionId("cookie/session-id.cookie"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  checkError(err)
  response, err := http.DefaultClient.Do(request)
	checkError(err)
	return readResponseBody(response)
}

func getAdvent(url string) *http.Response {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Cookie", collectSessionId("cookie/session-id.cookie"))
	response, err := http.DefaultClient.Do(request)
	checkError(err)
	return response
}

func collectSessionId(path string) string {
	cookieByteArray, err := os.ReadFile(path)
	checkError(err)
	return string(cookieByteArray)
}

func readBody(response *http.Response) string {
	data, err := io.ReadAll(response.Body)
	checkError(err)
	return string(data)
}

func readResponseBody(response *http.Response) string {
	token := html.NewTokenizer(response.Body)
	depth := 0
	for {
		switch token.Next(){
	  case html.ErrorToken: return "nil"
		case html.TextToken:
      if depth > 0 {
				text := string(token.Text())
        //fmt.Println("Text: " + text)
				if strings.Contains(text, "That's not the right answer."){ return "Wrong" }
				if strings.Contains(text, "You gave an answer too recently;"){ return "Wait" }
				if strings.Contains(text, "That's the right answer!"){ return "OK" }
				if strings.Contains(text, "You don't seem to be solving the right level."){ return "Already solved" }
			}
		case html.StartTagToken, html.EndTagToken: 
	    tn, _ := token.TagName()
			if len(tn) == 7 && tn[0] == 'a' {
        if token.Token().Type == html.StartTagToken {
					depth++
				} else { depth-- }
			}
		}
	}
}