package racer_event

import "github.com/gogoracer/racer/pkg/engine"

func CLICK(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("click", eventHandler)
}

func ANIMATIONEND(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("animationend", eventHandler)
}

func ANIMATIONCANCEL(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("animationcancel", eventHandler)
}

func MOUSEENTER(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("mouseEnter", eventHandler)
}

func MOUSELEAVE(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("mouseLeave", eventHandler)
}

func KEYUP(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("keyup", eventHandler)
}

func KEYDOWN(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("keydown", eventHandler)
}

func INPUT_INPUT(eventHandler engine.EventHandler) *engine.EventBinding {
	return engine.On("input", eventHandler)
}
