type A struct {
	a string
	b byte
	c byte
}

type B struct {
	a byte
	b string
	c byte
}

unsafe.Sizeof(A) => 24
unsafe.Sizeof(B) => 32