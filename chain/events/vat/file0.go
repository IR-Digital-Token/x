// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package vat

import (
	"errors"
	"github.com/IR-Digital-Token/x/chain/bindings/vat"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type File0Handler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFile0]
}

func (h *File0Handler) Signature() string {
	return "0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1"
}

func (h *File0Handler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFile0(log)
}

func (h *File0Handler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFile0)
	if !ok {
		return errors.New("event type is not VatFile0")
	}
	return h.callback(e)
}

func (h *File0Handler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFile0(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFile0Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFile0]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File0Handler{
		binding:  b,
		callback: callback,
	}
}
