package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	secret_keeper "secret-keeper"
)

var setCmd = &cobra.Command{
	Use: "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret_keeper.File(encodingKey, secretsPath())
		//fmt.Println(args)
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s = %s\n", key, value)
		fmt.Println("secret set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}