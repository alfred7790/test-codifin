package web

import (
	"codifin-challenge/config"
	_ "codifin-challenge/docs"
	"codifin-challenge/domain/repository"
	"codifin-challenge/domain/service"
	"codifin-challenge/infrastructure/web/controller"
	"codifin-challenge/infrastructure/web/database"
	"codifin-challenge/infrastructure/web/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	cfg          *config.Config
	router       *gin.Engine
	db           *gorm.DB
	middlewares  middlewares.MiddlewareService
	controllers  Controllers
	services     Services
	repositories Repositories
}

type Controllers struct {
	productCtrl      *controller.ProductController
	shoppingCartCtrl *controller.ShoppingCartController
}

type Services struct {
	productService      service.ProductService
	shoppingCartService service.ShoppingCartService
}

type Repositories struct {
	productRepository      repository.ProductRepository
	shoppingCartRepository repository.ShoppingCartRepository
}

func NewServer() *Server {
	s := &Server{}
	s.setup()
	return s
}

func (s *Server) Run() {
	err := s.router.Run(fmt.Sprintf(":%s", s.cfg.Host.Port))
	if err != nil {
		log.Fatalf("server can not run: %s", err.Error())
	}
}

func (s *Server) setup() {
	s.setConfig()
	s.setDataBase()
	s.setRouter()
	s.setRepositories()
	s.setServices()
	s.setControllers()
	s.setRoutes()
}

func (s *Server) setConfig() {
	s.cfg = config.GetConfig()
}

func (s *Server) setDataBase() {
	db, err := database.NewDataBase(s.cfg.DB.Host, s.cfg.DB.Port, s.cfg.DB.User, s.cfg.DB.Password, s.cfg.DB.Name, s.cfg.DB.Retries)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	s.db = db
}

func (s *Server) setRouter() {
	s.router = gin.Default()
}

func (s *Server) setRepositories() {
	s.repositories.productRepository = repository.NewProductRepository(s.db)
	s.repositories.shoppingCartRepository = repository.NewShoppingCartRepository(s.db)
}

func (s *Server) setServices() {
	s.middlewares = middlewares.NewMiddlewareService()

	s.services.productService = service.NewProductService(s.repositories.productRepository)
	s.services.shoppingCartService = service.NewShoppingCartService(s.repositories.shoppingCartRepository)
}

func (s *Server) setControllers() {
	s.controllers.productCtrl = controller.NewProductController(s.services.productService)
	s.controllers.shoppingCartCtrl = controller.NewShoppingCartController(s.services.shoppingCartService)
}
