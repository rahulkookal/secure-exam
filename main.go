package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rahul.com/secure-exam/internal/routes"
)

var ginLambda *ginadapter.GinLambda

func main() {
	server := gin.Default()
	server.Use(LogrusLogger(), gin.Recovery())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this for production
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register routes
	routes.RegisterExamRoutes(server)
	routes.RegisterUserRoutes(server)
	routes.RegisterEventRoutes(server)
	routes.RegisterOrganisationRoutes(server)

	if isLocal() {
		server.Run(":8080")
	} else {
		// Initialize ginLambda only for AWS Lambda
		ginLambda = ginadapter.New(server)
		lambda.Start(Handler)
	}
}

// AWS Lambda Handler
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(req)
}

// Logging Middleware
func LogrusLogger() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logger.WithFields(logrus.Fields{
			"time":       time.Now().Format(time.RFC3339),
			"status":     c.Writer.Status(),
			"latency":    time.Since(start).String(),
			"client_ip":  c.ClientIP(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"user_agent": c.Request.UserAgent(),
		}).Info("Request Processed")
	}
}

// Helper to check if running locally
func isLocal() bool {
	return os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == ""
}
