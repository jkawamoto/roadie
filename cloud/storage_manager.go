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
)

// StorageManager defines API which a storage service provider must have.
type StorageManager interface {

	// Upload a given stream with a given container and file name; returned string
	// represents a URI associated with the uploaded file.
	Upload(ctx context.Context, container, filename string, in io.Reader) (string, error)

	// Download a file associated with a given container and file name and write
	// it to a given writer.
	Download(ctx context.Context, container, filename string, out io.Writer) error

	// GetFileInfo gets file information of a given container and filename.
	GetFileInfo(ctx context.Context, container, filename string) (*FileInfo, error)

	// List up files matching a given prefix in a given container.
	// It takes a handler; information of found files are sent to it.
	List(ctx context.Context, container, prefix string, handler FileInfoHandler) error

	// Delete a given file in a given container.
	Delete(ctx context.Context, container, filename string) error
}

// FileInfoHandler is a handler to receive a file info.
type FileInfoHandler func(*FileInfo) error
