//
// cloud/storage_servicer.go
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
)

// StorageServicer defines API which a storage service provider must have.
type StorageServicer interface {

	// Upload a given stream with a given file name; returned string represents
	// a URI assosiated with the uploaded file.
	Upload(ctx context.Context, filename string, in io.Reader) (string, error)

	// Download a file associated with a given file name and write it to a given
	// writer.
	Download(ctx context.Context, filename string, out io.Writer) error

	// GetFileInfo gets file information of a given filename.
	GetFileInfo(ctx context.Context, filename string) (*FileInfo, error)

	// List up files matching a given prefix.
	// It takes a handler; information of found files are sent to it.
	List(ctx context.Context, prefix string, handler FileInfoHandler) error

	// Delete a given file.
	Delete(ctx context.Context, filename string) error

	// Close this service.
	Close() error
}

// FileInfoHandler is a handler to recieve a file info.
type FileInfoHandler func(*FileInfo) error
