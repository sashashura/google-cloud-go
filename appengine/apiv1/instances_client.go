// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package appengine

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newInstancesClientHook clientHook

// InstancesCallOptions contains the retry settings for each method of InstancesClient.
type InstancesCallOptions struct {
	ListInstances  []gax.CallOption
	GetInstance    []gax.CallOption
	DeleteInstance []gax.CallOption
	DebugInstance  []gax.CallOption
}

func defaultInstancesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultInstancesCallOptions() *InstancesCallOptions {
	return &InstancesCallOptions{
		ListInstances:  []gax.CallOption{},
		GetInstance:    []gax.CallOption{},
		DeleteInstance: []gax.CallOption{},
		DebugInstance:  []gax.CallOption{},
	}
}

// internalInstancesClient is an interface that defines the methods available from App Engine Admin API.
type internalInstancesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListInstances(context.Context, *appenginepb.ListInstancesRequest, ...gax.CallOption) *InstanceIterator
	GetInstance(context.Context, *appenginepb.GetInstanceRequest, ...gax.CallOption) (*appenginepb.Instance, error)
	DeleteInstance(context.Context, *appenginepb.DeleteInstanceRequest, ...gax.CallOption) (*DeleteInstanceOperation, error)
	DeleteInstanceOperation(name string) *DeleteInstanceOperation
	DebugInstance(context.Context, *appenginepb.DebugInstanceRequest, ...gax.CallOption) (*DebugInstanceOperation, error)
	DebugInstanceOperation(name string) *DebugInstanceOperation
}

// InstancesClient is a client for interacting with App Engine Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages instances of a version.
type InstancesClient struct {
	// The internal transport-dependent client.
	internalClient internalInstancesClient

	// The call options for this service.
	CallOptions *InstancesCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InstancesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InstancesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *InstancesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListInstances lists the instances of a version.
//
// Tip: To aggregate details about instances over time, see the
// Stackdriver Monitoring API (at https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list).
func (c *InstancesClient) ListInstances(ctx context.Context, req *appenginepb.ListInstancesRequest, opts ...gax.CallOption) *InstanceIterator {
	return c.internalClient.ListInstances(ctx, req, opts...)
}

// GetInstance gets instance information.
func (c *InstancesClient) GetInstance(ctx context.Context, req *appenginepb.GetInstanceRequest, opts ...gax.CallOption) (*appenginepb.Instance, error) {
	return c.internalClient.GetInstance(ctx, req, opts...)
}

// DeleteInstance stops a running instance.
//
// The instance might be automatically recreated based on the scaling settings
// of the version. For more information, see “How Instances are Managed”
// (standard environment (at https://cloud.google.com/appengine/docs/standard/python/how-instances-are-managed) |
// flexible environment (at https://cloud.google.com/appengine/docs/flexible/python/how-instances-are-managed)).
//
// To ensure that instances are not re-created and avoid getting billed, you
// can stop all instances within the target version by changing the serving
// status of the version to STOPPED with the
// apps.services.versions.patch (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions/patch)
// method.
func (c *InstancesClient) DeleteInstance(ctx context.Context, req *appenginepb.DeleteInstanceRequest, opts ...gax.CallOption) (*DeleteInstanceOperation, error) {
	return c.internalClient.DeleteInstance(ctx, req, opts...)
}

// DeleteInstanceOperation returns a new DeleteInstanceOperation from a given name.
// The name must be that of a previously created DeleteInstanceOperation, possibly from a different process.
func (c *InstancesClient) DeleteInstanceOperation(name string) *DeleteInstanceOperation {
	return c.internalClient.DeleteInstanceOperation(name)
}

// DebugInstance enables debugging on a VM instance. This allows you to use the SSH
// command to connect to the virtual machine where the instance lives.
// While in “debug mode”, the instance continues to serve live traffic.
// You should delete the instance when you are done debugging and then
// allow the system to take over and determine if another instance
// should be started.
//
// Only applicable for instances in App Engine flexible environment.
func (c *InstancesClient) DebugInstance(ctx context.Context, req *appenginepb.DebugInstanceRequest, opts ...gax.CallOption) (*DebugInstanceOperation, error) {
	return c.internalClient.DebugInstance(ctx, req, opts...)
}

// DebugInstanceOperation returns a new DebugInstanceOperation from a given name.
// The name must be that of a previously created DebugInstanceOperation, possibly from a different process.
func (c *InstancesClient) DebugInstanceOperation(name string) *DebugInstanceOperation {
	return c.internalClient.DebugInstanceOperation(name)
}

// instancesGRPCClient is a client for interacting with App Engine Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type instancesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing InstancesClient
	CallOptions **InstancesCallOptions

	// The gRPC API client.
	instancesClient appenginepb.InstancesClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewInstancesClient creates a new instances client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages instances of a version.
func NewInstancesClient(ctx context.Context, opts ...option.ClientOption) (*InstancesClient, error) {
	clientOpts := defaultInstancesGRPCClientOptions()
	if newInstancesClientHook != nil {
		hookOpts, err := newInstancesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := InstancesClient{CallOptions: defaultInstancesCallOptions()}

	c := &instancesGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		instancesClient:  appenginepb.NewInstancesClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *instancesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *instancesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *instancesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *instancesGRPCClient) ListInstances(ctx context.Context, req *appenginepb.ListInstancesRequest, opts ...gax.CallOption) *InstanceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListInstances[0:len((*c.CallOptions).ListInstances):len((*c.CallOptions).ListInstances)], opts...)
	it := &InstanceIterator{}
	req = proto.Clone(req).(*appenginepb.ListInstancesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.Instance, string, error) {
		resp := &appenginepb.ListInstancesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.instancesClient.ListInstances(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetInstances(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *instancesGRPCClient) GetInstance(ctx context.Context, req *appenginepb.GetInstanceRequest, opts ...gax.CallOption) (*appenginepb.Instance, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetInstance[0:len((*c.CallOptions).GetInstance):len((*c.CallOptions).GetInstance)], opts...)
	var resp *appenginepb.Instance
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.instancesClient.GetInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *instancesGRPCClient) DeleteInstance(ctx context.Context, req *appenginepb.DeleteInstanceRequest, opts ...gax.CallOption) (*DeleteInstanceOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteInstance[0:len((*c.CallOptions).DeleteInstance):len((*c.CallOptions).DeleteInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.instancesClient.DeleteInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *instancesGRPCClient) DebugInstance(ctx context.Context, req *appenginepb.DebugInstanceRequest, opts ...gax.CallOption) (*DebugInstanceOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DebugInstance[0:len((*c.CallOptions).DebugInstance):len((*c.CallOptions).DebugInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.instancesClient.DebugInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DebugInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// DebugInstanceOperation manages a long-running operation from DebugInstance.
type DebugInstanceOperation struct {
	lro *longrunning.Operation
}

// DebugInstanceOperation returns a new DebugInstanceOperation from a given name.
// The name must be that of a previously created DebugInstanceOperation, possibly from a different process.
func (c *instancesGRPCClient) DebugInstanceOperation(name string) *DebugInstanceOperation {
	return &DebugInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DebugInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Instance, error) {
	var resp appenginepb.Instance
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DebugInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Instance, error) {
	var resp appenginepb.Instance
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DebugInstanceOperation) Metadata() (*appenginepb.OperationMetadataV1, error) {
	var meta appenginepb.OperationMetadataV1
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DebugInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DebugInstanceOperation) Name() string {
	return op.lro.Name()
}

// DeleteInstanceOperation manages a long-running operation from DeleteInstance.
type DeleteInstanceOperation struct {
	lro *longrunning.Operation
}

// DeleteInstanceOperation returns a new DeleteInstanceOperation from a given name.
// The name must be that of a previously created DeleteInstanceOperation, possibly from a different process.
func (c *instancesGRPCClient) DeleteInstanceOperation(name string) *DeleteInstanceOperation {
	return &DeleteInstanceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteInstanceOperation) Metadata() (*appenginepb.OperationMetadataV1, error) {
	var meta appenginepb.OperationMetadataV1
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteInstanceOperation) Name() string {
	return op.lro.Name()
}

// InstanceIterator manages a stream of *appenginepb.Instance.
type InstanceIterator struct {
	items    []*appenginepb.Instance
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*appenginepb.Instance, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InstanceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InstanceIterator) Next() (*appenginepb.Instance, error) {
	var item *appenginepb.Instance
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InstanceIterator) bufLen() int {
	return len(it.items)
}

func (it *InstanceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
