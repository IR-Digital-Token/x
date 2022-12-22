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

type FluxHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFlux]
}

func (h *FluxHandler) Signature() string {
	return "0x5718eae79ffb8b6c98c497e5029a903705cf6a33a17aaab32de7fe198d8d8a0d"
}

func (h *FluxHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFlux(log)
}

func (h *FluxHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFlux)
	if !ok {
		return errors.New("event type is not VatFlux")
	}
	return h.callback(e)
}

func (h *FluxHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFlux(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFluxHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFlux]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FluxHandler{
		binding:  b,
		callback: callback,
	}
}
