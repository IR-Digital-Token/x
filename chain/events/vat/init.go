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

type InitHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatInit]
}

func (h *InitHandler) Signature() string {
	return "0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555"
}

func (h *InitHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseInit(log)
}

func (h *InitHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatInit)
	if !ok {
		return errors.New("event type is not VatInit")
	}
	return h.callback(e)
}

func (h *InitHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseInit(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewInitHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatInit]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &InitHandler{
		binding:  b,
		callback: callback,
	}
}
