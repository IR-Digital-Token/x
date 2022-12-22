package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

var eventHandlerTemplate = `// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package {{.Package}}

import (
	"errors"
	"{{.BindingPath}}"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type {{.BindingEventName}}Handler struct {
	binding  *{{.BindingPackage}}.{{.BindingContract}}
	callback events.CallbackFn[{{.BindingPackage}}.{{.BindingContract}}{{.BindingEventName}}]
}

func (h *{{.BindingEventName}}Handler) Signature() string {
	return "{{.BindingEventSignature}}"
}

func (h *{{.BindingEventName}}Handler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.Parse{{.BindingEventName}}(log)
}

func (h *{{.BindingEventName}}Handler) HandleEvent(event interface{}) error {
	e, ok := event.({{.BindingPackage}}.{{.BindingContract}}{{.BindingEventName}})
	if !ok {
		return errors.New("event type is not {{.BindingContract}}{{.BindingEventName}}")
	}
	return h.callback(e)
}

func (h *{{.BindingEventName}}Handler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.Parse{{.BindingEventName}}(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func New{{.BindingEventName}}Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[{{.BindingPackage}}.{{.BindingContract}}{{.BindingEventName}}]) events.Handler {
	b, err := {{.BindingPackage}}.New{{.BindingContract}}(addr, eth)
	if err != nil {
		panic(err)
	}
	return &{{.BindingEventName}}Handler{
		binding:  b,
		callback: callback,
	}
}
`

var rootCmd = &cobra.Command{
	Use:   "eh-gen",
	Short: "eh-gen generates event handlers from binding and events file",
	Run: func(cmd *cobra.Command, args []string) {
		contract, _ := cmd.Flags().GetString("contract")
		eventsFilePath, _ := cmd.Flags().GetString("events")
		bindingPackagePath, _ := cmd.Flags().GetString("binding-package")
		outputPath, _ := cmd.Flags().GetString("output-dir")

		eventsFile, err := os.Open(eventsFilePath)
		if err != nil {
			panic(err)
		}

		var events []Event
		err = json.NewDecoder(eventsFile).Decode(&events)
		if err != nil {
			panic(err)
		}

		for _, event := range events {
			data := Data{
				Package:               strings.ToLower(contract),
				BindingPath:           bindingPackagePath,
				BindingEventName:      event.Name,
				BindingEventSignature: event.Hash,
				BindingContract:       contract,
				BindingPackage:        strings.ToLower(contract),
			}
			codeGen(data, outputPath+"/"+strings.ToLower(event.Name)+".go")
		}
	},
}

type Event struct {
	Name      string `json:"name"`
	Signature string `json:"signature"`
	Hash      string `json:"hash"`
}

type Data struct {
	Package               string
	BindingPath           string
	BindingEventName      string
	BindingEventSignature string
	BindingContract       string
	BindingPackage        string
}

func codeGen(d Data, output string) {
	t := template.Must(template.New("event_handler").Parse(eventHandlerTemplate))
	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	err = t.Execute(f, d)
	if err != nil {
		panic(err)
	}
}

func Execute() {
	rootCmd.PersistentFlags().String("events", "", "")
	rootCmd.PersistentFlags().String("binding-package", "", "")
	rootCmd.PersistentFlags().String("output-dir", "", "")
	rootCmd.PersistentFlags().String("contract", "", "")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
