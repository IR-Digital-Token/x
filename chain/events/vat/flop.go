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

type FlopHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFlop]
}

func (h *FlopHandler) Signature() string {
	return "0xfa3d951cbf852d2a9cc2dfc9fc6b57914afbf57597ecf432c403ed2d74124b2c"
}

func (h *FlopHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFlop(log)
}

func (h *FlopHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFlop)
	if !ok {
		return errors.New("event type is not VatFlop")
	}
	return h.callback(e)
}

func (h *FlopHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFlop(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFlopHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFlop]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlopHandler{
		binding:  b,
		callback: callback,
	}
}
