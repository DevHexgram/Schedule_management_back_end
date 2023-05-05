module github.com/DevHexgram/Schedule_management_back_end

go 1.13

replace (
	github.com/DevHexgram/Schedule_management_back_end/middleware => ./middleware
	github.com/DevHexgram/Schedule_management_back_end/models => ./models
	github.com/DevHexgram/Schedule_management_back_end/pkg/e => ./pkg/e
	github.com/DevHexgram/Schedule_management_back_end/pkg/setting => ./pkg/setting
	github.com/DevHexgram/Schedule_management_back_end/routers => ./routers
	github.com/DevHexgram/Schedule_management_back_end/routers/api => ./routers/api
	github.com/DevHexgram/Schedule_management_back_end/util => ./pkg/util
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.9.0
	github.com/jinzhu/gorm v1.9.12
	github.com/ugorji/go v1.1.7 // indirect
)
