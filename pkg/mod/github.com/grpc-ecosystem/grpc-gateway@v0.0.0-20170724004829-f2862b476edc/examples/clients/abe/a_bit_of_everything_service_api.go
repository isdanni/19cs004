/* 
 * examples/examplepb/a_bit_of_everything.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: version not set
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package abe

import (
	"net/url"
	"strings"
	"time"
	"encoding/json"
	"fmt"
)

type ABitOfEverythingServiceApi struct {
	Configuration *Configuration
}

func NewABitOfEverythingServiceApi() *ABitOfEverythingServiceApi {
	configuration := NewConfiguration()
	return &ABitOfEverythingServiceApi{
		Configuration: configuration,
	}
}

func NewABitOfEverythingServiceApiWithBasePath(basePath string) *ABitOfEverythingServiceApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ABitOfEverythingServiceApi{
		Configuration: configuration,
	}
}

/**
 * 
 *
 * @param floatValue 
 * @param doubleValue 
 * @param int64Value 
 * @param uint64Value 
 * @param int32Value 
 * @param fixed64Value 
 * @param fixed32Value 
 * @param boolValue 
 * @param stringValue 
 * @param uint32Value 
 * @param sfixed32Value 
 * @param sfixed64Value 
 * @param sint32Value 
 * @param sint64Value 
 * @param nonConventionalNameValue 
 * @return *ExamplepbABitOfEverything
 */
func (a ABitOfEverythingServiceApi) Create(floatValue float32, doubleValue float64, int64Value string, uint64Value string, int32Value int32, fixed64Value string, fixed32Value int64, boolValue bool, stringValue string, uint32Value int64, sfixed32Value int32, sfixed64Value string, sint32Value int32, sint64Value string, nonConventionalNameValue string) (*ExamplepbABitOfEverything, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything/{float_value}/{double_value}/{int64_value}/separator/{uint64_value}/{int32_value}/{fixed64_value}/{fixed32_value}/{bool_value}/{string_value}/{uint32_value}/{sfixed32_value}/{sfixed64_value}/{sint32_value}/{sint64_value}/{nonConventionalNameValue}"
	localVarPath = strings.Replace(localVarPath, "{"+"float_value"+"}", fmt.Sprintf("%v", floatValue), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"double_value"+"}", fmt.Sprintf("%v", doubleValue), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"int64_value"+"}", fmt.Sprintf("%v", int64Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uint64_value"+"}", fmt.Sprintf("%v", uint64Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"int32_value"+"}", fmt.Sprintf("%v", int32Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fixed64_value"+"}", fmt.Sprintf("%v", fixed64Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fixed32_value"+"}", fmt.Sprintf("%v", fixed32Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bool_value"+"}", fmt.Sprintf("%v", boolValue), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"string_value"+"}", fmt.Sprintf("%v", stringValue), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uint32_value"+"}", fmt.Sprintf("%v", uint32Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sfixed32_value"+"}", fmt.Sprintf("%v", sfixed32Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sfixed64_value"+"}", fmt.Sprintf("%v", sfixed64Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sint32_value"+"}", fmt.Sprintf("%v", sint32Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sint64_value"+"}", fmt.Sprintf("%v", sint64Value), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nonConventionalNameValue"+"}", fmt.Sprintf("%v", nonConventionalNameValue), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbABitOfEverything)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Create", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param body 
 * @return *ExamplepbABitOfEverything
 */
func (a ABitOfEverythingServiceApi) CreateBody(body ExamplepbABitOfEverything) (*ExamplepbABitOfEverything, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(ExamplepbABitOfEverything)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreateBody", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param singleNestedName 
 * @param body 
 * @return *ExamplepbABitOfEverything
 */
func (a ABitOfEverythingServiceApi) DeepPathEcho(singleNestedName string, body ExamplepbABitOfEverything) (*ExamplepbABitOfEverything, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything/{single_nested.name}"
	localVarPath = strings.Replace(localVarPath, "{"+"single_nested.name"+"}", fmt.Sprintf("%v", singleNestedName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(ExamplepbABitOfEverything)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeepPathEcho", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param uuid 
 * @return *ProtobufEmpty
 */
func (a ABitOfEverythingServiceApi) Delete(uuid string) (*ProtobufEmpty, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", fmt.Sprintf("%v", uuid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ProtobufEmpty)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Delete", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param value 
 * @return *SubStringMessage
 */
func (a ABitOfEverythingServiceApi) Echo(value string) (*SubStringMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything/echo/{value}"
	localVarPath = strings.Replace(localVarPath, "{"+"value"+"}", fmt.Sprintf("%v", value), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SubStringMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param value 
 * @return *SubStringMessage
 */
func (a ABitOfEverythingServiceApi) Echo_1(value string) (*SubStringMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v2/example/echo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("value", a.Configuration.APIClient.ParameterToString(value, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(SubStringMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo_0", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param body 
 * @return *SubStringMessage
 */
func (a ABitOfEverythingServiceApi) Echo_2(body string) (*SubStringMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v2/example/echo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(SubStringMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo_1", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param uuid 
 * @param singleNestedName name is nested field.
 * @param singleNestedAmount 
 * @param singleNestedOk  - FALSE: FALSE is false.  - TRUE: TRUE is true.
 * @param floatValue 
 * @param doubleValue 
 * @param int64Value 
 * @param uint64Value 
 * @param int32Value 
 * @param fixed64Value 
 * @param fixed32Value 
 * @param boolValue 
 * @param stringValue 
 * @param uint32Value TODO(yugui) add bytes_value.
 * @param enumValue  - ZERO: ZERO means 0  - ONE: ONE means 1
 * @param sfixed32Value 
 * @param sfixed64Value 
 * @param sint32Value 
 * @param sint64Value 
 * @param repeatedStringValue 
 * @param oneofString 
 * @param nonConventionalNameValue 
 * @param timestampValue 
 * @param repeatedEnumValue repeated enum value. it is comma-separated in query.   - ZERO: ZERO means 0  - ONE: ONE means 1
 * @return *ProtobufEmpty
 */
func (a ABitOfEverythingServiceApi) GetQuery(uuid string, singleNestedName string, singleNestedAmount int64, singleNestedOk string, floatValue float32, doubleValue float64, int64Value string, uint64Value string, int32Value int32, fixed64Value string, fixed32Value int64, boolValue bool, stringValue string, uint32Value int64, enumValue string, sfixed32Value int32, sfixed64Value string, sint32Value int32, sint64Value string, repeatedStringValue []string, oneofString string, nonConventionalNameValue string, timestampValue time.Time, repeatedEnumValue []string) (*ProtobufEmpty, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything/query/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", fmt.Sprintf("%v", uuid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("single_nested.name", a.Configuration.APIClient.ParameterToString(singleNestedName, ""))
	localVarQueryParams.Add("single_nested.amount", a.Configuration.APIClient.ParameterToString(singleNestedAmount, ""))
	localVarQueryParams.Add("single_nested.ok", a.Configuration.APIClient.ParameterToString(singleNestedOk, ""))
	localVarQueryParams.Add("float_value", a.Configuration.APIClient.ParameterToString(floatValue, ""))
	localVarQueryParams.Add("double_value", a.Configuration.APIClient.ParameterToString(doubleValue, ""))
	localVarQueryParams.Add("int64_value", a.Configuration.APIClient.ParameterToString(int64Value, ""))
	localVarQueryParams.Add("uint64_value", a.Configuration.APIClient.ParameterToString(uint64Value, ""))
	localVarQueryParams.Add("int32_value", a.Configuration.APIClient.ParameterToString(int32Value, ""))
	localVarQueryParams.Add("fixed64_value", a.Configuration.APIClient.ParameterToString(fixed64Value, ""))
	localVarQueryParams.Add("fixed32_value", a.Configuration.APIClient.ParameterToString(fixed32Value, ""))
	localVarQueryParams.Add("bool_value", a.Configuration.APIClient.ParameterToString(boolValue, ""))
	localVarQueryParams.Add("string_value", a.Configuration.APIClient.ParameterToString(stringValue, ""))
	localVarQueryParams.Add("uint32_value", a.Configuration.APIClient.ParameterToString(uint32Value, ""))
	localVarQueryParams.Add("enum_value", a.Configuration.APIClient.ParameterToString(enumValue, ""))
	localVarQueryParams.Add("sfixed32_value", a.Configuration.APIClient.ParameterToString(sfixed32Value, ""))
	localVarQueryParams.Add("sfixed64_value", a.Configuration.APIClient.ParameterToString(sfixed64Value, ""))
	localVarQueryParams.Add("sint32_value", a.Configuration.APIClient.ParameterToString(sint32Value, ""))
	localVarQueryParams.Add("sint64_value", a.Configuration.APIClient.ParameterToString(sint64Value, ""))
	var repeatedStringValueCollectionFormat = "csv"
	localVarQueryParams.Add("repeated_string_value", a.Configuration.APIClient.ParameterToString(repeatedStringValue, repeatedStringValueCollectionFormat))

	localVarQueryParams.Add("oneof_string", a.Configuration.APIClient.ParameterToString(oneofString, ""))
	localVarQueryParams.Add("nonConventionalNameValue", a.Configuration.APIClient.ParameterToString(nonConventionalNameValue, ""))
	localVarQueryParams.Add("timestamp_value", a.Configuration.APIClient.ParameterToString(timestampValue, ""))
	var repeatedEnumValueCollectionFormat = "csv"
	localVarQueryParams.Add("repeated_enum_value", a.Configuration.APIClient.ParameterToString(repeatedEnumValue, repeatedEnumValueCollectionFormat))


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ProtobufEmpty)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetQuery", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param uuid 
 * @return *ExamplepbABitOfEverything
 */
func (a ABitOfEverythingServiceApi) Lookup(uuid string) (*ExamplepbABitOfEverything, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", fmt.Sprintf("%v", uuid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbABitOfEverything)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Lookup", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @return *ProtobufEmpty
 */
func (a ABitOfEverythingServiceApi) Timeout() (*ProtobufEmpty, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v2/example/timeout"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ProtobufEmpty)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Timeout", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 *
 * @param uuid 
 * @param body 
 * @return *ProtobufEmpty
 */
func (a ABitOfEverythingServiceApi) Update(uuid string, body ExamplepbABitOfEverything) (*ProtobufEmpty, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/a_bit_of_everything/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", fmt.Sprintf("%v", uuid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(ProtobufEmpty)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Update", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

