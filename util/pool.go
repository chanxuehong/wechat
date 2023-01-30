package util

import (
	"net/url"
	"strings"
	"sync"
)

var stringsBuilder = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func GetStringsBuilder() *strings.Builder {
	return stringsBuilder.Get().(*strings.Builder)
}

func PutStringsBuilder(b *strings.Builder) {
	b.Reset()
	stringsBuilder.Put(b)
}

func StringsJoin(strs ...string) string {
	var n int
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}
	if n <= 0 {
		return ""
	}
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	builder.Grow(n)
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

var urlValuesPool = sync.Pool{
	New: func() any {
		return make(url.Values)
	},
}

func GetUrlValues() url.Values {
	return urlValuesPool.Get().(url.Values)
}

func PutUrlValues(values url.Values) {
	for k := range values {
		values.Del(k)
	}
	urlValuesPool.Put(values)
}
