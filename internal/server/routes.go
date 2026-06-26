package server

import (
	"net/http"
	"spot-sync/internal/config"
	"spot-sync/internal/domain/user"
	"spot-sync/internal/domain/zone"
	"spot-sync/internal/httpresponse"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB, env *config.Env) {
	(func() {
		db.Exec(`
    		DO $$
    		BEGIN
        		IF NOT EXISTS (
            		SELECT 1 FROM pg_type WHERE typname = 'user_role'
        		) THEN
            		CREATE TYPE user_role AS ENUM (
                		'ADMIN',
                		'DRIVER'
            		);
        		ELSE
            		ALTER TYPE user_role ADD VALUE IF NOT EXISTS 'ADMIN';
            		ALTER TYPE user_role ADD VALUE IF NOT EXISTS 'DRIVER';
        		END IF;
    		END
    		$$;
		`)

		db.Exec(`
    		DO $$
    		BEGIN
        		IF NOT EXISTS (
            		SELECT 1 FROM pg_type WHERE typname = 'zone_type'
        		) THEN
            		CREATE TYPE zone_type AS ENUM (
                		'GENERAL',
                		'EV_CHARGING',
                		'COVERED'
            		);
        		ELSE
            		ALTER TYPE zone_type ADD VALUE IF NOT EXISTS 'GENERAL';
            		ALTER TYPE zone_type ADD VALUE IF NOT EXISTS 'EV_CHARGING';
            		ALTER TYPE zone_type ADD VALUE IF NOT EXISTS 'COVERED';
        		END IF;
    		END
    		$$;
		`)
	})()

	db.AutoMigrate(&user.User{}, &zone.Zone{})

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Response{
			Success: true,
			Message: "Server is running...",
		})
	})

	api := e.Group("/api/v1")

	user.Routes(db, api, env)
	zone.Routes(db, api)
}
