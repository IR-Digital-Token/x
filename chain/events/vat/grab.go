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

type GrabHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatGrab]
}

func (h *GrabHandler) Signature() string {
	return "0x1b2837fd40844c96cf39e52acaae7902fb74257fe20b1b7df5458b97d896c636"
}

func (h *GrabHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseGrab(log)
}

func (h *GrabHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatGrab)
	if !ok {
		return errors.New("event type is not VatGrab")
	}
	return h.callback(e)
}

func (h *GrabHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseGrab(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewGrabHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatGrab]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &GrabHandler{
		binding:  b,
		callback: callback,
	}
}
