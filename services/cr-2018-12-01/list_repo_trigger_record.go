package cr_20181201

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

// ListRepoTriggerRecord invokes the cr.ListRepoTriggerRecord API synchronously
// api document: https://help.aliyun.com/api/cr/listrepotriggerrecord.html
func (client *Client) ListRepoTriggerRecord(request *ListRepoTriggerRecordRequest) (response *ListRepoTriggerRecordResponse, err error) {
	response = CreateListRepoTriggerRecordResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepoTriggerRecordWithChan invokes the cr.ListRepoTriggerRecord API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotriggerrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTriggerRecordWithChan(request *ListRepoTriggerRecordRequest) (<-chan *ListRepoTriggerRecordResponse, <-chan error) {
	responseChan := make(chan *ListRepoTriggerRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepoTriggerRecord(request)
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

// ListRepoTriggerRecordWithCallback invokes the cr.ListRepoTriggerRecord API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotriggerrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTriggerRecordWithCallback(request *ListRepoTriggerRecordRequest, callback func(response *ListRepoTriggerRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepoTriggerRecordResponse
		var err error
		defer close(result)
		response, err = client.ListRepoTriggerRecord(request)
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

// ListRepoTriggerRecordRequest is the request struct for api ListRepoTriggerRecord
type ListRepoTriggerRecordRequest struct {
	*requests.RpcRequest
	InstanceId      string `position:"Query" name:"InstanceId"`
	TriggerRecordId string `position:"Query" name:"TriggerRecordId"`
}

// ListRepoTriggerRecordResponse is the response struct for api ListRepoTriggerRecord
type ListRepoTriggerRecordResponse struct {
	*responses.BaseResponse
	ListRepoTriggerRecordIsSuccess bool                     `json:"IsSuccess" xml:"IsSuccess"`
	Code                           string                   `json:"Code" xml:"Code"`
	RequestId                      string                   `json:"RequestId" xml:"RequestId"`
	RepoTriggerRecords             []RepoTriggerRecordsItem `json:"RepoTriggerRecords" xml:"RepoTriggerRecords"`
}

// CreateListRepoTriggerRecordRequest creates a request to invoke ListRepoTriggerRecord API
func CreateListRepoTriggerRecordRequest() (request *ListRepoTriggerRecordRequest) {
	request = &ListRepoTriggerRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "ListRepoTriggerRecord", "cr", "openAPI")
	return
}

// CreateListRepoTriggerRecordResponse creates a response to parse from ListRepoTriggerRecord response
func CreateListRepoTriggerRecordResponse() (response *ListRepoTriggerRecordResponse) {
	response = &ListRepoTriggerRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}