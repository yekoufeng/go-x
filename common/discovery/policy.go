package discovery

import "github.com/fananchong/gomap"

type Policy int

const (
	_          Policy = iota // 0
	Ordered                  // 1
	Random                   // 2
	RoundRobin               // 3
)

// ServersPolicyOrdered
type ServersPolicyOrdered struct {
	gomap.OrderedMap
}

func NewServersPolicyOrdered() *Servers {
	m := &ServersPolicyOrdered{}
	m.OrderedMap = *gomap.NewOrderedMap(less)
	return &Servers{
		ss:      make(map[int]IMap),
		creator: func() IMap { return m },
	}
}

func less(v1, v2 interface{}) bool {
	return v1.(*ServerInfo).Ordered >= v2.(*ServerInfo).Ordered
}

func (this *ServersPolicyOrdered) GetOne() (key, val interface{}, ok bool) {
	return this.Back()
}

// ServersPolicyRandom
type ServersPolicyRandom struct {
	gomap.RandomMap
}

func NewServersPolicyRandom() *Servers {
	return &Servers{
		ss:      make(map[int]IMap),
		creator: func() IMap { return &ServersPolicyRandom{} },
	}
}

func (this *ServersPolicyRandom) GetOne() (key, val interface{}, ok bool) {
	return this.Random()
}

// ServersPolicyRoundRobin
type ServersPolicyRoundRobin struct {
	gomap.RoundRobinMap
}

func NewServersPolicyRoundRobin() *Servers {
	return &Servers{
		ss:      make(map[int]IMap),
		creator: func() IMap { return &ServersPolicyRoundRobin{} },
	}
}

func (this *ServersPolicyRoundRobin) GetOne() (key, val interface{}, ok bool) {
	return this.RoundRobin()
}
