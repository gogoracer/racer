/* cSpell:disable */

package handlebars

import "github.com/gogoracer/racer/pkg/engine"

func OnClick(onClick engine.EventHandler) *engine.EventBinding {
	return engine.On("click", onClick)
}

func OnKeyDown(onKeyDown engine.EventHandler) *engine.EventBinding {
	return engine.On("keydown", onKeyDown)
}

func OnKeyUp(onKeyUp engine.EventHandler) *engine.EventBinding {
	return engine.On("keyup", onKeyUp)
}
