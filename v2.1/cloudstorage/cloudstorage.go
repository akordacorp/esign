// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package cloudstorage implements the DocuSign SDK
// category CloudStorage.
//
// Use the Cloud Storage category to list the user's cloud storage document contents.
//
// It is also used to manage the user's authentication/accounts with cloud storage service providers.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/reference/CloudStorage
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/v2.1/cloudstorage"
//       "github.com/jfcote87/esign/v2.1/model"
//   )
//   ...
//   cloudstorageService := cloudstorage.New(esignCredential)
package cloudstorage // import "github.com/jfcote87/esign/v2.1/cloudstorage"

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2.1/model"
)

// Service implements DocuSign CloudStorage Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a cloudstorage service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// List gets a list of items from a cloud storage provider.
//
// https://developers.docusign.com/esign-rest-api/reference/cloudstorage/cloudstorage/list
//
// SDK Method CloudStorage::list
func (s *Service) List(folderID string, serviceID string, userID string) *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "cloud_storage", serviceID, "folders", folderID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ListOp implements DocuSign API SDK CloudStorage::list
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.ExternalFolder, error) {
	var res *model.ExternalFolder
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CloudStorageFolderPath is the file path to a cloud storage folder.
func (op *ListOp) CloudStorageFolderPath(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("cloud_storage_folder_path", val)
	}
	return op
}

// CloudStorageFolderidPlain is a plain-text folder id that you can use as an alternative to the existing folder id. This property is mainly used for rooms. Enter multiple folder ids as a comma-separated list.
func (op *ListOp) CloudStorageFolderidPlain(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("cloud_storage_folderid_plain", val)
	}
	return op
}

// Count is an optional value that sets how many items are included in the response.
//
// The default setting for this is 25.
func (op *ListOp) Count(val int) *ListOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// Order (Optional) The order in which to sort the results.
//
// Valid values are:
//
//
// * `asc`: Ascending order.
// * `desc`: Descending order.
func (op *ListOp) Order(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("order", val)
	}
	return op
}

// OrderBy (Optional) The file attribute to use to sort the results.
//
// Valid values are:
//
// * `modified`
// * `name`
func (op *ListOp) OrderBy(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("order_by", val)
	}
	return op
}

// SearchText use this parameter to search for specific text.
func (op *ListOp) SearchText(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("search_text", val)
	}
	return op
}

// StartPosition is the starting index position in the result set from which to start returning values. The default setting is `0`.
func (op *ListOp) StartPosition(val int) *ListOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// ListFolders retrieves a list of all the items in a specified folder from the specified cloud storage provider.
//
// https://developers.docusign.com/esign-rest-api/reference/cloudstorage/cloudstorage/listfolders
//
// SDK Method CloudStorage::listFolders
func (s *Service) ListFolders(serviceID string, userID string) *ListFoldersOp {
	return &ListFoldersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "cloud_storage", serviceID, "folders"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ListFoldersOp implements DocuSign API SDK CloudStorage::listFolders
type ListFoldersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListFoldersOp) Do(ctx context.Context) (*model.ExternalFolder, error) {
	var res *model.ExternalFolder
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CloudStorageFolderPath is a comma separated list of folder IDs included in the request.
func (op *ListFoldersOp) CloudStorageFolderPath(val ...string) *ListFoldersOp {
	if op != nil {
		op.QueryOpts.Set("cloud_storage_folder_path", strings.Join(val, ","))
	}
	return op
}

// Count is an optional value that sets how many items are included in the response.
//
// The default setting for this is 25.
func (op *ListFoldersOp) Count(val int) *ListFoldersOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// Order (Optional) The order in which to sort the results.
//
// Valid values are:
//
//
// * `asc`: Ascending order.
// * `desc`: Descending order.
func (op *ListFoldersOp) Order(val string) *ListFoldersOp {
	if op != nil {
		op.QueryOpts.Set("order", val)
	}
	return op
}

// OrderBy (Optional) The file attribute to use to sort the results.
//
// Valid values are:
//
// * `modified`
// * `name`
func (op *ListFoldersOp) OrderBy(val string) *ListFoldersOp {
	if op != nil {
		op.QueryOpts.Set("order_by", val)
	}
	return op
}

// SearchText use this parameter to search for specific text.
func (op *ListFoldersOp) SearchText(val string) *ListFoldersOp {
	if op != nil {
		op.QueryOpts.Set("search_text", val)
	}
	return op
}

// StartPosition indicates the starting point of the first item included in the response set. It uses a 0-based index. The default setting for this is 0.
func (op *ListFoldersOp) StartPosition(val int) *ListFoldersOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// ProvidersCreate configures the redirect URL information  for one or more cloud storage providers for the specified user.
//
// https://developers.docusign.com/esign-rest-api/reference/cloudstorage/cloudstorageproviders/create
//
// SDK Method CloudStorage::createProvider
func (s *Service) ProvidersCreate(userID string, cloudStorageProviders *model.CloudStorageProviders) *ProvidersCreateOp {
	return &ProvidersCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"users", userID, "cloud_storage"}, "/"),
		Payload:    cloudStorageProviders,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ProvidersCreateOp implements DocuSign API SDK CloudStorage::createProvider
type ProvidersCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ProvidersCreateOp) Do(ctx context.Context) (*model.CloudStorageProviders, error) {
	var res *model.CloudStorageProviders
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ProvidersDelete deletes the user authentication information for the specified cloud storage provider.
//
// https://developers.docusign.com/esign-rest-api/reference/cloudstorage/cloudstorageproviders/delete
//
// SDK Method CloudStorage::deleteProvider
func (s *Service) ProvidersDelete(serviceID string, userID string) *ProvidersDeleteOp {
	return &ProvidersDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"users", userID, "cloud_storage", serviceID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ProvidersDeleteOp implements DocuSign API SDK CloudStorage::deleteProvider
type ProvidersDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ProvidersDeleteOp) Do(ctx context.Context) (*model.CloudStorageProviders, error) {
	var res *model.CloudStorageProviders
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ProvidersDeleteList deletes the user authentication information for one or more cloud storage providers.
//
// https://developers.docusign.com/esign-rest-api/reference/cloudstorage/cloudstorageproviders/deletelist
//
// SDK Method CloudStorage::deleteProviders
func (s *Service) ProvidersDeleteList(userID string, cloudStorageProviders *model.CloudStorageProviders) *ProvidersDeleteListOp {
	return &ProvidersDeleteListOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"users", userID, "cloud_storage"}, "/"),
		Payload:    cloudStorageProviders,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ProvidersDeleteListOp implements DocuSign API SDK CloudStorage::deleteProviders
type ProvidersDeleteListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ProvidersDeleteListOp) Do(ctx context.Context) (*model.CloudStorageProviders, error) {
	var res *model.CloudStorageProviders
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ProvidersGet gets the specified Cloud Storage Provider configuration for the User.
//
// https://developers.docusign.com/esign-rest-api/reference/cloudstorage/cloudstorageproviders/get
//
// SDK Method CloudStorage::getProvider
func (s *Service) ProvidersGet(serviceID string, userID string) *ProvidersGetOp {
	return &ProvidersGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "cloud_storage", serviceID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ProvidersGetOp implements DocuSign API SDK CloudStorage::getProvider
type ProvidersGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ProvidersGetOp) Do(ctx context.Context) (*model.CloudStorageProviders, error) {
	var res *model.CloudStorageProviders
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// RedirectURL is the URL the user is redirected to after the cloud storage provider authenticates the user. Using this will append the redirectUrl to the authenticationUrl.
//
// The redirectUrl is restricted to URLs in the docusign.com or docusign.net domains.
func (op *ProvidersGetOp) RedirectURL(val string) *ProvidersGetOp {
	if op != nil {
		op.QueryOpts.Set("redirectUrl", val)
	}
	return op
}

// ProvidersList get the Cloud Storage Provider configuration for the specified user.
//
// https://developers.docusign.com/esign-rest-api/reference/cloudstorage/cloudstorageproviders/list
//
// SDK Method CloudStorage::listProviders
func (s *Service) ProvidersList(userID string) *ProvidersListOp {
	return &ProvidersListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"users", userID, "cloud_storage"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ProvidersListOp implements DocuSign API SDK CloudStorage::listProviders
type ProvidersListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ProvidersListOp) Do(ctx context.Context) (*model.CloudStorageProviders, error) {
	var res *model.CloudStorageProviders
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// RedirectURL is the URL the user is redirected to after the cloud storage provider authenticates the user. Using this will append the redirectUrl to the authenticationUrl.
//
// The redirectUrl is restricted to URLs in the docusign.com or docusign.net domains.
func (op *ProvidersListOp) RedirectURL(val string) *ProvidersListOp {
	if op != nil {
		op.QueryOpts.Set("redirectUrl", val)
	}
	return op
}
