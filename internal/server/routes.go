package server

import (
	"net/http"
	"spot-sync/internal/domain/user"
	"spot-sync/internal/httpresponse"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	// (func() {
	// 	db.Exec(`
	// 		DO $$
	// 		BEGIN
	// 		IF NOT EXISTS (
	// 	    	SELECT 1 FROM pg_type WHERE typname = 'user_role'
	// 		) THEN
	// 	    CREATE TYPE user_role AS ENUM (
	// 	        'ADMIN',
	// 	        'DRIVER'
	// 	    );
	// 		END IF;
	// 		END
	// 		$$;
	// 	`)
	// })()

	db.AutoMigrate(&user.User{})

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Response{
			Success: true,
			Message: "Server is running...",
		})
	})

	api := e.Group("/api/v1")

	user.Routes(db, api)
}
