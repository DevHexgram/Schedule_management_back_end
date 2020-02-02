package routers

import (
	"github.com/DevHexgram/Schedule_management_back_end/middleware"
	"github.com/gin-gonic/gin"
	"github.com/DevHexgram/Schedule_management_back_end/routers/api"
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
		image.GET("", s.getImage)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", middleware.RequestEntryDefault(api.Login))
		auth.POST("/register", middleware.RequestEntryDefault(api.Register))
	}

	userStatus := r.Group("/userStatus")
	userStatus.Use(middleware.JWT())
	{
		userStatus.POST("", middleware.RequestEntryWithStatus(s.modifyUserStatus)) //改
		userStatus.GET("", middleware.RequestEntryWithStatus(s.getUserStatus))     //查
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
		operaDaily.POST("/add", middleware.RequestEntryWithStatus(s.addDailyEvents))
		operaDaily.DELETE("", middleware.RequestEntryWithStatus(s.deleteDailyEvents))
		operaDaily.PUT("", middleware.RequestEntryWithStatus(s.modifyDailyEvents))
		//operaDaily.GET()
	}

	opera := r.Group("/opera")
	opera.Use(middleware.JWT())
	{
		opera.POST("/add", middleware.RequestEntryWithStatus(s.addAffair))  //增
		opera.DELETE("", middleware.RequestEntryWithStatus(s.deleteAffair)) //删
		opera.PUT("", middleware.RequestEntryWithStatus(s.modifyAffair))    //改
		//opera.GET("/find", s.findAffair) //查
	}

	//s.Router = r
	//err := s.Router.Run(s.Config.Web.Port)
	//DealError(err)
	return r
}



