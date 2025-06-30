package router

import (
	psw "inibackend/config/middleware"
	"inibackend/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", psw.Middlewares("admin"))

	// Route untuk homepage
	api.Get("/", handler.Homepage)

	api.Get("/mahasiswa", handler.GetAllMahasiswa)

	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)

	// Route untuk menambah mahasiswa baru
	api.Post("/mahasiswa", handler.InsertMahasiswa)

	// Route untuk mengupdate data mahasiswa berdasarkan NPM
	api.Put("/mahasiswa/:npm", handler.UpdateMahasiswa)

	// Route untuk menghapus data mahasiswa berdasarkan NPM
	api.Delete("/mahasiswa/:npm", handler.DeleteMahasiswa)

	// Dokumentasi swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

}