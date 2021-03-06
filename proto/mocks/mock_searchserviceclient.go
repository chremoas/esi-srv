//// Automatically generated by pegomock. DO NOT EDIT!
//// Source: github.com/chremoas/esi-srv/proto (interfaces: SearchServiceClient)
//
package esi_mocks
//
//import (
//	context "golang.org/x/net/context"
//	proto "github.com/chremoas/esi-srv/proto"
//	client "github.com/micro/go-micro/client"
//	pegomock "github.com/petergtz/pegomock"
//	"reflect"
//)
//
//type MockSearchServiceClient struct {
//	fail func(message string, callerSkip ...int)
//}
//
//func NewMockSearchServiceClient() *MockSearchServiceClient {
//	return &MockSearchServiceClient{fail: pegomock.GlobalFailHandler}
//}
//
//func (mock *MockSearchServiceClient) Search(_param0 context.Context, _param1 *proto.SearchRequest, _param2 ...client.CallOption) (*proto.SearchResponse, error) {
//	params := []pegomock.Param{_param0, _param1}
//	for _, param := range _param2 {
//		params = append(params, param)
//	}
//	result := pegomock.GetGenericMockFrom(mock).Invoke("Search", params, []reflect.Type{reflect.TypeOf((**proto.SearchResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
//	var ret0 *proto.SearchResponse
//	var ret1 error
//	if len(result) != 0 {
//		if result[0] != nil {
//			ret0 = result[0].(*proto.SearchResponse)
//		}
//		if result[1] != nil {
//			ret1 = result[1].(error)
//		}
//	}
//	return ret0, ret1
//}
//
//func (mock *MockSearchServiceClient) VerifyWasCalledOnce() *VerifierSearchServiceClient {
//	return &VerifierSearchServiceClient{mock, pegomock.Times(1), nil}
//}
//
//func (mock *MockSearchServiceClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierSearchServiceClient {
//	return &VerifierSearchServiceClient{mock, invocationCountMatcher, nil}
//}
//
//func (mock *MockSearchServiceClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierSearchServiceClient {
//	return &VerifierSearchServiceClient{mock, invocationCountMatcher, inOrderContext}
//}
//
//type VerifierSearchServiceClient struct {
//	mock                   *MockSearchServiceClient
//	invocationCountMatcher pegomock.Matcher
//	inOrderContext         *pegomock.InOrderContext
//}
//
//func (verifier *VerifierSearchServiceClient) Search(_param0 context.Context, _param1 *proto.SearchRequest, _param2 ...client.CallOption) *SearchServiceClient_Search_OngoingVerification {
//	params := []pegomock.Param{_param0, _param1}
//	for _, param := range _param2 {
//		params = append(params, param)
//	}
//	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Search", params)
//	return &SearchServiceClient_Search_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
//}
//
//type SearchServiceClient_Search_OngoingVerification struct {
//	mock              *MockSearchServiceClient
//	methodInvocations []pegomock.MethodInvocation
//}
//
//func (c *SearchServiceClient_Search_OngoingVerification) GetCapturedArguments() (context.Context, *proto.SearchRequest, []client.CallOption) {
//	_param0, _param1, _param2 := c.GetAllCapturedArguments()
//	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
//}
//
//func (c *SearchServiceClient_Search_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.SearchRequest, _param2 [][]client.CallOption) {
//	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
//	if len(params) > 0 {
//		_param0 = make([]context.Context, len(params[0]))
//		for u, param := range params[0] {
//			_param0[u] = param.(context.Context)
//		}
//		_param1 = make([]*proto.SearchRequest, len(params[1]))
//		for u, param := range params[1] {
//			_param1[u] = param.(*proto.SearchRequest)
//		}
//		_param2 = make([][]client.CallOption, len(params[2]))
//		for u := range params[0] {
//			_param2[u] = make([]client.CallOption, len(params)-2)
//			for x := 2; x < len(params); x++ {
//				if params[x][u] != nil {
//					_param2[u][x-2] = params[x][u].(client.CallOption)
//				}
//			}
//		}
//	}
//	return
//}
