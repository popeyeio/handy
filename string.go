package handy

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/luci/go-render/render"
)

const (
	StrEmpty     = ""
	StrColon     = ":"
	StrLine      = "\n"
	StrAnd       = "&"
	StrQuestion  = "?"
	StrEqual     = "="
	StrSpace     = " "
	StrSlash     = "\\"
	StrBackslash = "/"
	StrHyphen    = "-"
	StrTilde     = "~"
	StrDot       = "."
)

func IsEmptyStr(s string) bool {
	return s == StrEmpty
}

var internerPool = sync.Pool{
	New: func() interface{} {
		return make(map[string]string)
	},
}

func InternString(s string) string {
	interner := internerPool.Get().(map[string]string)
	if val, exists := interner[s]; exists {
		internerPool.Put(interner)
		return val
	}

	interner[s] = s
	internerPool.Put(interner)
	return s
}

func InternBytes(b []byte) string {
	interner := internerPool.Get().(map[string]string)
	// The compiler recognizes m[string(byteSlice)] as a special
	// case, so a copy of a's bytes into a new string does not
	// happen in this map lookup.
	if val, exists := interner[string(b)]; exists {
		internerPool.Put(interner)
		return val
	}

	s := string(b)
	interner[s] = s
	internerPool.Put(interner)
	return s
}

func Bytes2Str(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := &reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(sh))
}

func Str2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := &reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bh))
}

func Stringify(obj interface{}) string {
	obj = dereference(obj)

	switch v := obj.(type) {
	case nil:
		return StrEmpty
	case bool:
		return strconv.FormatBool(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case int:
		strconv.Itoa(v)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case []byte:
		return string(v)
	case time.Time:
		return v.Format(time.RFC3339)
	case fmt.Stringer:
		return v.String()
	case error:
		return v.Error()
	}
	return render.Render(obj)
}

// dereference refers to html/template/content.go.
func dereference(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}

	rv := reflect.ValueOf(obj)
	et := reflect.TypeOf((*error)(nil)).Elem()
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	for !rv.Type().Implements(et) && !rv.Type().Implements(st) && rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	return rv.Interface()

}
