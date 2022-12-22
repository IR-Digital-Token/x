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

type HopeHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatHope]
}

func (h *HopeHandler) Signature() string {
	return "0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43"
}

func (h *HopeHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseHope(log)
}

func (h *HopeHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatHope)
	if !ok {
		return errors.New("event type is not VatHope")
	}
	return h.callback(e)
}

func (h *HopeHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseHope(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewHopeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatHope]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &HopeHandler{
		binding:  b,
		callback: callback,
	}
}
