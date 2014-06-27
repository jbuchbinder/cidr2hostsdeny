package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	svc   = flag.String("services", "sshd", "Service(s) to block")
	masks = map[int]string{
		4:  "240.0.0.0",
		5:  "248.0.0.0",
		6:  "252.0.0.0",
		7:  "254.0.0.0",
		8:  "255.0.0.0",
		9:  "255.128.0.0",
		10: "255.192.0.0",
		11: "255.224.0.0",
		12: "255.240.0.0",
		13: "255.248.0.0",
		14: "255.252.0.0",
		15: "255.254.0.0",
		16: "255.255.0.0",
		17: "255.255.128.0",
		18: "255.255.192.0",
		19: "255.255.224.0",
		20: "255.255.240.0",
		21: "255.255.248.0",
		22: "255.255.252.0",
		23: "255.255.254.0",
		24: "255.255.255.0",
		25: "255.255.255.128",
		26: "255.255.255.192",
		27: "255.255.255.224",
		28: "255.255.255.240",
		29: "255.255.255.248",
		30: "255.255.255.252",
		32: "255.255.255.255",
	}
)

func main() {
	flag.Parse()
	files := flag.Args()
	for _, file := range files {
		fmt.Println("## Source : " + file)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("## Error processing : " + err.Error())
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
			if line[0] == '#' {
				fmt.Println(line)
				continue
			}

			slashloc := strings.IndexAny(line, "/")
			if slashloc == -1 {
				continue
			}

			spaceloc := strings.IndexAny(line, " ")

			// Line header
			fmt.Print(*svc + ": ")

			// IP address
			fmt.Print(line[:slashloc] + "/")

			// Get netmask based on block
			nmask, _ := strconv.Atoi(line[slashloc+1 : spaceloc])
			fmt.Print(masks[nmask])

			// EOL
			fmt.Println("")
		}
	}
}
