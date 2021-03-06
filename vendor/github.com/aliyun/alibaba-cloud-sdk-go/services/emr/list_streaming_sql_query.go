package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListStreamingSqlQuery invokes the emr.ListStreamingSqlQuery API synchronously
// api document: https://help.aliyun.com/api/emr/liststreamingsqlquery.html
func (client *Client) ListStreamingSqlQuery(request *ListStreamingSqlQueryRequest) (response *ListStreamingSqlQueryResponse, err error) {
	response = CreateListStreamingSqlQueryResponse()
	err = client.DoAction(request, response)
	return
}

// ListStreamingSqlQueryWithChan invokes the emr.ListStreamingSqlQuery API asynchronously
// api document: https://help.aliyun.com/api/emr/liststreamingsqlquery.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStreamingSqlQueryWithChan(request *ListStreamingSqlQueryRequest) (<-chan *ListStreamingSqlQueryResponse, <-chan error) {
	responseChan := make(chan *ListStreamingSqlQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListStreamingSqlQuery(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListStreamingSqlQueryWithCallback invokes the emr.ListStreamingSqlQuery API asynchronously
// api document: https://help.aliyun.com/api/emr/liststreamingsqlquery.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStreamingSqlQueryWithCallback(request *ListStreamingSqlQueryRequest, callback func(response *ListStreamingSqlQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListStreamingSqlQueryResponse
		var err error
		defer close(result)
		response, err = client.ListStreamingSqlQuery(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListStreamingSqlQueryRequest is the request struct for api ListStreamingSqlQuery
type ListStreamingSqlQueryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
}

// ListStreamingSqlQueryResponse is the response struct for api ListStreamingSqlQuery
type ListStreamingSqlQueryResponse struct {
	*responses.BaseResponse
	RequestId  string                       `json:"RequestId" xml:"RequestId"`
	PageNumber int                          `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                          `json:"PageSize" xml:"PageSize"`
	TotalCount int                          `json:"TotalCount" xml:"TotalCount"`
	Items      ItemsInListStreamingSqlQuery `json:"Items" xml:"Items"`
}

// CreateListStreamingSqlQueryRequest creates a request to invoke ListStreamingSqlQuery API
func CreateListStreamingSqlQueryRequest() (request *ListStreamingSqlQueryRequest) {
	request = &ListStreamingSqlQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListStreamingSqlQuery", "emr", "openAPI")
	return
}

// CreateListStreamingSqlQueryResponse creates a response to parse from ListStreamingSqlQuery response
func CreateListStreamingSqlQueryResponse() (response *ListStreamingSqlQueryResponse) {
	response = &ListStreamingSqlQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
