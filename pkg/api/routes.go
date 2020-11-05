package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narakosen-festival-info-2020/clicker-back/pkg/facility"
)

func facilityRoute(engine *gin.Engine, app *App) {
	group := engine.Group("/facility")

	group.GET("", func(ctx *gin.Context) {
		type responce struct {
			Facilities []facility.JSONData `json:"facilities"`
		}
		ctx.JSON(http.StatusOK, responce{
			Facilities: app.clickerData.GetAllFacilityJSON(),
		})
	})

	group.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ret, err := app.clickerData.GetFacilityJSON(name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, ret)
	})

	group.POST("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		err := app.clickerData.PurchaseFacility(name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Accept",
		})
	})
}

func statementsRoute(engine *gin.Engine, app *App) {
	group := engine.Group("/statements")

	group.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, app.clickerData.GetStatements())
	})
}

func achievementsRoute(engine *gin.Engine, app *App) {
	group := engine.Group("/achievements")

	group.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, app.clickerData.GetAchievemnets())
	})
}
