// Code generated by go-bindata.
// sources:
// assets/fluentd.service
// assets/logcast.service
// assets/queue.service
// assets/roadie.service
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsFluentdService = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x53\x5d\x6f\xdb\x46\x10\x7c\xe7\xaf\x18\x40\x0f\x4e\x00\x8b\xb2\xd5\x8f\x00\x6e\x55\x80\x71\xa5\x44\xad\x2d\x0b\x24\x05\xc3\x08\xfc\x70\x22\x97\xe4\x35\xc7\x3b\xf6\x6e\x4f\x2c\x11\xe4\xbf\x17\x47\xc9\x0e\x9a\x00\x6e\xfb\x42\x90\xcb\x9d\xd9\xd9\xd9\xdd\x49\x34\x41\xa5\x3c\x69\x2e\x63\x47\xf6\x20\x0b\x8a\x26\xc8\xa9\xed\x94\x60\x82\xa9\x70\x8a\xc2\x6b\xc9\x28\xa9\x92\x5a\xb2\x34\x1a\x95\xb1\x4f\xc8\x28\xb0\x5c\x9b\x6e\xb0\xb2\x6e\x18\xaf\x8a\xd7\x98\x5f\x5c\xfe\x38\x9d\x5f\x5c\xbe\xc1\x6f\x5e\x77\x24\xf1\xbb\xe8\x45\x6b\xd8\x8c\xb9\x79\x23\x1d\x2a\xa9\x08\xd2\xa1\x13\x96\x43\xa1\xd4\x88\x52\x52\x3c\x26\x1c\xdf\xc3\xdf\xca\x12\xc1\x99\x8a\x7b\x61\xe9\x0a\x83\xf1\x28\x84\x86\xa5\x52\x3a\xb6\x72\xef\x99\x20\x19\x42\x97\x33\x63\xd1\x9a\x52\x56\x43\x34\x09\x21\xaf\x4b\xb2\xe0\x86\xc0\x64\x5b\x17\x4a\x84\x8f\x77\x9b\x1d\xde\x91\x26\x2b\x14\xb6\x7e\xaf\x64\x81\x1b\x59\x90\x76\x04\xe1\xd0\x85\x88\x6b\xa8\xc4\x3e\xd0\x04\xc0\x2a\x28\xc8\x4e\x0a\xb0\x32\x5e\x97\x22\x38\x70\x0e\x92\xdc\x90\xc5\x81\xac\x0b\x8e\x7c\xf7\x54\xe2\xc4\x77\x0e\x63\xa3\x09\x5e\x09\x0e\xb2\x2d\x4c\x17\x60\xaf\x21\xf4\x80\xe0\xee\x33\xf2\xeb\x9e\xbf\xb4\x56\x42\xea\x91\xb2\x31\x1d\x81\x1b\xc1\xa1\xb3\x5e\x2a\x85\x3d\xc1\x3b\xaa\xbc\x3a\x8f\x26\xd8\x7b\xc6\xfd\x3a\x7f\x7f\xb7\xcb\x91\x6c\x1e\x70\x9f\xa4\x69\xb2\xc9\x1f\x7e\x42\x2f\xb9\x31\x9e\x41\x07\x3a\x32\xc9\xb6\x53\x92\x4a\xf4\xc2\x5a\xa1\x79\x80\xa9\xa2\x09\x6e\x97\xe9\xf5\xfb\x64\x93\x27\x6f\xd7\x37\xeb\xfc\x01\xc6\x62\xb5\xce\x37\xcb\x2c\xc3\xea\x2e\x45\x82\x6d\x92\xe6\xeb\xeb\xdd\x4d\x92\x62\xbb\x4b\xb7\x77\xd9\x32\x06\x32\x0a\xa2\xc2\xc6\xbc\xe0\x6a\x35\xce\xc5\x12\x4a\x62\x21\x95\x3b\x76\xfb\x60\x3c\x5c\x63\xbc\x2a\xd1\x88\x03\xc1\x52\x41\xf2\x40\x25\x04\x0a\xd3\x0d\xff\x3e\xad\x68\x02\xa1\x8c\xae\xc7\x0e\x9f\x96\x07\x58\x57\xd0\x86\xcf\xe1\x88\xf0\x73\xc3\xdc\x5d\xcd\x66\x7d\xdf\xc7\xb5\xf6\xb1\xb1\xf5\x4c\x1d\xd1\x6e\xf6\x4b\xd0\xf1\x61\xa7\x25\x3f\x46\xbf\x92\x2b\xac\x1c\xc7\xb3\x58\x1d\x77\x1a\x37\xa6\xc6\xad\xd0\xa2\xa6\x96\x34\x23\x3b\x1d\x47\x4a\x7f\x7a\x69\xc9\x2d\x4a\x53\x7c\x24\xfb\x7c\x34\x49\xc5\x64\xbf\x0e\x46\x1f\x4e\xb0\xc7\x68\xf9\x17\x15\x19\x0b\xcb\x5b\x4b\x8b\xe9\xcc\x3b\x3b\xdb\x4b\x3d\x3b\x02\xe0\xd8\x74\xcf\xd7\xf4\x72\xaa\x6d\xff\x63\x62\xe7\x95\xc2\x1f\x1f\x4f\x77\x77\x8a\x4e\x6b\x63\x6a\x45\xd3\x13\xc5\x55\x49\x87\x2f\x34\x8b\x6f\x6a\x79\x8d\xe9\x54\x8b\x96\x9e\x6a\x62\x4a\x38\x5b\x6f\xb2\x3c\xd9\x5c\x2f\x17\x9f\x3e\xc5\x1b\xd1\xd2\xe7\xcf\x67\x63\x7c\x97\x2d\xd3\x4d\x72\xbb\x5c\xd8\x71\x1a\x67\x98\x76\xb8\x9c\xbf\x89\x2f\xe2\x8b\xf8\xf2\x6a\xfe\xfd\x7c\xfe\xc3\xf1\xf9\x3f\x64\x99\xee\x1b\x55\xff\x30\x2b\x25\x37\x4a\x17\xaa\x17\x83\x8b\xf2\xa1\xa3\x85\x0b\x3b\x1e\xdc\x5f\x6b\xc7\x42\xa9\xc7\xe8\x5e\x68\xa6\xf2\xed\xb0\x68\xbd\x62\x39\xf5\x8e\x6c\xcc\xc2\xd6\xc4\xd1\xdf\x01\x00\x00\xff\xff\x85\xdf\x1d\x83\x02\x05\x00\x00")

func assetsFluentdServiceBytes() ([]byte, error) {
	return bindataRead(
		_assetsFluentdService,
		"assets/fluentd.service",
	)
}

func assetsFluentdService() (*asset, error) {
	bytes, err := assetsFluentdServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/fluentd.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsLogcastService = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x53\x5d\x6f\xe3\x36\x10\x7c\xe7\xaf\x18\x9c\x5f\xee\x80\x58\x4e\xdc\x8f\x03\xae\x75\x01\x5d\x1a\xdf\xb9\x4d\x9d\x40\x96\x11\x04\x87\x3c\xd0\xd2\xca\xe2\x81\x26\xd5\xe5\xd2\xaa\x80\xfe\xf8\x82\x56\x7c\x29\xf2\xd0\x3e\x08\x20\x57\x9c\xd9\x99\xe5\x70\xa2\x26\xb0\x7e\x5f\xe9\x20\x59\x20\x3e\x9a\x8a\xd4\x04\x25\x1d\x3a\xab\x85\xe0\x1b\x3c\x57\x11\x9d\x11\xd4\xd4\x18\x67\xc4\x78\x87\xc6\x73\xfa\x7a\xcd\xb5\x71\xfb\x44\x82\x03\x85\xa0\xf7\x04\xf1\x58\xda\x48\x4e\xea\x4c\xa5\x06\xd7\xbe\x1b\xd8\xec\x5b\xc1\xdb\xea\x1d\xe6\x97\x57\x3f\x4e\xe7\x97\x57\xef\xf1\x5b\x74\x1d\x19\xfc\xae\x7b\x7d\xf0\xe2\x4f\x67\xcb\xd6\x04\x34\xc6\x12\x4c\x40\xa7\x59\x92\x86\xc2\xeb\xda\xd0\x48\x36\xae\xd3\xdf\x86\x89\x10\x7c\x23\xbd\x66\xfa\x80\xc1\x47\x54\xda\x81\xa9\x36\x41\xd8\xec\xa2\x10\x8c\x40\xbb\x7a\xe6\x19\x07\x5f\x9b\x66\x50\x93\x54\x8a\xae\x26\x86\xb4\x04\x21\x3e\x84\xd4\x22\x6d\x3e\xad\xb7\xf8\x44\x8e\x58\x5b\xdc\xc7\x9d\x35\x15\x6e\x4d\x45\x2e\x10\x74\x40\x97\x2a\xa1\xa5\x1a\xbb\x44\x93\x00\xcb\xa4\x60\xf3\xac\x00\x4b\x1f\x5d\xad\xd3\x70\x2e\x40\x46\x5a\x62\x1c\x89\x43\x1a\xd6\x77\xe7\x16\xcf\x7c\x17\xf0\xac\x26\x78\xab\x25\xc9\x66\xf8\x2e\xc1\xde\x41\xbb\x01\x69\xf0\xdf\x90\xaf\x3d\xbf\x58\xab\x61\xdc\x89\xb2\xf5\x1d\x41\x5a\x2d\xc9\x59\x6f\xac\xc5\x8e\x10\x03\x35\xd1\x5e\xa8\x09\x76\x51\xf0\xb0\x2a\x3f\xdf\x6d\x4b\xe4\xeb\x47\x3c\xe4\x45\x91\xaf\xcb\xc7\x9f\xd0\x1b\x69\x7d\x14\xd0\x91\x46\x26\x73\xe8\xac\xa1\x1a\xbd\x66\xd6\x4e\x06\xf8\x46\x4d\xf0\xc7\x4d\x71\xfd\x39\x5f\x97\xf9\xc7\xd5\xed\xaa\x7c\x84\x67\x2c\x57\xe5\xfa\x66\xb3\xc1\xf2\xae\x40\x8e\xfb\xbc\x28\x57\xd7\xdb\xdb\xbc\xc0\xfd\xb6\xb8\xbf\xdb\xdc\x64\xc0\x86\x92\xa8\x14\xa6\xff\x98\x6a\x73\xba\x17\x26\xd4\x24\xda\xd8\x30\xba\x7d\xf4\x11\xa1\xf5\xd1\xd6\x68\xf5\x91\xc0\x54\x91\x39\x52\x0d\x8d\xca\x77\xc3\xff\xdf\x96\x9a\x40\x5b\xef\xf6\x27\x87\xe7\xf0\x00\xab\x06\xce\xcb\x05\x02\x11\x7e\x6e\x45\xba\x0f\xb3\x59\xdf\xf7\xd9\xde\xc5\xcc\xf3\x7e\x66\x47\x74\x98\xfd\x92\x74\x7c\xd9\x3a\x23\x4f\xea\x57\x0a\x15\x9b\xd3\xf5\x2c\x96\x63\xdc\x71\xfb\x92\xf5\xf0\xaf\xb0\xab\x82\xfe\x8c\x86\x29\x2c\x9a\xe7\xf4\x9f\xdf\x0e\x8f\x12\xce\x0f\x2c\x6f\x84\xf8\xf5\x21\xa5\xbe\x6c\xc6\xd5\x93\xba\xf9\x8b\xaa\x8d\x68\x96\xc5\x6c\x67\xdc\x6c\xa7\x43\x8b\x69\x85\x37\xb3\x18\xf8\x54\xf9\xea\x23\x3b\x6d\x2b\xb1\x98\xc6\x57\xf4\x98\x7a\x7c\x0d\xde\x61\xda\xe0\x6f\x7c\x83\xb8\x4a\x0b\xae\xe6\xef\xb3\xcb\xec\x32\xbb\xc2\xfc\xfb\xf9\xfc\x87\x37\xaa\xa0\x70\xea\xa3\x6d\xaf\x87\x70\xde\x6e\xa8\x5a\xcc\x03\x55\xaa\x1c\x3a\x5a\x84\x94\x8b\xa4\x6f\xe5\x82\x68\x6b\x9f\xd4\x83\x76\x42\xf5\xc7\x61\x71\x88\x56\xcc\x34\x06\xe2\x4c\x34\xef\x49\xd4\x3f\x01\x00\x00\xff\xff\x8d\xe5\x17\x75\x51\x04\x00\x00")

func assetsLogcastServiceBytes() ([]byte, error) {
	return bindataRead(
		_assetsLogcastService,
		"assets/logcast.service",
	)
}

func assetsLogcastService() (*asset, error) {
	bytes, err := assetsLogcastServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/logcast.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsQueueService = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x54\x5d\x6f\xdb\x46\x10\x7c\xe7\xaf\x18\x40\x2f\x09\x60\x89\x71\x5a\xa4\x80\x5b\x15\x50\xfc\x91\xa8\x75\x24\x95\xa2\x6b\x18\x81\x61\x9c\xc8\xa5\x78\xe9\xf1\x8e\xd9\xdb\x13\xad\x18\xfa\xef\xc5\x51\x4a\x52\xb4\x6a\xf2\x26\xae\x76\x76\x67\x76\x77\x6e\x90\x0c\xf0\x31\x50\xa0\x91\x27\xde\xe8\x82\x92\x01\x72\x6a\x5a\xa3\x84\xe0\x2a\x1c\xa2\x08\x56\x0b\x4a\xaa\xb4\xd5\xa2\x9d\x45\xe5\x18\xec\x54\xa9\x69\xd8\xc3\x87\x8d\xb2\x6a\x4d\x9c\xc4\x82\xe7\xae\xdd\xb2\x5e\xd7\x82\x67\xc5\x73\xbc\x7c\x71\xfa\x6a\xf8\xf2\xc5\xe9\x4f\xf8\x2d\xd8\x96\x34\x7e\x57\x9d\x6a\x9c\xb8\x3e\x37\xaf\xb5\x47\xa5\x0d\x41\x7b\xb4\x8a\x25\x76\xcd\xfa\xca\xa3\x3e\x61\xff\x3b\xfe\x5b\x31\x11\xbc\xab\xa4\x53\x4c\x67\xd8\xba\x80\x42\x59\x30\x95\xda\x0b\xeb\x55\x10\x82\x16\x28\x5b\xa6\x8e\xd1\xb8\x52\x57\xdb\x64\x10\x43\xc1\x96\xc4\x90\x9a\x20\xc4\x8d\x8f\x2d\xe2\xc7\x9b\xd9\x0d\xde\x90\x25\x56\x06\x8b\xb0\x32\xba\xc0\xb5\x2e\xc8\x7a\x82\xf2\x68\x63\xc4\xd7\x54\x62\x15\xcb\x44\xc0\x55\x64\xb0\x3c\x30\xc0\x95\x0b\xb6\x54\x71\x1c\x27\x20\x2d\x35\x31\x36\xc4\x3e\x8e\xe7\x87\xcf\x2d\x0e\xf5\x4e\xe0\x38\x19\xe0\x99\x92\x48\x9b\xe1\xda\x08\x7b\x0e\x65\xb7\x88\xa3\xfe\x82\xfc\xb7\xe6\xaf\xd2\x4a\x68\xdb\x97\xac\x5d\x4b\x90\x5a\x49\x54\xd6\x69\x63\xb0\x22\x04\x4f\x55\x30\x27\xc9\x00\xab\x20\xb8\x9d\xe6\x6f\xe7\x37\x39\x26\xb3\x3b\xdc\x4e\xb2\x6c\x32\xcb\xef\x7e\x46\xa7\xa5\x76\x41\x40\x1b\xda\x57\xd2\x4d\x6b\x34\x95\xe8\x14\xb3\xb2\xb2\x85\xab\x92\x01\xde\x5d\x66\xe7\x6f\x27\xb3\x7c\xf2\x7a\x7a\x3d\xcd\xef\xe0\x18\x57\xd3\x7c\x76\xb9\x5c\xe2\x6a\x9e\x61\x82\xc5\x24\xcb\xa7\xe7\x37\xd7\x93\x0c\x8b\x9b\x6c\x31\x5f\x5e\x8e\x80\x25\x45\x52\xf1\x7c\xbe\x31\xd5\xaa\xdf\x0b\x13\x4a\x12\xa5\x8d\xdf\xab\xbd\x73\x01\xbe\x76\xc1\x94\xa8\xd5\x86\xc0\x54\x90\xde\x50\x09\x85\xc2\xb5\xdb\xef\x6f\x2b\x19\x40\x19\x67\xd7\xbd\xc2\xcf\xc7\x03\x4c\x2b\x58\x27\x27\xf0\x44\xf8\xa5\x16\x69\xcf\xd2\xb4\xeb\xba\xd1\xda\x86\x91\xe3\x75\x6a\xf6\x68\x9f\xfe\x1a\x79\xbc\xbf\xb1\x5a\xee\x93\x0b\xf2\x05\xeb\x7e\x3d\xe3\xc3\x1a\xfe\x88\x07\x8e\x77\xfb\x03\xc7\xf2\xe0\x93\x8c\x3e\x06\xcd\xe4\xc7\x95\x09\x64\xa5\xfc\x62\xa0\x49\x25\xc4\xff\x89\x26\xef\x0f\xc0\xfb\xe4\xf2\x91\x8a\xa5\x28\x96\x05\xd3\x38\x0d\x9e\xd3\x95\xb6\x69\x11\xd8\x60\x78\x3d\x47\x64\xea\xcf\xd2\x74\xad\xa5\x0e\xab\x51\xe1\x9a\xf4\xc3\x5f\x07\xcb\xa4\xc7\x3c\x97\x32\x19\x52\x51\x48\xe9\x3a\x6b\x9c\x2a\xd3\xcd\xd3\xd3\xe8\xcf\xfd\x45\xed\x76\x47\x41\x0f\xff\xcc\x78\x30\xda\x86\xc7\x07\xd5\x94\xaf\x7e\x1c\x89\xe2\xd1\xfa\xd3\xff\xb0\x14\xc5\x18\x7e\x7a\xdc\x54\x47\xdd\xff\xbd\xa2\x18\x9e\x23\x65\xe7\x04\xc3\x61\xbc\xeb\x76\x58\xb8\xa6\x75\x96\xac\x78\x9c\x7e\x6d\x39\xee\x93\x8e\xf2\xc6\xd3\xd3\x68\xc1\xee\x03\x15\x32\xbd\xd8\xed\xe2\x67\xbf\x9f\x99\x6a\x68\xb7\x8b\x0f\xca\xfc\x62\x7e\x76\x94\x1d\xb4\x78\x32\x15\x7c\x1d\x24\x4e\x6a\xef\x01\xeb\x45\xd9\xa2\x7f\x39\xb4\x07\x07\x6b\xb5\x5d\x23\x3a\x31\x23\xdf\x93\x51\xa6\x53\x5b\x9f\xe4\xdb\x96\xc6\x3e\x7a\x26\x6e\x73\x1a\x71\xc6\xdc\x27\xb7\xca\x0a\x95\xaf\xb7\xe3\x26\x18\xd1\xc3\xe0\x89\xa3\xda\x35\x49\xf2\x77\x00\x00\x00\xff\xff\x05\x94\x72\x6b\x5d\x05\x00\x00")

func assetsQueueServiceBytes() ([]byte, error) {
	return bindataRead(
		_assetsQueueService,
		"assets/queue.service",
	)
}

func assetsQueueService() (*asset, error) {
	bytes, err := assetsQueueServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/queue.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRoadieService = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xd1\x6e\xdb\x46\x10\x7c\xe7\x57\x0c\xac\x97\x04\x30\xc5\x38\x05\x5a\xc0\xad\x0a\x28\xae\x65\xab\x75\x64\x41\xa2\x60\x18\x86\x1f\x4e\xe4\x52\x3c\xf4\x78\xc7\xee\xed\x49\x25\x82\xfc\x7b\x71\x14\x6d\x03\x11\xd0\xe4\xed\xb8\x37\x33\x37\xb3\xbb\x1c\x25\x23\xb0\x53\xa5\xa6\xb1\x27\xde\xeb\x82\x92\x11\x72\x6a\x5a\xa3\x84\xe0\x2a\x0c\x55\x04\xab\x05\x25\x55\xda\x6a\xd1\xce\xa2\x72\x3c\x10\xd3\x5d\xd1\x26\x51\xe7\xca\xb5\x1d\xeb\x5d\x2d\x78\x57\xbc\xc7\xc7\x0f\x17\x3f\xa7\x1f\x3f\x5c\xfc\x82\x3f\x83\x6d\x49\xe3\x2f\x75\x50\x8d\x13\xd7\x63\xf3\x5a\x7b\x54\xda\x10\xb4\x47\xab\x58\xe2\x5b\xab\xa3\x91\x1e\x70\x3c\xc7\xdb\x8a\x89\xe0\x5d\x25\x07\xc5\x74\x89\xce\x05\x14\xca\x82\xa9\xd4\x5e\x58\x6f\x83\x10\xb4\x40\xd9\x32\x73\x8c\xc6\x95\xba\xea\x92\x51\x2c\x05\x5b\x12\x43\x6a\x82\x10\x37\x3e\x3e\x11\x3f\x6e\x16\x1b\xdc\x90\x25\x56\x06\xcb\xb0\x35\xba\xc0\x9d\x2e\xc8\x7a\x82\xf2\x68\x63\xc5\xd7\x54\x62\x1b\x65\x22\x61\x16\x1d\xac\x07\x07\x98\xb9\x60\x4b\x15\x9b\x70\x0e\xd2\x52\x13\x63\x4f\xec\x63\x53\x7e\x7a\x79\x62\xd0\x3b\x87\xe3\x64\x84\x77\x4a\xa2\x6d\x86\x6b\x23\xed\x3d\x94\xed\x10\x1b\xfc\xca\xfc\x36\xf3\x5b\xb4\x12\xda\xf6\x92\xb5\x6b\x09\x52\x2b\x89\xc9\x0e\xda\x18\x6c\x09\xc1\x53\x15\xcc\x79\x32\xc2\x36\x08\x1e\xe6\xf9\xed\xfd\x26\xc7\x74\xf1\x88\x87\xe9\x6a\x35\x5d\xe4\x8f\xbf\xe2\xa0\xa5\x76\x41\x40\x7b\x3a\x2a\xe9\xa6\x35\x9a\x4a\x1c\x14\xb3\xb2\xd2\xc1\x55\xc9\x08\x9f\xaf\x57\x57\xb7\xd3\x45\x3e\xfd\x34\xbf\x9b\xe7\x8f\x70\x8c\xd9\x3c\x5f\x5c\xaf\xd7\x98\xdd\xaf\x30\xc5\x72\xba\xca\xe7\x57\x9b\xbb\xe9\x0a\xcb\xcd\x6a\x79\xbf\xbe\x1e\x03\x6b\x8a\xa6\xe2\xd2\xfc\x4f\x57\xab\x7e\x2e\x4c\x28\x49\x94\x36\xfe\x98\xf6\xd1\x05\xf8\xda\x05\x53\xa2\x56\x7b\x02\x53\x41\x7a\x4f\x25\x14\x0a\xd7\x76\xdf\x9f\x56\x32\x82\x32\xce\xee\xfa\x84\x2f\xcb\x03\xcc\x2b\x58\x27\xe7\xf0\x44\xf8\xad\x16\x69\x2f\xb3\xec\x70\x38\x8c\x77\x36\x8c\x1d\xef\x32\x73\x64\xfb\xec\xf7\xe8\xe3\x69\x63\xb5\x3c\x27\x7f\x90\x2f\x58\xf7\xe3\x99\x0c\x63\x58\xf7\x05\x5c\xff\x4b\x45\x88\xf5\x64\x45\xff\x04\xcd\xe4\x27\x95\x09\x64\xa5\x7c\xfd\x63\xa6\x95\x10\x9f\x54\x93\xa7\xf5\xf1\xf4\x9c\x44\x8d\xb5\x28\x96\x25\xd3\x24\xcd\x82\xe7\x6c\xab\x6d\x56\xba\xe2\x6f\x62\x78\x71\x2d\xbe\x7c\x19\x2f\x54\x43\x5f\xbf\x7e\x07\xcc\xcd\x0f\x43\xdb\x60\x4c\x04\xcf\x1b\xb5\x3b\x41\xbf\x82\x8b\xc0\x06\xa9\x43\xc6\xce\x49\x76\xec\xc2\xb8\x6b\x0c\xce\x86\xde\x35\x24\xaa\x54\xa2\xc6\x3b\xe7\x76\x86\xc6\xda\x0a\xb1\x55\x26\x2b\x5c\xd3\x06\xa1\xcf\xc3\x7d\xb6\xbf\xc8\xb4\xf5\xa2\x6c\x41\x99\x92\x61\x7f\xfd\x20\x79\x86\xf4\x16\x67\x2f\xd8\x74\x66\xd4\xde\xf1\x25\x6e\x7a\xcd\xb3\x37\x6f\x93\x93\xc0\xc1\x22\xd5\x48\x53\xab\x1a\x7a\xcb\x8e\x74\x7f\x62\xf9\xf2\x24\xc3\x5b\xfa\x78\xbc\xef\xe7\xeb\x23\xb9\x38\x21\x0f\x16\x5c\x7b\xe2\xe0\x9b\xf9\xe4\x5d\x4b\x13\x67\xc9\xd7\x4e\x92\xe4\x69\x1e\x23\x1b\xf3\x9c\x3c\x28\x2b\x54\x7e\xea\x26\x4d\x30\xa2\xd3\xe0\x89\xc7\xa2\x78\x47\x92\xfc\x17\x00\x00\xff\xff\xce\x42\x74\x69\x66\x05\x00\x00")

func assetsRoadieServiceBytes() ([]byte, error) {
	return bindataRead(
		_assetsRoadieService,
		"assets/roadie.service",
	)
}

func assetsRoadieService() (*asset, error) {
	bytes, err := assetsRoadieServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/roadie.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/fluentd.service": assetsFluentdService,
	"assets/logcast.service": assetsLogcastService,
	"assets/queue.service": assetsQueueService,
	"assets/roadie.service": assetsRoadieService,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"fluentd.service": &bintree{assetsFluentdService, map[string]*bintree{}},
		"logcast.service": &bintree{assetsLogcastService, map[string]*bintree{}},
		"queue.service": &bintree{assetsQueueService, map[string]*bintree{}},
		"roadie.service": &bintree{assetsRoadieService, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

