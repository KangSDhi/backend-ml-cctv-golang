package routes

import (
	"backend-ml-cctv-golang/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCCTVRouter(router fiber.Router) {

	cctvRoutes := router.Group("/cctv")

	cctvRoutes.Post("", controllers.StoreCCTVData)
	cctvRoutes.Get("/last", controllers.GetLastCCTVData)
}
