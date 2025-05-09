/*
Service Manager

Service Manager provides REST APIs that are responsible for the creation and consumption of service instances in any connected runtime environment.   Use the Service Manager APIs to perform various operations related to your platforms, service brokers, service instances, and service bindings.  Get service plans and service offerings associated with your environment.    #### Platforms   Platforms are OSBAPI-enabled software systems on which applications and services are hosted.   With the Service Manager, you can now register your platform and enable it to consume the SAP BTP services from your native environment.   This registration results in a returned set of credentials that are needed to deploy the Service Manager agent.     #### Service Brokers   Service brokers act as brokers between the Service Manager and a platform’s marketplace to advertise catalogues of service offerings and service plans.  They also receive and process the requests from the marketplace to provision, bind, unbind, and deprovision these offerings and plans.    #### Service Instances   Service instances are instantiations of service plans that make the functionality of those service plans available for consumption.    #### Service Bindings   Service bindings provide access details to existing service instances.  The access details are part of the service bindings' ‘credentials’ property, and typically include access URLs and credentials.    #### Service Plans   Service plans represent sets of capabilities provided by a service offering.  For example, database service offerings provide different plans for different database versions or sizes, while the Service Manager plans offer different data access levels.    #### Service Offerings   Service offerings are advertisements of the services that are supported by a service broker.  For example, software that you can consume in the subaccount.  Service offerings are related to one or more service plans.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type OperationsAPI interface {

	/*
	GetSingleOperation Get operation status

	Check the status of the asynchronous operation you created after the POST, DELETE, or PATCH (where applicable) calls of the following API groups: Service Instances, and Service Bindings.</br> The APIs provide a boolean request parameter 'async' whose default value is TRUE.</br>Use the parameter to generate an asynchronous operation, if needed. </br>In the **Location** response headers of the APIs, you get the path parameters that are required for this API.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceType The type of the SAP Service Manager resource.
	@param resourceID The ID of the previously created entity of the specified resource type.
	@param operationID The ID of the operation for which to get status.
	@return ApiGetSingleOperationRequest
	*/
	GetSingleOperation(ctx context.Context, resourceType string, resourceID string, operationID string) ApiGetSingleOperationRequest

	// GetSingleOperationExecute executes the request
	//  @return OperationResponseObject
	GetSingleOperationExecute(r ApiGetSingleOperationRequest) (*OperationResponseObject, *http.Response, error)
}

// OperationsAPIService OperationsAPI service
type OperationsAPIService service

type ApiGetSingleOperationRequest struct {
	ctx context.Context
	ApiService OperationsAPI
	resourceType string
	resourceID string
	operationID string
}

func (r ApiGetSingleOperationRequest) Execute() (*OperationResponseObject, *http.Response, error) {
	return r.ApiService.GetSingleOperationExecute(r)
}

/*
GetSingleOperation Get operation status

Check the status of the asynchronous operation you created after the POST, DELETE, or PATCH (where applicable) calls of the following API groups: Service Instances, and Service Bindings.</br> The APIs provide a boolean request parameter 'async' whose default value is TRUE.</br>Use the parameter to generate an asynchronous operation, if needed. </br>In the **Location** response headers of the APIs, you get the path parameters that are required for this API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceType The type of the SAP Service Manager resource.
 @param resourceID The ID of the previously created entity of the specified resource type.
 @param operationID The ID of the operation for which to get status.
 @return ApiGetSingleOperationRequest
*/
func (a *OperationsAPIService) GetSingleOperation(ctx context.Context, resourceType string, resourceID string, operationID string) ApiGetSingleOperationRequest {
	return ApiGetSingleOperationRequest{
		ApiService: a,
		ctx: ctx,
		resourceType: resourceType,
		resourceID: resourceID,
		operationID: operationID,
	}
}

// Execute executes the request
//  @return OperationResponseObject
func (a *OperationsAPIService) GetSingleOperationExecute(r ApiGetSingleOperationRequest) (*OperationResponseObject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OperationResponseObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.GetSingleOperation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{resourceType}/{resourceID}/operations/{operationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceType"+"}", url.PathEscape(parameterValueToString(r.resourceType, "resourceType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceID"+"}", url.PathEscape(parameterValueToString(r.resourceID, "resourceID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"operationID"+"}", url.PathEscape(parameterValueToString(r.operationID, "operationID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
