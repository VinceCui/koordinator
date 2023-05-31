/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/koordlet/metriccache/metric_result.go

// Package mock_metriccache is a generated GoMock package.
package mock_metriccache

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	metriccache "github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	storage "github.com/prometheus/prometheus/storage"
)

// MockMetricResult is a mock of MetricResult interface.
type MockMetricResult struct {
	ctrl     *gomock.Controller
	recorder *MockMetricResultMockRecorder
}

// MockMetricResultMockRecorder is the mock recorder for MockMetricResult.
type MockMetricResultMockRecorder struct {
	mock *MockMetricResult
}

// NewMockMetricResult creates a new mock instance.
func NewMockMetricResult(ctrl *gomock.Controller) *MockMetricResult {
	mock := &MockMetricResult{ctrl: ctrl}
	mock.recorder = &MockMetricResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricResult) EXPECT() *MockMetricResultMockRecorder {
	return m.recorder
}

// AddSeries mocks base method.
func (m *MockMetricResult) AddSeries(arg0 storage.Series) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSeries", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSeries indicates an expected call of AddSeries.
func (mr *MockMetricResultMockRecorder) AddSeries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSeries", reflect.TypeOf((*MockMetricResult)(nil).AddSeries), arg0)
}

// GetKind mocks base method.
func (m *MockMetricResult) GetKind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKind")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetKind indicates an expected call of GetKind.
func (mr *MockMetricResultMockRecorder) GetKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKind", reflect.TypeOf((*MockMetricResult)(nil).GetKind))
}

// GetProperties mocks base method.
func (m *MockMetricResult) GetProperties() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperties")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetProperties indicates an expected call of GetProperties.
func (mr *MockMetricResultMockRecorder) GetProperties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperties", reflect.TypeOf((*MockMetricResult)(nil).GetProperties))
}

// MockAggregateResultFactory is a mock of AggregateResultFactory interface.
type MockAggregateResultFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAggregateResultFactoryMockRecorder
}

// MockAggregateResultFactoryMockRecorder is the mock recorder for MockAggregateResultFactory.
type MockAggregateResultFactoryMockRecorder struct {
	mock *MockAggregateResultFactory
}

// NewMockAggregateResultFactory creates a new mock instance.
func NewMockAggregateResultFactory(ctrl *gomock.Controller) *MockAggregateResultFactory {
	mock := &MockAggregateResultFactory{ctrl: ctrl}
	mock.recorder = &MockAggregateResultFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregateResultFactory) EXPECT() *MockAggregateResultFactoryMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockAggregateResultFactory) New(meta metriccache.MetricMeta) metriccache.AggregateResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", meta)
	ret0, _ := ret[0].(metriccache.AggregateResult)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockAggregateResultFactoryMockRecorder) New(meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockAggregateResultFactory)(nil).New), meta)
}

// MockAggregateResult is a mock of AggregateResult interface.
type MockAggregateResult struct {
	ctrl     *gomock.Controller
	recorder *MockAggregateResultMockRecorder
}

// MockAggregateResultMockRecorder is the mock recorder for MockAggregateResult.
type MockAggregateResultMockRecorder struct {
	mock *MockAggregateResult
}

// NewMockAggregateResult creates a new mock instance.
func NewMockAggregateResult(ctrl *gomock.Controller) *MockAggregateResult {
	mock := &MockAggregateResult{ctrl: ctrl}
	mock.recorder = &MockAggregateResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregateResult) EXPECT() *MockAggregateResultMockRecorder {
	return m.recorder
}

// AddSeries mocks base method.
func (m *MockAggregateResult) AddSeries(arg0 storage.Series) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSeries", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSeries indicates an expected call of AddSeries.
func (mr *MockAggregateResultMockRecorder) AddSeries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSeries", reflect.TypeOf((*MockAggregateResult)(nil).AddSeries), arg0)
}

// Count mocks base method.
func (m *MockAggregateResult) Count() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockAggregateResultMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockAggregateResult)(nil).Count))
}

// GetKind mocks base method.
func (m *MockAggregateResult) GetKind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKind")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetKind indicates an expected call of GetKind.
func (mr *MockAggregateResultMockRecorder) GetKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKind", reflect.TypeOf((*MockAggregateResult)(nil).GetKind))
}

// GetProperties mocks base method.
func (m *MockAggregateResult) GetProperties() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperties")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetProperties indicates an expected call of GetProperties.
func (mr *MockAggregateResultMockRecorder) GetProperties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperties", reflect.TypeOf((*MockAggregateResult)(nil).GetProperties))
}

// TimeRangeDuration mocks base method.
func (m *MockAggregateResult) TimeRangeDuration() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeRangeDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeRangeDuration indicates an expected call of TimeRangeDuration.
func (mr *MockAggregateResultMockRecorder) TimeRangeDuration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeRangeDuration", reflect.TypeOf((*MockAggregateResult)(nil).TimeRangeDuration))
}

// Value mocks base method.
func (m *MockAggregateResult) Value(t metriccache.AggregationType) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value", t)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Value indicates an expected call of Value.
func (mr *MockAggregateResultMockRecorder) Value(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockAggregateResult)(nil).Value), t)
}
