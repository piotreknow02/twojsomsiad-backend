package utils

import (
	"fmt"
	"strings"

	postman "github.com/rbretecher/go-postman-collection"
)

var (
	host     string
	port     string
	protocol string
)

func GenerateRequestCollections() {
	fmt.Println("Generate postman colletions")
}

func getUrl(endpoint string) *postman.URL {
	return &postman.URL{
		Raw:      fmt.Sprintf("%s://%s:%s%s", protocol, host, port, endpoint),
		Protocol: protocol,
		Host:     strings.Split(host, "."),
		Port:     port,
		Path:     strings.Split(endpoint, "/"),
	}
}
