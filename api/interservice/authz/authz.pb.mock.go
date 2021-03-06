// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: interservice/authz/authz.proto

package authz

import (
	"context"

	version "github.com/chef/automate/api/external/common/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the AuthorizationServiceServer interface (at compile time)
var _ AuthorizationServiceServer = &AuthorizationServiceServerMock{}

// NewAuthorizationServiceServerMock gives you a fresh instance of AuthorizationServiceServerMock.
func NewAuthorizationServiceServerMock() *AuthorizationServiceServerMock {
	return &AuthorizationServiceServerMock{validateRequests: true}
}

// NewAuthorizationServiceServerMockWithoutValidation gives you a fresh instance of
// AuthorizationServiceServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewAuthorizationServiceServerMockWithoutValidation() *AuthorizationServiceServerMock {
	return &AuthorizationServiceServerMock{}
}

// AuthorizationServiceServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type AuthorizationServiceServerMock struct {
	validateRequests              bool
	GetVersionFunc                func(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	FilterAuthorizedPairsFunc     func(context.Context, *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjectsFunc  func(context.Context, *FilterAuthorizedProjectsReq) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorizedFunc        func(context.Context, *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error)
	ValidateProjectAssignmentFunc func(context.Context, *ValidateProjectAssignmentReq) (*ValidateProjectAssignmentResp, error)
}

func (m *AuthorizationServiceServerMock) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetVersionFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetVersion' not implemented")
}

func (m *AuthorizationServiceServerMock) FilterAuthorizedPairs(ctx context.Context, req *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.FilterAuthorizedPairsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'FilterAuthorizedPairs' not implemented")
}

func (m *AuthorizationServiceServerMock) FilterAuthorizedProjects(ctx context.Context, req *FilterAuthorizedProjectsReq) (*FilterAuthorizedProjectsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.FilterAuthorizedProjectsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'FilterAuthorizedProjects' not implemented")
}

func (m *AuthorizationServiceServerMock) ProjectsAuthorized(ctx context.Context, req *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ProjectsAuthorizedFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ProjectsAuthorized' not implemented")
}

func (m *AuthorizationServiceServerMock) ValidateProjectAssignment(ctx context.Context, req *ValidateProjectAssignmentReq) (*ValidateProjectAssignmentResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ValidateProjectAssignmentFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ValidateProjectAssignment' not implemented")
}

// Reset resets all overridden functions
func (m *AuthorizationServiceServerMock) Reset() {
	m.GetVersionFunc = nil
	m.FilterAuthorizedPairsFunc = nil
	m.FilterAuthorizedProjectsFunc = nil
	m.ProjectsAuthorizedFunc = nil
	m.ValidateProjectAssignmentFunc = nil
}
