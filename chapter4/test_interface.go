package main

import (
	"errors"
	"fmt"
)

func UseCoder(c Coder) (err error) {
	c.WriteCode()
	//对于JavaCoder, salay:10000
	jc, ok := c.(JavaCoder)
	if !ok {
		return errors.New("not java coder")
	}
	fmt.Println("give jc $10000, jc:", jc)
	return
}

type Coder interface {
	WriteCode() (err error)
}

type BusinessMan interface {
	GetMoney() (err error)
}

type JavaCoder struct {
	Name string
}

func (jc JavaCoder) WriteCode() (err error) {
	fmt.Println("a java coder write it")
	return
}

func (jc JavaCoder) GetMoney() (err error) {
	return
}

type GoCoder struct {
	Name string
}

func (gc GoCoder) WriteCode() (err error) {
	fmt.Println("a go coder write it")
	return
}

func main() {
	jc := JavaCoder{
		Name: "xiaomin",
	}

	var bm BusinessMan
	bm = jc

	if _, ok := bm.(Coder); ok {
		fmt.Println("is a Coder")
	}
	return
	/**
	gc := GoCoder{
		Name: "laoA",
	}

	err := UseCoder(jc)
	if err != nil {
		fmt.Println(err)
	}
	err = UseCoder(gc)
	if err != nil {
		fmt.Println(err)
	}
	*/
}
