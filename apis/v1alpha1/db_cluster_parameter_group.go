// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DBClusterParameterGroupSpec defines the desired state of DBClusterParameterGroup.
//
// Contains the details of an Amazon RDS DB cluster parameter group.
//
// This data type is used as a response element in the DescribeDBClusterParameterGroups
// action.
type DBClusterParameterGroupSpec struct {

	// The description for the DB cluster parameter group.
	// +kubebuilder:validation:Required
	Description *string `json:"description"`
	// The DB cluster parameter group family name. A DB cluster parameter group
	// can be associated with one and only one DB cluster parameter group family,
	// and can be applied only to a DB cluster running a database engine and engine
	// version compatible with that DB cluster parameter group family.
	//
	// # Aurora MySQL
	//
	// Example: aurora5.6, aurora-mysql5.7, aurora-mysql8.0
	//
	// # Aurora PostgreSQL
	//
	// Example: aurora-postgresql9.6
	//
	// # RDS for MySQL
	//
	// Example: mysql8.0
	//
	// # RDS for PostgreSQL
	//
	// Example: postgres12
	//
	// To list all of the available parameter group families for a DB engine, use
	// the following command:
	//
	// aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"
	// --engine <engine>
	//
	// For example, to list all of the available parameter group families for the
	// Aurora PostgreSQL DB engine, use the following command:
	//
	// aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"
	// --engine aurora-postgresql
	//
	// The output contains duplicates.
	//
	// The following are the valid DB engine values:
	//
	//   - aurora (for MySQL 5.6-compatible Aurora)
	//
	//   - aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)
	//
	//   - aurora-postgresql
	//
	//   - mysql
	//
	//   - postgres
	//
	// +kubebuilder:validation:Required
	Family *string `json:"family"`
	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	//   - Must not match the name of an existing DB cluster parameter group.
	//
	// This value is stored as a lowercase string.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// These are ONLY user-defined parameter overrides for the DB cluster parameter group. This does not contain default or system parameters.
	ParameterOverrides map[string]*string `json:"parameterOverrides,omitempty"`
	// A list of parameters in the DB cluster parameter group to modify.
	//
	// Valid Values (for the application method): immediate | pending-reboot
	//
	// You can use the immediate value with dynamic parameters only. You can use
	// the pending-reboot value for both dynamic and static parameters.
	//
	// When the application method is immediate, changes to dynamic parameters are
	// applied immediately to the DB clusters associated with the parameter group.
	// When the application method is pending-reboot, changes to dynamic and static
	// parameters are applied after a reboot without failover to the DB clusters
	// associated with the parameter group.
	//
	// DEPRECATED - do not use.  Prefer ParameterOverrides instead.
	Parameters []*Parameter `json:"parameters,omitempty"`
	// Tags to assign to the DB cluster parameter group.
	Tags []*Tag `json:"tags,omitempty"`
}

// DBClusterParameterGroupStatus defines the observed state of DBClusterParameterGroup
type DBClusterParameterGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Provides a list of parameters for the DB cluster parameter group.
	// +kubebuilder:validation:Optional
	ParameterOverrideStatuses []*Parameter `json:"parameterOverrideStatuses,omitempty"`
}

// DBClusterParameterGroup is the Schema for the DBClusterParameterGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type DBClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBClusterParameterGroupSpec   `json:"spec,omitempty"`
	Status            DBClusterParameterGroupStatus `json:"status,omitempty"`
}

// DBClusterParameterGroupList contains a list of DBClusterParameterGroup
// +kubebuilder:object:root=true
type DBClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBClusterParameterGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBClusterParameterGroup{}, &DBClusterParameterGroupList{})
}
