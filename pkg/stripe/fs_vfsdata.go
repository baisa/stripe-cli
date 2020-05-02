// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package stripe

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS statically implements the virtual filesystem provided to vfsgen.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/DigiCertGlobalRootCA.crt.pem": &vfsgen۰CompressedFileInfo{
			name:             "DigiCertGlobalRootCA.crt.pem",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 1338,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd4\x4d\xcf\xba\xba\x12\x00\xf0\x3d\x9f\xe2\xee\xcd\x0d\x28\x28\xb2\xb8\x8b\xb6\x14\xac\x58\x90\x77\x71\x27\xef\xf2\xa2\x08\x48\x91\x4f\x7f\xe3\xf3\x24\x27\xff\x93\x73\x9a\x74\x33\xf9\xb5\x9d\x4c\x27\xf3\xdf\xef\x82\x58\x27\xe6\x7f\x10\x76\x3c\xa2\x11\x04\x3c\xfc\x13\xe5\x28\x21\x6a\xbf\x20\x04\xba\xac\x00\x8c\x40\x50\x10\x1b\xa9\x53\x11\x74\x10\x39\xbd\x5e\xa6\x61\x7f\x0c\xaf\x87\x83\x5b\x01\x13\x16\xf5\xab\xac\xef\xba\xc2\x04\x08\x6c\x5f\x03\x2a\x2c\x39\x6a\x0f\x0c\xd9\x91\x1a\xd8\xb6\x8e\xd9\x31\xf0\x17\x1c\x50\x48\x75\xb0\xf6\x31\x2a\x29\x75\xf4\xe6\x71\xf3\xcd\x26\x79\xd8\x85\x1b\x6e\x2b\xea\xd4\x4c\x63\x3f\xfe\x84\x67\x28\x72\xa9\x98\xbc\xaf\x5f\x14\xfe\xa0\x77\xb4\x51\x46\xea\x02\x76\x28\x7e\x90\x8a\xe7\x14\xdf\xc2\xb4\xb3\x37\xc1\x27\x45\xf0\x10\xeb\xca\x3d\x0a\x59\xe1\xb7\xca\x94\x22\xa8\x72\xb6\x07\x32\x8d\x09\xcc\xac\xf0\x4c\x3d\xc0\xa8\xfa\xdd\xf0\xa6\x31\x61\xa1\xde\xdf\x63\x54\xc7\x33\x5a\xc0\x11\x16\x66\x00\x41\xe4\x81\x26\xf0\x38\xea\xf8\x0c\xff\xa6\x64\x60\x36\xff\xf9\xda\x31\x6e\xe9\xac\x7b\xe0\xf2\x7b\x60\xf0\xf0\x21\x15\xd3\xcf\xb6\xfe\x9a\xe8\xc7\x6c\x2b\x2e\xde\x08\x33\x51\x41\xf6\x8b\xa8\xa7\x09\x4e\x77\xdd\x34\xea\xf5\x72\x14\x08\x4e\x87\x78\x73\x2c\x63\x04\xdd\x78\xa3\x08\x04\x9b\x90\x12\x02\xc9\x1f\x05\xe5\x7e\x2b\x8a\x35\x00\x2c\x04\xec\x3d\xf8\x02\x54\x18\x08\xd8\x18\x48\xd5\x54\xe2\xcb\x29\x7b\x19\x9e\xf7\x5c\x67\x2f\xdf\x30\xce\x48\xcc\xec\xcf\xcd\x68\xe4\xf2\x64\x35\xcd\x00\x39\xe4\xaa\x14\x5c\xad\x87\x57\x21\xd1\xe7\x53\x75\xd6\x6b\x10\x6c\xc5\x7b\xe5\x9e\xd2\x92\x5d\x01\x20\xf8\xb8\x0c\x52\x5c\xc8\x7c\xbe\x78\xe3\xec\xbc\x4f\xe1\x75\x48\xb4\x41\x8c\x1e\xda\x53\x91\xb9\x47\xb9\x0b\xf2\x6c\x27\xba\x06\x25\x9b\xf1\x36\x65\x05\xdb\xc2\x36\xe0\xdd\x46\xc8\x27\x98\x4b\x2f\x59\x7e\x1b\x66\x2a\xe4\x62\x27\xb5\x41\xab\xdd\xf4\x6d\x42\x96\xe3\x69\x12\x64\xb0\xd3\xba\x91\x93\x44\xc4\xa7\x33\xe2\x79\x70\xd8\x94\x69\xfb\x74\x20\x8c\xe8\xab\x59\xeb\xe6\xc5\x79\xf6\xdb\x83\x74\x4f\x5f\xca\xf1\xb9\xac\x70\x4d\x22\x32\xf9\x17\xd9\xde\x95\xa7\x55\xf9\xaa\x3b\x9a\x7b\xf2\x99\xf3\xd6\xca\x90\x36\xbb\xc2\x5d\x32\xe7\x31\xb2\xfb\xb6\x15\x2d\x0d\xbe\xac\xdb\x30\xad\x96\x98\xfa\x57\x98\x1f\xc2\x4f\x9b\xd1\x9e\xff\xc8\x53\xef\x21\xe1\xe4\xbf\xe4\x14\xd2\xf1\x49\xd7\x16\x2f\x71\x45\x1a\xca\x55\x50\xf0\xa3\x33\x3d\x5d\xf7\x7e\x4f\xcc\xe7\x0c\x4d\x51\x1c\xca\xf8\xe3\x81\xce\x82\xbb\x6a\x74\xab\x75\x36\x5e\x56\x55\x4d\xad\xe9\xc8\x88\x0a\x6c\x00\x9f\x1b\xca\x22\x0f\x58\xdc\xf7\xfb\x0e\xf6\x1e\x82\x7c\x8f\x21\xa0\x08\x44\x11\x53\xbf\x7d\xe1\x08\x1e\xb0\x0f\x3c\x04\x3e\x03\x0c\x43\x7e\x01\xe9\xaf\x95\xb0\x56\xd8\x3e\x50\xb6\xb6\x19\xc4\x0e\xe7\x9d\xc6\x76\x6f\x9c\xef\xfa\x3c\xa9\x8d\x4c\x14\x21\xf0\xd9\xe1\xf7\x82\x0a\xc2\x82\x69\x4f\xf0\x17\xfe\x37\xcb\xa9\x76\x74\x34\x9e\x57\x52\x4e\x89\xf9\xed\x07\x68\x03\xb5\x28\xbe\xc9\xbc\x13\x73\xd7\x11\x3c\x13\x63\x35\xae\xf1\x03\x2b\xee\x70\xf6\xf2\xbe\xf0\xd6\xd9\xa5\x26\xcf\x8f\x1d\xf1\x78\xe8\xb9\x92\x82\xf1\x9d\x5e\x0e\xfc\xe4\xc1\xc3\xba\x3a\xbd\xf5\x4d\x92\x3d\xbc\x47\x8b\xda\x1e\xc7\x97\x2a\x31\x50\xb9\xf8\x1f\xd2\x5e\x2d\x5a\x5f\xd4\xfb\x8b\xed\x93\xa9\xb3\x3a\x7e\x73\x0e\xb6\x20\x2d\x38\x61\x67\xf1\x8f\x60\x38\xee\xd3\xd0\x92\xd6\x67\xa1\x6a\xcf\xbb\xf3\x2e\x8f\x47\x3d\xce\xa3\x36\x0e\x85\x70\x0b\xab\x9c\x8c\x63\xd6\x89\x6e\xb7\x4a\x43\x8b\xf4\x61\x02\x01\x59\x09\xa3\x41\x8e\x1a\x77\x7e\x34\x7e\x7d\xbf\x45\x12\x81\xe4\xa5\xe6\xd3\xde\xbc\x6e\x23\x18\x67\xbd\x55\x58\x4b\xb8\x1b\x1c\x98\x48\x27\xe1\x71\x93\x7c\x7f\x65\xf4\xf5\xc6\xdf\xef\x77\x3e\x88\xc5\xd3\xbb\xc2\x81\xd0\x0c\x5c\xe4\xe2\x68\x6d\xbb\x63\xa6\xb2\xc1\x7a\xc2\xbe\x5b\xbd\x27\xcd\xf1\xba\x0d\x79\xc0\xb7\x57\x0e\x52\xa7\x0d\xf7\x49\xa9\xdf\x97\xa4\x09\x16\x15\xe8\x1f\xb7\x92\xd2\xa5\x13\x85\x74\x3f\xc6\x76\xcd\x21\xe0\x33\x19\x6d\x14\x24\x2b\xda\xb4\x46\xdb\x57\x7e\xee\x5b\x80\xdd\x3e\xb9\x93\xb9\x2b\x84\x8b\x24\x18\x67\x1a\x77\xeb\x6b\x18\xc4\xa9\xf4\x3f\xee\x67\x12\x62\x53\xfd\xe7\x74\xfc\x7f\x00\x00\x00\xff\xff\x55\xb3\x84\x99\x3a\x05\x00\x00"),
		},
		"/DigiCertHighAssuranceEVRootCA.crt.pem": &vfsgen۰CompressedFileInfo{
			name:             "DigiCertHighAssuranceEVRootCA.crt.pem",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 1367,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x94\xc9\xae\xab\x38\x18\x84\xf7\x3c\x45\xef\xa3\x56\x08\x64\x80\xa5\x6d\x0c\x61\x30\xc1\xcc\x66\x47\x80\x00\x61\xca\xcc\xf0\xf4\xad\x93\xb3\xea\x6e\xe9\x5a\xb2\x64\x95\x3e\x95\x6c\xd7\xaf\xfa\xfb\x67\x41\xac\xe9\xf6\x5f\x08\xbb\xbe\xae\xea\x08\xf8\xf8\xab\x72\x44\xd7\x95\xc9\x47\x08\xdc\x85\x12\x8c\x3a\x04\xa5\x4e\xc1\x7d\xca\x8c\x6e\xb0\xa8\xf1\x76\x90\xd8\xcf\x8f\x86\xb5\xf9\x02\x6c\x58\x36\xf7\xaa\xa9\x35\x79\xe4\x21\xa0\x81\x0a\x14\xf8\xe4\x08\x7d\x8e\x88\x32\x25\xa4\x54\xc3\xa3\x11\x06\x0b\x0e\x09\x24\x1a\xd8\x04\x18\x55\x84\xb8\x5a\xdb\xa7\x81\xdd\x66\x3d\x2d\xbd\x68\x77\x25\x6e\x33\xaa\xe3\x97\xb7\xf0\x04\x45\x2e\x17\xb3\x77\xf2\x03\x45\x5f\xe8\xcd\x04\xf9\x45\xbc\xe7\x68\xfe\x9a\x2a\x78\x36\x70\x1a\xe5\x37\x2a\x84\x73\x8e\xa0\x9e\x46\xf9\xa0\x63\x75\xc9\xc4\x70\x66\xd1\xee\xca\x25\x1e\x54\xc3\x1a\x7a\x67\x41\xe6\x75\x6c\x43\x02\xb7\xb1\xe2\x03\x81\xf8\x78\x22\x0a\x18\xbf\x5b\x1d\x62\xc5\x27\xd3\x7f\xb4\xf1\xac\x60\x8b\x23\xa0\xf9\x5e\x17\x56\x04\x85\x21\x99\x54\x1f\xf8\xb0\xb4\x43\x08\x06\x5f\xc1\xee\x2d\x11\x5a\x25\x89\x0d\x5e\xc7\xed\x9b\x2d\x38\x21\x30\xfb\x7d\xde\x44\x68\x2e\xe6\x22\x67\x75\x5f\xe8\xfa\x03\x59\x9d\xfd\x39\xfb\xf8\x41\xd0\xaf\x29\x98\x48\xfd\xaf\x3f\xf8\x39\x23\x08\x33\xd1\xde\x64\x9d\xfa\x66\x42\x50\x72\x6e\xc8\xca\xa0\x93\x3f\x39\x82\x0a\xfd\x49\xc3\xd3\x47\x85\x32\xc3\x1c\x12\xbd\xfa\x64\x36\xa0\x18\x42\x0a\x94\xb2\xc4\x0e\x50\x10\x02\x74\x40\x65\x89\x21\x20\x67\xb2\x8b\x9d\x8e\x5b\xc9\xde\x61\xe7\xf1\x2f\x72\x3f\x5f\x76\x0c\xaf\xe7\x8c\x6f\xbd\x73\x32\x99\x4f\x27\x6c\x15\xb7\x1f\xca\x21\x7b\xaa\xf2\xed\xd6\xa0\x69\xb2\x8a\xf9\x2a\x23\x76\x33\x5b\x18\xf9\x0f\x5f\x34\xfc\x88\x73\xec\x17\x7f\x32\x5d\x73\xc1\x7c\x5b\x7e\x72\xf3\x16\x12\xef\x74\x3a\x2c\x5e\xb4\x99\x9a\x78\x77\x7d\xdd\xdf\x5d\x2c\x9d\x9a\xca\xa9\x1c\xd6\x6a\xab\x15\x89\x9f\xc2\x52\x7b\xdb\xf1\xdc\x22\x03\x13\x6e\x42\x15\x0c\x2f\x1f\x2b\x1a\x9a\xf0\xd2\x1f\x07\xfb\x2c\xdb\x59\xd9\xc8\x9f\xeb\xb0\x0d\xd4\x97\x48\xdc\xb7\xfd\x94\xb2\xc6\x4d\xee\xfd\x43\xe3\x81\xaa\x0e\xf8\x75\x18\xfc\xfd\x06\x9b\x1d\x56\x21\xa7\x37\xbb\x96\xb1\x02\xd2\x10\x75\x45\x38\x1b\x62\xd5\x9a\xa1\x1c\xbc\x77\x2d\x9f\x05\xf3\xb4\xea\x08\x9f\xc2\x2a\x6d\xd2\xa3\x43\x6d\x40\xfd\xd8\x54\x27\x7e\x73\x93\xc2\xfc\x55\x24\x27\x2c\x72\xd5\x02\x23\x78\x0a\xdc\x17\xea\x00\xfe\xa8\xbb\x13\xab\x6b\x50\xa9\x92\x21\xa4\x62\x6d\xe5\x5b\xe9\x39\x98\x77\xa5\x7e\xa0\xce\x47\x1f\x21\xc9\x5b\xe6\xc3\xc1\x0b\x8a\x6a\xc3\xa7\x01\x78\x96\x1c\x7e\x4e\xf0\x2d\x6c\xad\xc0\xaf\xb7\x9e\xf4\x44\x60\xc4\x00\xa4\xf6\x95\x68\x78\x54\x4a\xa6\x84\x2e\xef\x00\x7a\x5c\x43\x40\x15\x50\x6a\x1a\x01\xd2\x4f\xc4\x39\x1e\x31\x5c\x8f\x54\xe5\x08\x20\x10\x5c\xa4\xf1\x48\xbf\xf0\x09\x42\x86\x55\x0b\xaf\x46\xa1\x51\x56\x96\x7c\x04\xb9\xc7\x8c\x6a\xd0\xc1\x5b\xbe\x26\xe8\xa3\x10\xf8\x6b\xa0\x8f\x94\x11\x98\x02\xee\x8f\x30\xe0\x35\xe4\xdd\x35\x4f\x3f\x8b\xca\x77\x1e\x02\x00\xb6\x3a\x04\x14\x64\x5a\x99\xc6\xa2\x5d\x64\x5c\xbf\xcc\x7a\x52\x32\x3d\x9c\x8f\x67\x3d\xb8\x6c\xcd\xae\xb8\x7f\xa6\x72\xce\x1b\x40\x43\x49\x33\x25\xf1\x91\xe0\x28\x3a\xd9\x97\x7b\xb1\xc6\xd1\xa6\x7f\xb5\x84\x04\xef\x6d\x53\x54\x8a\xa5\xef\x17\xae\x20\x87\xf3\x76\x63\xef\xb2\xfc\xdc\xea\x09\x85\x42\x1b\x1d\xbb\xda\x6d\xe4\xe1\xd6\x2d\xf6\x3e\xb3\x25\x61\xb0\x2d\xf5\xd6\xcd\x8e\xde\xf7\x65\x6d\x8a\x50\xd9\x6e\xc2\x23\x89\x70\x72\xd8\x5c\x55\xae\xf2\xe4\x13\x71\xd2\x92\xb8\xec\x3a\x9f\x2e\x75\xe2\xb2\x65\x3e\x48\xa9\xb6\x07\xf2\x8a\xdc\x8a\x7a\xd1\x2c\x06\x6a\xc3\xa2\xa3\x16\xab\xa6\x38\x39\x8d\xd9\xd9\x38\x8c\x77\x92\xf7\xe9\x47\x81\x63\x4b\x2d\xbb\xa6\xbb\xde\x21\xf6\x40\x4f\x2f\x4e\xa9\x78\xbb\x9e\x2c\x80\xd5\x62\x3b\x1f\x99\xd7\x84\xf1\xec\x69\x3d\xfb\xa0\x01\x45\xa3\x8c\x37\x08\x4c\xc2\xda\xdb\x67\x28\xc9\x1b\x0d\x15\xdc\x07\x3f\x63\xe4\xad\xf8\x79\xda\x29\x29\x69\x8e\x86\x74\xf4\x62\xe7\x72\xd7\xcf\xed\x80\x6f\xa3\xd4\x5b\xab\x62\xad\xc3\xac\x13\x1c\xfb\x80\x8b\xbb\xe1\xe5\xfd\xa0\x5c\x16\xa0\x1b\x72\x68\x17\x37\x6e\x75\x6a\xde\x78\x6f\x8b\x7b\x28\x9b\xdc\xb7\x26\xb1\xad\xfc\xbf\x3a\xff\x09\x00\x00\xff\xff\x61\x01\x51\x0d\x57\x05\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/DigiCertGlobalRootCA.crt.pem"].(os.FileInfo),
		fs["/DigiCertHighAssuranceEVRootCA.crt.pem"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
