package main

import (
	"strconv"
	"strings"
)

func splitHostPortFromURL(url string) (host string, port int) {

	// 192.168.102:8080

	split := strings.Split(url, ":")

	host = split[0]

	if len(split) == 1 {
		port = -1
		return
	}

	var err error

	port, err = strconv.Atoi(split[1])

	if err != nil {
		port = -1
		return
	}

	return
}
