package gls

import (
	"sync"
)

//this package is only used for tracing, don't use it for others
const shards = 16
const SpanKey = "spanKey"

var (
	//like map[int64]map[interface{}]interface{}
	globalMap []*sync.Map
)

func init() {
	globalMap = make([]*sync.Map, 16)
	for i := 0; i < 16; i++ {
		globalMap[i] = &sync.Map{}
	}
}

func getShard() int {
	return int(GoId() % shards)
}
func Set(key, val interface{}) {
	mp := globalMap[getShard()]
	v, _ := mp.LoadOrStore(GoId(), &sync.Map{})
	if imp, ok := v.(*sync.Map); ok {
		imp.Store(key, val)
	}
}

func Get(key interface{}) interface{} {
	mp := globalMap[getShard()]
	var val interface{}
	if tmp, ok := mp.Load(GoId()); ok {
		if dp, ok := tmp.(*sync.Map); ok {
			val, _ = dp.Load(key)
		}
	}
	return val
}

//for gls all
//rember clean gls ,this is important, maybe we will hack runtime.goexit()
func Clear() {
	mp := globalMap[getShard()]
	mp.Delete(GoId())
}

//get gls v is sync.Map
func GetGls() interface{} {
	mp := globalMap[getShard()]
	v, _ := mp.Load(GoId())
	return v
}

//set gls v is sync.Map
func SetGls(v interface{}) {
	mp := globalMap[getShard()]
	mp.Store(GoId(), v)
}

func RunGo(fn func()) {
	GoWithGls(GetGls(), fn)
}

func GoWithGls(glsV interface{}, fn func()) {
	go func() {
		SetGls(glsV)
		defer Clear()
		fn()
	}()
}
