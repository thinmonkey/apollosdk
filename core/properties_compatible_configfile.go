/**
 * @author Created by zhanghao on 2019-07-30 14:54, [haozhang@ecarx.com.cn]
 */
package core

type PropertiesCompatibleConfigFile interface {
	ConfigFile
	AsProperties() Properties
}
