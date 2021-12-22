//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloadmonitor

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
	"time"
)

// HealthMonitorsClient contains the methods for the HealthMonitors group.
// Don't use this type directly, use NewHealthMonitorsClient() instead.
type HealthMonitorsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewHealthMonitorsClient creates a new instance of HealthMonitorsClient with the specified values.
func NewHealthMonitorsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *HealthMonitorsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &HealthMonitorsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get the current health status of a monitor of a virtual machine. Optional parameter: $expand (retrieve the monitor's evidence and configuration).
// If the operation fails it returns the *ErrorResponse error type.
func (client *HealthMonitorsClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsGetOptions) (HealthMonitorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, options)
	if err != nil {
		return HealthMonitorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HealthMonitorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HealthMonitorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HealthMonitorsClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if monitorID == "" {
		return nil, errors.New("parameter monitorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorId}", url.PathEscape(monitorID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HealthMonitorsClient) getHandleResponse(resp *http.Response) (HealthMonitorsGetResponse, error) {
	result := HealthMonitorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitor); err != nil {
		return HealthMonitorsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *HealthMonitorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetStateChange - Get the health state change of a monitor of a virtual machine at the provided timestamp. Optional parameter: $expand (retrieve the monitor's
// evidence and configuration).
// If the operation fails it returns the *ErrorResponse error type.
func (client *HealthMonitorsClient) GetStateChange(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, timestampUnix string, options *HealthMonitorsGetStateChangeOptions) (HealthMonitorsGetStateChangeResponse, error) {
	req, err := client.getStateChangeCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, timestampUnix, options)
	if err != nil {
		return HealthMonitorsGetStateChangeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HealthMonitorsGetStateChangeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HealthMonitorsGetStateChangeResponse{}, client.getStateChangeHandleError(resp)
	}
	return client.getStateChangeHandleResponse(resp)
}

// getStateChangeCreateRequest creates the GetStateChange request.
func (client *HealthMonitorsClient) getStateChangeCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, timestampUnix string, options *HealthMonitorsGetStateChangeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history/{timestampUnix}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if monitorID == "" {
		return nil, errors.New("parameter monitorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorId}", url.PathEscape(monitorID))
	if timestampUnix == "" {
		return nil, errors.New("parameter timestampUnix cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{timestampUnix}", url.PathEscape(timestampUnix))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getStateChangeHandleResponse handles the GetStateChange response.
func (client *HealthMonitorsClient) getStateChangeHandleResponse(resp *http.Response) (HealthMonitorsGetStateChangeResponse, error) {
	result := HealthMonitorsGetStateChangeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitorStateChange); err != nil {
		return HealthMonitorsGetStateChangeResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getStateChangeHandleError handles the GetStateChange error response.
func (client *HealthMonitorsClient) getStateChangeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get the current health status of all monitors of a virtual machine. Optional parameters: $expand (retrieve the monitor's evidence and configuration)
// and $filter (filter by monitor name).
// If the operation fails it returns the *ErrorResponse error type.
func (client *HealthMonitorsClient) List(subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, options *HealthMonitorsListOptions) *HealthMonitorsListPager {
	return &HealthMonitorsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp HealthMonitorsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.HealthMonitorList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *HealthMonitorsClient) listCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, options *HealthMonitorsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HealthMonitorsClient) listHandleResponse(resp *http.Response) (HealthMonitorsListResponse, error) {
	result := HealthMonitorsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitorList); err != nil {
		return HealthMonitorsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *HealthMonitorsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListStateChanges - Get the health state changes of a monitor of a virtual machine within the provided time window (default is the last 24 hours). Optional
// parameters: $expand (retrieve the monitor's evidence and
// configuration) and $filter (filter by heartbeat condition).
// If the operation fails it returns the *ErrorResponse error type.
func (client *HealthMonitorsClient) ListStateChanges(subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsListStateChangesOptions) *HealthMonitorsListStateChangesPager {
	return &HealthMonitorsListStateChangesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listStateChangesCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, options)
		},
		advancer: func(ctx context.Context, resp HealthMonitorsListStateChangesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.HealthMonitorStateChangeList.NextLink)
		},
	}
}

// listStateChangesCreateRequest creates the ListStateChanges request.
func (client *HealthMonitorsClient) listStateChangesCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsListStateChangesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if monitorID == "" {
		return nil, errors.New("parameter monitorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorId}", url.PathEscape(monitorID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.StartTimestampUTC != nil {
		reqQP.Set("startTimestampUtc", options.StartTimestampUTC.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimestampUTC != nil {
		reqQP.Set("endTimestampUtc", options.EndTimestampUTC.Format(time.RFC3339Nano))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listStateChangesHandleResponse handles the ListStateChanges response.
func (client *HealthMonitorsClient) listStateChangesHandleResponse(resp *http.Response) (HealthMonitorsListStateChangesResponse, error) {
	result := HealthMonitorsListStateChangesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitorStateChangeList); err != nil {
		return HealthMonitorsListStateChangesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listStateChangesHandleError handles the ListStateChanges error response.
func (client *HealthMonitorsClient) listStateChangesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}