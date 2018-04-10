package compute

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// VirtualMachineExtensionsClient is the compute Client
type VirtualMachineExtensionsClient struct {
	ManagementClient
}

// NewVirtualMachineExtensionsClient creates an instance of the VirtualMachineExtensionsClient client.
func NewVirtualMachineExtensionsClient(subscriptionID string) VirtualMachineExtensionsClient {
	return NewVirtualMachineExtensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineExtensionsClientWithBaseURI creates an instance of the VirtualMachineExtensionsClient client.
func NewVirtualMachineExtensionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionsClient {
	return VirtualMachineExtensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update the extension. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine where the extension
// should be create or updated. VMExtensionName is the name of the virtual machine extension. extensionParameters is
// parameters supplied to the Create Virtual Machine Extension operation.
func (client VirtualMachineExtensionsClient) CreateOrUpdate(resourceGroupName string, VMName string, VMExtensionName string, extensionParameters VirtualMachineExtension, cancel <-chan struct{}) (<-chan VirtualMachineExtension, <-chan error) {
	resultChan := make(chan VirtualMachineExtension, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result VirtualMachineExtension
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, VMName, VMExtensionName, extensionParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachineExtensionsClient) CreateOrUpdatePreparer(resourceGroupName string, VMName string, VMExtensionName string, extensionParameters VirtualMachineExtension, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmExtensionName":   autorest.Encode("path", VMExtensionName),
		"vmName":            autorest.Encode("path", VMName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", pathParameters),
		autorest.WithJSON(extensionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualMachineExtension, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the extension. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine where the extension
// should be deleted. VMExtensionName is the name of the virtual machine extension.
func (client VirtualMachineExtensionsClient) Delete(resourceGroupName string, VMName string, VMExtensionName string, cancel <-chan struct{}) (<-chan OperationStatusResponse, <-chan error) {
	resultChan := make(chan OperationStatusResponse, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result OperationStatusResponse
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, VMName, VMExtensionName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineExtensionsClient) DeletePreparer(resourceGroupName string, VMName string, VMExtensionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmExtensionName":   autorest.Encode("path", VMExtensionName),
		"vmName":            autorest.Encode("path", VMName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionsClient) DeleteResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get the operation to get the extension.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine containing the
// extension. VMExtensionName is the name of the virtual machine extension. expand is the expand expression to apply on
// the operation.
func (client VirtualMachineExtensionsClient) Get(resourceGroupName string, VMName string, VMExtensionName string, expand string) (result VirtualMachineExtension, err error) {
	req, err := client.GetPreparer(resourceGroupName, VMName, VMExtensionName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineExtensionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineExtensionsClient) GetPreparer(resourceGroupName string, VMName string, VMExtensionName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmExtensionName":   autorest.Encode("path", VMExtensionName),
		"vmName":            autorest.Encode("path", VMName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionsClient) GetResponder(resp *http.Response) (result VirtualMachineExtension, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
