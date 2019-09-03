/**
 * (C) Copyright IBM Corp. 2019.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package issue998manual : Operations and models for the Issue998V1 service
package deser_problem

import (
	"encoding/json"
)

// Resource : Resource struct.
type Resource struct {

	// The resource id.
	ID string `json:"id" validate:"required"`

	// Information about a resource.
	Info InfoIntf `json:"info,omitempty"`
}

// UnmarshalJSON: A custom de-serializer for instances of Resource.
func (obj *Resource) UnmarshalJSON(data []byte) error {
	var deserTarget resourceDeserTarget
	if err := json.Unmarshal(data, &deserTarget); err != nil {
		return err
	}
	*obj = deserTarget.Resource()
	return nil
}

// ResourceDeserTarget : de-serialization target for the Resource struct.
type resourceDeserTarget struct {

	// The resource id.
	ID string `json:"id" validate:"required"`

	// Information about a resource.
	Info *Info `json:"info,omitempty"`
}

func (obj resourceDeserTarget) Resource() Resource {
	return Resource{
		obj.ID,
		obj.Info,
	}
}

// Info : Information about a resource.
// Models which "extend" this model:
// - InfoFooInfo
// - InfoBarInfo
type Info struct {

	// The foo of the resource.
	Foo string `json:"foo,omitempty"`

	// The bar of the resource.
	Bar string `json:"bar,omitempty"`
}

func (*Info) isaInfo() bool {
	return true
}

type InfoIntf interface {
	isaInfo() bool
}

// Bar : Bar Info.
// This model "extends" Info
type Bar struct {

	// The bar of the resource.
	Bar string `json:"bar" validate:"required"`
}

func (*Bar) isaInfo() bool {
	return true
}

// Foo : Foo Info.
// This model "extends" Info
type Foo struct {

	// The foo of the resource.
	Foo string `json:"foo" validate:"required"`
}

func (*Foo) isaInfo() bool {
	return true
}
