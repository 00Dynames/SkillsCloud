package main

import (
	"github.com/00Dynames/SkillsCloud/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
