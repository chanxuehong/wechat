encoding/json 默认会对 html 标记 <, >, & 做转换, 但是狗日的腾讯不识别, 所以只能 hack 标准的
json 库, 去掉这三个字符的转换(从 go1.4 encoding/json fork 而来).

修改的文件有:

encode.go:
注释掉 791-794行, 增加 795-798 行
注释掉 871-874行, 增加 875-878 行

indent.go:
注释掉 21-29 行
