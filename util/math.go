package util

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

func Min[V Ordered](left V, right V) V {
	if left < right {
		return left
	}
	return right
}

func Max[V Ordered](left V, right V) V {
	if left < right {
		return right
	}
	return left
}