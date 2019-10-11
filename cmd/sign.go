package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(args)
		msg, err := ReadStdMsgFromFile(args[0])
		name := viper.GetString("name")

		fmt.Println("message is " + msg)
		fmt.Println("message is " + name)
		return err
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.Flags().StringP("name", "n", "sahith", "Set your name")
	_ = viper.BindPFlag("name", signCmd.Flags().Lookup("name"))

	// Here you will define your flags and configuratioxn settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
