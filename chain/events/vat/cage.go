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

type CageHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatCage]
}

func (h *CageHandler) Signature() string {
	return "0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda"
}

func (h *CageHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseCage(log)
}

func (h *CageHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatCage)
	if !ok {
		return errors.New("event type is not VatCage")
	}
	return h.callback(e)
}

func (h *CageHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseCage(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewCageHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatCage]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &CageHandler{
		binding:  b,
		callback: callback,
	}
}
