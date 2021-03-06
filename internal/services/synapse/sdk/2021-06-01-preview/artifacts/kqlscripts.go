package artifacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
)

// KqlScriptsClient is the client for the KqlScripts methods of the Artifacts service.
type KqlScriptsClient struct {
	BaseClient
}

// NewKqlScriptsClient creates an instance of the KqlScriptsClient client.
func NewKqlScriptsClient(endpoint string) KqlScriptsClient {
	return KqlScriptsClient{New(endpoint)}
}

// GetAll get all KQL scripts
func (client KqlScriptsClient) GetAll(ctx context.Context) (result KqlScriptsResourceCollectionResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KqlScriptsClient.GetAll")
		defer func() {
			sc := -1
			if result.ksrcr.Response.Response != nil {
				sc = result.ksrcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllNextResults
	req, err := client.GetAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.KqlScriptsClient", "GetAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllSender(req)
	if err != nil {
		result.ksrcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.KqlScriptsClient", "GetAll", resp, "Failure sending request")
		return
	}

	result.ksrcr, err = client.GetAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.KqlScriptsClient", "GetAll", resp, "Failure responding to request")
		return
	}
	if result.ksrcr.hasNextLink() && result.ksrcr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetAllPreparer prepares the GetAll request.
func (client KqlScriptsClient) GetAllPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"endpoint": client.Endpoint,
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPath("/kqlScripts"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllSender sends the GetAll request. The method will close the
// http.Response Body if it receives an error.
func (client KqlScriptsClient) GetAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAllResponder handles the response to the GetAll request. The method always
// closes the http.Response Body.
func (client KqlScriptsClient) GetAllResponder(resp *http.Response) (result KqlScriptsResourceCollectionResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllNextResults retrieves the next set of results, if any.
func (client KqlScriptsClient) getAllNextResults(ctx context.Context, lastResults KqlScriptsResourceCollectionResponse) (result KqlScriptsResourceCollectionResponse, err error) {
	req, err := lastResults.kqlScriptsResourceCollectionResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "artifacts.KqlScriptsClient", "getAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "artifacts.KqlScriptsClient", "getAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.KqlScriptsClient", "getAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client KqlScriptsClient) GetAllComplete(ctx context.Context) (result KqlScriptsResourceCollectionResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KqlScriptsClient.GetAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAll(ctx)
	return
}
