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

type FlapHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatFlap]
}

func (h *FlapHandler) Signature() string {
	return "0x4c8a914f94eb278fed11934ac386da1b691d8c2e8a30291487140881038d13ce"
}

func (h *FlapHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFlap(log)
}

func (h *FlapHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatFlap)
	if !ok {
		return errors.New("event type is not VatFlap")
	}
	return h.callback(e)
}

func (h *FlapHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseFlap(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewFlapHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatFlap]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FlapHandler{
		binding:  b,
		callback: callback,
	}
}
