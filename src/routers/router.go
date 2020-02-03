package routers

import (
	"github.com/DevHexgram/Schedule_management_back_end/middleware"
	"github.com/DevHexgram/Schedule_management_back_end/routers/api"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	admin := r.Group("/@dmInSchM@n")
	admin.Use(middleware.JWT(), middleware.AuthorityControl(8))
	{
		admin.GET("/listInvitationCode", middleware.RequestEntryDefault(api.ListInvitationCode))
		admin.POST("/GenerateInvitationCode")
		admin.GET("/getPictureLink")
		admin.POST("/addPictureLink")
	}

	image := r.Group("/backgroundImage")
	//image.Use(JWT())
	{
		image.GET("", api.GetImage)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", middleware.RequestEntryDefault(api.Login))
		auth.POST("/register", middleware.RequestEntryDefault(api.Register))
	}

	userStatus := r.Group("/userSetting")
	userStatus.Use(middleware.JWT())
	{
		userStatus.POST("", middleware.RequestEntryWithStatus(api.ModifyUserSetting)) //改
		userStatus.GET("", middleware.RequestEntryWithStatus(api.GetUserSetting))     //查
	}

	all := r.Group("/all")
	all.Use(middleware.JWT())
	{
		all.GET("/affairs", middleware.RequestEntryWithStatus(api.GetAllAffairs))
		all.GET("/dailyAffairs", middleware.RequestEntryWithStatus(api.GetAllDailyAffairs))
	}

	operaDaily := r.Group("/operaDaily")
	operaDaily.Use(middleware.JWT())
	{
		operaDaily.POST("/add", middleware.RequestEntryWithStatus(api.AddDailyAffair))
		operaDaily.DELETE("", middleware.RequestEntryWithStatus(api.DeleteDailyAffair))
		operaDaily.PUT("", middleware.RequestEntryWithStatus(api.ModifyDailyAffair))
		//operaDaily.GET()
	}

	opera := r.Group("/opera")
	opera.Use(middleware.JWT())
	{
		opera.POST("/add", middleware.RequestEntryWithStatus(api.AddAffair))  //增
		opera.DELETE("", middleware.RequestEntryWithStatus(api.DeleteAffair)) //删
		opera.PUT("", middleware.RequestEntryWithStatus(api.ModifyAffair))    //改
		//opera.GET("/find", s.findAffair) //查
	}

	//s.Router = r
	//err := s.Router.Run(s.Config.Web.Port)
	//DealError(err)
	return r
}
