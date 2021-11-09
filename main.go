package main

import (
	"github.com/Hotchilli-Cloud/cli/cmd"
	// "io/ioutil"
	// "log"
	// "net/http"
)

func main() {
	cmd.Execute()
	// resp, err := http.Get("https://httpbin.org/get")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println(string(body))
}
