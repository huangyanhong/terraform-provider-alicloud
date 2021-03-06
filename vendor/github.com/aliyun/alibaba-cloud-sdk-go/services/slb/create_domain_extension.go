package slb

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

// CreateDomainExtension invokes the slb.CreateDomainExtension API synchronously
// api document: https://help.aliyun.com/api/slb/createdomainextension.html
func (client *Client) CreateDomainExtension(request *CreateDomainExtensionRequest) (response *CreateDomainExtensionResponse, err error) {
	response = CreateCreateDomainExtensionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDomainExtensionWithChan invokes the slb.CreateDomainExtension API asynchronously
// api document: https://help.aliyun.com/api/slb/createdomainextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDomainExtensionWithChan(request *CreateDomainExtensionRequest) (<-chan *CreateDomainExtensionResponse, <-chan error) {
	responseChan := make(chan *CreateDomainExtensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDomainExtension(request)
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

// CreateDomainExtensionWithCallback invokes the slb.CreateDomainExtension API asynchronously
// api document: https://help.aliyun.com/api/slb/createdomainextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDomainExtensionWithCallback(request *CreateDomainExtensionRequest, callback func(response *CreateDomainExtensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDomainExtensionResponse
		var err error
		defer close(result)
		response, err = client.CreateDomainExtension(request)
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

// CreateDomainExtensionRequest is the request struct for api CreateDomainExtension
type CreateDomainExtensionRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Domain               string           `position:"Query" name:"Domain"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ServerCertificateId  string           `position:"Query" name:"ServerCertificateId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// CreateDomainExtensionResponse is the response struct for api CreateDomainExtension
type CreateDomainExtensionResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ListenerPort      int    `json:"ListenerPort" xml:"ListenerPort"`
	DomainExtensionId string `json:"DomainExtensionId" xml:"DomainExtensionId"`
}

// CreateCreateDomainExtensionRequest creates a request to invoke CreateDomainExtension API
func CreateCreateDomainExtensionRequest() (request *CreateDomainExtensionRequest) {
	request = &CreateDomainExtensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateDomainExtension", "slb", "openAPI")
	return
}

// CreateCreateDomainExtensionResponse creates a response to parse from CreateDomainExtension response
func CreateCreateDomainExtensionResponse() (response *CreateDomainExtensionResponse) {
	response = &CreateDomainExtensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
