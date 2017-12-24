// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/chremoas/esi-srv/proto (interfaces: CharacterServiceClient)

package esi_mocks

import (
	context "context"
	proto "github.com/chremoas/esi-srv/proto"
	client "github.com/micro/go-micro/client"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockCharacterServiceClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCharacterServiceClient() *MockCharacterServiceClient {
	return &MockCharacterServiceClient{fail: pegomock.GlobalFailHandler}
}

func (mock *MockCharacterServiceClient) GetCharacterById(_param0 context.Context, _param1 *proto.GetCharacterByIdRequest, _param2 ...client.CallOption) (*proto.GetCharacterByIdResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetCharacterById", params, []reflect.Type{reflect.TypeOf((**proto.GetCharacterByIdResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.GetCharacterByIdResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.GetCharacterByIdResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCharacterServiceClient) VerifyWasCalledOnce() *VerifierCharacterServiceClient {
	return &VerifierCharacterServiceClient{mock, pegomock.Times(1), nil}
}

func (mock *MockCharacterServiceClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierCharacterServiceClient {
	return &VerifierCharacterServiceClient{mock, invocationCountMatcher, nil}
}

func (mock *MockCharacterServiceClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierCharacterServiceClient {
	return &VerifierCharacterServiceClient{mock, invocationCountMatcher, inOrderContext}
}

type VerifierCharacterServiceClient struct {
	mock                   *MockCharacterServiceClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierCharacterServiceClient) GetCharacterById(_param0 context.Context, _param1 *proto.GetCharacterByIdRequest, _param2 ...client.CallOption) *CharacterServiceClient_GetCharacterById_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetCharacterById", params)
	return &CharacterServiceClient_GetCharacterById_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CharacterServiceClient_GetCharacterById_OngoingVerification struct {
	mock              *MockCharacterServiceClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *CharacterServiceClient_GetCharacterById_OngoingVerification) GetCapturedArguments() (context.Context, *proto.GetCharacterByIdRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *CharacterServiceClient_GetCharacterById_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.GetCharacterByIdRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.GetCharacterByIdRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.GetCharacterByIdRequest)
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
