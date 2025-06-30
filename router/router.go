package router

import (
	psw "inibackend/config/middleware"
	"inibackend/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Route untuk homepage
	api.Get("/", handler.Homepage)

	api.Get("/mahasiswa", handler.GetAllMahasiswa)

	api.Get("/mahasiswa/:npm",psw.Middlewares("admin"), handler.GetMahasiswaByNPM)

	// Route untuk menambah mahasiswa baru
	api.Post("/mahasiswa", psw.Middlewares("admin") ,handler.InsertMahasiswa)

	// Route untuk mengupdate data mahasiswa berdasarkan NPM
	api.Put("/mahasiswa/:npm",psw.Middlewares("admin"), handler.UpdateMahasiswa)

	// Route untuk menghapus data mahasiswa berdasarkan NPM
	api.Delete("/mahasiswa/:npm",psw.Middlewares("admin"), handler.DeleteMahasiswa)

	// Dokumentasi swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

}