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

// DeleteTranscodeTemplateGroup invokes the vod.DeleteTranscodeTemplateGroup API synchronously
// api document: https://help.aliyun.com/api/vod/deletetranscodetemplategroup.html
func (client *Client) DeleteTranscodeTemplateGroup(request *DeleteTranscodeTemplateGroupRequest) (response *DeleteTranscodeTemplateGroupResponse, err error) {
	response = CreateDeleteTranscodeTemplateGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTranscodeTemplateGroupWithChan invokes the vod.DeleteTranscodeTemplateGroup API asynchronously
// api document: https://help.aliyun.com/api/vod/deletetranscodetemplategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTranscodeTemplateGroupWithChan(request *DeleteTranscodeTemplateGroupRequest) (<-chan *DeleteTranscodeTemplateGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteTranscodeTemplateGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTranscodeTemplateGroup(request)
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

// DeleteTranscodeTemplateGroupWithCallback invokes the vod.DeleteTranscodeTemplateGroup API asynchronously
// api document: https://help.aliyun.com/api/vod/deletetranscodetemplategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTranscodeTemplateGroupWithCallback(request *DeleteTranscodeTemplateGroupRequest, callback func(response *DeleteTranscodeTemplateGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTranscodeTemplateGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteTranscodeTemplateGroup(request)
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

// DeleteTranscodeTemplateGroupRequest is the request struct for api DeleteTranscodeTemplateGroup
type DeleteTranscodeTemplateGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	TranscodeTemplateGroupId string           `position:"Query" name:"TranscodeTemplateGroupId"`
}

// DeleteTranscodeTemplateGroupResponse is the response struct for api DeleteTranscodeTemplateGroup
type DeleteTranscodeTemplateGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTranscodeTemplateGroupRequest creates a request to invoke DeleteTranscodeTemplateGroup API
func CreateDeleteTranscodeTemplateGroupRequest() (request *DeleteTranscodeTemplateGroupRequest) {
	request = &DeleteTranscodeTemplateGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteTranscodeTemplateGroup", "vod", "openAPI")
	return
}

// CreateDeleteTranscodeTemplateGroupResponse creates a response to parse from DeleteTranscodeTemplateGroup response
func CreateDeleteTranscodeTemplateGroupResponse() (response *DeleteTranscodeTemplateGroupResponse) {
	response = &DeleteTranscodeTemplateGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
