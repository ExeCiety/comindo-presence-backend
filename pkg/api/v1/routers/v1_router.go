package routers

import (
	apiV1AuthRouters "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/routers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	v1Router := router.Group("/v1")

	apiV1AuthRouters.SetRouter(v1Router)
}