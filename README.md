> 如果在线上使用gid或者gls， 希望你已经清楚使用它们的后果...

# 背景    

- 项目已经微服务改造，服务之间通讯方式:HTTP, gRPC,
- 跨服务调用后不方便定位问题，需要实现trace功能

这样的话就需要在服务内的各func之间以及服务间传递trace\_id，没错，首先想到的就是Go推荐的context， 也就是在每个方法里加一个context参数。这时候要做的事情包括：  
1、修改所有旧代码，都加上context参数   
2、如果想把trace\_id写入日志, 在所有打印日志的地方，都需要修改   
想想这工作量，想想这满屏的context，真是受不了。。。。   
于是想到了gls，将trace\_id存入gls， 就不需要满屏携带，只需要根据gid获取即可。这样的话我们只需要修改项目的基础组件，业务方就可以无痛接入，想想还是蛮爽的, 所以我使用gls的目的是为了解决trace\_id的问题，当然你也可以有其他用法，前提是你已经清楚使用它们的后果...

# 实现
1、基于sync.Map实现，map做了sharding(shard-16)   
2、由于go并没有tid， 所以使用gid代替


# 使用
```shell
go get github.com/lanfang/gls

import "github.com/lanfang/gls"

//set 
gls.Set(kev, value)

//get 
value := gls.Get(key)

//start goroutine
用gls.RunGo(fun) 替代 go func()
```