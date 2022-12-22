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

type MoveHandler struct {
	binding  *vat.Vat
	callback events.CallbackFn[vat.VatMove]
}

func (h *MoveHandler) Signature() string {
	return "0xdeb3a6837278f6e9914a507e4d73f08e841d8fca434fb97d4307b3b0d3d6b105"
}

func (h *MoveHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseMove(log)
}

func (h *MoveHandler) HandleEvent(event interface{}) error {
	e, ok := event.(vat.VatMove)
	if !ok {
		return errors.New("event type is not VatMove")
	}
	return h.callback(e)
}

func (h *MoveHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseMove(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewMoveHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[vat.VatMove]) events.Handler {
	b, err := vat.NewVat(addr, eth)
	if err != nil {
		panic(err)
	}
	return &MoveHandler{
		binding:  b,
		callback: callback,
	}
}
