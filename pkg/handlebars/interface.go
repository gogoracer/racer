package handlebars

type any = interface{}

type IsHandlebarElement interface {
	HandlebarElement()
	GenerateVDOM() interface{}
}
