type S struct {
	a, b string
}

type B struct {
	S
}

func (s *S) String() {
	fmt.Println("In A: ", s.a, s.b)
}

func (b *B) String() {
	fmt.Println("In B: ", b.a, b.b)
}

var b B
b.a = "a"
b.b = "b"

b.S.String()
b.String() // 不等价于 b.S.String()
