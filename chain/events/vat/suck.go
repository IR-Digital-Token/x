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

type SuckHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatSuck]
}

func (h *SuckHandler) Signature() string {
	return "0x02d16dda43fd89f02e33ce23ecf0251cdc426807cc72ae74d37e8d3681dae7e5"
}

func (h *SuckHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseSuck(log)
}

func (h *SuckHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatSuck)
	if !ok {
		return errors.New("event type is not VatSuck")
	}
	return h.callback(e)
}

func (h *SuckHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseSuck(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewSuckHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatSuck]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &SuckHandler{
		binding:  b,
		callback: callback,
	}
}
