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

func (c CompareType) Bytes() {
  return [...][]byte{
    []byte{'='},
    []byte{'<','>'},
    []byte{'>'},
    []byte{'>','='},
    []byte{'<'},
    []byte{'<','='},
    []byte{'I','S',' ','N','U','L','L'},
    []byte{'I','S',' ','N','O','T',' ','N','U','L','L'},
    []byte{'I','N'},
    []byte{'N','O','T',' ','I','N'},
  }[c]
}
