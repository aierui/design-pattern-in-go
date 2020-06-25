package vistor

type shape interface {
	getType() string
	accept(visitor)
}
