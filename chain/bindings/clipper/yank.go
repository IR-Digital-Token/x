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

type YankHandler struct {
	binding  *Clipper
	callback events.CallbackFn[ClipperYank]
}

func (h *YankHandler) ID() string {
	return "0x2c5d2826eb5903b8fc201cf48094b858f42f61c7eaac9aaf43ebed490138144e"
}

func (h *YankHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseYank(log)
}

func (h *YankHandler) HandleEvent(header types.Header, event interface{}) error {
	e, ok := event.(ClipperYank)
	if !ok {
		return errors.New("event type is not ClipperYank")
	}
	return h.callback(header, e)
}

func (h *YankHandler) DecodeAndHandle(header types.Header, log types.Log) error {
	e, err := h.binding.ParseYank(log)
	if err != nil {
		return err
	}
	return h.callback(header, *e)
}

func NewYankHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperYank]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &YankHandler{
		binding:  b,
		callback: callback,
	}
}
