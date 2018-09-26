package apollosdk

import (
	"testing"
	. "github.com/zhhao226/apollosdk/core"
	"fmt"
)

func TestGetConfig(t *testing.T) {
	config := GetAppConfig()

	chanEvent := (*config).GetChangeKeyNotify()


	configNew := GetConfig("app.tc.mat.disable")

	t.Log((*configNew).GetStringProperty("mats", ""))

	configEvent := <-chanEvent
	t.Log(configEvent)


	//定义变化
	var varFunc OnChangeFunc =  func (changeEvent ConfigChangeEvent)  {
		fmt.Println(changeEvent)
	}

	//添加listener的变量
	addListener(varFunc)

	//添加func
	addListenerFunc(onChangeNofity)

	//添加interface的实现类型
	var  s Something= "22"
	addListener(s)
}


/***
 添加Listener的方式
 */
func addListener(changeListener ConfigChangeListener)  {
	if changeListener!=nil{
		fmt.Println(changeListener)
	}
}

/***
 添加Func的方式
 */
func addListenerFunc(f OnChangeFunc)  {
	if f!=nil{
		addListener(f)
	}
}


/***
 有事件来了,函数传进去
 */
func onChangeNofity(changeEvent ConfigChangeEvent)  {
	fmt.Println(changeEvent)
}


//某个实现了接口的类型
type Something string

func (s Something)OnChange(changeEvent ConfigChangeEvent)()  {

}


func TestGetAppConfig(t *testing.T) {

}
