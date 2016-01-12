var m map[string]int
m["k1"] = 7 // 错误: assignment to entry in nil map

m := make(map[string]int) // 初始化成一个空 map: [], 可以进行读写
m["k1"] = 7
m["k2"] = 13
delete(m, "k2")      // 删除 "k2" 对应的值
v, ok := m["k2"]     // 检查 "k2" 对应的值是否存在
fmt.Println(m["k3"]) // 不存在的 key 默认返回空值

for k, v := range m { // 迭代的结果每次可能不一致
	fmt.Println(k, v)
}
