package main

import (
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "os"
  "log"
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  "strings"
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
  log.Printf("%v", request.Path)

  splitPath := strings.Split(request.Path, "/")

  if (splitPath[0] == ".netlify") {
    splitPath = splitPath[2:]
  }

  if len(splitPath) < 3 {
    return events.APIGatewayProxyResponse{
    Body:       "Not found\n",
    StatusCode: 404,
  }, nil
  }

  id := splitPath[2]

  return events.APIGatewayProxyResponse{
    Body:       "",
    StatusCode: 302,
    Headers: map[string]string{
      "Location": fmt.Sprintf("https://%s.ngrok.id", id),
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
