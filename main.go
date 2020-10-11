package main

import (
	"fmt"
	"strings"
	"os"
	"flag"
	"bufio"
)
var delPort bool

func main() {
	sc := bufio.NewScanner(os.Stdin)
	portsMap := make(map[string][]string)
	flag.BoolVar(&delPort, "d", false,"删除80 443端口输出，默认全部输出")

	flag.Parse()
	for sc.Scan() {
		line := strings.ToLower(sc.Text())
		s := strings.Fields(line)
		ip := s[1]
		portNumber := strings.Split(s[4], "/")[0]

		portsMap[ip] = append(portsMap[ip], portNumber)


//		fmt.Println(ip)
//		fmt.Println(portNumber)
//		fmt.Println(line)
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}


	for ip, ports := range portsMap {
		if (len(ports)) <= 80 {
			for _, port := range ports {
				if !delPort {
					fmt.Println(ip + ":" + port)
					continue
				}

				if port == "80" || port == "443" {
					continue
				}

				fmt.Println(ip + ":" + port)
			}
		}
	}

}
