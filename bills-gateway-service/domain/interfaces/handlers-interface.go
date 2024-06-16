package interfaces

type Handlers interface {
	DatabaseHandle() (result interface{}, err error)
}
