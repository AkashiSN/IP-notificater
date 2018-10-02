package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	API string
)

func HttpPost(url, msg string) error {
	jsonStr := `{"text":"` + msg + `"}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func FetchGlobalIP() (string, error) {
	url := "http://inet-ip.info/ip"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byteArray), nil
}

func LoadGlobalIP() (string, error) {
	data, err := ioutil.ReadFile("ip")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteGlobalIP(globalIP string) error {
	content := []byte(globalIP)
	err := ioutil.WriteFile("ip", content, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	cron := flag.Bool("cron", false, "Cron execute")
	flag.Parse()

	GlobalIP, err := FetchGlobalIP()
	if err != nil {
		fmt.Println(err)
	}

	if *cron {
		now := time.Now().Format("2006/01/02")
		msg := fmt.Sprintf("Day : %s\nGlobal IP Address : %s", now, GlobalIP)
		fmt.Println(msg)
		HttpPost(API, msg)
		WriteGlobalIP(GlobalIP)
		os.Exit(0)
	}

	pGlobalIP, err := LoadGlobalIP()
	if err != nil {
		fmt.Println(err)
	}

	if GlobalIP != pGlobalIP {
		now := time.Now().Format("2006/01/02 15:4:5")
		msg := fmt.Sprintf("Time : %s\nChanged Global IP Address : %s -> %s", now, pGlobalIP, GlobalIP)
		fmt.Println(msg)
		HttpPost(API, msg)
	}

	WriteGlobalIP(GlobalIP)
}
