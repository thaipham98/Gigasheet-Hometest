package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	Ips []string
}

var (
	url                    = "https://raw.githubusercontent.com/stamparm/ipsum/master/levels/5.txt"
	latestFetchDate  int64 = 0
	latestGithubData       = make(map[string]int)
)

func readFile() (map[string]int, error) {
	currentTime := time.Now().Unix()

	if currentTime <= (latestFetchDate + 0.5*24*60*60) {
		return latestGithubData, nil
	}

	resp, err := http.Get(url)

	ipMap := make(map[string]int)
	if err != nil {
		return ipMap, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ipMap, err
	}

	dataString := string(data)
	ips := strings.Fields(dataString)
	ips = filter(ips)
	for i, element := range ips {
		ipMap[element] = i
	}
	latestFetchDate = currentTime
	latestGithubData = ipMap
	resp.Body.Close()

	return ipMap, nil
}

func filter(ips []string) []string {
	var filteredIps []string
	for _, ip := range ips {
		if checkIPAddress(ip) {
			filteredIps = append(filteredIps, ip)
		}
	}
	return filteredIps
}

func checkIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		return true
	}
}

func main() {
	http.HandleFunc("/count_ips_in_ipsum", ipsumHandler)
	http.ListenAndServe(":8080", nil)
}

func ipsumHandler(writer http.ResponseWriter, request *http.Request) {
	ips, err := readFile()
	if err != nil {
		fmt.Fprintln(writer, err)
		return
	}

	defer request.Body.Close()

	decoder := json.NewDecoder(request.Body)
	var data Data
	err = decoder.Decode(&data)

	if err != nil {
		fmt.Fprintln(writer, err)
		return
	}

	count := 0
	for _, ip := range data.Ips {
		if checkIPAddress(ip) {
			_, found := ips[ip]
			if found {
				count++
			}
		} else {
			http.Error(writer, ip+" is not a valid ip address", http.StatusBadRequest)
			return
		}
	}

	fmt.Fprintln(writer, count)
	request.Body.Close()
}
