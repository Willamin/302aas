package main

import (
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "os"
  "log"
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
)

const (
  AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
)

func main() {
  _, ok := os.LookupEnv(AWSLambdaFunctionVersion)
  if ok {
    log.Printf("Running in AWS lambda environment, starting lambda handler.")
    lambda.Start(AWSHandler)
    os.Exit(0)
  }

  log.Printf("Not running in AWS lambda environment, starting mock handler.")
  MockHandler()
  os.Exit(0)
}

func AWSHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  id := "a"

  log.Printf("%v", request)

  return events.APIGatewayProxyResponse{
    Body:       "",
    StatusCode: 302,
    Headers: map[string]string{
      "Location": fmt.Sprintf("https://%s.ngrok.io", id),
    },
  }, nil
}

func MockHandler() {
  r := gin.Default()

  r.GET("/:id", func(c *gin.Context) {
    path := fmt.Sprintf("https://%s.ngrok.io", c.Param("id"))
    c.Redirect(http.StatusFound, path)
  })

  r.Run()
}
