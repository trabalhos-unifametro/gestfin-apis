package routes

import (
	"gestfin-apis/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes() *fiber.App {
	crs := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, DELETE, PUT, OPTIONS",
		AllowHeaders:     "*",
		AllowCredentials: true,
	})
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // Limite de 10Mb para envio de arquivos
	})
	app.Use(crs)

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// AUTHENTICATION ROUTES ==============================================================
			v1.Post("/auth/signin", controllers.SignIn)
			v1.Post("/auth/generate_code", controllers.GenerateCodeRecoverPassword)
			v1.Post("/auth/confirm_code", controllers.ConfirmCode)
			v1.Put("/auth/recover_password", controllers.RecoverPassword)

			// USER ROUTES =============================================================================
			v1.Put("/user/update", controllers.UpdateDataUser)
			v1.Put("/user/change_password", controllers.ChangePassword)
			v1.Get("/user/my_data", controllers.GetUserDataByToken)
			v1.Put("/user/update_by_column", controllers.UpdateUserByColumn)
		}
	}

	return app
}
