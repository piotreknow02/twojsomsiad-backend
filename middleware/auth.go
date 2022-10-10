package middleware

import (
	"time"

	"twojsomsiad/config"
	"twojsomsiad/controller"
	"twojsomsiad/model"

	jwt "github.com/appleboy/gin-jwt/v2"
)

func Auth(api *controller.Controller) *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "twojsasiad",
		Key:             []byte(config.Conf.JwtSecret),
		Timeout:         time.Hour * time.Duration(config.Conf.JwtAccessExpireTime),
		MaxRefresh:      time.Hour * time.Duration(config.Conf.JwtRefreshExpireTime),
		IdentityKey:     "id",
		PayloadFunc:     PayloadFunc,
		IdentityHandler: api.IndentifyHandler,
		Authenticator:   api.Autheticator,
		Authorizator:    api.Authorizator,
		Unauthorized:    api.Unathorized,
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	if err != nil {
		panic(err)
	}

	err = authMiddleware.MiddlewareInit()
	if err != nil {
		panic(err)
	}

	return authMiddleware
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(model.User); ok {
		return jwt.MapClaims{
			"id":    v.ID,
			"email": v.Email,
		}
	}
	return jwt.MapClaims{}
}
