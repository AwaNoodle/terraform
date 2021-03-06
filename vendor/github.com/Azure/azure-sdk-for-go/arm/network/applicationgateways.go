package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// ApplicationGatewaysClient is the the Microsoft Azure Network management API
// provides a RESTful set of web services that interact with Microsoft Azure
// Networks service to manage your network resrources. The API has entities
// that capture the relationship between an end user and the Microsoft Azure
// Networks service.
type ApplicationGatewaysClient struct {
	ManagementClient
}

// NewApplicationGatewaysClient creates an instance of the
// ApplicationGatewaysClient client.
func NewApplicationGatewaysClient(subscriptionID string) ApplicationGatewaysClient {
	return NewApplicationGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGatewaysClientWithBaseURI creates an instance of the
// ApplicationGatewaysClient client.
func NewApplicationGatewaysClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewaysClient {
	return ApplicationGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put ApplicationGateway operation creates/updates a
// ApplicationGateway
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the ApplicationGateway. parameters is parameters supplied
// to the create/delete ApplicationGateway operation
func (client ApplicationGatewaysClient) CreateOrUpdate(resourceGroupName string, applicationGatewayName string, parameters ApplicationGateway) (result ApplicationGateway, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, applicationGatewayName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "CreateOrUpdate", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "CreateOrUpdate", resp.StatusCode, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "CreateOrUpdate", resp.StatusCode, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplicationGatewaysClient) CreateOrUpdatePreparer(resourceGroupName string, applicationGatewayName string, parameters ApplicationGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": url.QueryEscape(applicationGatewayName),
		"resourceGroupName":      url.QueryEscape(resourceGroupName),
		"subscriptionId":         url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the delete applicationgateway operation deletes the specified
// applicationgateway.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the applicationgateway.
func (client ApplicationGatewaysClient) Delete(resourceGroupName string, applicationGatewayName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, applicationGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Delete", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Delete", resp.StatusCode, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Delete", resp.StatusCode, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationGatewaysClient) DeletePreparer(resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": url.QueryEscape(applicationGatewayName),
		"resourceGroupName":      url.QueryEscape(resourceGroupName),
		"subscriptionId":         url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusAccepted, http.StatusNoContent, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get applicationgateway operation retreives information about the
// specified applicationgateway.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the applicationgateway.
func (client ApplicationGatewaysClient) Get(resourceGroupName string, applicationGatewayName string) (result ApplicationGateway, ae error) {
	req, err := client.GetPreparer(resourceGroupName, applicationGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Get", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Get", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Get", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationGatewaysClient) GetPreparer(resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": url.QueryEscape(applicationGatewayName),
		"resourceGroupName":      url.QueryEscape(resourceGroupName),
		"subscriptionId":         url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) GetResponder(resp *http.Response) (result ApplicationGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List ApplicationGateway opertion retrieves all the
// applicationgateways in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ApplicationGatewaysClient) List(resourceGroupName string) (result ApplicationGatewayListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "List", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "List", resp.StatusCode, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "List", resp.StatusCode, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationGatewaysClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListResponder(resp *http.Response) (result ApplicationGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ApplicationGatewaysClient) ListNextResults(lastResults ApplicationGatewayListResult) (result ApplicationGatewayListResult, ae error) {
	req, err := lastResults.ApplicationGatewayListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "List", autorest.UndefinedStatusCode, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "List", resp.StatusCode, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "List", resp.StatusCode, "Failure responding to next results request request")
	}

	return
}

// ListAll the List applicationgateway opertion retrieves all the
// applicationgateways in a subscription.
func (client ApplicationGatewaysClient) ListAll() (result ApplicationGatewayListResult, ae error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "ListAll", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "ListAll", resp.StatusCode, "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "ListAll", resp.StatusCode, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client ApplicationGatewaysClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGateways"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListAllResponder(resp *http.Response) (result ApplicationGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client ApplicationGatewaysClient) ListAllNextResults(lastResults ApplicationGatewayListResult) (result ApplicationGatewayListResult, ae error) {
	req, err := lastResults.ApplicationGatewayListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "ListAll", autorest.UndefinedStatusCode, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "ListAll", resp.StatusCode, "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "ListAll", resp.StatusCode, "Failure responding to next results request request")
	}

	return
}

// Start the Start ApplicationGateway operation starts application gatewayin
// the specified resource group through Network resource provider.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway.
func (client ApplicationGatewaysClient) Start(resourceGroupName string, applicationGatewayName string) (result autorest.Response, ae error) {
	req, err := client.StartPreparer(resourceGroupName, applicationGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Start", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Start", resp.StatusCode, "Failure sending request")
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Start", resp.StatusCode, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client ApplicationGatewaysClient) StartPreparer(resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": url.QueryEscape(applicationGatewayName),
		"resourceGroupName":      url.QueryEscape(resourceGroupName),
		"subscriptionId":         url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/start"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) StartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted)
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop the STOP ApplicationGateway operation stops application gatewayin the
// specified resource group through Network resource provider.
//
// resourceGroupName is the name of the resource group. applicationGatewayName
// is the name of the application gateway.
func (client ApplicationGatewaysClient) Stop(resourceGroupName string, applicationGatewayName string) (result autorest.Response, ae error) {
	req, err := client.StopPreparer(resourceGroupName, applicationGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Stop", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Stop", resp.StatusCode, "Failure sending request")
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ApplicationGatewaysClient", "Stop", resp.StatusCode, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client ApplicationGatewaysClient) StopPreparer(resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": url.QueryEscape(applicationGatewayName),
		"resourceGroupName":      url.QueryEscape(resourceGroupName),
		"subscriptionId":         url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/stop"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted)
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
