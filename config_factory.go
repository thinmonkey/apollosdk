package apollosdk

import (
	"fmt"
	"github.com/thinmonkey/apollosdk/core"
	"strings"
)

type ConfigFactory interface {
	CreateConfig(namespace string) core.Config
	CreateConfigFile(namespace string,configSourceType core.ConfigSourceType) core.ConfigFile
}

type DefaultConfigFactory struct {

}

func (*DefaultConfigFactory) CreateConfig(namespace string) core.Config {
	panic("implement me")
}

func (*DefaultConfigFactory) CreateConfigFile(namespace string, configSourceType core.ConfigSourceType) core.ConfigFile {
	panic("implement me")
}

func(*DefaultConfigFactory) determineFileFormat(namespaceName string) string  {
	lowerCase := strings.ToLower(namespaceName)
	for _,format := range core.ALL_FILEFORMAT{
		if strings.HasSuffix(lowerCase,fmt.Sprintf(".%s",format)){
			return format
		}
	}
	return core.PROPERTIES
}

func(*DefaultConfigFactory) trimNamespaceFormat(namespaceName string,format string) string{
	extension := fmt.Sprintf(".%s",format)
	if !strings.HasSuffix(strings.ToLower(namespaceName),extension){
		return namespaceName
	}
	return namespaceName[:len(namespaceName)-len(extension)]
}

