package task

type Task interface {
	Count() int
	Name() string
	En() []string
	Tj() []string
	Ru() []string
}
