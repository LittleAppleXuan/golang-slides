type Person struc {
	Name string
	Age int
	parent string
}

p := new(Person)

p.Name = "Lishi" // 首字母大写字段，可导出，可在其它包中进行读写
p.Age = "28"
p.parent = "He" // 首字母小写字段，不可导出，不可在其它包中进行读写