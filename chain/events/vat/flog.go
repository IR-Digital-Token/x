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

type FlogHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFlog]
}

func (h *FlogHandler) Signature() string {
	return "0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753"
}

func (h *FlogHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFlog(log)
}

func (h *FlogHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFlog)
	if !ok {
		return errors.New("event type is not VatFlog")
	}
	return h.callback(e)
}

func (h *FlogHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFlog(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFlogHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFlog]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlogHandler{
		binding:  b,
		callback: callback,
	}
}
