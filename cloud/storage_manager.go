//
// cloud/storage_manager.go
//
// Copyright (c) 2016-2017 Junpei Kawamoto
//
// This file is part of Roadie.
//
// Roadie is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Roadie is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Roadie.  If not, see <http://www.gnu.org/licenses/>.
//

package cloud

import (
	"context"
	"io"
	"net/url"
)

// StorageManager defines methods which a storage service provider must provides.
// Each method takes a URL to point a file stored in this storage.
// The URL should be
// - roadie://category/path
// where category is one of
// - script.SourcePrefix
// - script.DataPrefix
// - script.ResultPrefix
type StorageManager interface {

	// Upload a given stream to a given URL.
	Upload(ctx context.Context, loc *url.URL, in io.Reader) error

	// Download a file pointed by a given URL and write it to a given stream.
	Download(ctx context.Context, loc *url.URL, out io.Writer) error

	// GetFileInfo retrieves information of a file pointed by a given URL.
	GetFileInfo(ctx context.Context, loc *url.URL) (*FileInfo, error)

	// List up files of which URLs start with a given URL.
	// It takes a handler; information of found files are sent to it.
	List(ctx context.Context, loc *url.URL, handler FileInfoHandler) error

	// Delete a file pointed by a given URL.
	Delete(ctx context.Context, loc *url.URL) error
}

// FileInfoHandler is a handler to receive a file info.
type FileInfoHandler func(*FileInfo) error
