/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"rahul.com/secure-exam/internal/handlers"
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
	server.GET("/ping", handlers.GetExams)
	server.GET("/exams", handlers.GetExams)
	server.POST("/exam", handlers.CreateExams)
	server.POST("/exam/question", handlers.CreateQuestion)
	server.GET("/exam", handlers.GetExamByID)

	server.Run(":8080")
}
func init() {
	rootCmd.AddCommand(examCmd)
}
