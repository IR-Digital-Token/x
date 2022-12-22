package events

import (
	"html/template"
	"os"
)

type GenData struct {
	Package               string
	BindingPath           string
	BindingEventName      string
	BindingEventSignature string
	BindingContract       string
}

func CodeGen(d GenData, output string) error {
	t := template.Must(template.New("event_handler").Parse(eventHandlerTemplate))
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	return t.Execute(f, d)
}

var eventHandlerTemplate = `// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package {{.Package}}

import (
	"errors"

	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type {{.BindingEventName}}Handler struct {
	binding  *{{.BindingContract}}
	callback events.CallbackFn[{{.BindingContract}}{{.BindingEventName}}]
}

func (h *{{.BindingEventName}}Handler) ID() string {
	return "{{.BindingEventSignature}}"
}

func (h *{{.BindingEventName}}Handler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.Parse{{.BindingEventName}}(log)
}

func (h *{{.BindingEventName}}Handler) HandleEvent(event interface{}) error {
	e, ok := event.({{.BindingContract}}{{.BindingEventName}})
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

func New{{.BindingEventName}}Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[{{.BindingContract}}{{.BindingEventName}}]) events.Handler {
	b, err := New{{.BindingContract}}(addr, eth)
	if err != nil {
		panic(err)
	}
	return &{{.BindingEventName}}Handler{
		binding:  b,
		callback: callback,
	}
}
`
