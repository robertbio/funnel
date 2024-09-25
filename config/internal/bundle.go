// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/default-config.yaml
// config/gridengine-template.txt
// config/htcondor-template.txt
// config/kubernetes-executor-template.yaml
// config/kubernetes-template.yaml
// config/pbs-template.txt
// config/slurm-template.txt
package config

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configDefaultConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3a\x5d\x73\x1b\x39\x72\xef\xf3\x2b\x3a\xe2\x5e\xc5\x5b\xc5\x2f\xed\xd6\x5e\x72\xac\xd2\x83\xbe\xd6\x56\x2c\x7b\x19\x91\x8e\x93\x27\x17\x66\xd0\xc3\xc1\x72\x06\x98\x05\x30\xa4\x69\xc7\xff\x3d\xd5\x0d\xcc\x07\x25\x79\xed\xbb\x93\xab\x36\x75\xa7\x17\x71\x80\x46\xa3\xd1\xdf\xdd\xc0\x08\xd6\x05\x82\x16\x15\x82\xc9\xc1\x17\x08\x22\xf3\x6a\x87\xe0\xd0\xee\xd0\x82\x14\x5e\xa4\xc2\x21\xa4\x22\xdb\xa2\x96\xc9\x08\xce\x77\x42\x95\x22\x2d\xbb\x31\xb7\x80\xd4\x94\x5e\xa6\x63\x48\x85\xdc\xa0\x1d\xf3\x32\xe7\x8d\xc5\x31\xc8\x83\x16\x95\xa1\x49\x2c\x85\xf3\x2a\x1b\x43\x65\xf4\xc6\xc8\x34\xb9\x8a\xc8\xdb\xf5\x49\xf2\x59\x72\x32\x53\xd5\x8d\xff\x12\x19\xa5\xc9\x44\x39\x86\xc2\x67\x46\x4b\x63\xc7\xe0\xca\xc6\x56\x63\xa8\x53\x37\x86\x8d\x55\x12\xf5\x46\x69\x1c\x43\x25\x74\x43\x90\x62\xef\x26\xa9\xf0\x59\x91\x5c\x86\x0d\x22\x8e\xdf\xa1\x04\x77\xa8\x3d\xec\xad\xf2\x68\xdb\xad\x9f\xb9\xef\xa7\x9f\x25\x69\x33\xfe\xdb\xd8\x33\x86\xad\xc8\xb7\x22\xb9\xa6\x0d\xdf\xf2\x7e\x6e\x91\x00\x4c\x5a\x6e\xd1\xcf\xd2\x6c\x92\xe4\xd6\x6c\x36\x68\x69\x6e\x04\xf4\x5b\xe9\x0d\x94\xb8\xc3\xd2\x2d\x40\x62\xda\x6c\xc6\xa0\x74\x6e\xc6\x80\xd6\x1a\x9b\x00\xdc\xd2\xe4\x82\x07\x79\x11\x63\x27\x5c\x0e\xbc\x01\x5f\x28\x07\xb5\xf0\xc5\x14\x6e\x72\xc0\xaa\xf6\x87\x71\x98\x14\x16\xf9\xe4\x1e\x35\x01\x3a\x2f\xd1\xda\x69\x02\xf0\x4b\xe3\xeb\xc6\xff\xac\x4a\x5c\xc0\xc9\x49\x92\xac\x58\x7b\x02\x45\x2f\x8c\xf3\x43\x3e\xfe\xdc\x68\x8d\x65\x54\x30\x5a\x4c\x00\xaf\x45\xd5\xf2\xbe\x30\xce\x27\xbc\x72\x69\xac\x87\xc6\xa1\x84\xdc\x58\x78\xb1\x5e\x2f\x49\x0f\xaa\x46\xab\x4c\x78\x65\x34\x08\x2d\x19\xe5\x1e\x53\x90\xc2\x15\xa9\x11\x56\x32\xca\xf5\x7a\x49\xab\x17\xf0\xef\xf3\xf9\xfc\x31\x6c\x77\xcb\xcb\x63\x64\xb4\xec\x6e\x79\x19\x56\xfd\x65\xfe\x97\xb8\xea\x0e\x7f\x6b\x94\x25\x89\x3a\x95\x81\x68\x7c\x81\xda\xb7\xfb\x13\x22\xda\x3f\x1a\xcb\xf9\xf2\xc6\x41\xe3\x88\xfd\x02\x6a\xe1\xdc\xde\x04\x72\x46\xc4\x48\xda\x9a\x34\x6f\x8b\xe0\x1a\x8b\xc4\xc0\xda\x9a\x1a\x6d\x79\x00\x8b\xce\x5b\x95\x79\x10\x59\x86\x2e\x4a\x81\xb4\x5e\xe7\x6a\x03\xb9\x2a\x91\xb1\x3c\xc3\xe9\x66\x0a\x59\x51\x19\x09\x7f\x9e\xcf\x21\x67\x56\x4e\x03\xd8\xf4\x50\x95\xdf\x33\xd8\x05\xd1\x7a\xde\xf8\x22\x08\x80\xf4\xe4\x8d\x43\xbb\x20\x12\xec\x69\x1c\x03\x58\x46\x0a\x17\x20\xd2\xec\xf4\x87\x1f\x1f\x01\xfe\xe1\x11\xe0\xdc\x98\x54\xd8\xc0\x9d\x1b\x9d\x95\x8d\x44\x10\x70\x72\x29\xb2\x02\x27\x97\x46\x7b\x6b\xca\x05\x68\x33\x61\x0d\x3f\x09\x72\x2b\x50\x48\xb4\xa0\x34\x3c\x47\x3f\xbb\x55\xce\xd3\x99\x6b\xa3\x1d\x3a\xc6\xc4\xdc\x08\xb6\x95\x89\xac\x20\x1e\xa6\x07\x50\xda\xa3\xad\x50\x2a\x61\x0f\xcc\x65\x95\xa1\x23\x8e\x5e\x29\x47\x86\x46\xb8\x79\xe3\x05\x78\xdb\x60\x92\xdc\x2d\x2f\x2f\x4b\x85\xda\x87\xa3\x93\x94\xa3\x70\x84\x94\x16\x1d\x6d\x16\x94\xf3\x3c\x7c\x0f\x94\x6e\xd1\x4b\xfd\xd2\xa2\x24\x39\x8b\xd2\xb1\x90\x2f\xfe\x1f\x49\x3f\x48\x2f\x4c\x06\xc5\xbf\x2f\x67\x1e\xcd\x8c\xd6\x98\xf1\x41\xbc\xaa\xd0\x34\x9e\x68\x5d\x87\x9f\x0b\xf8\xf3\xdc\x05\x38\x72\x83\x95\x78\xaf\xaa\xa6\x02\xdd\x54\x29\x5a\x36\x64\x55\xa1\x03\x5f\x08\x0f\x02\x2c\xfe\xd6\xa0\xf3\xb0\x57\x65\x09\x29\x82\x45\x6f\x55\xb4\xb3\x5c\xa8\xb2\xb1\x41\x68\x23\x46\x0f\x29\xfa\x3d\xa2\x8e\x60\xc4\xe0\xb2\x34\x7b\x07\x42\x03\xbe\xaf\x8d\x0e\x8c\x67\x1f\x6a\xf2\x1c\x9c\x17\xd6\x33\x57\x3d\xfc\x04\x0e\xc9\xb7\x07\x9d\x69\x6a\x62\xd5\x29\x54\x4a\x37\x9e\x78\xf4\x4a\xbc\xbf\x0b\x48\x17\x70\x3a\x6f\x9d\xb8\xcb\x0a\x94\x4d\x49\xea\xe7\x7a\x07\x40\x1c\x7e\xc5\x61\xe0\x7e\x70\x99\x26\xab\x76\x45\xeb\xc1\xf6\x60\xf2\xe8\xf4\x6c\xa3\x41\x0c\x71\x7a\xb4\x9d\x03\x69\x17\xde\x09\x0a\x25\xa7\xae\x5b\x5e\x09\x7d\x00\x2f\xdc\x96\xa5\xdb\xae\x26\x7b\x30\x1a\x1f\xc7\x71\x59\x34\x7a\xcb\xe7\x68\x91\x94\x46\x6f\x68\xf9\x5e\x28\xdf\x71\xb1\xa9\xa5\xf0\xe8\x20\xc5\xdc\x58\x12\x95\xdd\x06\x15\xd4\x46\x22\x48\x14\xac\x83\xaf\x8d\xc4\xa5\xd2\x9b\x4e\xbe\xa7\xd5\xe3\x68\x89\x35\x71\x2d\x7b\x78\x61\xfd\xf8\x3e\x6e\x62\xdd\x03\xec\x37\x5a\xf9\x0e\xfb\x4f\x55\x92\xd0\xe0\xa2\x55\xff\x18\x45\x22\xe6\x9b\xab\x4e\x57\x44\xe3\x4d\x25\xc8\xa6\xca\xf2\x00\x1b\xd4\xc4\x09\x64\xac\x37\x57\x21\x98\x44\x14\xdd\xae\x85\xa0\xd3\xa2\x06\x25\x4b\x64\x82\xe9\x04\x48\x22\x15\x9a\xc1\xa2\x42\x8f\x41\x45\xa5\x74\x45\xe3\x41\x9a\xbd\x0e\x6a\x38\x39\x85\x0a\x85\x26\x05\x46\x8b\xa4\x16\xda\x74\x56\x00\xf3\x76\x32\x0c\x80\xaa\xd8\x03\x79\x2c\x0f\x20\x72\x8f\x41\x77\x72\x65\x9d\x67\x91\x1e\xd9\xcd\xe4\x34\xda\xcd\x39\x73\x25\x6c\x7f\x7c\x48\x6f\x0f\xc4\x5b\x89\x1e\x33\x0f\x7b\xb2\x21\x8b\xce\x34\x36\xc3\x10\x62\x45\x97\x47\x78\x03\xca\x07\x9a\xaf\x30\x57\x9a\xb8\x7f\xd7\xc1\xaa\x70\x5a\xde\x27\x78\x81\x26\xa8\x11\x98\x1d\x5a\x4a\x77\x5c\x08\xe6\x29\x16\x62\xa7\x0c\x47\xdb\x6e\x39\xc9\x86\x7d\xdd\xf2\x8d\xeb\xb7\x9c\xb6\xa3\x75\xe3\x16\xc0\xee\x90\x1d\xe9\xf9\xab\x1e\x66\xcc\x7e\xfc\xa2\x05\xbd\x13\xd5\xf3\x74\x01\xf3\x69\x07\x7d\xa5\xdc\x16\x5c\x2d\x32\xfc\xec\x22\x02\x19\xac\x1a\xc1\xcf\x2c\xc7\xfd\x84\x13\x17\xf0\x0d\x9d\x75\xfa\xd0\xfc\xdc\x41\x67\xb0\x57\xbe\x78\x3c\x97\x78\xc3\xd6\x10\xcc\xef\x27\x97\x24\x6f\x8d\xdd\xb6\x66\x4c\xe9\x89\x83\xcc\x22\x29\x18\xc8\xc6\x12\x37\x6b\x6b\xc8\xe7\xd2\xcf\x56\x25\xdb\x0c\x87\xd9\xab\x1c\x48\x65\x31\xf3\xc6\x1e\x68\x03\x42\x78\xa5\xec\x02\xa6\xb3\xe0\x62\x27\x7b\x63\xb7\x13\xa9\xec\x5f\x75\x8c\xda\x94\x25\xab\x6e\x26\x74\x46\x27\x50\x1b\x2d\x4a\xf2\x17\x4b\x53\x96\x4a\x6f\xfa\x23\xfc\x35\xcc\x41\x2d\x29\x2b\x33\x8d\x9f\xa1\xb5\xac\x9d\x94\xb9\x75\x4e\x22\x46\x96\x07\x6c\x1b\xc1\x0a\xbd\x0f\xb6\xad\x18\x6c\x1e\xd8\x61\xd1\x35\xa5\x8f\x9a\xe6\x48\xeb\xb1\x94\xa4\x50\x04\x1b\xb0\x4a\x72\x87\x4a\x6f\xca\x60\x77\x8c\xad\x37\x13\x7c\x8f\x59\xe3\x8d\x05\x7c\xaf\x3c\x47\x81\x5b\xb3\xb9\x2f\x25\x5a\xf2\x4a\xbc\x87\xf4\x10\x89\xe4\xd4\x81\xf9\x33\x38\x4d\x54\xf7\xf6\x50\x11\xd7\x5a\xa8\x72\xa5\x3e\x90\xbb\x9d\xcf\xe7\x73\x18\xc1\xe9\x1c\x5e\x5e\x04\xa4\xaf\x8d\xad\x82\xcd\x51\x96\xc8\xba\x00\x12\x4b\xa4\x6d\x94\x77\x3c\x44\x27\xe9\x44\x1c\x29\x0f\x54\x77\x4c\x5e\x13\x53\x4c\xcd\xa6\x25\x43\xe6\x11\x83\xdf\xd0\xb2\x6e\x51\xec\xb0\xd3\x8f\x5c\x94\x0e\x03\x15\xb7\xaa\x52\x3e\xf8\xaf\x2e\x80\x66\x46\x67\x8d\xb5\x94\xef\x90\x5f\x2a\x8d\x90\x6e\xd6\xd4\xfc\x3f\x44\xb2\xa5\xb0\xa2\x2c\xb1\x5c\x5b\xa1\x5d\x4e\xc9\x7f\x88\x69\x93\xa7\xfd\x4b\x46\xd0\xd6\x61\x14\x82\xe5\xcc\x58\xe0\x92\x03\x62\xcd\x31\x7b\x21\xb4\x2c\xd1\xba\xa7\xdf\x3a\xb9\x30\xa5\xbf\xba\x58\xc4\x34\x85\x8c\x3a\x28\x68\x57\x76\xc6\xe4\x87\xe6\x1e\x31\xb9\xf8\x3d\xa5\xd2\xf1\x8a\x0b\xa9\x16\xd9\x85\x70\xc8\x45\x87\x37\x14\xf2\x59\x95\xda\x52\x0b\x3c\x0b\x90\xbc\x36\xfd\x68\x41\x8f\x72\xa6\xf3\xb7\x2b\xb0\xb8\x51\x46\xb3\xcb\xa4\x1f\x1c\x8c\xda\xb9\xf3\x90\xa9\x6d\xf1\x00\x37\x57\x09\xc0\x4b\x3c\x1c\xcd\xaf\x30\xb3\xe8\x5b\xb0\x97\x78\xe0\xbc\x93\xc6\x42\x4c\xbb\x0e\xc5\x5e\x3c\xb9\xc5\x5c\xbd\x1f\x92\xaa\xb4\xc4\xf7\xe8\xe0\x19\x29\xfb\x38\xd4\x9c\x6e\xcc\xe1\xcf\x51\x96\x77\x43\xf3\x61\xd9\x11\xd9\x6f\xee\x6e\xdb\x32\x2b\x96\x93\x0e\x85\xcd\x8a\xa1\x87\xbc\xbb\x5d\x40\xe1\x7d\xbd\x98\xcd\x06\xa9\xef\x0f\x73\xce\x97\x9e\x1b\x43\x86\x7c\x59\x9a\x46\xb2\x5e\x04\x4b\x64\x9b\x6b\x85\x32\x4d\xba\x09\xa2\x7f\x69\xcd\xaf\x98\xf9\xee\xf8\xad\x1c\x45\x96\x99\x86\xd2\xf9\x61\x2e\xad\x42\x84\x19\xc1\x2f\x6c\x4d\xa2\xe4\x1a\xb3\x36\xce\x29\x8e\x11\x43\xe0\xc7\x53\x04\xa9\x5c\x46\xe1\x0d\x25\xe3\xc9\xad\xa9\xc2\x79\xf5\x4e\x59\xa3\x2b\xd4\x1c\x30\x07\x19\x7c\x5f\x96\xbe\xa2\xca\xba\x55\x11\x2a\x00\x1c\x14\x86\xfc\x19\xd7\xf7\xa1\x20\x40\x37\x48\xec\x51\x46\xce\xb1\xe7\xe2\x15\x21\x6c\x4e\x06\xb5\x2a\x87\xb3\x56\x61\x95\x3b\x56\x60\x56\x42\x4e\x3a\x29\x68\x29\x0d\x91\x86\x81\x97\x0b\xce\x98\x56\x70\x65\xd3\x35\x45\x06\x72\x5d\xb7\xa9\x48\x24\xb5\x62\xce\xc6\x3c\xfc\x5e\xee\x16\xcb\x11\x4a\x55\xb9\xc2\x92\xb0\x2f\x50\x07\x66\x71\xce\xd2\x66\xfe\x94\x25\x6a\x09\x5c\xc9\x50\x06\x4e\xf9\x34\x55\x24\x9c\x42\x74\xd9\x85\xa3\xe8\x68\x34\xc9\x29\x14\x6a\x3d\x29\x1f\xd0\x9a\x71\x28\x55\x44\x59\x42\x25\x0e\x90\x96\x26\xdb\x12\x21\x48\x34\x10\x55\xb4\x4d\x20\xac\xaf\x38\xda\x8a\x2f\x45\x40\x47\xd6\xa8\x5c\x11\x72\xbe\x61\x06\xd9\x16\x33\xcc\x42\xa2\xb4\x2d\x64\xb8\x71\x61\x83\xd8\xb3\x7b\x95\x1a\x8d\x29\xad\xb8\x82\x38\xae\xd7\x18\x9f\xa4\x74\xdb\xe8\x63\x19\x49\x4a\xad\x50\x52\xc1\x49\xe3\x57\xbd\xf3\xc1\x92\xa9\x6a\xa9\x88\x3a\xde\x17\x54\xa4\x54\x2f\x45\xbe\x15\x8b\xae\xae\xec\x14\x84\x41\xd7\xa6\x56\x59\x27\xca\x6f\xe1\xbc\x63\xe3\x0a\x2e\x62\xcb\xe9\x1b\x78\xe9\x17\xeb\x4b\xee\xa7\x05\xbb\x59\x37\x56\x03\x15\x66\xec\x14\x9c\x17\x9e\x2a\x3e\x8a\x67\xaa\x44\x3b\x85\xb7\x05\x6a\x40\x4d\xae\x55\x8e\xdb\x54\xa3\x6f\xb5\xa0\xeb\xd3\xb7\x17\xcb\x4b\x46\xd9\x57\x53\xde\x40\xae\xb4\x6c\x4b\x25\xae\x2f\x2d\x82\xf3\x4d\xb6\x25\x8d\x14\xf0\x5b\x83\x0d\x99\x25\xef\x4b\x79\x85\xb5\xc6\x52\x12\xa2\x65\x2c\x0a\x63\xaa\xd3\xa6\x0a\x01\x92\xbc\x91\x95\x94\xa6\x1c\x06\xfd\x83\xbb\x8e\xee\xd8\x40\x08\x2d\x9f\x38\x48\x09\x0a\xe9\x79\xd1\xe7\x58\xc5\x83\x56\x24\x7f\x0b\x8b\x2e\x6c\xc4\x76\x13\x0e\xfd\xaf\xae\x6b\x57\x46\x7d\xf7\x85\x71\xc4\xac\xda\x58\xdf\xeb\x5b\x0f\x74\xb4\x33\x85\x7b\x32\x82\x35\x56\x75\x29\x3c\x76\x6e\xac\x1f\x5a\xc0\xff\xb2\xb2\x35\x5a\x91\xe6\x21\x9c\xc1\x4e\x68\x55\x96\x82\x87\x37\xe8\x51\xef\xe0\x0c\xd6\xe1\x68\x10\xf3\x1a\x2e\x2d\xce\xe0\xe3\xc7\xe9\x75\xf7\xfd\xe9\x13\x03\x08\xbb\x69\xc8\x85\x3a\x38\x6b\xf3\x25\x2a\x78\x27\x93\xd8\x8c\xf8\xf8\x71\x7a\xc9\xbf\x3e\x7d\x82\xc9\x84\xf8\x3b\x51\x92\x46\xd7\xc2\x6d\x6f\x64\xc4\x42\x29\x27\xe3\x8f\xd9\xd0\xa7\x4f\xb3\xd0\x91\x9d\x70\x24\x9b\x94\x66\x13\xc8\x21\xd9\xdd\x87\x8c\x31\x3e\x34\x17\x19\xcc\x70\x77\xf1\xf3\x70\xa6\xf1\x0c\xe7\x0a\xd3\x94\xf2\x9d\x8f\x09\xd3\xbb\x9c\x13\xfe\x33\xf8\x9f\xeb\x15\xcf\x93\x1f\x7c\xe7\x4d\x0f\xd0\x21\xfe\xe5\xf5\xbb\xeb\xff\xbe\x59\xbf\xfb\xe5\xee\xdd\xf5\x7f\xdd\x5c\xae\x19\xfc\xe3\x47\x95\x83\x46\x98\x52\x2d\x04\x73\x98\xc4\xd3\x7d\xfc\x58\x5b\xa5\x7d\x0e\x27\xb1\xf1\xf1\x2e\x23\x80\x33\xf8\x93\x3c\x09\xc0\x1d\xe0\x04\x50\xcb\xee\x2b\xa2\xe3\x7a\x89\x0a\x9f\xdf\xc1\x58\x61\x45\xd9\xe8\x19\xfc\x69\x3a\xcf\xe1\xf9\xc5\x49\x5c\xf6\xfb\x98\x43\x51\xf5\x05\xd4\x92\x8a\xb3\x21\xe2\xb0\xea\x01\x66\xfe\x64\x6b\x4b\x92\xe5\xc5\xea\x9f\xc6\xff\x07\x35\xfe\xd1\xbf\xa4\x4a\xcf\x52\xe1\x8a\xf0\xb9\xbc\x58\xc1\xe4\xf5\x03\x9b\x0c\xe3\xe6\x4b\x36\x14\xc0\xf0\x4b\x26\xf9\x65\xdb\x08\x88\xca\x90\xb3\x9e\x9d\x2e\xea\x5a\x9f\x3d\x81\x81\xb4\x68\x2b\xac\xce\x48\x85\x37\xe9\x13\x98\x46\x8b\x94\x1c\x46\x8f\xf5\xf7\xec\xe2\x9e\xef\xfc\x4a\x5f\x79\x73\x75\x24\x96\xe4\xb9\x55\xf2\x9a\xaf\xa4\x16\x7f\x9b\xac\xbf\x7b\x54\xd2\xdf\x7d\x8d\x9c\xbf\xfb\x0a\x29\x13\x50\x27\xc1\xaf\x95\xfb\x77\x30\xa9\x11\xaa\x5a\x3d\x85\x3f\x0c\x14\x14\xef\x76\xad\xbc\x9f\x3f\x85\xb8\x23\xd2\xdc\xa9\x0f\xd8\x61\xfd\xf6\xe2\x5e\x95\x8d\xad\xfe\xe9\x47\xff\xb0\x7e\x74\x76\x6c\x5c\xab\x8b\xf3\xf5\xe5\x0b\x98\x4c\x7e\x35\xe9\x84\x8b\x90\x07\x96\xd6\x81\xe8\xc0\xeb\xd3\x7b\xc3\x21\xc1\xf9\x92\x95\x75\xe0\x31\x1f\xf9\x82\xe9\x7e\x85\x0d\x76\x18\x29\x33\x99\xd4\x68\x59\x1f\x9f\xc4\x20\x3b\xd4\x15\x56\x9c\x44\x3c\x49\x72\xd2\xa3\xf5\x55\xdd\xa3\xfd\xf6\x36\xc9\xed\x9a\x0b\xe1\xb3\x02\x24\xba\xcc\xaa\x34\xaa\xfd\x71\x43\xbd\x2d\x2d\xcf\xdf\xae\x20\x40\x3f\xb8\x34\x6a\xf1\x3c\xa9\x81\x77\xfb\xb5\xda\x7f\xdf\xb0\x35\x17\xda\x7c\x4f\x13\xec\xb7\xb7\xdd\x3f\xbc\xdd\x0e\x0f\xf7\xa8\xd5\x8e\xe0\x3f\x4c\x1a\x2e\x3e\x58\x0a\x99\xd0\xdc\x33\x50\xbe\x40\xbe\xa6\xe2\x67\x05\x51\x32\x95\xf8\x60\x74\x77\xbd\x01\xaf\x69\xee\xd9\xf9\xdd\xeb\xef\xe9\xc8\x47\x78\x16\x70\x12\xad\x8a\x2c\x5b\x62\x7e\xd2\xee\xf5\x9f\xe4\x17\xff\xbe\x6d\x18\xc5\xf1\x0e\xec\x6d\x4f\xee\xf5\x15\xdb\x3e\x9d\xab\x31\x53\xb9\x42\x09\xbf\x9a\x34\x38\xe6\xf0\xae\xc1\xc4\xab\x0a\x86\xa2\x39\xd9\x33\x42\x3d\x68\x4b\xf6\x0d\xc8\x61\x9b\x71\x04\x2f\x9b\x14\xad\xe6\x86\xf7\xd7\xe8\xf7\x00\xfc\x81\x82\xf7\x73\x8b\xee\x66\xb8\x6b\xef\x73\x8b\xcb\x9b\x38\x10\xa2\x86\x9b\x0e\x9e\xc4\xb4\x90\x6e\x01\xd2\x64\x5b\xb4\x63\xd8\x76\x08\x13\x80\xeb\x38\xbf\x80\x93\x30\x7f\x02\x4f\x69\x48\x2f\x87\x7b\x3d\x6a\x49\xff\xe8\x21\xf2\x48\x57\x48\xad\xc3\x05\x9e\x37\xe0\x6a\xb1\xd7\xa4\x82\x2e\x36\x31\x13\xe8\x01\xba\xbe\xef\x2b\xe1\xc2\x93\x28\xf2\x8e\xa4\xaf\x3e\x06\xd8\x29\xac\x10\x43\xcf\xd9\x2d\x66\xb3\x5e\xea\x53\x65\x66\xd2\x64\x6e\x66\x31\x47\x8b\x3a\xc3\x59\x77\x01\x3c\x00\x9b\x88\x5a\xcd\x76\xa7\xd3\xd3\x7f\x9b\x8d\xc8\x9c\x76\xa7\xf1\xcd\xd6\x30\x86\x47\x22\x86\x71\xbe\xed\x96\xd6\x83\x8b\x85\x2a\x10\x79\x44\xde\x63\xe9\x41\x78\x05\x72\x7c\xa7\x9a\xdf\xb7\x90\x56\xa5\xa7\x49\xeb\x3f\x3a\xa4\xfd\x3b\x83\xee\x2a\xa9\xd3\x14\xee\x25\xb7\xea\xfe\xe0\x04\xf7\x27\x3e\x77\x92\xce\xf2\x86\xe7\x78\x6c\x71\xf4\x04\x4f\xdd\x15\x1c\xc1\xca\x1b\x2b\x36\xf8\x0d\x1a\x8e\xa3\xbf\xe3\x52\xe0\x73\x57\x02\xc9\x08\x6e\x4d\x26\x42\x91\x07\xee\xe0\x3c\x56\xd3\x84\x87\xe2\x41\x82\x57\x7b\x5b\x28\x8f\xa5\x72\xdc\x5e\xe7\x26\xff\xe0\x76\x90\x04\xe0\x60\x5f\xa8\xac\x68\x5d\x8d\x72\x20\xca\xd2\xec\x83\xfb\x0b\x2f\x78\xf8\xb2\x20\x0c\x5e\xa9\xbe\x23\x3c\x9d\x11\x15\x2f\xd6\xeb\x65\xdc\xb1\x7b\x90\xe2\x4d\x77\x0b\x08\x75\x93\x96\x2a\x83\xd0\xbc\x8a\xad\xea\x3d\xa6\xb0\x53\x02\x04\x3c\xbf\x5e\xb7\xef\x6d\xa6\xc9\x00\xd5\xe2\xe8\x9e\x80\x14\x8f\xec\xed\x99\xfb\x7e\xb8\xc2\x1d\xb5\xd8\x7f\x9c\xbb\x24\x09\xf1\x6c\xf5\xe3\xa2\x77\x56\xb2\xbd\xbc\x7c\xe2\xe7\x3f\xf7\x5e\xe7\x3c\xd5\x55\x5a\x6f\xa7\xc8\x37\x2a\xc4\xd7\xf6\xfd\x25\xd3\xb0\xfa\x11\x6a\x6b\x76\x4a\xa2\x75\xe0\x9a\xac\x00\xe1\xe0\x95\xd2\xca\xb4\xf7\x9d\x97\x58\x17\xc9\x08\x9e\x93\xf3\x51\x19\x31\x23\x3c\x84\xeb\x19\xc2\x4e\x3b\xbc\x84\xbb\xd6\xb2\x36\x4a\x87\xdd\xc3\x50\x4b\x72\xf8\x1a\x12\x17\xee\xd3\x06\x32\x7a\x8c\xc7\x7f\xe0\x1b\xb3\xd5\x5e\xe5\xfe\x71\xba\xdf\x38\xb4\xaf\x3f\x73\x29\x02\x70\xde\xf8\x82\xaf\x1a\x63\xad\xa5\x85\xf6\x03\xe8\x30\x10\x5f\xf8\xb4\xa9\xcc\x60\x7e\x04\x3f\xcd\xe7\xf0\xea\x82\xe8\x2a\x1a\xbd\x5d\xa9\x0f\x78\x71\xa0\xd4\x83\x26\xc2\x5f\x92\xfc\x7c\xa4\xfe\x0f\x49\xec\x5f\x3b\xcd\x5d\xa4\x78\x01\x27\x42\x1b\x7d\xa8\x4c\xe3\xee\x91\x3d\x18\xff\xbf\x00\x00\x00\xff\xff\xb2\x17\xf6\x13\x17\x2d\x00\x00")

func configDefaultConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configDefaultConfigYaml,
		"config/default-config.yaml",
	)
}

func configDefaultConfigYaml() (*asset, error) {
	bytes, err := configDefaultConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/default-config.yaml", size: 11543, mode: os.FileMode(436), modTime: time.Unix(1702316879, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGridengineTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xcd\x4a\xc4\x30\x14\x85\xf7\x79\x8a\x6b\xc7\x59\x26\xed\x0b\xb8\xb2\x30\xb8\x71\x21\x82\x4b\x69\xc9\x0d\x13\x32\xf9\xe1\x26\x51\x30\xe4\xdd\xa5\x69\x11\x0a\x75\x76\x97\xc3\x77\x3e\xb8\xe7\xf4\xd0\xcf\xda\xf5\xf3\x14\xaf\xec\xf4\x08\xfc\x15\x4a\x11\xef\x53\x34\x2f\xb2\xd6\x96\xf8\x25\xf9\xf0\x64\x46\x4d\xb5\xf6\x2a\x3b\x87\x37\x1e\x93\xf4\x39\x35\x00\xff\x03\x90\x88\x95\xa2\x15\x38\x04\xf1\x1c\x72\x84\x01\x78\xad\xac\x94\x40\xda\x25\x05\xdd\x52\x0f\x08\x36\x68\x38\xcb\x6e\x85\x1a\xc0\x01\x9d\x6c\xd7\x56\x7f\x9b\xec\x65\x86\x41\x1c\x19\x6e\x70\xfd\xfc\xb2\x68\x9f\xce\x62\x50\x97\x6e\x83\x8f\x3d\xa3\x8e\xe6\xae\x48\x45\xfd\x83\x7f\xa6\x15\xdf\xa9\xd8\xfa\x20\x7c\x7b\x32\x48\x40\xd9\x01\xe7\x69\x59\x6c\xdc\x6d\xf7\x1b\x00\x00\xff\xff\xcf\x92\x30\x7f\x5a\x01\x00\x00")

func configGridengineTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configGridengineTemplateTxt,
		"config/gridengine-template.txt",
	)
}

func configGridengineTemplateTxt() (*asset, error) {
	bytes, err := configGridengineTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gridengine-template.txt", size: 346, mode: os.FileMode(436), modTime: time.Unix(1701431150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHtcondorTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xcd\x6e\xea\x30\x10\x85\xf7\x7e\x8a\x11\xd2\x5d\x26\x97\x17\xc8\xa6\x80\x10\x9b\x22\xd1\xa8\x3f\x2b\xcb\xe0\x49\xb0\xe2\x8c\x61\xec\x09\xad\xa2\xbc\x7b\x15\x40\xad\xa8\x4a\x77\x47\x3e\xdf\xf9\x3c\x42\xae\x43\x8e\x08\x05\x74\x86\x9c\xf7\x46\xd5\x98\x90\x3a\x28\xa0\x64\x41\x85\xef\xb8\x93\x64\xb6\x7e\x44\x2a\x21\x42\xaf\x0c\xd7\xd2\x22\xa5\x08\x05\x9c\x02\x37\xc8\xc0\x42\x90\x65\xc9\xc4\x66\x35\x87\xbe\xcf\xcb\x31\xd9\x61\x50\x3e\xd4\x50\x8c\x2f\x2f\x81\x9b\xb9\xe3\x61\xf8\xbf\x0b\x64\x03\x67\xd8\x21\xa5\xcc\x87\x5a\x21\x73\xe0\x9f\xd4\xe5\xaf\x2c\x26\x8b\xcc\x2a\x48\x3a\x48\xba\xcf\x04\x49\x2a\xee\x83\x78\xab\x13\x1b\x8a\x15\xb2\xae\x9c\xc7\xf1\xc4\xb7\xc5\x93\x3a\xed\x91\x74\x0a\xdf\xe5\x97\x70\xfd\xa8\x17\xaf\xab\x52\xaf\x37\x7a\xf1\xbc\x9a\x95\xaa\xef\x5d\x05\x84\x90\xcf\x0e\x12\x61\x0a\xd9\x30\xa8\xbe\x3f\xb0\xa3\x54\xc1\x84\xf1\x28\x18\x93\xde\x8d\x65\x01\xff\xec\xe4\x02\x9e\xa1\x0c\x90\xec\x39\x5d\x15\x1b\xd3\x2e\xb7\x30\xcd\xef\x59\x5a\x6c\x03\x7f\x8c\x9e\x7c\x5a\xc1\xf2\x61\x72\x9d\xfc\x6e\x9b\xbb\xd8\xfc\xa9\xb3\x2e\x36\x37\xb2\xcb\xe2\xc6\xa6\x8e\x82\x82\xea\x33\x00\x00\xff\xff\x61\xe0\x01\xde\xf9\x01\x00\x00")

func configHtcondorTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configHtcondorTemplateTxt,
		"config/htcondor-template.txt",
	)
}

func configHtcondorTemplateTxt() (*asset, error) {
	bytes, err := configHtcondorTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/htcondor-template.txt", size: 505, mode: os.FileMode(436), modTime: time.Unix(1701431150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configKubernetesExecutorTemplateYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x4d\x6f\xda\x40\x10\xbd\xfb\x57\x8c\x4c\x7a\x8b\x8d\xb9\xf4\x60\xa9\x87\x08\x2a\x42\xd4\x10\xd4\xa2\xf4\x50\xf5\xb0\x5e\xc6\xb0\xf5\x7e\x75\x67\x4d\x8a\x2c\xfe\x7b\xb5\xb6\xb1\x4c\xaa\x26\x7b\x1a\xcf\x9b\x79\x6f\xde\x68\x64\x66\xc5\x33\x3a\x12\x46\xe7\x50\x30\xcf\x0f\xd3\xe3\x2c\xaa\x84\xde\xe5\xf0\x60\x8a\x48\xa1\x67\x3b\xe6\x59\x1e\x01\x68\xa6\x30\x87\xa6\x49\xb7\x8c\xaa\xd5\xee\x7c\x4e\x9a\x26\x7d\x30\x45\x08\x7b\x98\x2c\xe3\x5d\xcd\xfa\xf2\xd5\x62\x92\x15\x28\x29\x90\x00\xfc\x32\x45\xf2\x26\x15\x59\xe4\x39\x44\x00\x05\xe3\x95\x29\xcb\x2f\x42\x09\x9f\x43\x16\x01\x70\xa3\xac\x44\x2f\x8c\xa6\x1c\x66\x11\x80\x47\x65\x25\xf3\xd8\x51\xb7\x9d\x6d\x04\xe0\x90\x3c\x73\x7e\x63\xa4\xe0\xa7\x1c\xd6\x78\x44\xd7\x43\xdc\x68\xcf\x84\x46\x47\xad\x4c\x78\x49\xef\xae\xac\xb5\x46\x99\xbc\x18\x57\xa1\x4b\x46\x03\xf6\x75\x00\x42\xb1\x7d\x37\xfb\x2a\x44\xaf\x91\x4d\x2d\xe5\x45\xf3\x4e\xbe\xb0\x13\x0d\x38\x37\x4a\xb1\xb0\xd8\x1f\xf1\xb4\x10\x7a\x4a\x87\xf8\x16\xe2\x84\xc7\x3f\x87\x12\xe6\xf6\xd4\x72\xcf\xbb\xda\x11\x7b\x18\x49\xe8\xfd\x42\xb8\xb6\xe0\xbb\x71\xd5\x4e\xb8\x51\x81\x43\x32\xb5\xe3\x48\xf9\x90\x0a\xc9\xdf\x35\x92\xbf\xca\x01\x70\x5b\x07\x12\x51\x82\x46\x48\xe7\xb6\x26\xc8\x20\x39\x9f\x83\xb0\xad\x29\x04\x80\x92\x10\x42\x14\xcf\xb2\x4c\xc5\x21\xc2\xab\x81\xc2\x53\xa8\x8c\x3b\x8d\xb8\xbe\x32\xb5\x2c\x20\x4b\x7b\x3a\xeb\x84\xf6\x25\xc4\x1f\xd2\xac\x5c\xc6\x3d\xdc\x52\x49\xc2\x8e\xfc\xe3\xe3\x7f\xb8\xd1\x1e\x50\xa1\x63\x32\x21\x6f\x5c\xbf\xf4\x5e\x66\x21\xa8\x7a\x4b\xa7\xc3\xaf\x85\xb2\x6c\xac\x34\x48\x1d\x8d\xac\x15\x3e\x9a\x5a\x8f\xb7\x34\x99\x4c\x60\xf1\x04\xeb\xa7\x2d\xcc\xef\xef\xd6\xcb\xcf\xb0\xbd\x5f\x7d\x1b\xe0\xa6\x71\x4c\xef\x11\x6e\xc4\xee\xcf\x2d\xdc\x08\x8f\x0a\xf2\x4f\x90\x3e\xb7\x64\x34\x72\x72\x39\xac\xde\x03\x8c\x2c\xaa\xa0\xb9\x61\xfe\x10\x8c\xb5\x1c\xe9\xfc\x72\x99\x21\x7d\xb5\x0f\xaa\x8b\xa1\x74\x38\xca\xf7\xda\x5e\x99\xed\xac\xfe\x7b\xf4\xfd\x6c\x43\x9b\x0d\x3f\x04\xf2\xa8\x7d\xe7\x67\x2e\x99\x50\xe3\x0b\xe2\x21\xb1\x1e\xf7\x26\xf6\xc8\xe1\x6f\x00\x00\x00\xff\xff\x1a\xb2\xe9\x0d\x4c\x04\x00\x00")

func configKubernetesExecutorTemplateYamlBytes() ([]byte, error) {
	return bindataRead(
		_configKubernetesExecutorTemplateYaml,
		"config/kubernetes-executor-template.yaml",
	)
}

func configKubernetesExecutorTemplateYaml() (*asset, error) {
	bytes, err := configKubernetesExecutorTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/kubernetes-executor-template.yaml", size: 1100, mode: os.FileMode(436), modTime: time.Unix(1702337819, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configKubernetesTemplateYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\x16\x0a\x7a\x94\xa5\x5c\x7a\xe0\xcd\x48\x82\xb4\x45\xad\x04\x45\x90\x3b\x45\xad\x62\xc2\x7c\x95\x4b\x39\x30\x0c\xfd\x7b\x41\x51\x6e\x2d\xbf\x0e\xe5\x89\xdc\x21\x67\x66\xd7\x63\x71\x27\xdf\xd1\x93\xb4\x86\x41\xc3\x83\x58\x97\xdb\xfb\x6c\x23\x4d\xcb\xe0\x87\x6d\x32\x8d\x81\xb7\x3c\x70\x96\x01\xdc\xdd\xc1\xe3\x0b\xd4\x2f\x6f\xf0\xf0\x6d\x59\x3f\x3f\x41\xbd\x5c\x3d\x65\x00\x86\x6b\x64\xb0\xdf\x2f\xde\x38\x6d\xbe\xb7\xc3\x30\xd5\xc8\x71\x91\x80\xfa\x70\x1a\x86\x8c\x1c\x0a\x06\x19\x40\xc3\xc5\xc6\x76\xdd\x4f\xa9\x65\x60\x50\x65\x00\xc2\x6a\xa7\x30\x48\x6b\x88\xc1\x7d\x06\x10\x50\x3b\xc5\x03\x46\x75\x80\xf1\xe5\xb8\x03\x20\xf4\x5b\x29\x70\x29\x84\xed\x4d\xa8\x47\x07\x5d\x6f\x0c\xaa\x82\xf8\x74\xc7\x23\x05\xee\xc3\xab\x55\x52\xec\x18\xd4\xb8\x45\x3f\x41\xc2\x9a\xc0\xa5\x41\x4f\xa3\x95\xb4\x8a\xa9\x93\x89\xe7\xd3\xfa\x0d\xfa\x62\xde\xd7\x61\x49\xcd\x3f\x90\x81\x5d\x53\x1f\x5d\x37\xd2\x96\xd3\xb3\x56\x9a\xb6\xf0\xd6\x06\x85\x44\x2c\xba\xa7\x70\xfa\xf0\xb5\x57\xea\x60\x6b\xa9\x3e\xf9\x8e\x8e\x6e\x70\xff\x41\xec\xe8\x1c\x8d\xe5\xc9\x4c\x7e\x5a\xf6\xbd\x39\xab\x15\x85\xb0\xa6\x93\x1f\x67\x40\x89\x41\x94\x09\x2b\xe7\x3d\xa6\xe2\x62\xa7\xd5\x05\xb6\x10\xbb\x7f\x3c\x05\x2e\x8f\xc5\x23\xd9\xde\x0b\x3c\x69\x20\x02\xbf\x7b\xa4\x70\x56\x07\x10\xae\x8f\x19\x91\x1d\x18\x84\xc5\x83\xeb\x09\x2a\x28\x86\x61\xbf\x1f\x0f\x71\x03\xa8\x08\x21\xee\xf2\xfb\xaa\xd2\x79\xdc\xa1\x99\x2b\xa7\xa5\x51\x5b\xbf\x3b\xe2\xfb\xc5\xf5\x73\x03\xd5\x62\xa2\x74\x5e\x9a\xd0\x41\xfe\x65\x51\x75\xcf\xf9\x04\x8f\x74\x8a\x30\x09\x7c\x5d\xdd\xe0\x47\xb7\x46\x8d\x9e\xab\x82\x82\xf5\x63\x04\xfe\x4a\x3d\x4a\xda\xdc\xd2\x4a\xf8\x5c\xac\xaa\x2e\xab\x6d\xad\xea\x35\xae\x62\xb6\xcf\xb2\x30\x0b\xe9\x64\xe3\x4a\x4a\xc7\x91\x44\x92\x57\x1e\xd6\x0c\x4a\xeb\xc2\xf4\xcb\x1f\x07\xa0\x68\xa5\x2f\xaf\x13\x50\xdf\xa4\xe7\xd7\xae\x1c\x2c\xa5\x14\x15\xc9\xfb\x0d\x17\xff\x62\x98\x65\xc7\xed\x5e\xfa\x2f\x5e\xe3\x4c\xf5\x15\x77\xf3\xe9\xcc\x66\x33\x69\xfc\xd7\xe4\x5c\xfc\x26\x52\x40\x13\xde\x47\xe9\x07\xc5\xa5\x9e\x6b\x89\x58\x4a\x5f\x9e\x03\x97\xdb\x0a\xf8\x13\x00\x00\xff\xff\x57\x40\xab\x27\x51\x05\x00\x00")

func configKubernetesTemplateYamlBytes() ([]byte, error) {
	return bindataRead(
		_configKubernetesTemplateYaml,
		"config/kubernetes-template.yaml",
	)
}

func configKubernetesTemplateYaml() (*asset, error) {
	bytes, err := configKubernetesTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/kubernetes-template.yaml", size: 1361, mode: os.FileMode(436), modTime: time.Unix(1702337830, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configPbsTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xd0\xc1\x4a\xc4\x30\x10\xc6\xf1\x7b\x9f\x62\xec\xb2\xc7\xb4\xf5\x2a\xf4\xa2\x05\xf1\x22\xa2\x82\xe7\x86\x4c\x34\xa4\x9d\x84\x49\x82\x87\x90\x77\x97\xdd\x04\x64\xc1\xf5\x16\x86\x3f\x3f\xc8\x77\xb8\x19\xa5\xa1\x51\xae\xe1\xab\x3b\xbc\xdc\xbf\x81\x78\x86\x9c\x87\xf7\x35\xd8\x27\x55\x4a\xbb\xb9\xd3\xed\xc3\xb1\x5d\x0c\x97\x32\xea\x44\x84\x9b\x08\x51\xb9\x14\x5b\x82\xd7\x12\x64\xee\x72\x36\x1a\x08\x61\x78\xf0\x29\xc0\x04\xa2\x94\x2e\x67\xcf\x86\xa2\x86\xbe\x02\x1b\x90\x53\x18\xe6\xdb\x3b\xef\x69\x3e\xaa\xbe\xd6\xe7\x52\x00\x92\x3a\xbf\x9a\xf3\xba\xee\x8f\x12\xa6\xe1\x1a\xb5\xe3\x3e\x1f\x87\x49\x7f\xca\xbe\xc5\x7f\x3b\x8b\x09\xf6\x5f\x48\x9b\x0d\x7f\xa5\x9a\x5f\x50\x5d\xfd\x29\x7c\x3b\xb6\xc8\xc0\x89\x40\x88\x78\x9a\x6f\xb9\x18\xf2\x27\x00\x00\xff\xff\x94\x91\x42\x0b\x69\x01\x00\x00")

func configPbsTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configPbsTemplateTxt,
		"config/pbs-template.txt",
	)
}

func configPbsTemplateTxt() (*asset, error) {
	bytes, err := configPbsTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/pbs-template.txt", size: 361, mode: os.FileMode(436), modTime: time.Unix(1701431150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configSlurmTemplateTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xc1\x6a\xc3\x30\x0c\x86\xef\x79\x0a\x2d\xa5\x47\x27\xd9\x23\xac\x0d\x74\xbb\x6e\x85\x9d\x9d\x45\x61\x9e\x67\xd9\x48\x36\x3b\x18\xbf\xfb\x48\x53\x68\x0a\x0b\xbb\x19\xeb\xfb\x3f\xa4\x7f\xf7\xd0\x0e\x86\xda\x41\xcb\x67\xb5\x7b\x3b\x3c\x9d\x8f\xcf\xa0\xd4\x97\x1f\x14\x69\x87\x90\x73\x73\xd6\x62\x5f\xc6\x52\x56\x63\x8a\x5a\xac\xc0\xe3\xea\x0b\x99\x3d\xcf\xf8\xbb\x67\xdb\x1b\x2e\xa5\x9d\x12\x11\x7e\x2b\x89\x23\x32\xaf\x50\x9f\x62\x48\x71\x8b\xf5\x29\x56\x39\x9b\x09\x08\xa1\x39\x86\x24\xd0\x81\x2a\xa5\xca\x39\xb0\xa1\x38\x41\x7d\x33\x7d\x84\x24\x2a\x20\xab\x79\x1f\xd8\x8f\xf5\x92\xb8\xd0\x0a\x90\xc6\xcb\xeb\xea\x7a\xd5\xee\x34\x40\xd7\x6c\xeb\x1c\x3a\xd8\x37\xdd\x74\x3a\xd4\x57\xfc\x6f\x53\x6f\xc4\xfe\xa3\x8a\x2e\xdc\x54\x0b\x7f\xe7\xaa\x96\x83\xe1\xc7\xb3\x45\x06\x4e\x34\x87\xe6\xa6\xfb\xbb\xce\x7f\x03\x00\x00\xff\xff\x71\x9e\x5b\xbd\x9f\x01\x00\x00")

func configSlurmTemplateTxtBytes() ([]byte, error) {
	return bindataRead(
		_configSlurmTemplateTxt,
		"config/slurm-template.txt",
	)
}

func configSlurmTemplateTxt() (*asset, error) {
	bytes, err := configSlurmTemplateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/slurm-template.txt", size: 415, mode: os.FileMode(436), modTime: time.Unix(1701431150, 0)}
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
	"config/default-config.yaml":               configDefaultConfigYaml,
	"config/gridengine-template.txt":           configGridengineTemplateTxt,
	"config/htcondor-template.txt":             configHtcondorTemplateTxt,
	"config/kubernetes-executor-template.yaml": configKubernetesExecutorTemplateYaml,
	"config/kubernetes-template.yaml":          configKubernetesTemplateYaml,
	"config/pbs-template.txt":                  configPbsTemplateTxt,
	"config/slurm-template.txt":                configSlurmTemplateTxt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"config": &bintree{nil, map[string]*bintree{
		"default-config.yaml":               &bintree{configDefaultConfigYaml, map[string]*bintree{}},
		"gridengine-template.txt":           &bintree{configGridengineTemplateTxt, map[string]*bintree{}},
		"htcondor-template.txt":             &bintree{configHtcondorTemplateTxt, map[string]*bintree{}},
		"kubernetes-executor-template.yaml": &bintree{configKubernetesExecutorTemplateYaml, map[string]*bintree{}},
		"kubernetes-template.yaml":          &bintree{configKubernetesTemplateYaml, map[string]*bintree{}},
		"pbs-template.txt":                  &bintree{configPbsTemplateTxt, map[string]*bintree{}},
		"slurm-template.txt":                &bintree{configSlurmTemplateTxt, map[string]*bintree{}},
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
