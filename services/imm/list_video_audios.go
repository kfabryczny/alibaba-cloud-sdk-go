package imm

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

// ListVideoAudios invokes the imm.ListVideoAudios API synchronously
// api document: https://help.aliyun.com/api/imm/listvideoaudios.html
func (client *Client) ListVideoAudios(request *ListVideoAudiosRequest) (response *ListVideoAudiosResponse, err error) {
	response = CreateListVideoAudiosResponse()
	err = client.DoAction(request, response)
	return
}

// ListVideoAudiosWithChan invokes the imm.ListVideoAudios API asynchronously
// api document: https://help.aliyun.com/api/imm/listvideoaudios.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVideoAudiosWithChan(request *ListVideoAudiosRequest) (<-chan *ListVideoAudiosResponse, <-chan error) {
	responseChan := make(chan *ListVideoAudiosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVideoAudios(request)
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

// ListVideoAudiosWithCallback invokes the imm.ListVideoAudios API asynchronously
// api document: https://help.aliyun.com/api/imm/listvideoaudios.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVideoAudiosWithCallback(request *ListVideoAudiosRequest, callback func(response *ListVideoAudiosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVideoAudiosResponse
		var err error
		defer close(result)
		response, err = client.ListVideoAudios(request)
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

// ListVideoAudiosRequest is the request struct for api ListVideoAudios
type ListVideoAudiosRequest struct {
	*requests.RpcRequest
	VideoUri string `position:"Query" name:"VideoUri"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
}

// ListVideoAudiosResponse is the response struct for api ListVideoAudios
type ListVideoAudiosResponse struct {
	*responses.BaseResponse
	SetId      string       `json:"SetId" xml:"SetId"`
	VideoUri   string       `json:"VideoUri" xml:"VideoUri"`
	NextMarker string       `json:"NextMarker" xml:"NextMarker"`
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	Audios     []AudiosItem `json:"Audios" xml:"Audios"`
}

// CreateListVideoAudiosRequest creates a request to invoke ListVideoAudios API
func CreateListVideoAudiosRequest() (request *ListVideoAudiosRequest) {
	request = &ListVideoAudiosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListVideoAudios", "2017-09-06", "openAPI")
	return
}

// CreateListVideoAudiosResponse creates a response to parse from ListVideoAudios response
func CreateListVideoAudiosResponse() (response *ListVideoAudiosResponse) {
	response = &ListVideoAudiosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}