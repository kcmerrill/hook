package hook

import (
	"reflect"
	"sync"
)

// Contains our triggers AND priority mapping
var triggers map[string]map[int][]interface{}

// Trigger lock, for those pesky race conditions
var tl *sync.Mutex

// Filter alias to trigger, for optics and code readability
var Filter func(string, ...interface{})

// Register a trigger with priority 0
func Register(trigger string, args ...interface{}) {
	tl.Lock()
	defer tl.Unlock()

	var f interface{}
	var priority int
	if len(args) >= 2 {
		priority, _ = args[0].(int)
		f = args[1]
	} else {
		priority = 0
		f = args[0]
	}

	_, exists := triggers[trigger]
	if !exists {
		triggers[trigger] = make(map[int][]interface{})
		for p := 0; p <= 100; p++ {
			triggers[trigger][p] = make([]interface{}, 0)
		}
	}
	triggers[trigger][priority] = append(triggers[trigger][priority], f)
}

// Trigger a, uh, trigger
func Trigger(trigger string, args ...interface{}) {
	tl.Lock()
	priorities, exists := triggers[trigger]
	tl.Unlock()

	if exists {
		for p := 0; p <= 100; p++ {
			for _, f := range priorities[p] {
				params := make([]reflect.Value, len(args))
				for idx := range args {
					params[idx] = reflect.ValueOf(args[idx])
				}
				reflect.ValueOf(f).Call(params)
			}
		}
	}
}

// Giddy Up
func init() {
	Filter = Trigger
	triggers = make(map[string]map[int][]interface{})
	tl = &sync.Mutex{}
}
