//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the DigitalTwins group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Check if a DigitalTwinsInstance name is available.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Location of DigitalTwinsInstance.
// digitalTwinsInstanceCheckName - Set the name parameter in the DigitalTwinsInstanceCheckName structure to the name of the
// DigitalTwinsInstance to check.
// options - ClientCheckNameAvailabilityOptions contains the optional parameters for the Client.CheckNameAvailability method.
func (client *Client) CheckNameAvailability(ctx context.Context, location string, digitalTwinsInstanceCheckName CheckNameRequest, options *ClientCheckNameAvailabilityOptions) (ClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, digitalTwinsInstanceCheckName, options)
	if err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *Client) checkNameAvailabilityCreateRequest(ctx context.Context, location string, digitalTwinsInstanceCheckName CheckNameRequest, options *ClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DigitalTwins/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, digitalTwinsInstanceCheckName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *Client) checkNameAvailabilityHandleResponse(resp *http.Response) (ClientCheckNameAvailabilityResponse, error) {
	result := ClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create or update the metadata of a DigitalTwinsInstance. The usual pattern to modify a property is
// to retrieve the DigitalTwinsInstance and security metadata, and then combine them with the modified
// values in a new body to update the DigitalTwinsInstance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// digitalTwinsCreate - The DigitalTwinsInstance and security metadata.
// options - ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
func (client *Client) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsCreate Description, options *ClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, digitalTwinsCreate, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the metadata of a DigitalTwinsInstance. The usual pattern to modify a property is to
// retrieve the DigitalTwinsInstance and security metadata, and then combine them with the modified
// values in a new body to update the DigitalTwinsInstance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsCreate Description, options *ClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, digitalTwinsCreate, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *Client) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsCreate Description, options *ClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, digitalTwinsCreate)
}

// BeginDelete - Delete a DigitalTwinsInstance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// options - ClientBeginDeleteOptions contains the optional parameters for the Client.BeginDelete method.
func (client *Client) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *ClientBeginDeleteOptions) (*armruntime.Poller[ClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a DigitalTwinsInstance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *ClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *Client) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get DigitalTwinsInstances resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// options - ClientGetOptions contains the optional parameters for the Client.Get method.
func (client *Client) Get(ctx context.Context, resourceGroupName string, resourceName string, options *ClientGetOptions) (ClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *Client) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Client) getHandleResponse(resp *http.Response) (ClientGetResponse, error) {
	result := ClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Description); err != nil {
		return ClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all the DigitalTwinsInstances in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ClientListOptions contains the optional parameters for the Client.List method.
func (client *Client) NewListPager(options *ClientListOptions) *runtime.Pager[ClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ClientListResponse]{
		More: func(page ClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListResponse) (ClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *Client) listCreateRequest(ctx context.Context, options *ClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DigitalTwins/digitalTwinsInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *Client) listHandleResponse(resp *http.Response) (ClientListResponse, error) {
	result := ClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DescriptionListResult); err != nil {
		return ClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all the DigitalTwinsInstances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// options - ClientListByResourceGroupOptions contains the optional parameters for the Client.ListByResourceGroup method.
func (client *Client) NewListByResourceGroupPager(resourceGroupName string, options *ClientListByResourceGroupOptions) *runtime.Pager[ClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[ClientListByResourceGroupResponse]{
		More: func(page ClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListByResourceGroupResponse) (ClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *Client) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *Client) listByResourceGroupHandleResponse(resp *http.Response) (ClientListByResourceGroupResponse, error) {
	result := ClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DescriptionListResult); err != nil {
		return ClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update metadata of DigitalTwinsInstance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// digitalTwinsPatchDescription - The DigitalTwinsInstance and security metadata.
// options - ClientBeginUpdateOptions contains the optional parameters for the Client.BeginUpdate method.
func (client *Client) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsPatchDescription PatchDescription, options *ClientBeginUpdateOptions) (*armruntime.Poller[ClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, digitalTwinsPatchDescription, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update metadata of DigitalTwinsInstance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) update(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsPatchDescription PatchDescription, options *ClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, digitalTwinsPatchDescription, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *Client) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, digitalTwinsPatchDescription PatchDescription, options *ClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, digitalTwinsPatchDescription)
}
