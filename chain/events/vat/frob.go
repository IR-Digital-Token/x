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

type FrobHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFrob]
}

func (h *FrobHandler) Signature() string {
	return "0xe37707842c8387f7c3c357f1d6c5bf57084e681573bdda024fae70cf0ecde80e"
}

func (h *FrobHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFrob(log)
}

func (h *FrobHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFrob)
	if !ok {
		return errors.New("event type is not VatFrob")
	}
	return h.callback(e)
}

func (h *FrobHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFrob(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFrobHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFrob]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FrobHandler{
		binding:  b,
		callback: callback,
	}
}
