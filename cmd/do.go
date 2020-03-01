package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/lexty/go-dns-yandex-do-migrate/api"
	"github.com/spf13/cobra"
)

func init() {
	doCmd.Flags().StringVarP(&token, "token", "t", "", "Access token")
	doCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain name")
	// doCmd.Flags().BoolVarP(&prettyPrint, "pretty-print", "p", false, "Pretty print JSON")
	doCmd.MarkFlagRequired("token")
	doCmd.MarkFlagRequired("domain")
	rootCmd.AddCommand(doCmd)
}

func readPipeInput(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Loads DNS records to DigitalOcean",
	Run: func(cmd *cobra.Command, args []string) {

		recordsJSON := readPipeInput(os.Stdin)
		var records []api.DNSRecord
		err := json.Unmarshal(recordsJSON, &records)
		if err != nil {
			panic(err)
		}

		client := api.DOClient{}
		for _, record := range records {
			recordRequest := api.DOAPICreateRecordRequest{
				Type:     record.Type,
				Data:     record.Content,
				TTL:      record.TTL,
				Priority: record.Priority,
				Name:     record.Subdomain,
			}
			if record.Priority == "" {
				recordRequest.Priority = nil
			} else {
				recordRequest.Priority = record.Priority
			}
			recordRequest.Weight = nil
			recordRequest.Flags = nil
			recordRequest.Tag = nil
			recordRequest.Port = nil

			doRecord, err := client.CreateRecord(recordRequest, token, domain)
			if err != nil {
				panic(err)
			}
			if doRecord != nil {
				log.Printf("[DO] Created record with ID %d\n", doRecord.ID)
			}
		}
	},
}
