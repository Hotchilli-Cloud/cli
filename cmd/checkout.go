package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"

	"io/ioutil"
	"log"
	"net/http"

	"bytes"
	"encoding/json"
)

func init() {
	rootCmd.AddCommand(checkoutCmd)

	checkoutCmd.AddCommand(sessionCmd)
	sessionCmd.AddCommand(getSessionCmd)
	getSessionCmd.Flags().String("id", "", "Checkout session ID")
	getSessionCmd.Flags().String("extend", "", "Extend API Request - Go deeper in struct")

}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Deploy artifacts (web, api or database)",
	Long:  `This command can be used together with web, api or database sub-commands to deploy respective artifacts`,
}

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Deploy web artifacts",
	Long:  `This command can be used to deploy web artifacts`,
}

var getSessionCmd = &cobra.Command{
	Use:   "get",
	Short: "Deploy web artifacts",
	Long:  `This command can be used to deploy web artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		id, _ := cmd.Flags().GetString("id")
		extend, _ := cmd.Flags().GetString("extend")

		getCheckoutSession(id, extend)
	},
}

func getCheckoutSession(id string, extend string) {
	var urlQuery string
	if extend != "" {
		urlQuery = "?extend=" + extend
	} else {
		urlQuery = ""
	}
	resp, err := http.Get("http://localhost:3500/checkout/session/" + id + urlQuery)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")

	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Println(prettyJSON.String())
}
