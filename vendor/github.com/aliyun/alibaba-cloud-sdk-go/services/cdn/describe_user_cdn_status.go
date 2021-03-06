package cdn

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

// DescribeUserCdnStatus invokes the cdn.DescribeUserCdnStatus API synchronously
// api document: https://help.aliyun.com/api/cdn/describeusercdnstatus.html
func (client *Client) DescribeUserCdnStatus(request *DescribeUserCdnStatusRequest) (response *DescribeUserCdnStatusResponse, err error) {
	response = CreateDescribeUserCdnStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserCdnStatusWithChan invokes the cdn.DescribeUserCdnStatus API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeusercdnstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserCdnStatusWithChan(request *DescribeUserCdnStatusRequest) (<-chan *DescribeUserCdnStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeUserCdnStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserCdnStatus(request)
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

// DescribeUserCdnStatusWithCallback invokes the cdn.DescribeUserCdnStatus API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeusercdnstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserCdnStatusWithCallback(request *DescribeUserCdnStatusRequest, callback func(response *DescribeUserCdnStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserCdnStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserCdnStatus(request)
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

// DescribeUserCdnStatusRequest is the request struct for api DescribeUserCdnStatus
type DescribeUserCdnStatusRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeUserCdnStatusResponse is the response struct for api DescribeUserCdnStatus
type DescribeUserCdnStatusResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Enabled       bool   `json:"Enabled" xml:"Enabled"`
	OnService     bool   `json:"OnService" xml:"OnService"`
	InDebt        bool   `json:"InDebt" xml:"InDebt"`
	InDebtOverdue bool   `json:"InDebtOverdue" xml:"InDebtOverdue"`
}

// CreateDescribeUserCdnStatusRequest creates a request to invoke DescribeUserCdnStatus API
func CreateDescribeUserCdnStatusRequest() (request *DescribeUserCdnStatusRequest) {
	request = &DescribeUserCdnStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeUserCdnStatus", "", "")
	return
}

// CreateDescribeUserCdnStatusResponse creates a response to parse from DescribeUserCdnStatus response
func CreateDescribeUserCdnStatusResponse() (response *DescribeUserCdnStatusResponse) {
	response = &DescribeUserCdnStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
