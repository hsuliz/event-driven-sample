package kafka

type EngineMsg struct {
	MatrixHash string
	Done       bool
	Value      int
}
