package router

import (
	"crypto/rsa"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-architecture/config"
	"go-architecture/i18n"
	util "go-architecture/pkg"
	"reflect"
	"time"
)

const JWT_IDENTITY_KEY = "id"

type Router struct {
	Engine             *gin.Engine
	AuthMiddleware     *jwt.GinJWTMiddleware
	I18n               *i18n.I18n
	PrivateKey         *rsa.PrivateKey
	LongRefreshExpTime time.Duration
}

func NewRouterWithAuthMw(c config.Config, i18n *i18n.I18n) (*Router, error) {
	return NewRouter(c, i18n, jwt.GinJWTMiddleware{})
}

func NewRouter(c config.Config, i18n *i18n.I18n, jwtMdw jwt.GinJWTMiddleware) (*Router, error) {
	e := gin.New()

	e.RedirectTrailingSlash = true
	e.RedirectFixedPath = true

	//e.Use(gin.Logger())
	e.Use(logger.SetLogger())

	// CORS
	corsMiddleware, err := initCorsMiddleware(c)
	if err != nil {
		return nil, err
	}
	e.Use(corsMiddleware)

	authMiddleware, err := initJWTMiddleware(c, jwtMdw)
	if err != nil {
		log.Fatal().Err(err).Msg("JWT Error:" + err.Error())
		return nil, err
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal().Err(errInit).Msg("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		return nil, errInit
	}

	refreshExpTime, err := time.ParseDuration(c.Jwt.RefreshExpTime)
	if err != nil {
		return nil, err
	}
	return &Router{
		Engine:             e,
		AuthMiddleware:     authMiddleware,
		I18n:               i18n,
		LongRefreshExpTime: refreshExpTime,
		PrivateKey:         privateKey(authMiddleware),
	}, nil
}

func initJWTMiddleware(c config.Config, jwtMdw jwt.GinJWTMiddleware) (*jwt.GinJWTMiddleware, error) {
	expiredTime, err := time.ParseDuration(c.Jwt.ExpiredTime)
	if err != nil {
		return nil, err
	}
	refreshExpTime, err := time.ParseDuration(c.Jwt.RefreshExpTime)
	if err != nil {
		return nil, err
	}
	// the jwt router
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: c.Jwt.SigningAlg,
		Key:              []byte(c.Jwt.Secret),
		Timeout:          expiredTime,
		MaxRefresh:       refreshExpTime,
		IdentityKey:      JWT_IDENTITY_KEY,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	return authMiddleware, nil
}

func initCorsMiddleware(c config.Config) (gin.HandlerFunc, error) {
	maxAge, err := time.ParseDuration(c.CORS.MaxAge)
	if err != nil {
		return nil, nil
	}

	return cors.New(cors.Config{
		AllowOrigins:     c.CORS.AllowOrigins,
		AllowMethods:     c.CORS.AllowMethods,
		AllowHeaders:     c.CORS.AllowHeaders,
		ExposeHeaders:    c.CORS.ExposeHeaders,
		AllowCredentials: c.CORS.AllowCredentials,
		MaxAge:           maxAge,
	}), nil
}

func privateKey(privKeyFile *jwt.GinJWTMiddleware) *rsa.PrivateKey {
	value := util.GetUnexportedField(reflect.ValueOf(privKeyFile).Elem().FieldByName("privKey"))
	if value == nil {
		return nil
	}
	return value.(*rsa.PrivateKey)
}

func (r *Router) InitSwagger(c config.Config) {
	url := ginSwagger.URL(fmt.Sprintf("%v/swagger/doc.json", c.Swagger.Url))
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}


