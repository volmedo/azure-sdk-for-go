//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ClusterPrincipalAssignmentsClient contains the methods for the ClusterPrincipalAssignments group.
// Don't use this type directly, use NewClusterPrincipalAssignmentsClient() instead.
type ClusterPrincipalAssignmentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewClusterPrincipalAssignmentsClient creates a new instance of ClusterPrincipalAssignmentsClient with the specified values.
func NewClusterPrincipalAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ClusterPrincipalAssignmentsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ClusterPrincipalAssignmentsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - Checks that the principal assignment name is valid and is not already in use.
// If the operation fails it returns the *CloudError error type.
func (client *ClusterPrincipalAssignmentsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName ClusterPrincipalAssignmentCheckNameRequest, options *ClusterPrincipalAssignmentsCheckNameAvailabilityOptions) (ClusterPrincipalAssignmentsCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, principalAssignmentName, options)
	if err != nil {
		return ClusterPrincipalAssignmentsCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterPrincipalAssignmentsCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterPrincipalAssignmentsCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ClusterPrincipalAssignmentsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName ClusterPrincipalAssignmentCheckNameRequest, options *ClusterPrincipalAssignmentsCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/checkPrincipalAssignmentNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, principalAssignmentName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ClusterPrincipalAssignmentsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ClusterPrincipalAssignmentsCheckNameAvailabilityResponse, error) {
	result := ClusterPrincipalAssignmentsCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return ClusterPrincipalAssignmentsCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *ClusterPrincipalAssignmentsClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdate - Create a Kusto cluster principalAssignment.
// If the operation fails it returns the *CloudError error type.
func (client *ClusterPrincipalAssignmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, parameters ClusterPrincipalAssignment, options *ClusterPrincipalAssignmentsBeginCreateOrUpdateOptions) (ClusterPrincipalAssignmentsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, principalAssignmentName, parameters, options)
	if err != nil {
		return ClusterPrincipalAssignmentsCreateOrUpdatePollerResponse{}, err
	}
	result := ClusterPrincipalAssignmentsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ClusterPrincipalAssignmentsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ClusterPrincipalAssignmentsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ClusterPrincipalAssignmentsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create a Kusto cluster principalAssignment.
// If the operation fails it returns the *CloudError error type.
func (client *ClusterPrincipalAssignmentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, parameters ClusterPrincipalAssignment, options *ClusterPrincipalAssignmentsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, principalAssignmentName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ClusterPrincipalAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, parameters ClusterPrincipalAssignment, options *ClusterPrincipalAssignmentsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments/{principalAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if principalAssignmentName == "" {
		return nil, errors.New("parameter principalAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{principalAssignmentName}", url.PathEscape(principalAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ClusterPrincipalAssignmentsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a Kusto cluster principalAssignment.
// If the operation fails it returns the *CloudError error type.
func (client *ClusterPrincipalAssignmentsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, options *ClusterPrincipalAssignmentsBeginDeleteOptions) (ClusterPrincipalAssignmentsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, principalAssignmentName, options)
	if err != nil {
		return ClusterPrincipalAssignmentsDeletePollerResponse{}, err
	}
	result := ClusterPrincipalAssignmentsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ClusterPrincipalAssignmentsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ClusterPrincipalAssignmentsDeletePollerResponse{}, err
	}
	result.Poller = &ClusterPrincipalAssignmentsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a Kusto cluster principalAssignment.
// If the operation fails it returns the *CloudError error type.
func (client *ClusterPrincipalAssignmentsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, options *ClusterPrincipalAssignmentsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, principalAssignmentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ClusterPrincipalAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, options *ClusterPrincipalAssignmentsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments/{principalAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if principalAssignmentName == "" {
		return nil, errors.New("parameter principalAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{principalAssignmentName}", url.PathEscape(principalAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ClusterPrincipalAssignmentsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a Kusto cluster principalAssignment.
// If the operation fails it returns the *CloudError error type.
func (client *ClusterPrincipalAssignmentsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, options *ClusterPrincipalAssignmentsGetOptions) (ClusterPrincipalAssignmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, principalAssignmentName, options)
	if err != nil {
		return ClusterPrincipalAssignmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterPrincipalAssignmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterPrincipalAssignmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ClusterPrincipalAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, options *ClusterPrincipalAssignmentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments/{principalAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if principalAssignmentName == "" {
		return nil, errors.New("parameter principalAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{principalAssignmentName}", url.PathEscape(principalAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClusterPrincipalAssignmentsClient) getHandleResponse(resp *http.Response) (ClusterPrincipalAssignmentsGetResponse, error) {
	result := ClusterPrincipalAssignmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterPrincipalAssignment); err != nil {
		return ClusterPrincipalAssignmentsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ClusterPrincipalAssignmentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all Kusto cluster principalAssignments.
// If the operation fails it returns the *CloudError error type.
func (client *ClusterPrincipalAssignmentsClient) List(ctx context.Context, resourceGroupName string, clusterName string, options *ClusterPrincipalAssignmentsListOptions) (ClusterPrincipalAssignmentsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ClusterPrincipalAssignmentsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterPrincipalAssignmentsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterPrincipalAssignmentsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ClusterPrincipalAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClusterPrincipalAssignmentsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ClusterPrincipalAssignmentsClient) listHandleResponse(resp *http.Response) (ClusterPrincipalAssignmentsListResponse, error) {
	result := ClusterPrincipalAssignmentsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterPrincipalAssignmentListResult); err != nil {
		return ClusterPrincipalAssignmentsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ClusterPrincipalAssignmentsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}