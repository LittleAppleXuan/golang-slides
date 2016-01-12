type S struct {
	a string
	b string
}

type B struct {
	S       // 匿名字段，只有类型 S, 字段名是 S
	b   int // 字段名是 b
	int     // 匿名字段，只有类型 int, 无意义
}

var b B
b.S.a = "a" // 给匿名字段 S 赋值
b.a = "a"   // 同上

fmt.Println(b)
// 输出结果: {{a } }
不是{a}

type String string

b.S.b = "b" // 给匿名字段 S 赋值
b.b = 10    // 名字冲突，不同上，不能这样给匿名字段中的相同字段赋值