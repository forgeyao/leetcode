// https://leetcode.cn/problems/lru-cache/
package main

import (
	"fmt"
	"strconv"
)

/**
 * 缓存，直接用O(1) 的 map
 * 要保证淘汰机制O(1)效率，使用双向链表
 *
 * 参考答案更简洁
 * 一是拆分了很多node 相关子函数；
 * 二是多加了一个节点来简化头节点为空的问题（涉及链表操作都可以这么玩）
 */
type LinkedList struct {
	Key  int
	Val  int
	Prev *LinkedList
	Next *LinkedList
}

type LRUCache struct {
	Kv       map[int]*LinkedList
	Root     *LinkedList
	Tail     *LinkedList
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Kv:       map[int]*LinkedList{},
		Root:     nil,
		Tail:     nil,
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Kv[key]; ok {
		if node != this.Root { // 头节点不用更新
			node.Prev.Next = node.Next
			if node == this.Tail { // 尾节点特殊处理
				this.Tail = node.Prev
			} else {
				node.Next.Prev = node.Prev
			}
			node.Prev = nil
			node.Next = this.Root
			this.Root.Prev = node
			this.Root = node
		}
		//fmt.Println("get:", key, node.Val)
		//this.Log()
		return node.Val
	}
	//fmt.Println("get:", key, -1)
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	val := this.Get(key)
	if val != -1 {
		this.Root.Val = value
	} else {
		var node *LinkedList
		if len(this.Kv) >= this.Capacity {
			//fmt.Println("put new full:", key, value)
			node = this.Tail
			delete(this.Kv, node.Key)
			if node == this.Root { // 只有一个元素
				this.Root.Key = key
				this.Root.Val = value
				this.Kv[key] = this.Root
				return
			} else {
				this.Tail = node.Prev
				this.Tail.Next = nil
			}
		} else {
			//fmt.Println("put new:", key, value)
			node = new(LinkedList)
		}
		node.Key = key
		node.Val = value
		node.Prev = nil
		node.Next = this.Root
		if this.Root != nil {
			this.Root.Prev = node
		} else {
			this.Tail = node
		}
		this.Root = node
		this.Kv[key] = this.Root
	}
	//this.Log()
}

func (this *LRUCache) Log() {
	fmt.Print("list:")
	for node := this.Root; node != nil; node = node.Next {
		fmt.Print("<", node.Key, ",", node.Val, "> ")
	}
	fmt.Println()
}

func main() {
	ss := [][]string{
		{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
		{"LRUCache", "put", "put", "get", "put", "put", "get"},
		{"LRUCache", "get", "get", "put", "get", "put", "put", "put", "put", "get", "put"},
		{"LRUCache", "put", "get", "put", "get", "get"},
	}
	nums := [][][]int{
		{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
		{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}},
		{{1}, {6}, {8}, {12, 1}, {2}, {15, 11}, {5, 2}, {1, 15}, {4, 2}, {5}, {15, 15}},
		{{1}, {2, 1}, {2}, {3, 2}, {2}, {3}},
	}
	answers := [][]string{
		{"null", "null", "null", "1", "null", "-1", "null", "-1", "3", "4"},
		{"null", "null", "null", "2", "null", "null", "-1"},
		{"null", "-1", "-1", "null", "-1", "null", "null", "null", "null", "-1", "null"},
		{"null", "null", "1", "null", "-1", "2"},
	}

	for j := 0; j < len(ss) && j < len(nums) && j < len(answers); j++ {
		var lRUCache LRUCache
		rets := []string{}
		for i := 0; i < len(ss[j]) && i < len(nums[j]); i++ {
			if ss[j][i] == "LRUCache" {
				lRUCache = Constructor(nums[j][i][0])
				rets = append(rets, "null")
			} else if ss[j][i] == "put" {
				lRUCache.Put(nums[j][i][0], nums[j][i][1])
				rets = append(rets, "null")
			} else if ss[j][i] == "get" {
				rets = append(rets, strconv.Itoa(lRUCache.Get(nums[j][i][0])))
			}
		}
		if len(answers[j]) != len(rets) {
			fmt.Println("error: len is different")
		} else {
			for i := 0; i < len(answers[j]); i++ {
				if answers[j][i] != rets[i] {
					fmt.Println("error: ")
					fmt.Println(rets)
					fmt.Println(answers[j])
					return
				}
			}
			fmt.Println("right!")
		}
	}
}
