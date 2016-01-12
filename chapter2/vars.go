var i int = 10
i = 20 // 可修改
j := 20 // 默认类型 int，验证方法 fmt.Println(j+1.1)

var x interface{} // 初始值为空值 nil
var p *int // 初始值为空值 nil, 
var x int // 初始值为空值 0