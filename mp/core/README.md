## 微信公众号 SDK 核心 package
微信公众号的处理逻辑都在这个 package 里面, 其他的模块都是在这个 package 基础上的再封装!

### 回调请求处理
一个回调地址(多个公众号可以共用一个回调地址)的 http 请求对应了一个 http handler(http.Handler, gin.HandlerFunc…), 
这个 http handler 里面的主要逻辑是调用对应公众号的 core.Server 的 ServeHTTP 方法来处理回调请求, 
core.Server.ServeHTTP 做签名的验证和消息解密, 然后调用 core.Server 的 core.Handler 属性的 ServeMsg 方法来处理消息(事件).  
![回调请求处理逻辑图](https://github.com/chanxuehong/wechat/blob/v2/mp/core/callback20160118.png)
