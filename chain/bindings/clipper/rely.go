// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package clipper

import (
	"errors"

	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RelyHandler struct {
	binding  *Clipper
	callback events.CallbackFn[ClipperRely]
}

func (h *RelyHandler) ID() string {
	return "0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60"
}

func (h *RelyHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseRely(log)
}

func (h *RelyHandler) HandleEvent(header types.Header, event interface{}) error {
	e, ok := event.(ClipperRely)
	if !ok {
		return errors.New("event type is not ClipperRely")
	}
	return h.callback(header, e)
}

func (h *RelyHandler) DecodeAndHandle(header types.Header, log types.Log) error {
	e, err := h.binding.ParseRely(log)
	if err != nil {
		return err
	}
	return h.callback(header, *e)
}

func NewRelyHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperRely]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RelyHandler{
		binding:  b,
		callback: callback,
	}
}
