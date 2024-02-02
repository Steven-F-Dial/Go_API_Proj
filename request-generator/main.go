package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const url = "http://0.0.0.0:9090"

func main() {
	nextID := getNextID(url)
	_, d := buildRequest(nextID, url)

	postAlbum(d, url)

}

func buildRequest(id int, url string) (curl string, data string) {

	fmt.Println("input Title, Author, Price:")
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		if len(lines) == 3 {
			break
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}

	curl = fmt.Sprintf(`curl "%s/albums" \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{"id": "%d","title": "%s","artist": "%s","price": %s}'`, url, id, strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]), strings.TrimSpace(lines[2]))
	data = fmt.Sprintf(`{"id": "%d","title": "%s","artist": "%s","price": %s}`, id, strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]), strings.TrimSpace(lines[2]))
	return curl, data

}

func getNextID(url string) int {
	c, err := http.Get(url + "/secretAPI")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, c.Body)
	s := buf.String()
	num, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return num
}

func postAlbum(data string, url string) {
	d := []byte(data)
	u := url + "/albums"
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(d))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Println("Successfully Posted")
}
