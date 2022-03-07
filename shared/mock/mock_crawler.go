// Code generated by MockGen. DO NOT EDIT.
// Source: license/shared/crawler.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCrawler is a mock of Crawler interface
type MockCrawler struct {
	ctrl     *gomock.Controller
	recorder *MockCrawlerMockRecorder
}

// MockCrawlerMockRecorder is the mock recorder for MockCrawler
type MockCrawlerMockRecorder struct {
	mock *MockCrawler
}

// NewMockCrawler creates a new mock instance
func NewMockCrawler(ctrl *gomock.Controller) *MockCrawler {
	mock := &MockCrawler{ctrl: ctrl}
	mock.recorder = &MockCrawlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrawler) EXPECT() *MockCrawlerMockRecorder {
	return m.recorder
}

// Crawl mocks base method
func (m *MockCrawler) Crawl(url string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Crawl", url)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Crawl indicates an expected call of Crawl
func (mr *MockCrawlerMockRecorder) Crawl(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Crawl", reflect.TypeOf((*MockCrawler)(nil).Crawl), url)
}