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

type HealHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatHeal]
}

func (h *HealHandler) Signature() string {
	return "0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7"
}

func (h *HealHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseHeal(log)
}

func (h *HealHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatHeal)
	if !ok {
		return errors.New("event type is not VatHeal")
	}
	return h.callback(e)
}

func (h *HealHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseHeal(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewHealHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatHeal]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &HealHandler{
		binding:  b,
		callback: callback,
	}
}
