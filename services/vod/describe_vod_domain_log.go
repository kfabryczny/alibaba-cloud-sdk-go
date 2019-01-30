package vod

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

// DescribeVodDomainLog invokes the vod.DescribeVodDomainLog API synchronously
// api document: https://help.aliyun.com/api/vod/describevoddomainlog.html
func (client *Client) DescribeVodDomainLog(request *DescribeVodDomainLogRequest) (response *DescribeVodDomainLogResponse, err error) {
	response = CreateDescribeVodDomainLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainLogWithChan invokes the vod.DescribeVodDomainLog API asynchronously
// api document: https://help.aliyun.com/api/vod/describevoddomainlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVodDomainLogWithChan(request *DescribeVodDomainLogRequest) (<-chan *DescribeVodDomainLogResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainLog(request)
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

// DescribeVodDomainLogWithCallback invokes the vod.DescribeVodDomainLog API asynchronously
// api document: https://help.aliyun.com/api/vod/describevoddomainlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVodDomainLogWithCallback(request *DescribeVodDomainLogRequest, callback func(response *DescribeVodDomainLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainLog(request)
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

// DescribeVodDomainLogRequest is the request struct for api DescribeVodDomainLog
type DescribeVodDomainLogRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVodDomainLogResponse is the response struct for api DescribeVodDomainLog
type DescribeVodDomainLogResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	DomainLogDetails DomainLogDetails `json:"DomainLogDetails" xml:"DomainLogDetails"`
}

// CreateDescribeVodDomainLogRequest creates a request to invoke DescribeVodDomainLog API
func CreateDescribeVodDomainLogRequest() (request *DescribeVodDomainLogRequest) {
	request = &DescribeVodDomainLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainLog", "vod", "openAPI")
	return
}

// CreateDescribeVodDomainLogResponse creates a response to parse from DescribeVodDomainLog response
func CreateDescribeVodDomainLogResponse() (response *DescribeVodDomainLogResponse) {
	response = &DescribeVodDomainLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
