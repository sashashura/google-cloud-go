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

// [START dataflow_v1beta3_generated_JobsV1Beta3_CreateJob_sync]

package main

import (
	"context"

	dataflow "cloud.google.com/go/dataflow/apiv1beta3"
	dataflowpb "google.golang.org/genproto/googleapis/dataflow/v1beta3"
)

func main() {
	ctx := context.Background()
	c, err := dataflow.NewJobsV1Beta3Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataflowpb.CreateJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/dataflow/v1beta3#CreateJobRequest.
	}
	resp, err := c.CreateJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END dataflow_v1beta3_generated_JobsV1Beta3_CreateJob_sync]
