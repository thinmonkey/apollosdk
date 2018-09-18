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
## 初始化配置信息
启动配置文件在config.properties文件中。
```
{
    "appId":"app-capability",
    "cluster":"default",
    "metaServer":"http://10.160.2.153:8083",
    "refreshInterval":"300",
    "connectTimeout":"20",
    "onErrorRetryInterval":"1",
    "maxConfigCacheSize":52428800,
    "configCacheExpireTime":60,
    "longPollingInitialDelayInMills":"2"
}
```

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
