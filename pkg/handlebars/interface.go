package handlebars

type any = interface{}

type IsHandlebarElement interface {
	HandlebarElement()
	GenerateVDOM() interface{}
}

type TextWrapper string

func Text(text string) TextWrapper {
	return TextWrapper(text)
}

func (t TextWrapper) HandlebarElement() {}

func (t TextWrapper) GenerateVDOM() interface{} {
	return string(t)
}

type BoundTextWrapper struct {
	Value *string
}

func BoundText(value *string) BoundTextWrapper {
	return BoundTextWrapper{Value: value}
}

func (t BoundTextWrapper) HandlebarElement() {}

func (t BoundTextWrapper) GenerateVDOM() interface{} {
	return t.Value
}
