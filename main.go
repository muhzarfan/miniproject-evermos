package main

import (
	"evermos/apps/auth"
	"evermos/apps/category"
	"evermos/apps/produk"
	"evermos/apps/toko"
	"evermos/apps/transaction"
	"evermos/apps/user"
	"evermos/apps/wilayah"
	"evermos/config"
	"evermos/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Koneksi ke database
	config.InitDB()

	// Inisialisasi Fiber untuk API
	app := fiber.New()

	// Register/Login/Logout
	api := app.Group("/api")

	authGroup := api.Group("/auth")
	authGroup.Post("/register", auth.RegisterHandler)
	authGroup.Post("/login", auth.LoginHandler)
	authGroup.Post("/logout", auth.LogoutHandler)

	// API Wilayah
	api.Get("/provinsi", wilayah.GetProvincesHandler)
	api.Get("/kota/:id_provinsi", wilayah.GetRegenciesHandler)

	// User
	userGroup := api.Group("/user", middleware.AuthMiddleware)
	userGroup.Get("/me", user.GetMyProfileHandler)
	userGroup.Put("/me", user.UpdateProfileHandler)
	userGroup.Put("/change-password", user.ChangePasswordHandler)

	// Alamat
	userGroup.Get("/alamat", user.GetMyAlamatHandler)
	userGroup.Post("/alamat", user.CreateAlamatHandler)
	userGroup.Put("/alamat/:id", user.UpdateAlamatHandler)
	userGroup.Delete("/alamat/:id", user.DeleteAlamatHandler)

	// Toko
	tokoGroup := api.Group("/toko", middleware.AuthMiddleware)
	tokoGroup.Get("/my", toko.GetMyTokoHandler)
	tokoGroup.Put("/:id_toko", toko.UpdateTokoHandler)

	// Kategori
	api.Get("/category", middleware.AuthMiddleware, category.GetAllCategoryHandler)

	// CRUD Kategori (Admin)
	adminCat := api.Group("/category", middleware.AuthMiddleware, middleware.AdminMiddleware)
	adminCat.Post("/", category.CreateCategoryHandler)
	adminCat.Put("/:id", category.UpdateCategoryHandler)
	adminCat.Delete("/:id", category.DeleteCategoryHandler)

	// Produk
	api.Get("/produk", produk.GetAllProdukHandler)
	api.Get("/produk/:id", produk.GetProdukByIDHandler)

	produkOwner := api.Group("/produk", middleware.AuthMiddleware)
	produkOwner.Post("/", produk.CreateProdukHandler)
	produkOwner.Put("/:id", produk.UpdateProdukHandler)
	produkOwner.Delete("/:id", produk.DeleteProdukHandler)

	// Transaksi
	trxGroup := api.Group("/api/trx", middleware.AuthMiddleware)
	trxGroup.Get("/", transaction.GetAllTrxHandler)
	trxGroup.Post("/", transaction.CreateTrxHandler)

	// Jalankan Server
	log.Fatal(app.Listen(":8000"))
}
