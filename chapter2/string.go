var s string = "Hello, World"
s[0] = 'h' // error, 不可修改
fmt.Println(s[0]) // 72, 输出 Unicode
fmt.Println("Hello, World"[0]) // 72, 输出 Unicode