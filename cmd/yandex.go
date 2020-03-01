package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/lexty/go-dns-yandex-do-migrate/api"
	"github.com/spf13/cobra"
)

var token string
var domain string
var prettyPrint bool

func init() {
	yandexCmd.Flags().StringVarP(&token, "token", "t", "", "Admin token")
	yandexCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain name")
	yandexCmd.Flags().BoolVarP(&prettyPrint, "pretty-print", "p", false, "Pretty print JSON")
	yandexCmd.MarkFlagRequired("token")
	yandexCmd.MarkFlagRequired("domain")
	rootCmd.AddCommand(yandexCmd)
}

var yandexCmd = &cobra.Command{
	Use:   "yandex",
	Short: "Extract yandex DNS records",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.YandexClient{}
		records, err := client.ListRecords(token, domain)
		if err != nil {
			panic(err)
		}

		var result []byte
		if prettyPrint {
			result, err = json.MarshalIndent(records, "", "    ")
		} else {
			result, err = json.Marshal(records)
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", result)
	},
}
