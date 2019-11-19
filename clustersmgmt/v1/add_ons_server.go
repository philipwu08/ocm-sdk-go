/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// AddOnsServer represents the interface the manages the 'add_ons' resource.
type AddOnsServer interface {

	// Add handles a request for the 'add' method.
	//
	// Create a new add-on and add it to the collection of add-ons.
	Add(ctx context.Context, request *AddOnsAddServerRequest, response *AddOnsAddServerResponse) error

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of add-ons.
	List(ctx context.Context, request *AddOnsListServerRequest, response *AddOnsListServerResponse) error

	// Addon returns the target 'add_on' server for the given identifier.
	//
	// Returns a reference to the service that manages a specific add-on.
	Addon(id string) AddOnServer
}

// AddOnsAddServerRequest is the request for the 'add' method.
type AddOnsAddServerRequest struct {
	body *AddOn
}

// Body returns the value of the 'body' parameter.
//
// Description of the add-on.
func (r *AddOnsAddServerRequest) Body() *AddOn {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Description of the add-on.
func (r *AddOnsAddServerRequest) GetBody() (value *AddOn, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *AddOnsAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(addOnData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// AddOnsAddServerResponse is the response for the 'add' method.
type AddOnsAddServerResponse struct {
	status int
	err    *errors.Error
	body   *AddOn
}

// Body sets the value of the 'body' parameter.
//
// Description of the add-on.
func (r *AddOnsAddServerResponse) Body(value *AddOn) *AddOnsAddServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *AddOnsAddServerResponse) Status(value int) *AddOnsAddServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *AddOnsAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// AddOnsListServerRequest is the request for the 'list' method.
type AddOnsListServerRequest struct {
	order  *string
	page   *int
	search *string
	size   *int
	total  *int
}

// Order returns the value of the 'order' parameter.
//
// Order criteria.
//
// The syntax of this parameter is similar to the syntax of the _order by_ clause of
// a SQL statement, but using the names of the attributes of the add-on instead of
// the names of the columns of a table. For example, in order to sort the add-ons
// descending by name the value should be:
//
// [source,sql]
// ----
// name desc
// ----
//
// If the parameter isn't provided, or if the value is empty, then the order of the
// results is undefined.
func (r *AddOnsListServerRequest) Order() string {
	if r != nil && r.order != nil {
		return *r.order
	}
	return ""
}

// GetOrder returns the value of the 'order' parameter and
// a flag indicating if the parameter has a value.
//
// Order criteria.
//
// The syntax of this parameter is similar to the syntax of the _order by_ clause of
// a SQL statement, but using the names of the attributes of the add-on instead of
// the names of the columns of a table. For example, in order to sort the add-ons
// descending by name the value should be:
//
// [source,sql]
// ----
// name desc
// ----
//
// If the parameter isn't provided, or if the value is empty, then the order of the
// results is undefined.
func (r *AddOnsListServerRequest) GetOrder() (value string, ok bool) {
	ok = r != nil && r.order != nil
	if ok {
		value = *r.order
	}
	return
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *AddOnsListServerRequest) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *AddOnsListServerRequest) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Search returns the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause of an
// SQL statement, but using the names of the attributes of the add-on instead of
// the names of the columns of a table. For example, in order to retrieve all the
// add-ons with a name starting with `my` the value should be:
//
// [source,sql]
// ----
// name like 'my%'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the add-ons
// that the user has permission to see will be returned.
func (r *AddOnsListServerRequest) Search() string {
	if r != nil && r.search != nil {
		return *r.search
	}
	return ""
}

// GetSearch returns the value of the 'search' parameter and
// a flag indicating if the parameter has a value.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause of an
// SQL statement, but using the names of the attributes of the add-on instead of
// the names of the columns of a table. For example, in order to retrieve all the
// add-ons with a name starting with `my` the value should be:
//
// [source,sql]
// ----
// name like 'my%'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the add-ons
// that the user has permission to see will be returned.
func (r *AddOnsListServerRequest) GetSearch() (value string, ok bool) {
	ok = r != nil && r.search != nil
	if ok {
		value = *r.search
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *AddOnsListServerRequest) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *AddOnsListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *AddOnsListServerRequest) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *AddOnsListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// AddOnsListServerResponse is the response for the 'list' method.
type AddOnsListServerResponse struct {
	status int
	err    *errors.Error
	items  *AddOnList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of add-ons.
func (r *AddOnsListServerResponse) Items(value *AddOnList) *AddOnsListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *AddOnsListServerResponse) Page(value int) *AddOnsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *AddOnsListServerResponse) Size(value int) *AddOnsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *AddOnsListServerResponse) Total(value int) *AddOnsListServerResponse {
	r.total = &value
	return r
}

// Status sets the status code.
func (r *AddOnsListServerResponse) Status(value int) *AddOnsListServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *AddOnsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(addOnsListServerResponseData)
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	err = encoder.Encode(data)
	return err
}

// addOnsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type addOnsListServerResponseData struct {
	Items addOnListData "json:\"items,omitempty\""
	Page  *int          "json:\"page,omitempty\""
	Size  *int          "json:\"size,omitempty\""
	Total *int          "json:\"total,omitempty\""
}

// dispatchAddOns navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchAddOns(w http.ResponseWriter, r *http.Request, server AddOnsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "POST":
			adaptAddOnsAddRequest(w, r, server)
		case "GET":
			adaptAddOnsListRequest(w, r, server)
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	} else {
		switch segments[0] {
		default:
			target := server.Addon(segments[0])
			if target == nil {
				errors.SendNotFound(w, r)
				return
			}
			dispatchAddOn(w, r, target, segments[1:])
		}
	}
}

// readAddOnsAddRequest reads the given HTTP requests and translates it
// into an object of type AddOnsAddServerRequest.
func readAddOnsAddRequest(r *http.Request) (*AddOnsAddServerRequest, error) {
	var err error
	result := new(AddOnsAddServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}

// writeAddOnsAddResponse translates the given request object into an
// HTTP response.
func writeAddOnsAddResponse(w http.ResponseWriter, r *AddOnsAddServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}

// adaptAddOnsAddRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptAddOnsAddRequest(w http.ResponseWriter, r *http.Request, server AddOnsServer) {
	request, err := readAddOnsAddRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(AddOnsAddServerResponse)
	response.status = 201
	err = server.Add(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeAddOnsAddResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// readAddOnsListRequest reads the given HTTP requests and translates it
// into an object of type AddOnsListServerRequest.
func readAddOnsListRequest(r *http.Request) (*AddOnsListServerRequest, error) {
	var err error
	result := new(AddOnsListServerRequest)
	query := r.URL.Query()
	result.order, err = helpers.ParseString(query, "order")
	if err != nil {
		return nil, err
	}
	result.page, err = helpers.ParseInteger(query, "page")
	if err != nil {
		return nil, err
	}
	result.search, err = helpers.ParseString(query, "search")
	if err != nil {
		return nil, err
	}
	result.size, err = helpers.ParseInteger(query, "size")
	if err != nil {
		return nil, err
	}
	result.total, err = helpers.ParseInteger(query, "total")
	if err != nil {
		return nil, err
	}
	return result, err
}

// writeAddOnsListResponse translates the given request object into an
// HTTP response.
func writeAddOnsListResponse(w http.ResponseWriter, r *AddOnsListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}

// adaptAddOnsListRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptAddOnsListRequest(w http.ResponseWriter, r *http.Request, server AddOnsServer) {
	request, err := readAddOnsListRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(AddOnsListServerResponse)
	response.status = 200
	err = server.List(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeAddOnsListResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}