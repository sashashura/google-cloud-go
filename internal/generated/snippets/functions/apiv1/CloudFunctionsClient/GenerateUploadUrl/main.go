// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by cloud.google.com/go/internal/gapicgen/gensnippets. DO NOT EDIT.

// [START cloudfunctions_v1_generated_CloudFunctionsService_GenerateUploadUrl_sync]

package main

import (
	"context"

	functions "cloud.google.com/go/functions/apiv1"
	functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"
)

func main() {
	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &functionspb.GenerateUploadUrlRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/functions/v1#GenerateUploadUrlRequest.
	}
	resp, err := c.GenerateUploadUrl(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END cloudfunctions_v1_generated_CloudFunctionsService_GenerateUploadUrl_sync]
