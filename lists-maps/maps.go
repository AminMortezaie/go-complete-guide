package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "google.com",
		"Amazon": "aws.com",
	}
	websites["LinkedIn"] = "linkedin.com"

	delete(websites, "Google")
	fmt.Printf("%v\n", websites)
}
