package taskmaster

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//https://adventofcode.com/2023/day/1/answer
//body: level: <number> answer: <answer>

func GetPussleInput(year string, day string) (input string) {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	return readBody(callAdvent(url))
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func callAdvent(url string) (response *http.Response) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Cookie", collectSessionId("cookie/session-id.cookie"))
	response, err := http.DefaultClient.Do(request)
	checkError(err)
	return response
}

func collectSessionId(path string) (sessionId string) {
	cookieByteArray, err := os.ReadFile(path)
	checkError(err)
	return string(cookieByteArray)
}

func readBody(response *http.Response) (body string) {
	data, err := io.ReadAll(response.Body)
	checkError(err)
	return string(data)
}
