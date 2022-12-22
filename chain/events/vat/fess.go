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

type FessHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFess]
}

func (h *FessHandler) Signature() string {
	return "0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba"
}

func (h *FessHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFess(log)
}

func (h *FessHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFess)
	if !ok {
		return errors.New("event type is not VatFess")
	}
	return h.callback(e)
}

func (h *FessHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFess(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFessHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFess]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FessHandler{
		binding:  b,
		callback: callback,
	}
}
