package common

type AppConfig struct {
	appID     int64
	apiKey    string
	secretKey string
}

func WithAppConfig(appId int64, apiKey, secretKey string) *AppConfig {
	return &AppConfig{
		appID:     appId,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (a *AppConfig) Apply(c *AccessToken) error {
	c.appID = a.appID
	c.apiKey = a.apiKey
	c.secretKey = a.secretKey
	return nil
}

type CacheConfig struct {
	cache Cache
}

func WithCacheConfig(cache Cache) *CacheConfig {
	if cache != nil {
		return &CacheConfig{cache: cache}
	}
	return nil
}

func (a *CacheConfig) Apply(c *AccessToken) error {
	if a.cache != nil {
		c.cache = a.cache
	}
	return nil
}
