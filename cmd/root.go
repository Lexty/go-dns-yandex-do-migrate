package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "dns-yandex-do-migrate",
		Short: "Migrate DNS from Yandex to DigitalOcean",
		Long:  `Extract DNS records from Yandex.Connect and load to DigitalOcean.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	viper.AutomaticEnv()
}
