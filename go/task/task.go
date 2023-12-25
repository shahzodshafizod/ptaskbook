package task

type Task interface {
	GetCount() int
	GetName() string
	SetData(taskNo int, testNo int) error
	ScannerEOF() bool
	CheckerEOF() bool
	CleanData()
	GetFn(int) func() bool
}
