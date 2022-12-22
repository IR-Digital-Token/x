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

type KickHandler struct {
	binding  *clipper.Clipper
	callback events.CallbackFn[clipper.ClipperKick]
}

func (h *KickHandler) Signature() string {
	return "0x7c5bfdc0a5e8192f6cd4972f382cec69116862fb62e6abff8003874c58e064b8"
}

func (h *KickHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseKick(log)
}

func (h *KickHandler) HandleEvent(event interface{}) error {
	e, ok := event.(clipper.ClipperKick)
	if !ok {
		return errors.New("event type is not ClipperKick")
	}
	return h.callback(e)
}

func (h *KickHandler) DecodeAndHandle(log types.Log) error {
	e, err := h.binding.ParseKick(log)
	if err != nil {
		return err
	}
	return h.callback(*e)
}

func NewKickHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[clipper.ClipperKick]) events.Handler {
	b, err := clipper.NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &KickHandler{
		binding:  b,
		callback: callback,
	}
}
