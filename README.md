Agollo - Go Client for Apollo
================


方便Golang接入配置中心框架 [Apollo](https://github.com/ctripcorp/apollo) 所开发的Golang版本客户端。

Installation
------------

如果还没有安装Go开发环境，请参考以下文档[Getting Started](http://golang.org/doc/install.html) ，安装完成后，请执行以下命令：

``` shell
go get -u github.com/cihub/seelog
go get -u github.com/coocood/freecache
go get -u github.com/zhhao226/apollosdk
```


*请注意*: 最好使用Go 1.8进行开发

# Features
* 实时同步配置
* 灰度配置
* 客户端容灾
* 支持多namespace,多集群
* 支持配置变更实时变化

# Use
## 初始化配置信息（必须）
启动配置文件在config.properties文件中，可配置内容如下：
```
{
    "appId":"app-capability",//对应apollo里的应用配置appId
    "cluster":"default",//对应apollo里的集群配置
    "metaServer":"http://10.160.1.128:8083",//对应apollo里的负载ip地址
    "refreshInterval":"300",//对应的轮询刷新间隔时间，单位是秒
    "connectTimeout":"20",//http连接超时时间，单位是秒
    "onErrorRetryInterval":"1",//某次http请求错误重试次数
    "maxConfigCacheSize":52428800,//内存缓存的最大内存，单位字节 50 * 1024 *1024
    "configCacheExpireTime":60,//缓存配置失效时间，单位分钟
    "longPollingInitialDelayInMills":"2"//启动通知长链接的延迟时间，单位秒。
}
```
1. appId可以多个地方获取，获取的优先级：
- 系统环境变量配置了appId优先获取环境变量配置：os.Getenv("apollo.appId")
- 系统环境变量没有配置则从config.properties配置文件中加载
- 否则默认返回"application"
2. cluster可以多个地方获取，获取的优先级：
- 系统环境变量配置了cluster优先获取环境变量配置：os.Getenv("apollo.cluster")
- 系统环境变量没有配置则从config.properties配置文件中加载
- 否则默认返回“default”
3. metaServer可以从多个地方获取，获取的优先级：
- 系统环境变量配置了metaServer优先获取环境变量配置：os.Getenv("DOCKER_SERVER")
- 系统环境变量没有配置则从config.properties配置文件中加载

## 默认的namespace获取
```
config := apollosdk.GetAppConfig()
(*config).GetStringProperty("mats", "")
go func(){
  chanEvent := (*config).GetChangeKeyNotify()
  configEvent := <-chanEvent
}

```
## 自定义的namespace获取
```
config := apollosdk.GetConfig(""app.tc.mat.disable"")
(*config).GetStringProperty("mats", "")
go func(){
  chanEvent := (*config).GetChangeKeyNotify()
  configEvent := <-chanEvent
}
```
