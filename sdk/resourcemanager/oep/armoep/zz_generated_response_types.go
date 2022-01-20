//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoep

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// EnergyServicesClientCreatePollerResponse contains the response from method EnergyServicesClient.Create.
type EnergyServicesClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *EnergyServicesClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l EnergyServicesClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (EnergyServicesClientCreateResponse, error) {
	respType := EnergyServicesClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.EnergyService)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a EnergyServicesClientCreatePollerResponse from the provided client and resume token.
func (l *EnergyServicesClientCreatePollerResponse) Resume(ctx context.Context, client *EnergyServicesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("EnergyServicesClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &EnergyServicesClientCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// EnergyServicesClientCreateResponse contains the response from method EnergyServicesClient.Create.
type EnergyServicesClientCreateResponse struct {
	EnergyServicesClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnergyServicesClientCreateResult contains the result from method EnergyServicesClient.Create.
type EnergyServicesClientCreateResult struct {
	EnergyService
}

// EnergyServicesClientDeletePollerResponse contains the response from method EnergyServicesClient.Delete.
type EnergyServicesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *EnergyServicesClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l EnergyServicesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (EnergyServicesClientDeleteResponse, error) {
	respType := EnergyServicesClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a EnergyServicesClientDeletePollerResponse from the provided client and resume token.
func (l *EnergyServicesClientDeletePollerResponse) Resume(ctx context.Context, client *EnergyServicesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("EnergyServicesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &EnergyServicesClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// EnergyServicesClientDeleteResponse contains the response from method EnergyServicesClient.Delete.
type EnergyServicesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnergyServicesClientGetResponse contains the response from method EnergyServicesClient.Get.
type EnergyServicesClientGetResponse struct {
	EnergyServicesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnergyServicesClientGetResult contains the result from method EnergyServicesClient.Get.
type EnergyServicesClientGetResult struct {
	EnergyService
}

// EnergyServicesClientListByResourceGroupResponse contains the response from method EnergyServicesClient.ListByResourceGroup.
type EnergyServicesClientListByResourceGroupResponse struct {
	EnergyServicesClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnergyServicesClientListByResourceGroupResult contains the result from method EnergyServicesClient.ListByResourceGroup.
type EnergyServicesClientListByResourceGroupResult struct {
	EnergyServiceList
}

// EnergyServicesClientListBySubscriptionResponse contains the response from method EnergyServicesClient.ListBySubscription.
type EnergyServicesClientListBySubscriptionResponse struct {
	EnergyServicesClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnergyServicesClientListBySubscriptionResult contains the result from method EnergyServicesClient.ListBySubscription.
type EnergyServicesClientListBySubscriptionResult struct {
	EnergyServiceList
}

// EnergyServicesClientUpdateResponse contains the response from method EnergyServicesClient.Update.
type EnergyServicesClientUpdateResponse struct {
	EnergyServicesClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EnergyServicesClientUpdateResult contains the result from method EnergyServicesClient.Update.
type EnergyServicesClientUpdateResult struct {
	EnergyService
}

// LocationsClientCheckNameAvailabilityResponse contains the response from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResponse struct {
	LocationsClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LocationsClientCheckNameAvailabilityResult contains the result from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}