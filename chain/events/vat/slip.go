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

type SlipHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatSlip]
}

func (h *SlipHandler) Signature() string {
	return "0x0d5f62756a04d37a9bb68fd20b97c7c6a03e96ab87385a99f99c2463157dba4e"
}

func (h *SlipHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseSlip(log)
}

func (h *SlipHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatSlip)
	if !ok {
		return errors.New("event type is not VatSlip")
	}
	return h.callback(e)
}

func (h *SlipHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseSlip(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewSlipHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatSlip]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SlipHandler{
		binding:  b,
		callback: callback,
	}
}
