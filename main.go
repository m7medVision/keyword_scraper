package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	client := &http.Client{}
	sites, _ := os.ReadFile("sites.txt")
	spilt_sites := strings.Split(string(sites), "\n")
	for _, site := range spilt_sites {
		fmt.Println(site)
		fmt.Println("-----------------")
		site_trim := strings.TrimSpace(site)
		req, _ := http.NewRequest("GET", site_trim, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
		resp, _ := client.Do(req)
		bodybytes, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		keywords_regex := `<meta name="keywords" content="(.*?)"`
		compiled_regex, _ := regexp.Compile(keywords_regex)
		keywords := compiled_regex.FindStringSubmatch(string(bodybytes))[1]
		split_keywords := strings.Split(keywords, ", ")
		f, err := os.OpenFile("keywords.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		for _, keyword := range split_keywords {
			if _, err = f.WriteString(keyword + "\n"); err != nil {
				fmt.Println(err)
			}
		}
	}

}
