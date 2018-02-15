package a

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/csvikram/test-packaging/b"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//index, err := ioutil.ReadFile("public/index.html")
	//if err != nil {
	//	return events.APIGatewayProxyResponse{}, err
	//}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       b.GetString(),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(handler)
}
