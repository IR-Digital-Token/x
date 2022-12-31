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

type RedoHandler struct {
	binding  *Clipper
	callback events.CallbackFn[ClipperRedo]
}

func (h *RedoHandler) ID() string {
	return "0x275de7ecdd375b5e8049319f8b350686131c219dd4dc450a08e9cf83b03c865f"
}

func (h *RedoHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseRedo(log)
}

func (h *RedoHandler) HandleEvent(header types.Header, event interface{}) error {
	e, ok := event.(ClipperRedo)
	if !ok {
		return errors.New("event type is not ClipperRedo")
	}
	return h.callback(header, e)
}

func (h *RedoHandler) DecodeAndHandle(header types.Header, log types.Log) error {
	e, err := h.binding.ParseRedo(log)
	if err != nil {
		return err
	}
	return h.callback(header, *e)
}

func NewRedoHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperRedo]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RedoHandler{
		binding:  b,
		callback: callback,
	}
}
