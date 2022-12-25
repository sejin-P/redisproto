/*
Copyright © 2022 Sejin-P chunpark37@gmail.com
*/
package cmd

import (
	"os"

	"github.com/sejin-P/redisproto/redis"
	"github.com/spf13/cobra"
)

var (
	address  string
	password string
	action   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "redisproto",
	Short: "Redis viewer with proto binary format",
	Long:  `You can use this tool to view redis data with proto binary format. At first, I'll support viewer feature by only cli. But I'll add web viewer feature later.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		redisCli := redis.New(address, password)
		if err := redisCli.Ping(cmd.Context()).Err(); err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.redisproto.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVar(&address, "address", "localhost:6379", "redis host")
	rootCmd.Flags().StringVar(&password, "password", "", "redis password")
	rootCmd.Flags().StringVar(&action, "action", "view", "action")
}
