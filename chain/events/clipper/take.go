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

type TakeHandler struct {
	binding  *clipper.Clipper
	callback events.CallbackFn[clipper.ClipperTake]
}

func (h *TakeHandler) Signature() string {
	return "0x05e309fd6ce72f2ab888a20056bb4210df08daed86f21f95053deb19964d86b1"
}

func (h *TakeHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseTake(log)
}

func (h *TakeHandler) HandleEvent(event interface{}) error {
	e, ok := event.(clipper.ClipperTake)
	if !ok {
		return errors.New("event type is not ClipperTake")
	}
	return h.callback(e)
}

func (h *TakeHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseTake(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewTakeHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[clipper.ClipperTake]) events.Handler {
	b, err := clipper.NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &TakeHandler{
		binding:  b,
		callback: callback,
	}
}
