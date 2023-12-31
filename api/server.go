package api 

import (
	"log"
	"os"
	"rest-book/storage"

	"github.com/gofiber/fiber/v2"
)
type Server struct {
	Address			string 
	Store			storage.Store // Storage Dependency
}

func NewServer(address string ) *Server {
	uri := os.Getenv("MONGO_URI")
	store := storage.NewMongoStore(uri)
	store.RestaurantCollection("rest-book", "restaurant_details")
	store.UserCollection("rest-book", "user_details")

	return &Server {
		Address: address, 
		Store: store,
	}
}

func (s *Server) Start() {
	app := fiber.New()
	v1 := app.Group("api/v1")

	// Handlers for Restaurant
	v1.Get("/restaurant", s.GetAllRestaurantDetails)
	v1.Get("/restaurant/:id", s.GetRestaurantDetailsById)
	v1.Put("/restaurant/:id", s.UpdateRestaurantDetails)
	v1.Delete("/restaurant/:id", s.DeleteRestaurantDetails)
	v1.Post("/restaurant", s.AddRestaurantDetails)

	// Handlers for Users
	v1.Post("/user", s.CreateUser)
	v1.Get("/user", s.GetUsers)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}

