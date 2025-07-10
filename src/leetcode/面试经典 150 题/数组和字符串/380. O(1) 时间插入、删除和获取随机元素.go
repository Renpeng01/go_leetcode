package main

import "math/rand"

type RandomizedSet struct {
	dataMap map[int]int
	data    []int
}

func Constructor() RandomizedSet {
	e := RandomizedSet{}
	e.dataMap = make(map[int]int, 128)
	e.data = make([]int, 0, 128)
	return e
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		return false
	}
	this.dataMap[val] = len(this.data)
	this.data = append(this.data, val)
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		last := len(this.data) - 1
		this.dataMap[this.data[last]] = this.dataMap[val]
		this.data[this.dataMap[val]] = this.data[last]
		this.data = this.data[:last]
		delete(this.dataMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}
