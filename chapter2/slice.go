s := make([]string, 3)
s[0] = "a"
s[1] = "b"
s[2] = "c"
fmt.Println("切片元素: ", s)
fmt.Println("长度: ", len(s))
s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("Append 之后: ", s)

l = s[:5] // s[0]...s[4], 左闭右合
l = s[1:3] // s[1]..s[2], 左闭右合