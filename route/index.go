package route

import (
	"otaqu/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/api/get-kupon", controller.GetKupon)
	r.Get("/api/show-kupon", controller.ShowKupon)
}
