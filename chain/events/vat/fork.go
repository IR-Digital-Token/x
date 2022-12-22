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

type ForkHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFork]
}

func (h *ForkHandler) Signature() string {
	return "0x4b67161d2a4293b296b2f023c52ea4214353fa4f03e58973572faa00097dbd1e"
}

func (h *ForkHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFork(log)
}

func (h *ForkHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFork)
	if !ok {
		return errors.New("event type is not VatFork")
	}
	return h.callback(e)
}

func (h *ForkHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFork(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewForkHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFork]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &ForkHandler{
		binding:  b,
		callback: callback,
	}
}
