package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/spf13/cobra"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var rootCmd = &cobra.Command{
	Use:   "eh-gen",
	Short: "eh-gen generates event handlers from binding and events file",
	Run: func(cmd *cobra.Command, args []string) {
		contract, _ := cmd.Flags().GetString("contract")
		outputPath, _ := cmd.Flags().GetString("output-dir")
		abiPath, _ := cmd.Flags().GetString("abi")

		abiFile, err := os.Open(abiPath)
		if err != nil {
			panic(err)
		}
		abi, err := abi.JSON(abiFile)
		if err != nil {
			panic(err)
		}
		for _, event := range abi.Events {
			data := events.GenData{
				Package:               toSnakeCase(contract),
				BindingEventName:      event.Name,
				BindingEventSignature: event.ID.Hex(),
				BindingContract:       contract,
			}
			err = events.CodeGen(data, outputPath+"/"+strings.ToLower(event.Name)+".go")
			if err != nil {
				panic(err)
			}
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().String("output-dir", "", "")
	rootCmd.PersistentFlags().String("contract", "", "")
	rootCmd.PersistentFlags().String("abi", "", "")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
