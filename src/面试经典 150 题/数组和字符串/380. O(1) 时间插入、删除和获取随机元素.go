package main

type RandomizedSet struct {
	dataMap map[int]struct{}
	data    []int
}

func Constructor() RandomizedSet {
	e := RandomizedSet{}
	e.dataMap = make(map[int]struct{}, 128)
	e.data = make([]int, 0, 128)
	return e
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		return false
	}
	this.dataMap[val] = struct{}{}
	this.data = append(this.data, val)
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		delete(this.dataMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {

}
