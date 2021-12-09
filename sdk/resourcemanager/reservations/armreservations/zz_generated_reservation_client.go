//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

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
	"strconv"
	"strings"
)

// ReservationClient contains the methods for the Reservation group.
// Don't use this type directly, use NewReservationClient() instead.
type ReservationClient struct {
	ep string
	pl runtime.Pipeline
}

// NewReservationClient creates a new instance of ReservationClient with the specified values.
func NewReservationClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ReservationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReservationClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginAvailableScopes - Get Available Scopes for Reservation.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) BeginAvailableScopes(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest, options *ReservationBeginAvailableScopesOptions) (ReservationAvailableScopesPollerResponse, error) {
	resp, err := client.availableScopes(ctx, reservationOrderID, reservationID, body, options)
	if err != nil {
		return ReservationAvailableScopesPollerResponse{}, err
	}
	result := ReservationAvailableScopesPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.AvailableScopes", "", resp, client.pl, client.availableScopesHandleError)
	if err != nil {
		return ReservationAvailableScopesPollerResponse{}, err
	}
	result.Poller = &ReservationAvailableScopesPoller{
		pt: pt,
	}
	return result, nil
}

// AvailableScopes - Get Available Scopes for Reservation.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) availableScopes(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest, options *ReservationBeginAvailableScopesOptions) (*http.Response, error) {
	req, err := client.availableScopesCreateRequest(ctx, reservationOrderID, reservationID, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.availableScopesHandleError(resp)
	}
	return resp, nil
}

// availableScopesCreateRequest creates the AvailableScopes request.
func (client *ReservationClient) availableScopesCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest, options *ReservationBeginAvailableScopesOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/availableScopes"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// availableScopesHandleError handles the AvailableScopes error response.
func (client *ReservationClient) availableScopesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get specific Reservation details.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) Get(ctx context.Context, reservationID string, reservationOrderID string, options *ReservationGetOptions) (ReservationGetResponse, error) {
	req, err := client.getCreateRequest(ctx, reservationID, reservationOrderID, options)
	if err != nil {
		return ReservationGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReservationGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReservationGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReservationClient) getCreateRequest(ctx context.Context, reservationID string, reservationOrderID string, options *ReservationGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}"
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReservationClient) getHandleResponse(resp *http.Response) (ReservationGetResponse, error) {
	result := ReservationGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationResponse); err != nil {
		return ReservationGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ReservationClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List Reservations within a single ReservationOrder.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) List(reservationOrderID string, options *ReservationListOptions) *ReservationListPager {
	return &ReservationListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, reservationOrderID, options)
		},
		advancer: func(ctx context.Context, resp ReservationListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReservationClient) listCreateRequest(ctx context.Context, reservationOrderID string, options *ReservationListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationClient) listHandleResponse(resp *http.Response) (ReservationListResponse, error) {
	result := ReservationListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationList); err != nil {
		return ReservationListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReservationClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAll - List the reservations and the roll up counts of reservations group by provisioning states that the user has access to in the current tenant.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ReservationClient) ListAll(options *ReservationListAllOptions) *ReservationListAllPager {
	return &ReservationListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ReservationListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationsListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *ReservationClient) listAllCreateRequest(ctx context.Context, options *ReservationListAllOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.RefreshSummary != nil {
		reqQP.Set("refreshSummary", *options.RefreshSummary)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", strconv.FormatFloat(float64(*options.Skiptoken), 'f', -1, 32))
	}
	if options != nil && options.SelectedState != nil {
		reqQP.Set("selectedState", *options.SelectedState)
	}
	if options != nil && options.Take != nil {
		reqQP.Set("take", strconv.FormatFloat(float64(*options.Take), 'f', -1, 32))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *ReservationClient) listAllHandleResponse(resp *http.Response) (ReservationListAllResponse, error) {
	result := ReservationListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationsListResult); err != nil {
		return ReservationListAllResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *ReservationClient) listAllHandleError(resp *http.Response) error {
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

// ListRevisions - List of all the revisions for the Reservation.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) ListRevisions(reservationID string, reservationOrderID string, options *ReservationListRevisionsOptions) *ReservationListRevisionsPager {
	return &ReservationListRevisionsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listRevisionsCreateRequest(ctx, reservationID, reservationOrderID, options)
		},
		advancer: func(ctx context.Context, resp ReservationListRevisionsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationList.NextLink)
		},
	}
}

// listRevisionsCreateRequest creates the ListRevisions request.
func (client *ReservationClient) listRevisionsCreateRequest(ctx context.Context, reservationID string, reservationOrderID string, options *ReservationListRevisionsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/revisions"
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRevisionsHandleResponse handles the ListRevisions response.
func (client *ReservationClient) listRevisionsHandleResponse(resp *http.Response) (ReservationListRevisionsResponse, error) {
	result := ReservationListRevisionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationList); err != nil {
		return ReservationListRevisionsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listRevisionsHandleError handles the ListRevisions error response.
func (client *ReservationClient) listRevisionsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginMerge - Merge the specified Reservations into a new Reservation. The two Reservations being merged must have same properties.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) BeginMerge(ctx context.Context, reservationOrderID string, body MergeRequest, options *ReservationBeginMergeOptions) (ReservationMergePollerResponse, error) {
	resp, err := client.merge(ctx, reservationOrderID, body, options)
	if err != nil {
		return ReservationMergePollerResponse{}, err
	}
	result := ReservationMergePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.Merge", "location", resp, client.pl, client.mergeHandleError)
	if err != nil {
		return ReservationMergePollerResponse{}, err
	}
	result.Poller = &ReservationMergePoller{
		pt: pt,
	}
	return result, nil
}

// Merge - Merge the specified Reservations into a new Reservation. The two Reservations being merged must have same properties.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) merge(ctx context.Context, reservationOrderID string, body MergeRequest, options *ReservationBeginMergeOptions) (*http.Response, error) {
	req, err := client.mergeCreateRequest(ctx, reservationOrderID, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.mergeHandleError(resp)
	}
	return resp, nil
}

// mergeCreateRequest creates the Merge request.
func (client *ReservationClient) mergeCreateRequest(ctx context.Context, reservationOrderID string, body MergeRequest, options *ReservationBeginMergeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/merge"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// mergeHandleError handles the Merge error response.
func (client *ReservationClient) mergeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginSplit - Split a Reservation into two Reservations with specified quantity distribution.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) BeginSplit(ctx context.Context, reservationOrderID string, body SplitRequest, options *ReservationBeginSplitOptions) (ReservationSplitPollerResponse, error) {
	resp, err := client.split(ctx, reservationOrderID, body, options)
	if err != nil {
		return ReservationSplitPollerResponse{}, err
	}
	result := ReservationSplitPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.Split", "location", resp, client.pl, client.splitHandleError)
	if err != nil {
		return ReservationSplitPollerResponse{}, err
	}
	result.Poller = &ReservationSplitPoller{
		pt: pt,
	}
	return result, nil
}

// Split - Split a Reservation into two Reservations with specified quantity distribution.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) split(ctx context.Context, reservationOrderID string, body SplitRequest, options *ReservationBeginSplitOptions) (*http.Response, error) {
	req, err := client.splitCreateRequest(ctx, reservationOrderID, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.splitHandleError(resp)
	}
	return resp, nil
}

// splitCreateRequest creates the Split request.
func (client *ReservationClient) splitCreateRequest(ctx context.Context, reservationOrderID string, body SplitRequest, options *ReservationBeginSplitOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/split"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// splitHandleError handles the Split error response.
func (client *ReservationClient) splitHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates the applied scopes of the Reservation.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) BeginUpdate(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch, options *ReservationBeginUpdateOptions) (ReservationUpdatePollerResponse, error) {
	resp, err := client.update(ctx, reservationOrderID, reservationID, parameters, options)
	if err != nil {
		return ReservationUpdatePollerResponse{}, err
	}
	result := ReservationUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ReservationUpdatePollerResponse{}, err
	}
	result.Poller = &ReservationUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates the applied scopes of the Reservation.
// If the operation fails it returns the *Error error type.
func (client *ReservationClient) update(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch, options *ReservationBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, reservationOrderID, reservationID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReservationClient) updateCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch, options *ReservationBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *ReservationClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}