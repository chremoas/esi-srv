// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/chremoas/esi-srv/proto (interfaces: CorporationServiceClient)

package esi_mocks

import (
	context "context"
	proto "github.com/chremoas/esi-srv/proto"
	client "github.com/micro/go-micro/client"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockCorporationServiceClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCorporationServiceClient() *MockCorporationServiceClient {
	return &MockCorporationServiceClient{fail: pegomock.GlobalFailHandler}
}

func (mock *MockCorporationServiceClient) GetCorporationById(_param0 context.Context, _param1 *proto.GetCorporationByIdRequest, _param2 ...client.CallOption) (*proto.GetCorporationByIdResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetCorporationById", params, []reflect.Type{reflect.TypeOf((**proto.GetCorporationByIdResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.GetCorporationByIdResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.GetCorporationByIdResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCorporationServiceClient) VerifyWasCalledOnce() *VerifierCorporationServiceClient {
	return &VerifierCorporationServiceClient{mock, pegomock.Times(1), nil}
}

func (mock *MockCorporationServiceClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierCorporationServiceClient {
	return &VerifierCorporationServiceClient{mock, invocationCountMatcher, nil}
}

func (mock *MockCorporationServiceClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierCorporationServiceClient {
	return &VerifierCorporationServiceClient{mock, invocationCountMatcher, inOrderContext}
}

type VerifierCorporationServiceClient struct {
	mock                   *MockCorporationServiceClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierCorporationServiceClient) GetCorporationById(_param0 context.Context, _param1 *proto.GetCorporationByIdRequest, _param2 ...client.CallOption) *CorporationServiceClient_GetCorporationById_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetCorporationById", params)
	return &CorporationServiceClient_GetCorporationById_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CorporationServiceClient_GetCorporationById_OngoingVerification struct {
	mock              *MockCorporationServiceClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *CorporationServiceClient_GetCorporationById_OngoingVerification) GetCapturedArguments() (context.Context, *proto.GetCorporationByIdRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *CorporationServiceClient_GetCorporationById_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.GetCorporationByIdRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.GetCorporationByIdRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.GetCorporationByIdRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}
