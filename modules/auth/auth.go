package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

//FlnowAuthMiddware to auth
func FlnowAuthMiddware() *jwt.GinJWTMiddleware {
	authMiddware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "FLNOW",
		Key:         []byte("FLNOW SECRET"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		// 增加默认的payload信息给前端
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					"id":    v.UserName,
					"email": v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["id"].(string),
			}
		},
		// 必要项, 这个函数用来判断 User 信息是否合法，如果合法就反馈 true，否则就是 false, 认证的逻辑就在这里
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals User
			if err := c.Bind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.UserName
			password := loginVals.Password
			if (userID == "admin" && password == "123123") || (userID == "test" && password == "test") {
				return &User{
					UserName: userID,
					Email:    "test@flnow.com",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		// 可选项，用来在 Authenticator 认证成功的基础上进一步的检验用户是否有权限，默认为 success
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		// 如果认证不成功的的处理函数
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TimeFunc: time.Now,
	})
	if err != nil {
		panic(err)
	}
	return authMiddware
}
