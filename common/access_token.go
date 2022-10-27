package common

import (
	"fmt"
	"github.com/gmars/go-baidu-ai/util"
)

type Option interface {
	Apply(c *AccessToken) error
}

type AccessToken struct {
	appID     int64
	apiKey    string
	secretKey string
	cache     Cache
}

func NewAccessToken(opts ...Option) (*AccessToken, error) {
	accessToken := new(AccessToken)
	err := initAccessToken(accessToken, opts)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func initAccessToken(a *AccessToken, opts []Option) error {
	for _, opt := range opts {
		err := opt.Apply(a)
		if err != nil {
			return err
		}
	}

	//设置默认缓存
	if a.cache == nil {
		fmt.Println("本包默认缓存为文件缓存，生产环境中建议您设置更高性能的缓存。您只需要实现本包中的Cache接口")
		fc, err := util.NewFileCache("")
		if err != nil {
			return err
		}
		a.cache = fc
	}
	return nil
}
