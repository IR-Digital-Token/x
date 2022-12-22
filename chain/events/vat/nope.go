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

type NopeHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatNope]
}

func (h *NopeHandler) Signature() string {
	return "0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc"
}

func (h *NopeHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseNope(log)
}

func (h *NopeHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatNope)
	if !ok {
		return errors.New("event type is not VatNope")
	}
	return h.callback(e)
}

func (h *NopeHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseNope(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewNopeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatNope]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &NopeHandler{
		binding:  b,
		callback: callback,
	}
}
