package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type allCache struct {
	user     *cache.Cache
	password *cache.Cache
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

func NewCache() *allCache {
	Cache := cache.New(defaultExpiration, purgeTime)
	return &allCache{
		user:     Cache,
		password: Cache,
	}
}

// TODO: Hacer funcionalidades de actualizar y leer caché con el usuario y contraseña presentes
