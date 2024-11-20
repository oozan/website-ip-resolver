package main

import (
	"fmt"
	"strings"
	errorhandler "website-ip-resolver/internal/handler"
	input "website-ip-resolver/internal/reader"
	"website-ip-resolver/internal/resolver"
)

func main() {
	fmt.Println("Website to IP Resolver")
	fmt.Println("Type a website name (e.g., google.com) and press Enter. Type 'exit' to quit.\n")

	resolver := resolver.NewResolver()
	reader := input.NewReader()

	for {
		fmt.Print("Enter website: ")
		domain, err := reader.ReadInput()
		if err != nil {
			errorhandler.Handle(err)
			continue
		}

		if strings.ToLower(domain) == "exit" {
			fmt.Println("Exiting application. Goodbye!")
			break
		}

		ips, err := resolver.Resolve(domain)
		if err != nil {
			errorhandler.Handle(err)
			continue
		}

		fmt.Printf("IP addresses for %s:\n", domain)
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}
}
