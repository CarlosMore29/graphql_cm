package graphql_cm

import (
	"context"
	"os"

	"github.com/CarlosMore29/env_cm"
	"github.com/machinebox/graphql"
)

type InputVariable struct {
	Key   string
	Value interface{}
}

type MutationResponse struct {
	AffectedRows int    `json:"affects_rows"`
	Error        string `json:"error"`
	Id           int    `json:"id"`
}

var APIQL_URL string
var APIQL_SECRET string

func init() {
	envGLobals()
}

// type queryRequest func (string, string, []InputVariable) (map[string]interface{}, error)
type APIQL struct {
	apiQLUrl    string
	apiQLSecret string
}

func InitAPIQL() *APIQL {
	apiql := APIQL{
		apiQLUrl:    APIQL_URL,
		apiQLSecret: APIQL_SECRET,
	}

	return &apiql
}

func (f *APIQL) SetCredentials(url, secret string) {
	f.apiQLUrl = url
	f.apiQLSecret = secret
}

func (f *APIQL) QueryRequest(query string, variables []InputVariable) (map[string]interface{}, error) {
	// Type Response
	var respData map[string]interface{}

	// Client
	client := graphql.NewClient(f.apiQLUrl)

	// make a request
	req := graphql.NewRequest(query)

	// Variables
	for _, input := range variables {
		req.Var(input.Key, input.Value)
	}

	// set header fields
	req.Header.Set("x-hasura-admin-secret", f.apiQLSecret)

	// define a Context for the request
	ctx := context.Background()

	err := client.Run(ctx, req, &respData)
	return respData, err
}

func (f *APIQL) MutationRequest(query, inputRootName string, variables interface{}) (map[string]MutationResponse, error) {
	// Type Request
	var respData map[string]MutationResponse
	var errorReq error

	// Client
	client := graphql.NewClient(f.apiQLUrl)

	// make a request
	req := graphql.NewRequest(query)

	// Variables
	req.Var(inputRootName, variables)

	// set header fields
	req.Header.Set("x-hasura-admin-secret", f.apiQLSecret)

	// define a Context for the request
	ctx := context.Background()

	errorReq = client.Run(ctx, req, &respData)

	return respData, errorReq
}

func envGLobals() {
	env_cm.GetEnvFile()
	APIQL_URL = os.Getenv("APIQL_URL")
	APIQL_SECRET = os.Getenv("APIQL_SECRET")
}
