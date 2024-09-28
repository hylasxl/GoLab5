package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/internal/db"
)

func main() {
	r := gin.Default()

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	err = db.MigrateDB(database)
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
		return
	}

	r.POST("/students", func(c *gin.Context) { controllers.CreateStudent(c, database) })
	r.GET("/students", func(c *gin.Context) { controllers.GetStudents(c, database) })
	r.POST("/students/id", func(c *gin.Context) { controllers.GetStudentByID(c, database) })
	r.PUT("/students", func(c *gin.Context) { controllers.UpdateStudent(c, database) })
	r.DELETE("/students", func(c *gin.Context) { controllers.DeleteStudent(c, database) })

	r.POST("/classes", func(c *gin.Context) { controllers.CreateClass(c, database) })
	r.GET("/classes", func(c *gin.Context) { controllers.GetClasses(c, database) })
	r.POST("/classes/id", func(c *gin.Context) { controllers.GetClassByID(c, database) })
	r.PUT("/classes", func(c *gin.Context) { controllers.UpdateClass(c, database) })
	r.DELETE("/classes", func(c *gin.Context) { controllers.DeleteClass(c, database) })

	err = r.Run()
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
		return
	}
}
