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

type FoldHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFold]
}

func (h *FoldHandler) Signature() string {
	return "0x8e03d1ac49b6d4e90dd1c4e641ecc5ca76b7c07a487690b6897c0e5e374b19d2"
}

func (h *FoldHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFold(log)
}

func (h *FoldHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFold)
	if !ok {
		return errors.New("event type is not VatFold")
	}
	return h.callback(e)
}

func (h *FoldHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFold(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFoldHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFold]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FoldHandler{
		binding:  b,
		callback: callback,
	}
}
