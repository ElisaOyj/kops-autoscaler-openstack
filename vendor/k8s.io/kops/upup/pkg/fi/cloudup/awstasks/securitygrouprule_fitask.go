/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by ""fitask" -type=SecurityGroupRule"; DO NOT EDIT

package awstasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// SecurityGroupRule

// JSON marshaling boilerplate
type realSecurityGroupRule SecurityGroupRule

// UnmarshalJSON implements conversion to JSON, supporting an alternate specification of the object as a string
func (o *SecurityGroupRule) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realSecurityGroupRule
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = SecurityGroupRule(r)
	return nil
}

var _ fi.HasLifecycle = &SecurityGroupRule{}

// GetLifecycle returns the Lifecycle of the object, implementing fi.HasLifecycle
func (o *SecurityGroupRule) GetLifecycle() *fi.Lifecycle {
	return o.Lifecycle
}

// SetLifecycle sets the Lifecycle of the object, implementing fi.SetLifecycle
func (o *SecurityGroupRule) SetLifecycle(lifecycle fi.Lifecycle) {
	o.Lifecycle = &lifecycle
}

var _ fi.HasName = &SecurityGroupRule{}

// GetName returns the Name of the object, implementing fi.HasName
func (o *SecurityGroupRule) GetName() *string {
	return o.Name
}

// SetName sets the Name of the object, implementing fi.SetName
func (o *SecurityGroupRule) SetName(name string) {
	o.Name = &name
}

// String is the stringer function for the task, producing readable output using fi.TaskAsString
func (o *SecurityGroupRule) String() string {
	return fi.TaskAsString(o)
}
