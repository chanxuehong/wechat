copy 了 go1.3.1 的 encoding/json 包, 只是修改了几个函数, 使得不转义 '<', '>', '&'

修改了下面的函数:

```Go
encode.go:
    func HTMLEscape(dst *bytes.Buffer, src []byte)
    func (e *encodeState) string(s string) (int, error)
    func (e *encodeState) stringBytes(s []byte) (int, error)

indent.go:
    func compact(dst *bytes.Buffer, src []byte, escape bool) error
```