package genericsds

type node[T any] struct {
	Datum    T
	Previous *node[T]
	Next     *node[T]
}
