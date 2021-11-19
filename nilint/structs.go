package nilint

type NilInt struct {
	isNil bool
	val   int
}

func NewNil() NilInt {
	return NilInt{
		isNil: true,
	}
}

func NewInt(val int) NilInt {
	return NilInt{
		val: val,
	}
}

func (ni *NilInt) IsNil() bool {
	return ni.isNil
}

func (ni *NilInt) Val() int {
	return ni.val
}
