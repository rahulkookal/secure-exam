/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"rahul.com/secure-exam/internal/routes"
)

// examCmd represents the exam command
var examCmd = &cobra.Command{
	Use:   "exam",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exam called")
		runServer()
	},
}

func runServer() {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Update with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.RegisterExamRoutes(server)
	routes.RegisterUserRoutes(server)
	routes.RegisterEventRoutes(server)

	server.Run(":8080")
}
func init() {
	rootCmd.AddCommand(examCmd)
}
