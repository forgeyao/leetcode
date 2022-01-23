// https://leetcode-cn.com/problems/stock-price-fluctuation/
package main

import (
	"container/heap"
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

/**
 * 方法一
 * 哈希表 + 有序集合
 */
type StockPrice struct {
	record map[int]int
	curKey int
	prices *redblacktree.Tree
}

func Constructor() StockPrice {
	return StockPrice{record: map[int]int{}, prices: redblacktree.NewWithIntComparator()}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if prevPrice := this.record[timestamp]; prevPrice > 0 {
		if times, _ := this.prices.Get(prevPrice); times.(int) > 1 {
			this.prices.Put(prevPrice, times.(int)-1)
		} else {
			this.prices.Remove(prevPrice)
		}
	}
	times := 0
	if val, ok := this.prices.Get(price); ok {
		times = val.(int)
	}
	this.prices.Put(price, times+1)
	this.record[timestamp] = price
	if timestamp > this.curKey {
		this.curKey = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.record[this.curKey]
}

func (this *StockPrice) Maximum() int {
	return this.prices.Right().Key.(int)
}

func (this *StockPrice) Minimum() int {
	return this.prices.Left().Key.(int)
}

/**
 * 方法二
 * 哈希表 + 两个优先队列
 */
type StockPrice2 struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	maxTimestamp       int
}

func Constructor2() StockPrice2 {
	return StockPrice2{timePriceMap: map[int]int{}}
}

func (sp *StockPrice2) Update(timestamp, price int) {
	heap.Push(&sp.maxPrice, pair{-price, timestamp})
	heap.Push(&sp.minPrice, pair{price, timestamp})
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice2) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}
func (sp *StockPrice2) Maximum() int {
	for {
		if p := sp.maxPrice[0]; -p.price == sp.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&sp.maxPrice)
	}
}
func (sp *StockPrice2) Minimum() int {
	for {
		if p := sp.minPrice[0]; p.price == sp.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&sp.minPrice)
	}
}

type pair struct{ price, timestamp int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	cmd := []string{"StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"}
	val := [][]int{{}, {1, 10}, {2, 5}, {}, {}, {1, 3}, {}, {4, 2}, {}}
	//cmd := []string{"StockPrice", "update", "maximum", "current", "minimum", "maximum", "maximum", "maximum", "minimum", "minimum", "maximum", "update", "maximum", "minimum", "update", "maximum", "minimum", "current", "maximum", "update", "minimum", "maximum", "update", "maximum", "maximum"}
	//cmd := []string{"StockPrice", "update", "update", "update", "update", "update", "update", "update", "update", "update", "update", "update"}
	//val := [][]int{{}, {38, 2308}, {47, 7876}, {58, 1866}, {43, 121}, {40, 5339}, {32, 5339}, {43, 6414}, {49, 9369}, {36, 3192}, {48, 1006}, {53, 8013}}
	var stock StockPrice
	var stock2 StockPrice2
	for i := 0; i < len(cmd) && i < len(val); i++ {
		switch cmd[i] {
		case "StockPrice":
			stock = Constructor()
			stock2 = Constructor2()
			fmt.Print("null ")
		case "update":
			if len(val[i]) != 2 {
				fmt.Print("size error. i,val[i]:", i, val[i])
				return
			}
			stock.Update(val[i][0], val[i][1])
			stock2.Update(val[i][0], val[i][1])
			//fmt.Println("val:", val[i], " map: ", stock.record, " max,min,cur:", stock.Maximum(), stock.Minimum(), stock.Current())
		case "current":
			fmt.Print(stock.Current(), " ")
			fmt.Print(stock2.Current(), " ")
		case "maximum":
			fmt.Print(stock.Maximum(), " ")
			fmt.Print(stock2.Maximum(), " ")
		case "minimum":
			fmt.Print(stock.Minimum(), " ")
			fmt.Print(stock2.Minimum(), " ")
		}
	}
}
