package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// you can define a named type for empty struct
// No fields
// you can define methods for this type
type Empty struct{}

func main() {
	// anonymous struct has zero size
	// unsafe.Sizeof only shows the static representation of the value ifself
	fmt.Println(unsafe.Sizeof(struct{}{}))
	a := Empty{}
	b := Empty{}
	// Comparable and all equal
	fmt.Println(a == b)

	var value struct{}
	// equivalent to value := struct{}{}
	var pointer *struct{}
	var iface any = struct{}{}
	var slice []struct{}
	var channel chan struct{}
	var mapping map[string]struct{}

	// 0
	fmt.Println(unsafe.Sizeof(value))
	// 8
	fmt.Println(unsafe.Sizeof(pointer))
	// 16 type + data
	fmt.Println(unsafe.Sizeof(iface))
	// 24 Data pointer + Len + Cap
	fmt.Println(unsafe.Sizeof(slice))
	// 8 *runtime.hchan
	fmt.Println(unsafe.Sizeof(channel))
	// *runtime.hmap
	fmt.Println(unsafe.Sizeof(mapping))

	c := struct{}{}
	d := struct{}{}

	fmt.Printf("the address of c: %p\n", &c)
	fmt.Printf("the address of d: %p\n", &d)

	// the array in go is a value type, which means it's the array itself, not a pointer to the head of the array
	//
	var values [1_000_000]struct{}
	// 1000000 * 0 = 0
	fmt.Println(unsafe.Sizeof(values))
	fmt.Println(len(values))

	values2 := make([]struct{}, 1_000_000)
	fmt.Println(unsafe.Sizeof(values2))
	fmt.Println(len(values2))

	// why do not use map[string]bool?
	languages := map[string]struct{}{
		"Go":   {},
		"Java": {},
	}

	languages["Rust"] = struct{}{}

	if _, exists := languages["Go"]; exists {
		fmt.Println("Go exists")
	}

	delete(languages, "Java")
	// will only return the value(zero-value if it doesn't exist)
	v := languages["Java"]
	// comma-ok idiom will return the value and the bool result to show if the key exists
	v, ok := languages["Java"]
	fmt.Println(v, ok)

	// when the channel only respond to pass signal without passing value
	done := make(chan struct{}, 0)
	go func() {
		done <- struct{}{}
	}()
	<-done
	// broadcast the signal
	for i := 0; i < 3; i++ {
		go func() {
			<-done
			fmt.Println("stopped")
		}()
	}
	// when the channel is closed, all the goroutine waiting for the data from this channel will not be blocked
	close(done)

    var encoder Encoder = JsonEncoder{}
    if bs, err := encoder.Encode("go"); err == nil {
        fmt.Println(bs)
    }
}

// you can even capsule the Set like this
type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) contains(value T) bool {
	_, exists := s[value]
	return exists
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// When you only need something to provide a fixed behaviour without the configs, you can use empty struct to implement the interface
type Encoder interface {
	Encode(value any) ([]byte, error)
}

type JsonEncoder struct{}

// if the method needs to access the receiver, you have to give a name to it. but here no need
func (JsonEncoder) Encode(value any) ([]byte, error) {
    return json.Marshal(value)
}

// the empty struct may affect the size of the outer struct
type A struct {
    Empty struct{}
    Value int64
}

type B struct {
    Value int64
    Empty struct{}
}
