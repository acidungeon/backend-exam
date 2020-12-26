package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
import "strconv"
const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers,Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}

type Pokers []Poker

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PokerSelf()," ")
	}
	fmt.Println()
}

func shuffle(p Pokers) Pokers {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(p); i > 0; i-- {
		last := i - 1
		idx := rand.Intn(i)
		p[last], p[idx] = p[idx], p[last]
	}
	return p
}
func (p Pokers) Len() int{
	return  len(p)
}
func (p Pokers) Less(i,j int) bool{
	if p[i].Flower<p[j].Flower{
		return true
	}
	if p[i].Flower>p[j].Flower{
		return false
	}
	return p[i].Num<p[j].Num//Less的逻辑为，以Flower排序，当Flower相同时，以Num排序
}
func (p Pokers) Swap(i,j int) {
	p[i],p[j]=p[j],p[i]
}
func (p *Pokers) Sort(){
	sort.Sort(p)
}
func main()  {
	var p1 Pokers
	p1=CreatePokers()
	fmt.Println("洗牌前：")
	p1.Print()
	var randp Pokers
	randp=shuffle(p1)
	fmt.Println("洗牌后：")
	randp.Print()
	randp.Sort()
	fmt.Println("同花色排序后：")
	randp.Print()
}
//洗牌算法：使rand.Seed随时间更改，保证随机性，再反向遍历slice，用rand.Intn取当前位数的随机前一位，交换两Pocker，完成洗牌；
