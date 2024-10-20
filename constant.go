package demo

type CompareType int

const (
  EQ CompareType = iota
  NE
  GT
  GE
  LT
  LE
  IsNull
  IsNotNull
  In
  NotIn
)

func (c CompareType) Bytes() []byte {
  return [...][]byte{
    {'='},
    {'<','>'},
    {'>'},
    {'>','='},
    {'<'},
    {'<','='},
    {'I','S',' ','N','U','L','L'},
    {'I','S',' ','N','O','T',' ','N','U','L','L'},
    {'I','N'},
    {'N','O','T',' ','I','N'},
  }[c]
}
