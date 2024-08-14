package router

import (
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/Delivery/controller"
	infrastructure "github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/Infrastructure"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/repositories"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(db mongo.Database, group *gin.RouterGroup) {
	tr := repositories.NewTaskRepository(db, domain.TaskCollection)
	tuc := usecase.NewTaskUsecase(tr)
	tc := controller.NewTaskController(tuc)

	group.Use(infrastructure.AuthMiddleware("admin"))

	group.GET("/", tc.GetTasks)
	group.GET("/:id", tc.GetTaskById)
	group.POST("/", tc.CreateTask)
	group.PUT("/:id", tc.UpdateTask)
	group.DELETE("/:id", tc.DeleteTask)
}

func NewUserRouter(db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, "user")
	pis := infrastructure.NewPasswordInfrastructureService()
	jwtS := infrastructure.NewJwtService()
	uuc := usecase.NewUserUsecase(ur, pis, jwtS)
	uc := controller.NewUserController(uuc)

	group.POST("/register", uc.SignUp)
	group.POST("/login", uc.LoginUser)
	group.PUT("/promote/:id", infrastructure.AuthMiddleware("admin"), uc.PromoteUser)
}
