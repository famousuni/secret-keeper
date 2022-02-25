package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	secret_keeper "secret-keeper"
)

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Gets a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret_keeper.File(encodingKey, secretsPath())
		//fmt.Println(args)
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("no value set")
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}