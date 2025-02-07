package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.xom/xorima/example-repo/internal/app"
)

var rootCmd = &cobra.Command{
	Use:   "example-repo",
	Short: "A very simple webserver we can use",
	Long:  `This is a webserver which returns some static data`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(app.NewApp().Run())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
