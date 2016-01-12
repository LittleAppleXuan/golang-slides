package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	om := &OrderMarket{}
	om.Init()

	cm := &CoderMarket{}
	cm.Init()

	totalMoney := 100.0 //单位：w

	for {
		if totalMoney > 1000.0 || totalMoney < 0 {
			fmt.Println("创业结束. totalMoney:%v", totalMoney)
			return
		}

		fmt.Println("新一轮工作开始")
		order, err := om.GetOrder()
		if err != nil {
			fmt.Println("没有拉到订单，休息一天后继续")
			time.Sleep(time.Second * 1) //1秒模拟1天
			totalMoney -= 0.1
			continue
		}

		coder, err := cm.GetCoder()
		if err != nil {
			fmt.Println("没有找到程序员，退回订单后继续")
			om.ReturnOrder(order)
			totalMoney -= order.Money * 2
			continue
		}

		software, err := coder.WriteCode(order)
		if err != nil {
			fmt.Println("开发出错，公司倒闭")
			return
		}

		om.DeliverOrder(order, software)
		totalMoney += order.Money

		moneyToCoder := order.Money / 10.0
		coder.GetMoney(moneyToCoder)
		totalMoney -= moneyToCoder

		fmt.Println("给程序员买盒饭闪了腰，休息两天\n")
		totalMoney -= 0.1 * 2
		time.Sleep(time.Second * 2) //1秒模拟1天
	}
}

type Order struct {
	Money float64 //w
	Name  string
}

type OrderMarket struct {
	orderPool []Order
}

func (om *OrderMarket) Init() {

	rand.Seed(time.Now().Unix())
	var orderPool []Order
	for i := 0; i < 10000; i++ {
		order := Order{
			Money: float64(rand.Int()%50 + 50),
			Name:  "order:" + strconv.Itoa(i),
		}
		orderPool = append(orderPool, order)
	}

	om.orderPool = orderPool
}

func (om *OrderMarket) GetOrder() (order Order, err error) {
	order = om.orderPool[0]
	om.orderPool = append(om.orderPool[:1], om.orderPool[:1]...)
	fmt.Printf("get an order:%v money:%vw\n", order.Name, order.Money)
	return
}

func (om *OrderMarket) ReturnOrder(order Order) {
	om.orderPool = append(om.orderPool, order)
	fmt.Println("ret an no complete order:", order.Name)
}

func (om *OrderMarket) DeliverOrder(order Order, sf SoftWare) {
	fmt.Println("deliver order success. software info:", sf.Name)
}

type CoderMarket struct {
	coderPool []Coder
}

func (cm *CoderMarket) Init() {

	gc := &GoCoder{
		Name: "xiaomin",
	}
	cm.coderPool = append(cm.coderPool, gc)

	gc = &GoCoder{
		Name: "laoA",
	}
	cm.coderPool = append(cm.coderPool, gc)

	dc := &Dentist{
		Name: "yu_ba_ya",
	}
	cm.coderPool = append(cm.coderPool, dc)
}

func (cm *CoderMarket) GetCoder() (coder Coder, err error) {

	index := rand.Int() % len(cm.coderPool)
	coder = cm.coderPool[index]
	fmt.Println("success get a order.")
	return
}

type SoftWare struct {
	Name string
}

type Coder interface {
	WriteCode(order Order) (SoftWare, error)
	GetMoney(money float64)
}

type GoCoder struct {
	Name string
}

func (gc *GoCoder) WriteCode(order Order) (sf SoftWare, err error) {

	fmt.Printf("begin to write code for %v\n", order.Name)
	time.Sleep(time.Second * 3)
	fmt.Println("write code ok.")

	sf.Name = fmt.Sprintf("hello, world! by %v", gc.Name)
	return
}

func (gc *GoCoder) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", gc.Name, money)
}

type Doctor interface {
	Cure() (err error)
	GetMoney(money float64)
}

type Dentist struct {
	Name string
}

func (d *Dentist) Cure() (err error) {
	//拔牙
	return
}

func (d *Dentist) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", d.Name, money)
}

func (d *Dentist) WriteCode(order Order) (sf SoftWare, err error) {
	fmt.Printf("i'm a dentist. begin to write code for %v\n", order.Name)
	time.Sleep(time.Second * 3)
	fmt.Println("write code ok.")

	sf.Name = fmt.Sprintf("hello, world! by dentist:%v", d.Name)
	return
}
