package eas

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

// GetResourcePods invokes the eas.GetResourcePods API synchronously
// api document: https://help.aliyun.com/api/eas/getresourcepods.html
func (client *Client) GetResourcePods(request *GetResourcePodsRequest) (response *GetResourcePodsResponse, err error) {
	response = CreateGetResourcePodsResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourcePodsWithChan invokes the eas.GetResourcePods API asynchronously
// api document: https://help.aliyun.com/api/eas/getresourcepods.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourcePodsWithChan(request *GetResourcePodsRequest) (<-chan *GetResourcePodsResponse, <-chan error) {
	responseChan := make(chan *GetResourcePodsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourcePods(request)
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

// GetResourcePodsWithCallback invokes the eas.GetResourcePods API asynchronously
// api document: https://help.aliyun.com/api/eas/getresourcepods.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourcePodsWithCallback(request *GetResourcePodsRequest, callback func(response *GetResourcePodsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourcePodsResponse
		var err error
		defer close(result)
		response, err = client.GetResourcePods(request)
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

// GetResourcePodsRequest is the request struct for api GetResourcePods
type GetResourcePodsRequest struct {
	*requests.RoaRequest
	ClusterId    string `position:"Path" name:"cluster_id"`
	InstanceName string `position:"Path" name:"instance_name"`
	ResourceName string `position:"Path" name:"resource_name"`
}

// GetResourcePodsResponse is the response struct for api GetResourcePods
type GetResourcePodsResponse struct {
	*responses.BaseResponse
}

// CreateGetResourcePodsRequest creates a request to invoke GetResourcePods API
func CreateGetResourcePodsRequest() (request *GetResourcePodsRequest) {
	request = &GetResourcePodsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "GetResourcePods", "/api/resources/[cluster_id]/[resource_name]/instances/[instance_name]/pods", "", "")
	request.Method = requests.GET
	return
}

// CreateGetResourcePodsResponse creates a response to parse from GetResourcePods response
func CreateGetResourcePodsResponse() (response *GetResourcePodsResponse) {
	response = &GetResourcePodsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
