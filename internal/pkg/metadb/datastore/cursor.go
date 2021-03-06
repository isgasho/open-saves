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

package datastore

import (
	"errors"

	ds "cloud.google.com/go/datastore"
	"github.com/googleforgames/open-saves/internal/pkg/metadb"
)

var _ metadb.BlobRefCursor = new(blobRefCursor)

type blobRefCursor struct {
	iter *ds.Iterator
}

func newBlobRefIterator(i *ds.Iterator) *blobRefCursor {
	return &blobRefCursor{iter: i}
}

func (i *blobRefCursor) Next() (*metadb.BlobRef, error) {
	if i == nil {
		return nil, errors.New("BlobRefIterator.Next was called on nil")
	}
	if i.iter == nil {
		return nil, errors.New("Iterator is nil")
	}
	var blob metadb.BlobRef
	_, err := i.iter.Next(&blob)
	if err != nil {
		return nil, err
	}
	return &blob, nil
}
