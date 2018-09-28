Apollosdk - Go Client for Apollo
================================


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
* 支持多namespace,多集群,多数据中心配置
* 支持监听配置变更实时变化
* 支持容器化（docker）多环境配置

# 使用规则
### 根据apollo官方client使用规则，apollosdk启动时需要配置相应的apollo相关参数，和一些运行时配置参数（可选配）。
### apollo启动参数有多种获取方式：
1. 配置文件config.properties（文件位置可以动态设置）。
2. os系统环境变量
3. apollosdk启动时动态传递apollo配置参数。

#### 配置文件config.properties获取启动参数
配置文件默认寻找当前项目的根目录下的config.json文件，配置内容格式为json，可配置内容如下：
```
{
    "appId":"app-capability",
    "cluster":"default",
    "metaServer":"http://10.160.1.153:8083",
    "httpRefreshInterval":"300s",
    "httpTimeout":"20s",
    "onErrorRetryInterval":"1s",
    "maxConfigCacheSize":52428800,
    "configCacheExpireTime":60,
    "longPollingRefreshInterval":"2s",
    "longPollingTimeout":"60s"
}
```
#### os系统环境变量获取配置，只支持必要apollo参数，其他运行时参数为默认值（方便支持docker容器注入不同环境参数）
- appId获取：os.Getenv("apollo.appId")
- cluster获取：os.Getenv("apollo.cluster")
- metaServer获取：os.Getenv("DOCKER_SERVER")
- dataCenter获取：os.Getenv("apollo.dataCenter")
#### 启动时动态注入apollo参数和配置文件地址
```
使用前调用方法：
apollosdk.Start("/opt/config.json","appId","cluster","metaServer","dataCenter")
```
### 名词解释：
- appId对应于apollo appId,当前应用的appId
- cluster对应于apollo cluster，当前应用所在的集群
- metaServer对应于apollo配置地址，当前应用连接的apollo环境配置地址
- dataCenter对应于apollo dataCenter，当前应用连接的数据中心。
- httpRefreshInterval获取配置接口的调用间隔时间，默认五分钟
- httpTimeout获取配置接口的超时时间，默认20s
- onErrorRetryInterval获取配置接口错误之后延迟尝试间隔，默认1s
- maxConfigCacheSize缓存最大的存储大小，默认50M
- configCacheExpireTime缓存失效时间，默认60秒
- longPollingRefreshInterval通知接口的刷新间隔时间，默认2s
- longPollingTimeout通知接口的keep-alive时间，默认60s

### 获取apollo配置的优先级：
第3种方式 > 第2种方式 > 第1种方式,优先选择动态传递的参数，其次是系统环境变量，最后是配置文件。

# Use
### 初始化（可选）,不设置则按照默认规则
```
apollosdk.Start("/opt/config.json","appId","cluster","metaServer","dataCenter")
```

### 默认的namespace配置获取
```
config := apollosdk.GetAppConfig()
config.GetStringProperty("mats", "")

```
### 自定义的namespace配置获取
```
config := apollosdk.GetConfig(""app.tc.mat.disable"")
config.GetStringProperty("mats", "")
```
### 配置改变实时监听(支持多种监听回调方式)
```
configNew := apollosdk.GetConfig(""app.tc.mat.disable"")
//方式一：定义变量来监听
var varFunc OnChangeFunc =  func (changeEvent ConfigChangeEvent)  {
		fmt.Println("variable onChange",changeEvent)
}
configNew.AddChangeListenerFunc(varFunc)


//方式二：定义普通函数来监听
func onTestFunc(changeEvent ConfigChangeEvent) {
	fmt.Println("func onChange",changeEvent)
}
configNew.AddChangeListenerFunc(onTestFunc)

//方式三：定义结构实现接口来监听
type SomeThing string

func (s SomeThing)OnChange (changeEvent ConfigChangeEvent)  {
	fmt.Println("struct onChange",changeEvent)
}
var s SomeThing = "s"
configNew.AddChangeListener(s)

//移除监听器
configNew.RemoveChangeListener(s)
```
