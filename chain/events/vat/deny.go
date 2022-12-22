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

type DenyHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatDeny]
}

func (h *DenyHandler) Signature() string {
	return "0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b"
}

func (h *DenyHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseDeny(log)
}

func (h *DenyHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatDeny)
	if !ok {
		return errors.New("event type is not VatDeny")
	}
	return h.callback(e)
}

func (h *DenyHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseDeny(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewDenyHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatDeny]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &DenyHandler{
		binding:  b,
		callback: callback,
	}
}
