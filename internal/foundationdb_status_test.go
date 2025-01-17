/*
 * foundationdb_status_test.go
 *
 * This source file is part of the FoundationDB open source project
 *
 * Copyright 2021 Apple Inc. and the FoundationDB project authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package internal

import (
	"github.com/FoundationDB/fdb-kubernetes-operator/api/v1beta2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Internal FoundationDBStatus", func() {
	When("parsing the status for coordinators", func() {
		type testCase struct {
			status   *v1beta2.FoundationDBStatus
			expected map[string]struct{}
		}

		DescribeTable("parse the status",
			func(tc testCase) {
				coordinators := GetCoordinatorsFromStatus(tc.status)
				Expect(coordinators).To(Equal(tc.expected))
			},
			Entry("no coordinators",
				testCase{
					status:   &v1beta2.FoundationDBStatus{},
					expected: map[string]struct{}{},
				}),
			Entry("single coordinators",
				testCase{
					status: &v1beta2.FoundationDBStatus{
						Cluster: v1beta2.FoundationDBStatusClusterInfo{
							Processes: map[string]v1beta2.FoundationDBStatusProcessInfo{
								"foo": {
									Locality: map[string]string{
										v1beta2.FDBLocalityInstanceIDKey: "foo",
									},
									Roles: []v1beta2.FoundationDBStatusProcessRoleInfo{
										{
											Role: "coordinator",
										},
									},
								},
								"bar": {
									Locality: map[string]string{
										v1beta2.FDBLocalityInstanceIDKey: "bar",
									},
								},
							},
						},
					},
					expected: map[string]struct{}{
						"foo": {},
					},
				}),
			Entry("multiple coordinators",
				testCase{
					status: &v1beta2.FoundationDBStatus{
						Cluster: v1beta2.FoundationDBStatusClusterInfo{
							Processes: map[string]v1beta2.FoundationDBStatusProcessInfo{
								"foo": {
									Locality: map[string]string{
										v1beta2.FDBLocalityInstanceIDKey: "foo",
									},
									Roles: []v1beta2.FoundationDBStatusProcessRoleInfo{
										{
											Role: "coordinator",
										},
									},
								},
								"bar": {
									Locality: map[string]string{
										v1beta2.FDBLocalityInstanceIDKey: "bar",
									},
									Roles: []v1beta2.FoundationDBStatusProcessRoleInfo{
										{
											Role: "coordinator",
										},
									},
								},
							},
						},
					},
					expected: map[string]struct{}{
						"foo": {},
						"bar": {},
					},
				}),
		)
	})
})
