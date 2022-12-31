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

type FileHandler struct {
	binding  *Clipper
	callback events.CallbackFn[ClipperFile]
}

func (h *FileHandler) ID() string {
	return "0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7"
}

func (h *FileHandler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFile(log)
}

func (h *FileHandler) HandleEvent(header types.Header, event interface{}) error {
	e, ok := event.(ClipperFile)
	if !ok {
		return errors.New("event type is not ClipperFile")
	}
	return h.callback(header, e)
}

func (h *FileHandler) DecodeAndHandle(header types.Header, log types.Log) error {
	e, err := h.binding.ParseFile(log)
	if err != nil {
		return err
	}
	return h.callback(header, *e)
}

func NewFileHandler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperFile]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &FileHandler{
		binding:  b,
		callback: callback,
	}
}
