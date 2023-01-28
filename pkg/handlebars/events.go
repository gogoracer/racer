package handlebars

import "github.com/gogoracer/racer/pkg/gas"

func OnClick(onClick gas.EventHandler) *gas.EventBinding {
	return gas.On("click", onClick)
}

func OnKeyDown(onKeyDown gas.EventHandler) *gas.EventBinding {
	return gas.On("keydown", onKeyDown)
}

func OnKeyUp(onKeyUp gas.EventHandler) *gas.EventBinding {
	return gas.On("keyup", onKeyUp)
}
