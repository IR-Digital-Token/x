// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package clipper

import (
	"errors"
	"github.com/IR-Digital-Token/x/chain/bindings/clipper"
	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RedoHandler struct {
	binding  *clipper.Clipper
	callback events.CallbackFn[clipper.ClipperRedo]
}

func (h *RedoHandler) Signature() string {
	return "0x275de7ecdd375b5e8049319f8b350686131c219dd4dc450a08e9cf83b03c865f"
}

func (h *RedoHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseRedo(log)
}

func (h *RedoHandler) HandleEvent(event interface{}) error {
	e, ok := event.(clipper.ClipperRedo)
	if !ok {
		return errors.New("event type is not ClipperRedo")
	}
	return h.callback(e)
}

func (h *RedoHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseRedo(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewRedoHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[clipper.ClipperRedo]) events.Handler {
	b, err := clipper.NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &RedoHandler{
		binding:  b,
		callback: callback,
	}
}
