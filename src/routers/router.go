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

	userSetting := r.Group("/userSetting")
	userSetting.Use(middleware.JWT())
	{
		userSetting.POST("", middleware.RequestEntryWithStatus(api.ModifyUserSetting)) //改
		userSetting.GET("", middleware.RequestEntryWithStatus(api.GetUserSetting))     //查
	}

	article := r.Group("/article")
	article.Use(middleware.JWT())
	{
		article.GET("/allTag",middleware.RequestEntryWithStatus(api.GetAllTag))//获取所有文章标签
		article.GET("",middleware.RequestEntryWithStatus(api.GetArticle))//获取指定文章
		article.POST("",middleware.RequestEntryWithStatus(api.AddArticle))//添加文章
		//article.DELETE("",middleware.RequestEntryWithStatus(api.DeleteArticle))//删除文章
		//article.PUT("",middleware.RequestEntryWithStatus(api.ModifyArticle))//修改文章
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
