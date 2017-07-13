// Code generated by go-bindata.
// sources:
// js/event.js
// js/job.js
// js/mock8s.js
// js/run.js
// js/run_mock.js
// js/runner.js
// js/waitgroup.js
// DO NOT EDIT!

package lib

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

var _jsEventJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x4a\xc6\x30\x10\x85\xf7\x39\xc5\xdb\x55\x41\xd2\x13\x74\xe1\x42\x70\xed\x0d\xa6\xc9\xd4\x04\x62\x52\x33\x53\x6b\x91\xde\x5d\x42\x94\x9f\xf2\xef\x86\x79\x8f\xef\x9b\x19\x47\xbc\x7c\x71\xd6\x57\xca\x3e\x71\x85\x67\x71\x35\xce\x2c\xd0\xc0\x48\x51\x14\x65\x01\xb7\x4a\x5b\x91\xe2\xd9\x45\x8f\x28\xa0\x9d\x2a\xa3\x2c\xd6\x2c\x5b\x76\x1a\x4b\xbe\x90\x1e\x1e\xf1\x63\x80\xce\xaf\x47\x47\x20\xfc\x69\xde\x59\xbb\x61\xa5\x4a\x1f\x18\x3c\x29\x0d\x4f\xd8\x43\x74\xa1\xc1\x5b\x34\x17\x7f\x34\x79\x9b\x2b\x7f\x6e\x2c\x6a\x0d\xa0\x21\x8a\x5d\x37\x09\x98\xf0\x6f\x6e\xb2\xf3\x96\xa5\xf4\xd6\xfb\x77\x95\xd3\x18\xfe\x5e\x4b\x55\xb1\x97\xb7\xa7\xcb\xed\xe6\x37\x00\x00\xff\xff\x38\xd9\xc5\x1d\x16\x01\x00\x00")

func jsEventJsBytes() ([]byte, error) {
	return bindataRead(
		_jsEventJs,
		"js/event.js",
	)
}

func jsEventJs() (*asset, error) {
	bytes, err := jsEventJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/event.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsJobJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x54\xc1\x6e\x1b\x37\x10\xbd\xef\x57\xbc\xea\x24\xb5\x8a\xd6\x41\xda\x1e\x62\xec\xc1\xa8\x1d\x20\x29\x50\x04\x8d\x8b\x1e\x8a\xa2\x98\xe5\xce\x6a\x19\x73\x49\x95\x1c\x4a\x16\x02\xff\x7b\x41\x72\x25\xaf\x95\xf4\x62\x83\x23\xce\x7b\xb3\xef\x3d\x4e\x5d\xe3\x7e\x60\x74\xdc\x53\x34\x02\x3d\xd2\x96\xa1\x03\x82\x38\xf5\x80\xd8\x46\x2b\x11\xaf\x7f\xde\x5c\xfd\x88\x1f\x30\xd2\x03\x83\x6c\x87\xad\x96\x4d\xb5\x27\x0f\x52\xba\x7b\x9f\x7b\x1a\x2c\xd2\xe1\x55\x69\x79\x6b\x48\x38\xc8\xa2\xaa\x2e\x08\x84\xfd\xa8\x2d\x19\xf0\x18\x0d\x89\xf3\x90\x81\x04\x9f\x5d\x0b\xa1\xf0\x10\x70\xd0\xc6\xa0\x65\xf0\x23\xab\x28\xdc\x21\xda\x8e\x7d\x61\x9b\x40\xee\x4f\x18\x0d\x16\x75\xab\x6d\x1d\x86\x42\xf4\xd1\x3b\x71\x72\xdc\x31\x7a\xe7\xf1\xc1\xb5\x9b\xaa\x8f\x56\x89\x76\x36\x9d\x96\x96\x46\x5e\x17\x9e\x15\xbe\x54\x40\x02\x1d\x8f\x68\x20\x83\x0e\x55\x05\xe8\x1e\xcb\xef\xf8\x71\xe7\xbc\x84\xcd\x3f\xbc\x67\x2b\xe5\x26\x20\x83\x77\x07\x2c\x72\x0d\xd6\x09\x7a\x17\x6d\xb7\xa8\x80\xa7\xd4\x59\xd7\xf8\x8d\x46\x3e\xcd\xaf\xdc\xc8\x90\x81\xb1\xf3\xdc\xeb\xc7\x3c\x50\x3e\xba\xae\x56\xce\xf6\x7a\x3b\xd2\x0e\x69\xa0\xb0\xa9\x90\xf9\x37\xe9\x84\x26\x17\xaf\x27\xc8\x4f\x03\x1b\x93\xfc\x48\xbd\xc2\x17\xd2\x1d\x06\xad\x86\xb9\x6e\x3e\xda\x93\x5e\x13\x66\xc8\x00\xcd\xa5\x74\xd7\x15\x50\x18\xee\x73\xf7\xc4\x60\x74\x10\xb8\x7e\x82\x14\x97\x00\x37\xc9\xbf\x23\xc8\xcf\x3c\xd1\x16\x81\xff\x8d\x6c\x15\x43\xdb\xa0\x3b\x86\xeb\x0b\x1e\xa1\x50\x2e\x27\x67\x56\xe7\x51\x0a\x68\x83\xbf\xfe\xbe\x9e\x94\x9e\x39\x71\x71\x27\xff\x9f\x49\x7b\x03\xe5\x8c\xe1\x62\xa5\xeb\xb3\x46\xf5\x9e\x4c\x64\xec\x48\xfb\x90\x6a\x6c\xf7\xda\x3b\x3b\x26\x7f\xf6\xe4\x35\xb5\x66\x26\x2e\xdb\x3d\x1a\x7c\x79\x3a\x29\x9b\x42\x59\xd2\x9e\x02\x4d\x16\x6e\x97\xc0\xc9\x40\x68\x7b\xee\xd2\x53\xb6\xcf\x39\x3f\xb5\xef\x5c\x97\xed\x4e\x4f\x85\x05\xed\x31\x49\xb5\x5c\xad\x33\x9a\x72\x56\x48\xdb\xa2\x69\x76\x35\x69\x5a\xdc\x87\xf2\x4c\xc2\xdd\x99\x62\x42\x9a\x70\x93\x83\x81\x6d\x17\xf2\x8f\xe5\x5d\x38\xfc\x1a\x5b\xf6\x96\x65\xf6\x3d\xe9\x62\x83\x53\xbc\x97\x2f\x54\x6c\x49\x3d\x6c\x7d\x8a\xe7\xf2\x22\xcb\xcf\x77\x0e\xa4\x65\xb9\xaa\x72\xc1\xb3\x44\x6f\xcb\x23\x00\x8a\x44\x17\x40\xff\xcb\x75\x12\xa2\xc9\x02\xa4\xd2\x1a\x17\xa4\xd7\x67\xd0\xba\x46\xe2\xfd\xc3\x8a\x36\xb7\xce\x66\xf9\x06\xf6\xe5\xbd\x26\xb2\x03\xf9\x2e\x40\xb9\x71\x47\xa2\x5b\x6d\xb4\x1c\xd7\x68\xa3\xa0\x73\x1c\xd2\xab\x1b\xb4\xcd\xee\xd4\x35\x6e\xef\x3e\xfe\x7e\xf7\xcb\xcd\xfd\xdd\xed\x5b\xfc\x39\x6d\x0d\xcf\xa3\xdb\x73\x87\x2e\x7a\x6d\xb7\xb8\x31\xbb\x81\xaa\xd9\x17\x3f\x33\x5f\x7c\xcf\xb7\x74\x98\x4d\x9c\xff\x04\xc4\xd4\x0e\xca\x3e\x0e\xa9\xb0\xf8\x14\x95\x62\xee\x38\x2f\x82\xba\x7e\xee\x58\xae\xa0\xc8\xa6\x99\x14\x19\xc3\x1d\x9c\xc5\xb3\x9c\xe9\xdc\x7e\x66\x25\x61\xf3\x75\x9f\x0e\xa0\x28\x6e\x24\xd1\xa9\xf7\x78\x42\x68\x8f\x67\xeb\xe7\x5d\xef\xfb\x3c\x4f\x19\x3e\x60\xf1\x8e\xb4\xe1\x6e\xb1\x2e\x09\xca\x6b\x2b\xa4\x80\xf3\xa3\xe2\x9c\xf1\xcd\xcb\xbe\x68\x43\xd6\x7f\x74\x3e\xad\x2c\xb2\x78\xfd\x13\x46\x6d\xa3\x70\xc0\xf2\xcd\xd5\x15\xbe\xc7\x9b\x57\x81\x95\xb3\xe9\xe5\x0b\xfb\x3d\x99\xb0\x5a\x9f\xa1\x21\x7a\x64\x17\xe5\x25\xc3\x59\xf3\x6f\x45\x27\xd5\xdf\x39\x9f\x96\xf2\x78\xfc\x2a\x30\x25\x2f\x4f\x55\x75\xaa\x7f\x70\x2d\x9a\xb4\xc2\xab\xff\x02\x00\x00\xff\xff\xe6\xf7\xc2\x46\xb0\x06\x00\x00")

func jsJobJsBytes() ([]byte, error) {
	return bindataRead(
		_jsJobJs,
		"js/job.js",
	)
}

func jsJobJs() (*asset, error) {
	bytes, err := jsJobJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/job.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsMock8sJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x8a\xdb\x30\x14\x85\xf7\x7e\x8a\x43\xe8\x62\x02\xc1\x43\xb7\x81\xac\x0a\xdd\xb5\x94\x0e\x74\x53\x4a\xb9\x91\x8e\x13\x75\x6c\xc9\xe8\x67\x3a\x50\xfc\xee\x45\xae\xa5\xd8\xd9\xb4\x8b\x6e\x94\x9b\xfb\xf3\x59\xe7\xea\x3c\x3e\xe2\x29\x3a\x4f\x0c\x4e\x3d\x63\x74\x1a\x9a\x9d\xb1\x26\x1a\x67\x03\xae\xf4\x84\x04\xec\xac\x0c\xdc\xe1\x88\x37\x9e\xc1\x25\xaf\xd8\xbc\x88\x9f\x67\x3e\x39\x1d\x70\xc2\xaf\x66\x6a\x9a\x0a\x53\xce\x76\xe6\x32\xc8\xf8\x4f\x88\x77\x73\xf7\x87\xdc\xbd\x80\x72\xe5\x39\x9d\xe9\x2d\x23\xff\x64\x81\x9f\x26\x5e\x3f\x3e\x1d\xd1\x25\xab\xf2\xf5\x1e\x6c\xd8\xcf\x05\xc0\x33\x26\x6f\x17\x98\x67\x03\x4c\x05\x53\x72\x0b\x44\x39\xcf\x2f\x6f\x8f\xcb\xdc\xe8\x74\x09\x81\x0b\xe3\x1a\x2e\x03\xf7\xb5\x56\x3f\xf1\xbd\xed\x8c\xd5\x0f\x45\xf9\xe1\x36\xc0\x7e\xd3\x5f\x27\xd8\xb3\x1d\x18\x45\x4b\x94\x36\x53\x71\x3a\x21\xff\xd6\xd6\x69\xbf\x84\xd3\x61\x09\x94\xa7\x44\xae\x6e\xa3\xd9\xad\xe1\x79\xd1\x49\x29\x52\x53\x1f\xf0\x39\x59\x6b\xec\xe5\x00\xb1\x1a\xef\xc5\xf4\xd4\x10\x4f\x04\x37\x10\x2f\xd2\x1b\x9d\xcf\xc4\xd0\x56\x80\x66\xd7\x86\x28\x31\xcd\xcb\xc5\x78\x95\xc0\x23\x76\x15\xba\xc3\x54\x7b\x8b\xd6\xaf\x79\xa8\x2a\xe9\xe5\xcc\x3e\xb4\x3f\xdc\x39\x6b\xf9\x86\x53\x66\x36\xf7\xeb\xca\xb9\x45\x5b\x53\xce\x59\x25\x5f\x23\x6d\xc8\x26\x2b\x0f\x50\x3d\x73\x7b\x91\xbf\xad\x61\x6b\x9e\xff\x76\xbf\x6c\x1d\xbe\x8e\xce\xc7\xd0\xae\x2c\x5e\xc2\x4d\x6d\xe3\xdd\x6d\xa2\xf6\x6d\x9c\x7c\xfb\x73\xc7\x99\x2d\x5a\x1d\xfc\x3b\x00\x00\xff\xff\xc6\x15\x83\xe4\x98\x03\x00\x00")

func jsMock8sJsBytes() ([]byte, error) {
	return bindataRead(
		_jsMock8sJs,
		"js/mock8s.js",
	)
}

func jsMock8sJs() (*asset, error) {
	bytes, err := jsMock8sJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/mock8s.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRunJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x6b\x6f\x1b\xb9\xd5\xfe\xae\x5f\x71\x96\x6f\xb0\x91\x90\xf1\x28\xc1\x02\x2f\x16\x4a\x85\xc2\x2b\x3b\x1b\x25\xf1\xa5\xb1\x37\x6d\xe1\x18\x5e\x6a\xe6\x48\xc3\x88\x43\x0e\x48\x8e\xbc\x6a\xd6\xff\xbd\x38\x24\xe7\x22\x4b\x4e\xd0\xa2\x5f\x92\x11\x79\xee\xe7\x39\x17\x7a\x3c\x86\xeb\x02\x41\x28\x87\x46\x71\x69\x61\xa9\x0d\x98\x5a\x29\xa1\x56\xe0\xb8\x5d\xdb\x14\xae\x0b\x61\xa1\xac\xad\x83\x05\x82\xd4\x3c\xc7\x1c\x16\xb8\xd4\x06\x81\xab\x2d\xe8\x25\xb8\x02\x07\xe3\x31\xe8\xc5\x17\xcc\x9c\x05\x57\x70\x07\xb5\x45\x12\x34\x1c\xa5\x83\xf1\x98\x6e\x8f\xa5\x84\xf7\xf5\x02\x8d\x42\x87\x16\x8e\x2f\xe7\x90\x71\x29\x2d\xd8\x42\xd7\x32\x0f\xc2\x33\x2e\xc5\xbf\x30\x87\x02\x0d\xa6\x70\xe1\x0a\x34\x50\xea\xbc\x96\xd8\xd2\x29\xed\x48\x1c\xf1\xc2\xf3\x75\x2b\xf0\x39\xe4\xc2\x60\xe6\xe4\x36\x1d\xd0\xfd\x3d\x17\x0e\x1c\x5f\xa3\x05\x0e\x5f\xf4\xe2\x48\x8a\x35\x46\x13\x81\xab\xdc\x13\x58\xa8\x95\x13\x12\x84\x03\x5b\x67\x19\x62\x6e\x41\x1b\x58\x72\x21\x6d\x3a\x58\xd6\x2a\x73\x42\x2b\x4f\xfa\x46\x9b\x77\x7a\x31\xfc\xa2\x17\x09\xe0\x08\xbe\x0e\x00\x36\xdc\x40\xb9\x85\x29\x89\x1f\x0c\xc0\xc7\x6e\x48\x87\x02\xa6\xf0\xf2\x35\x08\xf8\x0b\xfc\xf4\x92\x3e\x5e\xbc\x08\x1c\x00\x99\x56\x56\x4b\x4c\xa5\x5e\x0d\x59\x56\x60\xb6\xa6\x48\x5b\xc7\x5d\x6d\x29\x94\x0c\x5e\x40\xb9\x4d\x2b\x9d\x9f\xf3\x12\x47\x9e\x87\x44\xae\x61\x0a\x9d\xb3\xe9\xbd\x70\xc5\xf9\xd5\x10\xd3\xde\x99\xe2\x25\xda\x8a\x67\x3d\xae\x72\x5b\xe9\x9c\x38\xd3\x4c\x1b\xfc\xf4\x8a\xe4\xa6\x2b\x74\xc3\xbe\x8a\x3d\xbb\xde\x5d\x5d\x9c\xa7\xd6\x19\xa1\x56\x62\xb9\x1d\x7a\x21\xa3\xd1\xbe\xf9\x97\x3a\x7f\x64\x2f\xbc\x00\x06\xc2\x82\x50\xde\x25\x8c\xd7\xa4\x35\xb8\x98\x56\x05\xb7\x8d\x4e\xb1\x84\xe1\xfe\x25\x4c\xa7\xc0\xde\x70\x21\x31\x67\x4d\xd4\x00\x5c\x61\xf4\x3d\x3c\xa5\x72\xe9\xc9\xc1\x69\x82\x1c\xfd\x97\xe9\xb2\x92\x48\xc9\x63\xaf\xbd\x84\x87\xef\x69\xbc\x0a\xf9\xdf\x51\x6a\xd0\xd5\x46\x81\x33\x35\xf6\x84\x8c\xc7\x70\x25\x11\x2b\x9f\x6f\x0e\x39\x2e\x85\xc2\x1c\x78\xa9\x6b\xe5\x7c\x39\x88\x12\x53\x4f\x6a\x89\x6e\xf8\x13\xc5\x8e\x78\xa3\x13\x74\x9f\x83\xae\x9d\x07\x16\xe5\x9f\x24\x55\x87\x5d\x0b\x3e\xb1\xc1\x83\x47\x35\xb9\x67\x6a\x15\x31\x1d\x2c\x50\x80\x1b\x54\x0e\x86\x38\x4a\x61\xee\x20\xd7\x68\xa9\x4a\x42\x09\x10\x89\x2b\xd0\x93\x77\x81\x41\xaa\xca\x50\xd9\x82\x84\x49\x7d\x7f\x24\x71\x83\x12\x2a\x23\x4a\xe1\xc4\x06\x7b\xf8\xa7\x3a\xee\x03\x7f\x3c\x86\x67\xef\x2e\x7e\x39\x7a\x76\x3d\x3f\x3b\x3d\x7a\xf6\xeb\xfc\xfa\xea\xed\x71\x2c\x88\xf5\xcf\xd6\xdb\xee\xab\xc2\x63\x92\xdc\x38\x22\xcf\x4e\xb8\xc3\x54\xe9\xfb\xe1\xa8\x3d\xc2\x34\xd3\x65\x29\x5c\x6a\xeb\x45\x40\xdc\xf0\x65\x02\x3f\x8f\x5e\x47\x69\x59\x19\x85\x45\xb1\xf1\x98\x3a\x14\x9a\x83\x57\x59\x09\x53\x50\x78\x3f\x3b\x1b\x06\xe6\xd1\x0e\x4f\xb8\xfc\xe8\xbf\x2f\x75\x3e\xec\x24\x25\xde\x60\x51\xf2\x55\xc0\x67\xb8\x49\x4b\x74\x3c\xe7\x8e\xa7\x92\x2f\x50\xda\xf4\x8b\x5e\xa8\x5d\xff\x9e\xa6\x5d\xa0\xd4\x6a\x65\x9d\x86\x29\x60\x6a\xb0\xd2\x9e\x81\xbe\x24\xcf\x70\xc8\xc6\x2c\xa1\x40\x8c\x9e\x16\x11\xc2\xe3\xf9\xc3\xe7\x20\xc4\xff\x38\xcf\x01\xd5\x86\x3c\xb3\x69\xf4\x10\xd5\xe6\x13\x37\x16\xa6\x70\x73\x4b\x64\x77\x29\xf2\xac\xa0\xcc\xa5\xa8\x36\x09\x34\xf9\x1c\x6e\xb8\x4c\x60\x8d\xdb\x06\xe9\x91\x31\xad\x6a\x5b\x0c\xbf\x92\x89\x13\xba\x4e\x60\xc3\x65\x8d\x13\xfa\xef\xc1\xa7\x84\xfe\x0d\xfa\x4f\x34\xdc\x23\x58\x27\xa4\x84\x7b\xae\x1c\x61\x8b\xe7\x39\x38\x82\x94\xd3\x1e\x72\x3e\x96\x6d\x5b\xfe\x2b\xfc\xbd\x10\x12\xa9\xd7\x7a\xc8\x59\xcc\x6a\x23\xdc\x36\xc8\x73\x85\x50\xab\xc4\xc3\x96\xe7\x39\xd5\x84\x70\x70\xef\x1b\xbe\x41\x5b\x4b\x47\x3d\xa5\xb6\x68\x02\xb4\x17\x48\x24\x7c\x21\x91\x94\x91\xdd\x34\x8b\xbc\x90\x20\xaf\xae\xac\x33\xc8\x4b\x1a\x6d\xc1\x1a\xaa\xfa\x41\x68\x01\x31\x17\xd6\x16\xef\x9f\x0a\x42\xec\x00\x21\x16\xec\x78\x36\x3f\xb9\xfb\x78\x7a\x79\x71\xf7\xfe\xf4\x9f\x2c\x89\x97\x31\x3a\x3b\xd2\x42\x9f\x08\x05\xdf\x65\xca\xe9\x2a\x96\x57\x9b\x33\x9a\xbb\x16\xc3\x5c\xd5\x1b\x34\x46\xe4\x61\x9e\x72\xe7\xb0\xac\x7c\x40\x2d\x3a\x6f\xba\x57\x64\x63\xa0\x34\x58\x5d\xa2\x77\x15\x50\x5a\xdf\x69\x76\x6d\x6f\xac\x9e\x7d\xb8\x38\x3f\xbd\xfb\xed\xe3\x07\x96\x3c\xb2\x35\x93\x5a\xe1\x6f\x1f\x3f\x04\x4b\x0f\x73\x5f\x5d\xbd\x3d\xc8\x6b\x6d\xf1\x1d\xce\x5f\xe7\xd7\x07\x39\x57\xc2\x7d\x87\xf3\xed\xe9\xf1\xc9\xdd\xec\xe2\xec\x6c\x7e\x7d\x37\x3f\xe9\x0b\x88\x75\xf0\x0d\xde\xd9\xbc\xa3\x67\xd4\xb2\x59\xa0\x8e\x85\x65\x2b\xcc\xd2\x4c\x2b\xc7\x85\x42\x63\x6f\x5e\xde\x52\x51\x50\x61\x05\x69\x83\x66\xa0\x53\x17\xbf\xe4\xae\x88\x15\xde\xfd\xfe\xf3\x4f\x60\x63\x6b\x32\xd6\x4b\x6c\xa6\xd5\x52\xac\xa0\xe4\x15\x6c\xb4\xac\xfb\xcd\xc0\x2b\x0c\x87\xbe\x24\x3d\x32\x1a\x6b\x43\x6b\x4a\x22\xff\x19\xaf\x26\xf0\xb5\x7f\x03\x0f\x0f\xc9\x0e\x03\xdb\x64\xf6\xc8\x8a\x1c\x33\x6e\x58\x02\x84\x90\xed\x89\x30\x13\xf8\xfa\xd0\x4e\xa5\xa4\xa3\x16\xb9\xb1\x9c\x25\x54\x64\x06\xdd\x04\xbe\xc6\xaf\x73\x7f\x6d\x31\x8b\x4a\x06\x00\xb7\xaf\xbf\x19\xa4\xe0\xc1\x19\x45\xe1\x69\x37\xda\x20\x4d\x80\x8d\x0b\xad\xd7\xec\xdb\xe6\xf7\xe8\xdb\xcf\x6f\xb9\xb1\x27\x7f\x6c\x6d\xc1\x12\x30\xc8\xf3\x0b\x25\xb7\x13\x3f\xa3\xa3\x37\xbd\xb2\x2b\x10\xa2\xd2\xa6\x47\xc6\x9f\x30\x6d\xbe\xae\x2a\xcc\x86\x98\xc4\xd4\x26\xb0\xb3\x52\x6d\x32\x7b\x15\xc8\x46\x51\xec\xf5\xc5\xc9\xc5\x84\xd2\xb6\x41\xe3\xda\x6e\xc7\x15\x08\x25\x1c\xb4\xa1\x03\x5a\xd0\xfc\xba\x0b\xaf\xd2\xff\x0f\xac\xfd\x18\x13\xf5\xac\x8d\x33\xc5\x35\x9a\x73\x7b\x60\x14\x70\xa5\xb4\xe3\xd4\xb9\xed\x0d\xa3\x36\xb6\x40\xc7\xfb\x66\x0a\x3d\x26\x81\x47\x5d\xe6\xd8\x2d\x4c\x81\xdd\xd0\x88\x7d\xb4\xcd\x45\x3d\x7e\x04\xdf\xb2\x06\xf5\x34\x30\x4b\xda\x14\xd9\xff\xfd\x40\x4c\x84\x7c\x5b\xa0\x94\x44\xf6\x59\x7d\x56\x0d\xe8\xc5\x12\xfc\x79\xdc\xf2\x6d\x5d\x55\xda\x84\x2e\xf5\xbb\x45\xf7\x3b\x2d\x18\x25\xad\xd7\xdc\x20\x58\x94\x98\x39\xcc\x13\x90\xe8\x9e\x5b\x3f\x20\xa8\x79\x81\xe5\x0a\x69\x71\xe2\xb5\x74\x36\x36\xe5\x4e\x25\xad\x63\xe3\x85\x50\x63\x5b\x30\xaa\xba\x03\x37\x0b\x6e\x8b\x76\x51\x8b\xc6\xbf\x98\x02\xa3\x8e\x79\x84\xc1\xe0\xae\x05\xbf\xd3\x42\x79\x13\xfd\x8b\x86\x12\x56\xf2\x35\x02\x27\xce\xc6\xe0\x49\xcf\x0c\x4f\xb6\x2f\xbd\xbd\x4a\xbf\x68\xa1\x86\xec\xb3\x62\x6d\xa3\xcf\xca\x94\x72\x75\xc3\x4a\x2e\x54\x6a\x0b\x9f\x80\xc0\x3a\x78\xb4\x36\xd7\x96\x5a\x77\xa4\x9b\x7c\x56\x14\xef\x7d\xf6\x51\x93\x9a\xff\x6c\xf3\x7f\xac\x6c\x66\x90\xfb\x15\x33\x74\x1a\x6a\x54\x51\x5f\x0b\x2f\x15\x97\xa4\x75\x8a\x7f\x38\x54\x96\x70\x96\xb6\xe4\x69\x46\x12\x70\x98\x95\xa3\x27\x65\x37\xab\xeb\xd3\x5d\xa4\xd3\xd1\x7b\x88\x44\xc9\x81\x6b\x4f\x7a\x7c\x85\xa6\x69\xca\xc2\x3e\x16\x96\xf1\x6e\x61\x7b\x4d\x3b\x71\xbb\xa3\x3e\x2a\x67\xff\x94\x4c\xc2\x0a\xd2\xbd\xd6\xfc\xcf\x6b\xbe\x82\x69\xf8\x6c\x56\x43\xac\x34\x4d\xa6\xe9\xe3\xf9\x38\x88\xa0\xf8\xa1\x61\x6c\x40\xd1\x13\xc4\x78\x26\xf2\x71\xaf\xbf\x4d\x24\x77\x68\x5d\x0b\xc0\xbd\x8d\x03\x7e\x98\x02\x6b\xd1\xbb\xa7\x3c\x0c\xd8\x86\xdb\x37\xad\x0a\x33\x98\x36\x78\x0c\xad\x91\xb4\x1e\xed\x74\xd5\x66\x8b\x99\xc4\x26\xdd\xeb\xbe\x9f\x66\x57\x7e\x7d\xe9\x26\x64\xa3\xf5\x21\x39\x44\xfb\xe1\x62\x76\xfc\xe1\xee\xf2\xf8\xfa\x6d\xc7\xe1\x23\xfa\x04\xfd\xc7\xd3\x4f\xf3\xab\xf9\xc5\xf9\xa1\x89\x1d\x18\x6e\x93\x2e\x6e\x93\x36\x7c\x49\x7c\x54\x86\x0a\x84\x1b\xd6\x8f\x23\xeb\xf3\x5c\xd6\x52\x5e\x6a\x29\xb2\x2d\xad\x63\xf2\x9e\x6f\x6d\xf4\xb8\x3f\xa0\x0e\xb8\xfe\xe4\xe0\xf1\xfe\x34\xd6\x3d\x99\xab\x98\x25\x0f\x69\x54\x9b\xff\xed\x76\x18\x21\x4d\xc2\x77\xb0\xbc\xf3\x42\xa9\x74\xae\xfc\x90\xa5\x8c\xcf\x3b\x3c\x47\xe6\x60\x0b\x5b\x0b\x95\xb3\x89\x7f\x20\x47\x0b\x18\xaf\xc4\x27\x34\x54\xce\x74\xb1\x79\xd5\x9c\x37\x85\xcf\x26\xed\x43\x97\x91\x06\x36\x81\x46\x57\x73\x1c\x1e\x21\x3d\x42\x00\x56\xa0\x11\x8e\xaf\x88\x9c\xfd\xad\xd6\xeb\x35\x6f\x5d\x26\xe1\x5c\xf1\x15\xe6\xbf\x6c\x59\xc4\x28\x8b\x77\x2d\x70\x58\x6f\x9c\x91\xe4\x30\xf7\xe3\x35\xa3\x58\xf4\x0d\x33\x68\x1d\x37\x2e\xa4\x9e\x64\x9e\xe3\x06\x4d\xab\x91\xf5\x46\x5e\x97\x7c\xe8\xd9\xdb\x39\xe7\xcd\xa1\xb7\x74\xd2\xbf\xf4\xe8\x62\x93\x2e\xba\x3b\xb7\x11\x9a\x3b\xb2\xfd\x45\x33\x9f\x92\x47\xc7\x7e\x39\x69\x9a\x78\xef\xee\xb6\x4f\x38\x1e\xc3\x9b\xf9\x3f\xce\x4e\x27\x30\x2b\xb8\x5a\xf9\x37\x0c\x9b\x2f\xcf\xb5\xbb\x34\x68\x51\x39\xb6\x67\x61\x87\x7f\xd6\x15\x40\x4b\xf5\x10\xbf\x6e\xdb\x3f\x53\x3c\xf8\xee\xb8\x03\xa9\xd9\xd9\xd0\xb7\xe1\x6f\xa0\x67\xd6\xac\xa3\xff\x05\x86\xba\x40\xf7\x21\x74\x18\x44\xdf\x07\x52\xc4\x43\x03\x8b\x7d\x5d\x4d\x8c\x27\xc0\x30\x2b\x34\xd0\x7e\xa0\xe1\xc7\x1f\xc1\xff\x5a\x69\x9d\x2f\xb6\xc8\x3a\x21\x21\x20\xf8\x07\xad\x2b\x36\x35\xb5\x82\x29\xcd\x92\xf6\xa4\xfb\x1b\x1f\x4c\x7b\x7f\xf0\x1b\xfc\x3b\x00\x00\xff\xff\xe1\x82\x9b\xca\x22\x15\x00\x00")

func jsRunJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRunJs,
		"js/run.js",
	)
}

func jsRunJs() (*asset, error) {
	bytes, err := jsRunJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/run.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRun_mockJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xcd\x4a\x03\x51\x0c\x85\xf7\xf7\x29\x0e\x5d\x75\x68\x4d\x5d\x0a\xd2\x9d\x08\xba\xf6\x05\xee\x5c\xe2\xf4\xb6\x33\xc9\x90\xe4\x52\x41\x7c\x77\x19\xb4\x3f\xdb\x93\x2f\xe7\x7c\x69\xb7\xc3\xc7\xa1\x3a\xe6\x5c\x4e\x79\x60\x4c\x5a\x4e\x8e\x38\x30\xac\x09\x1d\x1d\x9f\x4d\x4a\x54\x15\x27\xbc\x05\x4a\x16\x18\xcf\x63\x2e\x37\x40\x0d\xc1\x1e\x55\x06\x4a\xe9\x82\xe3\x9c\x6b\xbc\xaa\xbd\x6b\xbf\x3e\x6a\xdf\xe1\x3b\x01\x45\xc5\x75\x64\x1a\x75\x58\x42\x92\x3c\x71\x97\x12\x60\x1c\xcd\x04\x61\x8d\xd3\xcf\x5d\x89\x35\x59\xc0\x2d\xf8\xaf\xe0\x9f\xbb\xfc\x62\x83\xd5\xc3\x0a\x1b\xbc\xe4\x60\x12\x3d\xaf\xbb\x6b\xc4\x54\x74\x9a\x6a\x90\xb7\xde\xc3\xaa\x0c\xeb\xc7\x2d\x9e\xba\xe7\x65\x80\xbf\x66\xb5\x70\xba\x49\x62\x7f\x67\x7c\xbd\x5b\x13\xec\x17\x8b\xf4\x1b\x00\x00\xff\xff\xf2\x0f\x52\xe5\x29\x01\x00\x00")

func jsRun_mockJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRun_mockJs,
		"js/run_mock.js",
	)
}

func jsRun_mockJs() (*asset, error) {
	bytes, err := jsRun_mockJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/run_mock.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRunnerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x8a\x14\x31\x10\x86\xef\x79\x8a\xdf\x3e\x75\xe3\xd2\xfb\x04\x73\x10\x5d\x51\xf0\xe8\x4d\x44\x42\x52\xdd\x1d\x08\x55\x6d\x55\x7a\x7b\x17\xd9\x77\x97\x74\xdc\x71\x66\xf0\xe0\x31\xf9\xc9\x97\xbf\x3e\xea\xfe\x1e\x5f\x97\x64\x48\x86\xb2\x10\x74\x63\x26\xc5\xae\x7e\x5d\x13\xcf\xb0\xa0\x69\x2d\xa3\x0b\xc2\x26\x99\xc6\x2c\x73\xdf\x7d\x11\x1f\x6b\xf8\xee\xfd\xe7\x0f\x08\xa2\xd4\x0d\xce\x4d\x1b\x87\x92\x84\x31\x25\xa5\x87\x47\xe2\xd2\xd3\x80\x5f\x0e\x78\xf4\x0a\xaa\x17\x9f\x3c\xc7\x4c\x8a\x13\x98\x76\x3c\x5c\x5c\xf5\x83\x73\xc0\x6b\x97\x59\xc8\x90\xf8\x28\x64\x41\x56\xaa\x87\x7d\x49\x61\xc1\xd4\x0f\xb5\x2a\x3d\x51\xd8\x0a\x45\x98\xa0\x2c\xbe\x60\x27\x04\xcf\xf0\x21\x90\x59\x43\xd5\xd7\xa2\x69\x4e\xec\x73\xfb\x1f\x7b\x2a\x8b\x6c\x05\x4a\x3f\xb7\xa4\x75\x84\x5e\xf4\x0e\x4a\x3e\xe7\xe7\x3b\xf8\x9c\x65\x4f\x3c\x0f\xc7\xdb\xcd\x48\x51\xa4\xc1\x56\x6f\x4d\x50\x03\x79\x95\x8d\xe3\xe8\x00\x7a\x5a\x45\x8b\x8d\x3f\x5a\x70\x02\xd5\x49\xd2\x84\xfe\x8d\xd2\x9c\xac\x90\x1e\x83\x5a\x73\x01\x5c\x99\x64\xf9\x03\x5c\x9a\x07\x43\xa4\x29\x31\xc5\xee\x10\x02\x28\x95\x4d\xd9\x01\x2f\xf5\x7c\x4d\xec\x2f\xa5\x0e\xee\x06\x7d\x84\x86\x2c\x3e\x52\x1c\xf1\xb1\xcd\xdb\xe1\x2d\x68\x2c\xcf\x2b\x0d\xe7\x9e\x97\x98\x6f\x2d\xfc\xfe\x5f\x6d\xcf\x7d\x28\x62\x12\xbd\x85\xdf\xb6\xaf\x7b\x30\x55\x43\xff\xf8\xaf\xe6\x53\x4f\x83\x7b\x71\xee\x55\xe9\x79\x91\x70\xfa\xbb\x54\xee\x77\x00\x00\x00\xff\xff\x0b\x8b\xea\xff\xb1\x02\x00\x00")

func jsRunnerJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRunnerJs,
		"js/runner.js",
	)
}

func jsRunnerJs() (*asset, error) {
	bytes, err := jsRunnerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/runner.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsWaitgroupJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xb1\x8e\x9c\x30\x10\x86\x7b\x3f\xc5\x5f\x82\x14\x99\x27\xd8\x32\x8a\x92\x22\x4d\x22\xa5\x88\x52\x18\x18\x62\x73\xc6\x83\xec\xe1\xd8\xd5\x89\x77\x3f\xd9\xbb\x2c\x48\x57\x5e\x61\xad\xd6\x78\xbe\xff\x9b\xb1\x9b\x06\x7f\x8c\x93\x6f\x91\x97\x19\xab\x71\x92\x30\x70\xc4\xb4\x78\x71\xb3\x27\x8c\xdc\x26\x08\x63\x70\xc1\x25\xab\xf1\x5d\xb0\x3a\xef\x21\x36\xf2\x0a\x13\x40\x31\x72\x54\x4d\x03\x93\x90\x98\x43\xfe\x35\xb9\x0c\x91\x66\x8e\x92\x9e\x87\xb4\x1a\x96\xd0\x89\xe3\x70\x44\x56\x35\xde\x14\x20\xd6\x25\x5d\xa2\x2e\xf8\xfb\x4f\x29\x20\x03\xfb\x3e\xaf\x8c\x0b\xb4\x16\xa4\x30\xc4\x52\xf1\xfc\x9f\xcb\xf7\xd2\x7c\xf4\x82\x1d\x5f\x8d\xdc\xde\xb9\x27\xb2\x9e\x97\x64\xcb\x17\x05\x6c\x8f\x88\xb8\x84\xbc\x12\xe8\x95\xe2\xad\x44\xb8\x50\x22\x0a\xfe\x0b\x4c\xe8\xf3\xdf\x70\x1a\x8d\x58\x9a\xb2\x48\xc7\xd3\xec\x49\x48\xef\x12\x19\x76\x92\xf8\x68\x30\x70\xfc\x6a\x3a\x5b\x3d\xc7\x50\x8d\xfb\x21\x60\xd4\xad\xe9\x5e\x72\x6c\xe8\xab\xba\x6c\x6e\xf5\x01\xc8\xf9\xd5\x59\x3d\x6f\x3c\xac\x96\x20\xce\xdf\x6f\xca\x44\x3a\xc4\xf0\x93\x85\x20\xd6\x48\x61\xa0\x67\x4a\x08\x2c\xa5\xed\xdc\x64\xb1\xc2\x6f\x4b\xb7\x3b\x73\x5a\x92\xa0\x25\x24\x31\x51\xa8\x07\x5d\x85\x62\x30\xde\xdf\x34\xaa\x5f\x44\xc7\xbd\x95\x66\x39\xe2\x07\xb7\x27\xef\x5a\x9d\x6c\x3f\x35\x8b\x67\xbb\x8f\x29\x6c\x6a\x53\x8a\xae\xe5\x41\xe9\xe3\xc1\x5e\x0e\x23\xf5\x1e\x00\x00\xff\xff\x60\x6f\x0e\xe7\xca\x02\x00\x00")

func jsWaitgroupJsBytes() ([]byte, error) {
	return bindataRead(
		_jsWaitgroupJs,
		"js/waitgroup.js",
	)
}

func jsWaitgroupJs() (*asset, error) {
	bytes, err := jsWaitgroupJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/waitgroup.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"js/event.js": jsEventJs,
	"js/job.js": jsJobJs,
	"js/mock8s.js": jsMock8sJs,
	"js/run.js": jsRunJs,
	"js/run_mock.js": jsRun_mockJs,
	"js/runner.js": jsRunnerJs,
	"js/waitgroup.js": jsWaitgroupJs,
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
	"js": &bintree{nil, map[string]*bintree{
		"event.js": &bintree{jsEventJs, map[string]*bintree{}},
		"job.js": &bintree{jsJobJs, map[string]*bintree{}},
		"mock8s.js": &bintree{jsMock8sJs, map[string]*bintree{}},
		"run.js": &bintree{jsRunJs, map[string]*bintree{}},
		"run_mock.js": &bintree{jsRun_mockJs, map[string]*bintree{}},
		"runner.js": &bintree{jsRunnerJs, map[string]*bintree{}},
		"waitgroup.js": &bintree{jsWaitgroupJs, map[string]*bintree{}},
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
