//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListContainerAppsBySubscription.json
func ExampleContainerAppsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListContainerAppsByResourceGroup.json
func ExampleContainerAppsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetContainerApp.json
func ExampleContainerAppsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ContainerAppsClientGetResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateContainerApp.json
func ExampleContainerAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.ContainerApp{
			Kind:     to.StringPtr("<kind>"),
			Location: to.StringPtr("<location>"),
			Properties: &armappservice.ContainerAppProperties{
				Configuration: &armappservice.Configuration{
					Ingress: &armappservice.Ingress{
						External:   to.BoolPtr(true),
						TargetPort: to.Int32Ptr(3000),
					},
				},
				KubeEnvironmentID: to.StringPtr("<kube-environment-id>"),
				Template: &armappservice.Template{
					Containers: []*armappservice.Container{
						{
							Name:  to.StringPtr("<name>"),
							Image: to.StringPtr("<image>"),
						}},
					Dapr: &armappservice.Dapr{
						AppPort: to.Int32Ptr(3000),
						Enabled: to.BoolPtr(true),
					},
					Scale: &armappservice.Scale{
						MaxReplicas: to.Int32Ptr(5),
						MinReplicas: to.Int32Ptr(1),
						Rules: []*armappservice.ScaleRule{
							{
								Name: to.StringPtr("<name>"),
								Custom: &armappservice.CustomScaleRule{
									Type: to.StringPtr("<type>"),
									Metadata: map[string]*string{
										"concurrentRequests": to.StringPtr("50"),
									},
								},
							}},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ContainerAppsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteContainerApp.json
func ExampleContainerAppsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListContainerAppSecrets.json
func ExampleContainerAppsClient_ListSecrets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	res, err := client.ListSecrets(ctx,
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ContainerAppsClientListSecretsResult)
}