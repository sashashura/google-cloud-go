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

package webrisk

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	webriskpb "google.golang.org/genproto/googleapis/cloud/webrisk/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ComputeThreatListDiff []gax.CallOption
	SearchUris            []gax.CallOption
	SearchHashes          []gax.CallOption
	CreateSubmission      []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("webrisk.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("webrisk.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://webrisk.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ComputeThreatListDiff: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchUris: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchHashes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateSubmission: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Web Risk API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ComputeThreatListDiff(context.Context, *webriskpb.ComputeThreatListDiffRequest, ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error)
	SearchUris(context.Context, *webriskpb.SearchUrisRequest, ...gax.CallOption) (*webriskpb.SearchUrisResponse, error)
	SearchHashes(context.Context, *webriskpb.SearchHashesRequest, ...gax.CallOption) (*webriskpb.SearchHashesResponse, error)
	CreateSubmission(context.Context, *webriskpb.CreateSubmissionRequest, ...gax.CallOption) (*webriskpb.Submission, error)
}

// Client is a client for interacting with Web Risk API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Web Risk API defines an interface to detect malicious URLs on your
// website and in client applications.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ComputeThreatListDiff gets the most recent threat list diffs. These diffs should be applied to
// a local database of hashes to keep it up-to-date. If the local database is
// empty or excessively out-of-date, a complete snapshot of the database will
// be returned. This Method only updates a single ThreatList at a time. To
// update multiple ThreatList databases, this method needs to be called once
// for each list.
func (c *Client) ComputeThreatListDiff(ctx context.Context, req *webriskpb.ComputeThreatListDiffRequest, opts ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error) {
	return c.internalClient.ComputeThreatListDiff(ctx, req, opts...)
}

// SearchUris this method is used to check whether a URI is on a given threatList.
// Multiple threatLists may be searched in a single query.
// The response will list all requested threatLists the URI was found to
// match. If the URI is not found on any of the requested ThreatList an
// empty response will be returned.
func (c *Client) SearchUris(ctx context.Context, req *webriskpb.SearchUrisRequest, opts ...gax.CallOption) (*webriskpb.SearchUrisResponse, error) {
	return c.internalClient.SearchUris(ctx, req, opts...)
}

// SearchHashes gets the full hashes that match the requested hash prefix.
// This is used after a hash prefix is looked up in a threatList
// and there is a match. The client side threatList only holds partial hashes
// so the client must query this method to determine if there is a full
// hash match of a threat.
func (c *Client) SearchHashes(ctx context.Context, req *webriskpb.SearchHashesRequest, opts ...gax.CallOption) (*webriskpb.SearchHashesResponse, error) {
	return c.internalClient.SearchHashes(ctx, req, opts...)
}

// CreateSubmission creates a Submission of a URI suspected of containing phishing content to
// be reviewed. If the result verifies the existence of malicious phishing
// content, the site will be added to the Google’s Social Engineering
// lists (at https://support.google.com/webmasters/answer/6350487/) in order to
// protect users that could get exposed to this threat in the future. Only
// allowlisted projects can use this method during Early Access. Please reach
// out to Sales or your customer engineer to obtain access.
func (c *Client) CreateSubmission(ctx context.Context, req *webriskpb.CreateSubmissionRequest, opts ...gax.CallOption) (*webriskpb.Submission, error) {
	return c.internalClient.CreateSubmission(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Web Risk API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client webriskpb.WebRiskServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new web risk service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Web Risk API defines an interface to detect malicious URLs on your
// website and in client applications.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           webriskpb.NewWebRiskServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ComputeThreatListDiff(ctx context.Context, req *webriskpb.ComputeThreatListDiffRequest, opts ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ComputeThreatListDiff[0:len((*c.CallOptions).ComputeThreatListDiff):len((*c.CallOptions).ComputeThreatListDiff)], opts...)
	var resp *webriskpb.ComputeThreatListDiffResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ComputeThreatListDiff(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) SearchUris(ctx context.Context, req *webriskpb.SearchUrisRequest, opts ...gax.CallOption) (*webriskpb.SearchUrisResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SearchUris[0:len((*c.CallOptions).SearchUris):len((*c.CallOptions).SearchUris)], opts...)
	var resp *webriskpb.SearchUrisResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SearchUris(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) SearchHashes(ctx context.Context, req *webriskpb.SearchHashesRequest, opts ...gax.CallOption) (*webriskpb.SearchHashesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SearchHashes[0:len((*c.CallOptions).SearchHashes):len((*c.CallOptions).SearchHashes)], opts...)
	var resp *webriskpb.SearchHashesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SearchHashes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateSubmission(ctx context.Context, req *webriskpb.CreateSubmissionRequest, opts ...gax.CallOption) (*webriskpb.Submission, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateSubmission[0:len((*c.CallOptions).CreateSubmission):len((*c.CallOptions).CreateSubmission)], opts...)
	var resp *webriskpb.Submission
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateSubmission(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
