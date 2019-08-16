package main

import (
	"log"
	"regexp"
)

func main() {
	str := "projects/_/instances/redish-staging/refs/messages/u51075_mu50045/-LmO2R_dl-kB7fo6Xj2T"
	reg := regexp.MustCompile(`(u\d*_mu\d*)`)
	log.Printf("Match: %v", reg.FindAllString(str, -1))
}
