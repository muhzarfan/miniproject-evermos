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
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Koneksi ke database
	config.InitDB()

	// Routes requests
	r := mux.NewRouter()

	// Route Register/Login/Logout
	r.HandleFunc("/api/auth/register", auth.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/auth/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/api/auth/logout", auth.LogoutHandler).Methods("POST")

	// Route API Wilayah
	r.HandleFunc("/api/provinsi", wilayah.GetProvincesHandler).Methods("GET")
	r.HandleFunc("/api/kota/{id_provinsi}", wilayah.GetRegenciesHandler).Methods("GET")

	// Route User (Akun dan Alamat)
	userRouter := r.PathPrefix("/api/user").Subrouter()
	userRouter.Use(middleware.AuthMiddleware)
	userRouter.HandleFunc("/me", user.GetMyProfileHandler).Methods("GET")
	userRouter.HandleFunc("/me", user.UpdateProfileHandler).Methods("PUT")
	userRouter.HandleFunc("/change-pwd", user.ChangePasswordHandler).Methods("PUT")
	userRouter.HandleFunc("/alamat", user.GetMyAlamatHandler).Methods("GET")
	userRouter.HandleFunc("/alamat", user.CreateAlamatHandler).Methods("POST")
	userRouter.HandleFunc("/alamat/{id}", user.UpdateAlamatHandler).Methods("PUT")
	userRouter.HandleFunc("/alamat/{id}", user.DeleteAlamatHandler).Methods("DELETE")

	// Route Toko
	tokoRouter := r.PathPrefix("/api/toko").Subrouter()
	tokoRouter.Use(middleware.AuthMiddleware)
	tokoRouter.HandleFunc("/my", toko.GetMyTokoHandler).Methods("GET")
	tokoRouter.HandleFunc("/{id_toko}", toko.UpdateTokoHandler).Methods("PUT")

	// Route Kategori
	// GET Kategori (semua user bisa lihat)
	r.Handle("/api/category", middleware.AuthMiddleware(http.HandlerFunc(category.GetAllCategoryHandler))).Methods("GET")

	// CRUD Kategori (bagi admin)
	adminRouter := r.PathPrefix("/api/category").Subrouter()
	adminRouter.Use(middleware.AuthMiddleware)
	adminRouter.Use(middleware.AdminMiddleware)
	adminRouter.HandleFunc("", category.CreateCategoryHandler).Methods("POST")
	adminRouter.HandleFunc("/{id}", category.UpdateCategoryHandler).Methods("PUT")
	adminRouter.HandleFunc("/{id}", category.DeleteCategoryHandler).Methods("DELETE")

	// Route Produk
	// GET Produk (semua user bisa lihat)
	r.HandleFunc("/api/produk", produk.GetAllProdukHandler).Methods("GET")
	r.HandleFunc("/api/produk/{id}", produk.GetProdukByIDHandler).Methods("GET")

	// Route CRUD hanya untuk Pemilik Toko
	produkRouter := r.PathPrefix("/api/produk").Subrouter()
	produkRouter.Use(middleware.AuthMiddleware)
	produkRouter.HandleFunc("", produk.CreateProdukHandler).Methods("POST")
	produkRouter.HandleFunc("/{id}", produk.UpdateProdukHandler).Methods("PUT")
	produkRouter.HandleFunc("/{id}", produk.DeleteProdukHandler).Methods("DELETE")

	// Route Transaksi
	trxRouter := r.PathPrefix("/api/trx").Subrouter()
	trxRouter.Use(middleware.AuthMiddleware)
	trxRouter.HandleFunc("", transaction.GetAllTrxHandler).Methods("GET")
	trxRouter.HandleFunc("", transaction.CreateTrxHandler).Methods("POST")

	// Jalankan Server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
