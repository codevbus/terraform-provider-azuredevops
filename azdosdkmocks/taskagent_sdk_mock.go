// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/taskagent (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	taskagent "github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
	reflect "reflect"
)

// MockTaskagentClient is a mock of Client interface
type MockTaskagentClient struct {
	ctrl     *gomock.Controller
	recorder *MockTaskagentClientMockRecorder
}

// MockTaskagentClientMockRecorder is the mock recorder for MockTaskagentClient
type MockTaskagentClientMockRecorder struct {
	mock *MockTaskagentClient
}

// NewMockTaskagentClient creates a new mock instance
func NewMockTaskagentClient(ctrl *gomock.Controller) *MockTaskagentClient {
	mock := &MockTaskagentClient{ctrl: ctrl}
	mock.recorder = &MockTaskagentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskagentClient) EXPECT() *MockTaskagentClientMockRecorder {
	return m.recorder
}

// AddAgent mocks base method
func (m *MockTaskagentClient) AddAgent(arg0 context.Context, arg1 taskagent.AddAgentArgs) (*taskagent.TaskAgent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAgent", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAgent indicates an expected call of AddAgent
func (mr *MockTaskagentClientMockRecorder) AddAgent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAgent", reflect.TypeOf((*MockTaskagentClient)(nil).AddAgent), arg0, arg1)
}

// AddAgentCloud mocks base method
func (m *MockTaskagentClient) AddAgentCloud(arg0 context.Context, arg1 taskagent.AddAgentCloudArgs) (*taskagent.TaskAgentCloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAgentCloud", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentCloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAgentCloud indicates an expected call of AddAgentCloud
func (mr *MockTaskagentClientMockRecorder) AddAgentCloud(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAgentCloud", reflect.TypeOf((*MockTaskagentClient)(nil).AddAgentCloud), arg0, arg1)
}

// AddAgentPool mocks base method
func (m *MockTaskagentClient) AddAgentPool(arg0 context.Context, arg1 taskagent.AddAgentPoolArgs) (*taskagent.TaskAgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAgentPool", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAgentPool indicates an expected call of AddAgentPool
func (mr *MockTaskagentClientMockRecorder) AddAgentPool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAgentPool", reflect.TypeOf((*MockTaskagentClient)(nil).AddAgentPool), arg0, arg1)
}

// AddAgentQueue mocks base method
func (m *MockTaskagentClient) AddAgentQueue(arg0 context.Context, arg1 taskagent.AddAgentQueueArgs) (*taskagent.TaskAgentQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAgentQueue", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAgentQueue indicates an expected call of AddAgentQueue
func (mr *MockTaskagentClientMockRecorder) AddAgentQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAgentQueue", reflect.TypeOf((*MockTaskagentClient)(nil).AddAgentQueue), arg0, arg1)
}

// AddDeploymentGroup mocks base method
func (m *MockTaskagentClient) AddDeploymentGroup(arg0 context.Context, arg1 taskagent.AddDeploymentGroupArgs) (*taskagent.DeploymentGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeploymentGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.DeploymentGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDeploymentGroup indicates an expected call of AddDeploymentGroup
func (mr *MockTaskagentClientMockRecorder) AddDeploymentGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeploymentGroup", reflect.TypeOf((*MockTaskagentClient)(nil).AddDeploymentGroup), arg0, arg1)
}

// AddTaskGroup mocks base method
func (m *MockTaskagentClient) AddTaskGroup(arg0 context.Context, arg1 taskagent.AddTaskGroupArgs) (*taskagent.TaskGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTaskGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTaskGroup indicates an expected call of AddTaskGroup
func (mr *MockTaskagentClientMockRecorder) AddTaskGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTaskGroup", reflect.TypeOf((*MockTaskagentClient)(nil).AddTaskGroup), arg0, arg1)
}

// AddVariableGroup mocks base method
func (m *MockTaskagentClient) AddVariableGroup(arg0 context.Context, arg1 taskagent.AddVariableGroupArgs) (*taskagent.VariableGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVariableGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.VariableGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddVariableGroup indicates an expected call of AddVariableGroup
func (mr *MockTaskagentClientMockRecorder) AddVariableGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVariableGroup", reflect.TypeOf((*MockTaskagentClient)(nil).AddVariableGroup), arg0, arg1)
}

// DeleteAgent mocks base method
func (m *MockTaskagentClient) DeleteAgent(arg0 context.Context, arg1 taskagent.DeleteAgentArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgent indicates an expected call of DeleteAgent
func (mr *MockTaskagentClientMockRecorder) DeleteAgent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgent", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteAgent), arg0, arg1)
}

// DeleteAgentCloud mocks base method
func (m *MockTaskagentClient) DeleteAgentCloud(arg0 context.Context, arg1 taskagent.DeleteAgentCloudArgs) (*taskagent.TaskAgentCloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgentCloud", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentCloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAgentCloud indicates an expected call of DeleteAgentCloud
func (mr *MockTaskagentClientMockRecorder) DeleteAgentCloud(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentCloud", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteAgentCloud), arg0, arg1)
}

// DeleteAgentPool mocks base method
func (m *MockTaskagentClient) DeleteAgentPool(arg0 context.Context, arg1 taskagent.DeleteAgentPoolArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgentPool", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgentPool indicates an expected call of DeleteAgentPool
func (mr *MockTaskagentClientMockRecorder) DeleteAgentPool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentPool", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteAgentPool), arg0, arg1)
}

// DeleteAgentQueue mocks base method
func (m *MockTaskagentClient) DeleteAgentQueue(arg0 context.Context, arg1 taskagent.DeleteAgentQueueArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgentQueue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgentQueue indicates an expected call of DeleteAgentQueue
func (mr *MockTaskagentClientMockRecorder) DeleteAgentQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgentQueue", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteAgentQueue), arg0, arg1)
}

// DeleteDeploymentGroup mocks base method
func (m *MockTaskagentClient) DeleteDeploymentGroup(arg0 context.Context, arg1 taskagent.DeleteDeploymentGroupArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeploymentGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeploymentGroup indicates an expected call of DeleteDeploymentGroup
func (mr *MockTaskagentClientMockRecorder) DeleteDeploymentGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeploymentGroup", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteDeploymentGroup), arg0, arg1)
}

// DeleteDeploymentTarget mocks base method
func (m *MockTaskagentClient) DeleteDeploymentTarget(arg0 context.Context, arg1 taskagent.DeleteDeploymentTargetArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeploymentTarget", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeploymentTarget indicates an expected call of DeleteDeploymentTarget
func (mr *MockTaskagentClientMockRecorder) DeleteDeploymentTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeploymentTarget", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteDeploymentTarget), arg0, arg1)
}

// DeleteTaskGroup mocks base method
func (m *MockTaskagentClient) DeleteTaskGroup(arg0 context.Context, arg1 taskagent.DeleteTaskGroupArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTaskGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTaskGroup indicates an expected call of DeleteTaskGroup
func (mr *MockTaskagentClientMockRecorder) DeleteTaskGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTaskGroup", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteTaskGroup), arg0, arg1)
}

// DeleteVariableGroup mocks base method
func (m *MockTaskagentClient) DeleteVariableGroup(arg0 context.Context, arg1 taskagent.DeleteVariableGroupArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVariableGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVariableGroup indicates an expected call of DeleteVariableGroup
func (mr *MockTaskagentClientMockRecorder) DeleteVariableGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVariableGroup", reflect.TypeOf((*MockTaskagentClient)(nil).DeleteVariableGroup), arg0, arg1)
}

// GetAgent mocks base method
func (m *MockTaskagentClient) GetAgent(arg0 context.Context, arg1 taskagent.GetAgentArgs) (*taskagent.TaskAgent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgent", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgent indicates an expected call of GetAgent
func (mr *MockTaskagentClientMockRecorder) GetAgent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgent), arg0, arg1)
}

// GetAgentCloud mocks base method
func (m *MockTaskagentClient) GetAgentCloud(arg0 context.Context, arg1 taskagent.GetAgentCloudArgs) (*taskagent.TaskAgentCloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentCloud", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentCloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentCloud indicates an expected call of GetAgentCloud
func (mr *MockTaskagentClientMockRecorder) GetAgentCloud(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentCloud", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentCloud), arg0, arg1)
}

// GetAgentCloudRequests mocks base method
func (m *MockTaskagentClient) GetAgentCloudRequests(arg0 context.Context, arg1 taskagent.GetAgentCloudRequestsArgs) (*[]taskagent.TaskAgentCloudRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentCloudRequests", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentCloudRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentCloudRequests indicates an expected call of GetAgentCloudRequests
func (mr *MockTaskagentClientMockRecorder) GetAgentCloudRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentCloudRequests", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentCloudRequests), arg0, arg1)
}

// GetAgentCloudTypes mocks base method
func (m *MockTaskagentClient) GetAgentCloudTypes(arg0 context.Context, arg1 taskagent.GetAgentCloudTypesArgs) (*[]taskagent.TaskAgentCloudType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentCloudTypes", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentCloudType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentCloudTypes indicates an expected call of GetAgentCloudTypes
func (mr *MockTaskagentClientMockRecorder) GetAgentCloudTypes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentCloudTypes", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentCloudTypes), arg0, arg1)
}

// GetAgentClouds mocks base method
func (m *MockTaskagentClient) GetAgentClouds(arg0 context.Context, arg1 taskagent.GetAgentCloudsArgs) (*[]taskagent.TaskAgentCloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentClouds", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentCloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentClouds indicates an expected call of GetAgentClouds
func (mr *MockTaskagentClientMockRecorder) GetAgentClouds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentClouds", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentClouds), arg0, arg1)
}

// GetAgentPool mocks base method
func (m *MockTaskagentClient) GetAgentPool(arg0 context.Context, arg1 taskagent.GetAgentPoolArgs) (*taskagent.TaskAgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentPool", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentPool indicates an expected call of GetAgentPool
func (mr *MockTaskagentClientMockRecorder) GetAgentPool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentPool", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentPool), arg0, arg1)
}

// GetAgentPools mocks base method
func (m *MockTaskagentClient) GetAgentPools(arg0 context.Context, arg1 taskagent.GetAgentPoolsArgs) (*[]taskagent.TaskAgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentPools", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentPools indicates an expected call of GetAgentPools
func (mr *MockTaskagentClientMockRecorder) GetAgentPools(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentPools", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentPools), arg0, arg1)
}

// GetAgentPoolsByIds mocks base method
func (m *MockTaskagentClient) GetAgentPoolsByIds(arg0 context.Context, arg1 taskagent.GetAgentPoolsByIdsArgs) (*[]taskagent.TaskAgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentPoolsByIds", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentPoolsByIds indicates an expected call of GetAgentPoolsByIds
func (mr *MockTaskagentClientMockRecorder) GetAgentPoolsByIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentPoolsByIds", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentPoolsByIds), arg0, arg1)
}

// GetAgentQueue mocks base method
func (m *MockTaskagentClient) GetAgentQueue(arg0 context.Context, arg1 taskagent.GetAgentQueueArgs) (*taskagent.TaskAgentQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentQueue", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentQueue indicates an expected call of GetAgentQueue
func (mr *MockTaskagentClientMockRecorder) GetAgentQueue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentQueue", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentQueue), arg0, arg1)
}

// GetAgentQueues mocks base method
func (m *MockTaskagentClient) GetAgentQueues(arg0 context.Context, arg1 taskagent.GetAgentQueuesArgs) (*[]taskagent.TaskAgentQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentQueues", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentQueues indicates an expected call of GetAgentQueues
func (mr *MockTaskagentClientMockRecorder) GetAgentQueues(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentQueues", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentQueues), arg0, arg1)
}

// GetAgentQueuesByIds mocks base method
func (m *MockTaskagentClient) GetAgentQueuesByIds(arg0 context.Context, arg1 taskagent.GetAgentQueuesByIdsArgs) (*[]taskagent.TaskAgentQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentQueuesByIds", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentQueuesByIds indicates an expected call of GetAgentQueuesByIds
func (mr *MockTaskagentClientMockRecorder) GetAgentQueuesByIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentQueuesByIds", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentQueuesByIds), arg0, arg1)
}

// GetAgentQueuesByNames mocks base method
func (m *MockTaskagentClient) GetAgentQueuesByNames(arg0 context.Context, arg1 taskagent.GetAgentQueuesByNamesArgs) (*[]taskagent.TaskAgentQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentQueuesByNames", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgentQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentQueuesByNames indicates an expected call of GetAgentQueuesByNames
func (mr *MockTaskagentClientMockRecorder) GetAgentQueuesByNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentQueuesByNames", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgentQueuesByNames), arg0, arg1)
}

// GetAgents mocks base method
func (m *MockTaskagentClient) GetAgents(arg0 context.Context, arg1 taskagent.GetAgentsArgs) (*[]taskagent.TaskAgent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgents", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgents indicates an expected call of GetAgents
func (mr *MockTaskagentClientMockRecorder) GetAgents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgents", reflect.TypeOf((*MockTaskagentClient)(nil).GetAgents), arg0, arg1)
}

// GetDeploymentGroup mocks base method
func (m *MockTaskagentClient) GetDeploymentGroup(arg0 context.Context, arg1 taskagent.GetDeploymentGroupArgs) (*taskagent.DeploymentGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.DeploymentGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentGroup indicates an expected call of GetDeploymentGroup
func (mr *MockTaskagentClientMockRecorder) GetDeploymentGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentGroup", reflect.TypeOf((*MockTaskagentClient)(nil).GetDeploymentGroup), arg0, arg1)
}

// GetDeploymentGroups mocks base method
func (m *MockTaskagentClient) GetDeploymentGroups(arg0 context.Context, arg1 taskagent.GetDeploymentGroupsArgs) (*taskagent.GetDeploymentGroupsResponseValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentGroups", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.GetDeploymentGroupsResponseValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentGroups indicates an expected call of GetDeploymentGroups
func (mr *MockTaskagentClientMockRecorder) GetDeploymentGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentGroups", reflect.TypeOf((*MockTaskagentClient)(nil).GetDeploymentGroups), arg0, arg1)
}

// GetDeploymentTarget mocks base method
func (m *MockTaskagentClient) GetDeploymentTarget(arg0 context.Context, arg1 taskagent.GetDeploymentTargetArgs) (*taskagent.DeploymentMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentTarget", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.DeploymentMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentTarget indicates an expected call of GetDeploymentTarget
func (mr *MockTaskagentClientMockRecorder) GetDeploymentTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentTarget", reflect.TypeOf((*MockTaskagentClient)(nil).GetDeploymentTarget), arg0, arg1)
}

// GetDeploymentTargets mocks base method
func (m *MockTaskagentClient) GetDeploymentTargets(arg0 context.Context, arg1 taskagent.GetDeploymentTargetsArgs) (*taskagent.GetDeploymentTargetsResponseValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentTargets", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.GetDeploymentTargetsResponseValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentTargets indicates an expected call of GetDeploymentTargets
func (mr *MockTaskagentClientMockRecorder) GetDeploymentTargets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentTargets", reflect.TypeOf((*MockTaskagentClient)(nil).GetDeploymentTargets), arg0, arg1)
}

// GetTaskGroups mocks base method
func (m *MockTaskagentClient) GetTaskGroups(arg0 context.Context, arg1 taskagent.GetTaskGroupsArgs) (*[]taskagent.TaskGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskGroups", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.TaskGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskGroups indicates an expected call of GetTaskGroups
func (mr *MockTaskagentClientMockRecorder) GetTaskGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskGroups", reflect.TypeOf((*MockTaskagentClient)(nil).GetTaskGroups), arg0, arg1)
}

// GetVariableGroup mocks base method
func (m *MockTaskagentClient) GetVariableGroup(arg0 context.Context, arg1 taskagent.GetVariableGroupArgs) (*taskagent.VariableGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVariableGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.VariableGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVariableGroup indicates an expected call of GetVariableGroup
func (mr *MockTaskagentClientMockRecorder) GetVariableGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVariableGroup", reflect.TypeOf((*MockTaskagentClient)(nil).GetVariableGroup), arg0, arg1)
}

// GetVariableGroups mocks base method
func (m *MockTaskagentClient) GetVariableGroups(arg0 context.Context, arg1 taskagent.GetVariableGroupsArgs) (*[]taskagent.VariableGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVariableGroups", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.VariableGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVariableGroups indicates an expected call of GetVariableGroups
func (mr *MockTaskagentClientMockRecorder) GetVariableGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVariableGroups", reflect.TypeOf((*MockTaskagentClient)(nil).GetVariableGroups), arg0, arg1)
}

// GetVariableGroupsById mocks base method
func (m *MockTaskagentClient) GetVariableGroupsById(arg0 context.Context, arg1 taskagent.GetVariableGroupsByIdArgs) (*[]taskagent.VariableGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVariableGroupsById", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.VariableGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVariableGroupsById indicates an expected call of GetVariableGroupsById
func (mr *MockTaskagentClientMockRecorder) GetVariableGroupsById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVariableGroupsById", reflect.TypeOf((*MockTaskagentClient)(nil).GetVariableGroupsById), arg0, arg1)
}

// GetYamlSchema mocks base method
func (m *MockTaskagentClient) GetYamlSchema(arg0 context.Context, arg1 taskagent.GetYamlSchemaArgs) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetYamlSchema", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetYamlSchema indicates an expected call of GetYamlSchema
func (mr *MockTaskagentClientMockRecorder) GetYamlSchema(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetYamlSchema", reflect.TypeOf((*MockTaskagentClient)(nil).GetYamlSchema), arg0, arg1)
}

// ReplaceAgent mocks base method
func (m *MockTaskagentClient) ReplaceAgent(arg0 context.Context, arg1 taskagent.ReplaceAgentArgs) (*taskagent.TaskAgent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceAgent", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceAgent indicates an expected call of ReplaceAgent
func (mr *MockTaskagentClientMockRecorder) ReplaceAgent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceAgent", reflect.TypeOf((*MockTaskagentClient)(nil).ReplaceAgent), arg0, arg1)
}

// UpdateAgent mocks base method
func (m *MockTaskagentClient) UpdateAgent(arg0 context.Context, arg1 taskagent.UpdateAgentArgs) (*taskagent.TaskAgent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgent", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAgent indicates an expected call of UpdateAgent
func (mr *MockTaskagentClientMockRecorder) UpdateAgent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgent", reflect.TypeOf((*MockTaskagentClient)(nil).UpdateAgent), arg0, arg1)
}

// UpdateAgentPool mocks base method
func (m *MockTaskagentClient) UpdateAgentPool(arg0 context.Context, arg1 taskagent.UpdateAgentPoolArgs) (*taskagent.TaskAgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgentPool", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskAgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAgentPool indicates an expected call of UpdateAgentPool
func (mr *MockTaskagentClientMockRecorder) UpdateAgentPool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgentPool", reflect.TypeOf((*MockTaskagentClient)(nil).UpdateAgentPool), arg0, arg1)
}

// UpdateDeploymentGroup mocks base method
func (m *MockTaskagentClient) UpdateDeploymentGroup(arg0 context.Context, arg1 taskagent.UpdateDeploymentGroupArgs) (*taskagent.DeploymentGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeploymentGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.DeploymentGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDeploymentGroup indicates an expected call of UpdateDeploymentGroup
func (mr *MockTaskagentClientMockRecorder) UpdateDeploymentGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeploymentGroup", reflect.TypeOf((*MockTaskagentClient)(nil).UpdateDeploymentGroup), arg0, arg1)
}

// UpdateDeploymentTargets mocks base method
func (m *MockTaskagentClient) UpdateDeploymentTargets(arg0 context.Context, arg1 taskagent.UpdateDeploymentTargetsArgs) (*[]taskagent.DeploymentMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeploymentTargets", arg0, arg1)
	ret0, _ := ret[0].(*[]taskagent.DeploymentMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDeploymentTargets indicates an expected call of UpdateDeploymentTargets
func (mr *MockTaskagentClientMockRecorder) UpdateDeploymentTargets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeploymentTargets", reflect.TypeOf((*MockTaskagentClient)(nil).UpdateDeploymentTargets), arg0, arg1)
}

// UpdateTaskGroup mocks base method
func (m *MockTaskagentClient) UpdateTaskGroup(arg0 context.Context, arg1 taskagent.UpdateTaskGroupArgs) (*taskagent.TaskGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.TaskGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskGroup indicates an expected call of UpdateTaskGroup
func (mr *MockTaskagentClientMockRecorder) UpdateTaskGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskGroup", reflect.TypeOf((*MockTaskagentClient)(nil).UpdateTaskGroup), arg0, arg1)
}

// UpdateVariableGroup mocks base method
func (m *MockTaskagentClient) UpdateVariableGroup(arg0 context.Context, arg1 taskagent.UpdateVariableGroupArgs) (*taskagent.VariableGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVariableGroup", arg0, arg1)
	ret0, _ := ret[0].(*taskagent.VariableGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVariableGroup indicates an expected call of UpdateVariableGroup
func (mr *MockTaskagentClientMockRecorder) UpdateVariableGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVariableGroup", reflect.TypeOf((*MockTaskagentClient)(nil).UpdateVariableGroup), arg0, arg1)
}
