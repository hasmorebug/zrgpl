package reflection

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"
)

func RfExample() {
	rf3()
	//rf2()
	//rf1()
}

///////////////////////////
func rf3() {
	var x int64 = 1
	var d = 1 * time.Nanosecond
	var i interface{}

	fmt.Println(formatAny(x))
	fmt.Println(formatAny(d))
	fmt.Println(formatAny([]int64{}))
	fmt.Println(formatAny([]time.Duration{d}))
	fmt.Println(formatAny(&i))
}

func formatAny(i interface{}) string {
	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array reflect.Struct reflect.Interface
		return v.Type().String() + "value"
	}
}

///////////////////////////
func rf2() {
	// TypeOf
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	fmt.Printf("%T\n", 3)

	// ValueOf
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	s := reflect.ValueOf("Hello World")
	fmt.Println(s)
	fmt.Printf("%v\n", s)
	fmt.Println(s.String())

	// Value.Type()
	vt := v.Type()
	fmt.Println(vt.String())

	// Value.Interface()
	vi := v.Interface()
	i := vi.(int)
	fmt.Println(i)
}

///////////////////////////
type zrStr struct {
}

func (s zrStr) String() string {
	return "Hello World!"
}

type zrs struct {
}

func rf1() {
	s := "Hello World!!"
	i := 1024
	b := true

	fmt.Println(Sprint(zrStr{}))
	fmt.Println(Sprint(s))
	fmt.Println(Sprint(i))
	fmt.Println(Sprint(b))
	fmt.Println(Sprint(zrs{}))
}

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		} else {
			return "false"
		}
	default:
		return "unknown type"
	}
}
