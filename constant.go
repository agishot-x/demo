package demo

type CompareType int

const (
  Eq CompareType = iota
  GT
  GE
  LT
  LE
  IsNull
  IsNotNull
  In
  NotIn
)
