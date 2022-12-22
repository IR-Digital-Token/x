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

type KissHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatKiss]
}

func (h *KissHandler) Signature() string {
	return "0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b"
}

func (h *KissHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseKiss(log)
}

func (h *KissHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatKiss)
	if !ok {
		return errors.New("event type is not VatKiss")
	}
	return h.callback(e)
}

func (h *KissHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseKiss(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewKissHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatKiss]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &KissHandler{
		binding:  b,
		callback: callback,
	}
}
