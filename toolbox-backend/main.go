package main

import (
	"net/http"
	"github.com/gin-contrib/cors"
	"time"
	"github.com/gin-gonic/gin"
)

type Slot struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	Tags    string `json:"tags"`
	Slots   []Slot `json:"slots"`
}

type Toolbox struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	Slots   []Slot `json:"slots"`
}

func main() {
	r := gin.Default()

	// Add the CORS middleware here
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allowed origins (modify as needed)
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		data := []Toolbox{
			{
				ID:      0,
				Address: "0",
				Slots: []Slot{
					{
						ID:      0,
						Address: "0.0",
						Tags:    "pencils",
						Slots: []Slot{
							{
								ID:      1,
								Address: "0.0.0",
								Tags:    "pencil sharpeners",
								Slots:   []Slot{},
							},
						},
					},
				},
			},
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080") // Start the Gin server on port 8080
}
