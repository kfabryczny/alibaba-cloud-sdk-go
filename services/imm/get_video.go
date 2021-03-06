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

// GetVideo invokes the imm.GetVideo API synchronously
// api document: https://help.aliyun.com/api/imm/getvideo.html
func (client *Client) GetVideo(request *GetVideoRequest) (response *GetVideoResponse, err error) {
	response = CreateGetVideoResponse()
	err = client.DoAction(request, response)
	return
}

// GetVideoWithChan invokes the imm.GetVideo API asynchronously
// api document: https://help.aliyun.com/api/imm/getvideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoWithChan(request *GetVideoRequest) (<-chan *GetVideoResponse, <-chan error) {
	responseChan := make(chan *GetVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVideo(request)
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

// GetVideoWithCallback invokes the imm.GetVideo API asynchronously
// api document: https://help.aliyun.com/api/imm/getvideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoWithCallback(request *GetVideoRequest, callback func(response *GetVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVideoResponse
		var err error
		defer close(result)
		response, err = client.GetVideo(request)
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

// GetVideoRequest is the request struct for api GetVideo
type GetVideoRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	VideoUri string `position:"Query" name:"VideoUri"`
	SetId    string `position:"Query" name:"SetId"`
}

// GetVideoResponse is the response struct for api GetVideo
type GetVideoResponse struct {
	*responses.BaseResponse
	RequestId           string          `json:"RequestId" xml:"RequestId"`
	SetId               string          `json:"SetId" xml:"SetId"`
	VideoUri            string          `json:"VideoUri" xml:"VideoUri"`
	RemarksA            string          `json:"RemarksA" xml:"RemarksA"`
	RemarksB            string          `json:"RemarksB" xml:"RemarksB"`
	CreateTime          string          `json:"CreateTime" xml:"CreateTime"`
	ModifyTime          string          `json:"ModifyTime" xml:"ModifyTime"`
	VideoWidth          int             `json:"VideoWidth" xml:"VideoWidth"`
	VideoHeight         int             `json:"VideoHeight" xml:"VideoHeight"`
	VideoFormat         string          `json:"VideoFormat" xml:"VideoFormat"`
	VideoDuration       float64         `json:"VideoDuration" xml:"VideoDuration"`
	FileSize            int             `json:"FileSize" xml:"FileSize"`
	VideoFrames         int             `json:"VideoFrames" xml:"VideoFrames"`
	SourceType          string          `json:"SourceType" xml:"SourceType"`
	SourceUri           string          `json:"SourceUri" xml:"SourceUri"`
	SourcePosition      string          `json:"SourcePosition" xml:"SourcePosition"`
	ProcessStatus       string          `json:"ProcessStatus" xml:"ProcessStatus"`
	ProcessModifyTime   string          `json:"ProcessModifyTime" xml:"ProcessModifyTime"`
	VideoTagsStatus     string          `json:"VideoTagsStatus" xml:"VideoTagsStatus"`
	VideoTagsModifyTime string          `json:"VideoTagsModifyTime" xml:"VideoTagsModifyTime"`
	CelebrityStatus     string          `json:"CelebrityStatus" xml:"CelebrityStatus"`
	CelebrityModifyTime string          `json:"CelebrityModifyTime" xml:"CelebrityModifyTime"`
	ProcessFailReason   string          `json:"ProcessFailReason" xml:"ProcessFailReason"`
	VideoTagsFailReason string          `json:"VideoTagsFailReason" xml:"VideoTagsFailReason"`
	CelebrityFailReason string          `json:"CelebrityFailReason" xml:"CelebrityFailReason"`
	RemarksC            string          `json:"RemarksC" xml:"RemarksC"`
	RemarksD            string          `json:"RemarksD" xml:"RemarksD"`
	ExternalId          string          `json:"ExternalId" xml:"ExternalId"`
	Celebrity           []CelebrityItem `json:"Celebrity" xml:"Celebrity"`
	VideoTags           []VideoTagsItem `json:"VideoTags" xml:"VideoTags"`
	Persons             []PersonsItem   `json:"Persons" xml:"Persons"`
}

// CreateGetVideoRequest creates a request to invoke GetVideo API
func CreateGetVideoRequest() (request *GetVideoRequest) {
	request = &GetVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetVideo", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVideoResponse creates a response to parse from GetVideo response
func CreateGetVideoResponse() (response *GetVideoResponse) {
	response = &GetVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
