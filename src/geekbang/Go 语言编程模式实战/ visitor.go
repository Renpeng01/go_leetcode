package visitor

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// 将算法与操作对象分离的方法。
// 能够在不修改结构的情况下向现有对象结构添加新操作

type Visitor func(shape Shape)

type Shape interface {
	accetp(Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accetp(v Visitor) {
	v(c)
}

type Rectangle struct {
	Witch, Height int
}

func (r Rectangle) accetp(v Visitor) {
	v(r)
}

func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}

	shapes := []Shape{c, r}
	for _, s := range shapes {
		s.accetp(JsonVisitor)
		r.accetp(XmlVisitor)
	}
}
