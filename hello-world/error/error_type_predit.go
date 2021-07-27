package main

import (
	"fmt"
	"net"
)

func main() {
	//file, err := os.Open("/text.txt")
	//if err, ok := err.(*os.PathError); ok { // predit the error is PathError.
	//	fmt.Println(" file at ", err.Path, " doesn't exists.")
	//	return
	//}
	//fmt.Println("file at ", file.Name(), " open successfully.")

	fmt.Println("---------------")
	addrs, err := net.LookupHost("chenzuoli111111.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation time out.")
		} else if err.Temporary() {
			fmt.Println("temporary error.")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addrs)
}
