// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package scheme generated by go-bindata.// sources:
// out/mysql/scheme/dds_area_info.sql
// out/mysql/scheme/dds_dictionary_info.sql
// out/mysql/scheme/sso_data_permission.sql
// out/mysql/scheme/sso_operate_log.sql
// out/mysql/scheme/sso_role_datapermission.sql
// out/mysql/scheme/sso_role_info.sql
// out/mysql/scheme/sso_role_menu.sql
// out/mysql/scheme/sso_system_info.sql
// out/mysql/scheme/sso_system_menu.sql
// out/mysql/scheme/sso_user_info.sql
// out/mysql/scheme/sso_user_role.sql
package scheme

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _outMysqlSchemeDds_area_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd2\xcf\x6b\xd4\x40\x14\x07\xf0\x73\xf6\xaf\x78\xb7\x24\x20\xb8\xbb\x5e\x4a\x65\x0f\xe9\x76\xd4\xc5\xee\x2a\x31\x1e\x7a\x1a\xc6\x64\xd2\x06\x27\x33\x4b\x32\x11\x7a\xdb\x1e\xa4\xc5\xae\xb6\xc8\xb2\x27\xa1\xf4\xe2\x06\x44\xad\x97\x2a\xa1\xe0\x3f\xe3\x64\xf6\xcf\x90\x10\xb4\x29\x04\x3d\xce\x8f\xcf\x7c\xdf\xf0\x5e\xc7\x18\xba\xc8\xf1\x10\x78\xce\xd6\x0e\x02\x08\x82\x14\x93\x84\x12\x1c\xf1\x50\x80\xd5\x31\x0c\x9f\x70\x29\x38\xf6\x45\x40\xe1\x15\x49\xfc\x7d\x92\x58\xf7\xfa\x36\x00\x17\x12\x78\xc6\x18\x80\x2f\xe2\x98\x72\x09\xa6\x9a\x17\xea\xfc\x5c\x5f\x2f\xd5\xe9\x77\x13\xee\x54\x7a\x3f\xe2\x34\xa5\x98\x93\xf8\x86\xf7\xfa\x1b\xed\xfe\xd7\x8f\xcf\xe5\xf2\x48\x9d\xbd\xd5\xab\xcb\xda\x4f\x49\x42\xb9\x6c\x49\x6f\x28\x7d\x7c\xa5\x8b\x55\x33\x75\x2f\x21\x01\x05\x19\xf1\x83\x88\x4b\xab\xd7\x1e\xb6\xbe\x98\x97\x8b\x9f\xba\x58\xa9\xe3\x4f\x35\x0b\x33\xc6\x70\x3a\xa5\x8c\xfd\xcd\xea\x77\x6f\x67\xad\x4f\xbe\x95\xcb\xa3\xbb\xea\x75\x5e\x9e\x5c\xd7\x2a\x8d\xe2\x29\xa3\xff\x71\xfa\xcb\xec\x06\x88\x44\x62\x2e\xe0\x45\xb4\x57\x95\x57\x5d\x0d\x68\x48\x32\x26\xa1\xdb\x56\x68\xf9\xee\xbd\x2a\x4e\xd5\xec\x0f\x97\x44\x66\x69\xf3\x77\xff\xd4\xfa\xcd\x55\x39\x3b\x84\xde\x26\xe8\x8f\x87\x7a\x91\x43\x77\x13\xd4\xd9\x57\xbd\xc8\xeb\xe7\x9e\xba\xa3\xb1\xe3\xee\xc2\x63\xb4\x0b\x56\xa3\xd9\x76\x75\x58\x6d\xbe\xa4\x07\xf8\xd6\x58\xe0\x66\x53\xac\xc6\xc2\xee\x18\x86\x0d\x68\xf2\x70\x34\x41\x83\x11\xe7\x62\x7b\x0b\x60\x1b\x3d\x70\x9e\xef\x78\x30\x7c\xe4\xb8\xcf\x90\x37\xc8\x64\xb8\x01\xc3\x27\xe3\x31\x9a\x78\x03\x53\x7d\xb8\x54\xf3\x62\x7d\x91\x9b\xf7\x3b\xbf\x03\x00\x00\xff\xff\x9c\x77\x2c\xf0\x8c\x02\x00\x00")

func outMysqlSchemeDds_area_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeDds_area_infoSql,
		"out/mysql/scheme/dds_area_info.sql",
	)
}

func outMysqlSchemeDds_area_infoSql() (*asset, error) {
	bytes, err := outMysqlSchemeDds_area_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/dds_area_info.sql", size: 652, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeDds_dictionary_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\xb1\x4a\xc3\x40\x1c\xc7\xf1\x39\x79\x8a\xdf\x96\x14\x1c\xda\x2a\x22\x95\x0e\x69\x7a\x6a\xb0\x4d\x25\xa6\x43\xa7\x10\x7b\xad\x1e\xa4\x77\x92\x5c\x0a\xdd\xda\xc1\x45\x51\x8b\x14\xdc\x04\x17\xed\x22\xea\x22\xa2\xbe\x4e\x1a\x1f\x43\x4a\xc0\x76\xb0\xb8\x7f\x3f\xf7\xe3\xfe\xaa\x62\x3a\xc4\x70\x09\x5c\xa3\x52\x23\x00\xa5\x91\x47\x59\x5b\x32\xc1\xfd\x70\xe0\x31\xde\x15\xd0\x55\x45\x61\x14\x47\xec\x98\x71\xa9\x17\xf3\x39\x80\x0b\x09\x1e\x07\x01\x8c\xa6\xdb\xf0\x2c\xdb\x74\x48\x9d\xd8\x2e\xda\xa2\xd7\xeb\x70\x09\x8d\x51\x0d\x6b\xaa\xa2\x70\xbf\xd7\x41\xdf\x0f\xdb\x27\x7e\xa8\x6f\x6e\x2c\xd3\x45\x9c\x8c\x2f\xd3\xc7\x97\x0c\xf4\xfd\x20\x5e\x88\xf5\xe2\x0a\x31\xfc\xca\x72\x39\x38\xfd\xbf\x4e\x5f\x3f\x93\xbb\x8b\x0c\x44\x22\x94\x1e\x17\xcb\xbf\xa1\x9d\xae\x1f\x07\x12\xf9\xbf\xec\xec\xea\x26\xf9\xb8\xfe\xdd\x8b\xa4\x2f\xe3\x08\x92\xf1\xc1\x9c\x17\x56\x0c\x9e\xbf\xcd\x86\x23\x14\x4a\x48\x1f\x46\xe9\x64\x8a\x7c\x09\xc9\xf8\x39\x9d\x4c\xb3\x57\x0e\x1c\xab\x6e\x38\x2d\xec\x93\x16\x74\x46\x73\x39\x10\x7b\xd7\xb2\x49\xd9\xe2\x5c\x54\x2b\x40\x95\xec\x18\xcd\x9a\x0b\x73\xcf\x70\x0e\x89\x5b\x8e\x65\x77\x0b\x66\xa3\x3e\x3f\x72\x59\x4b\x9e\x6e\x93\xb3\xf7\xef\xfb\xa9\xb6\xad\xfe\x04\x00\x00\xff\xff\x70\xb6\x2c\x96\xc1\x01\x00\x00")

func outMysqlSchemeDds_dictionary_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeDds_dictionary_infoSql,
		"out/mysql/scheme/dds_dictionary_info.sql",
	)
}

func outMysqlSchemeDds_dictionary_infoSql() (*asset, error) {
	bytes, err := outMysqlSchemeDds_dictionary_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/dds_dictionary_info.sql", size: 449, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_data_permissionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x4f\xd4\x40\x14\xc7\xcf\xed\xa7\x78\xb7\x6d\x13\x0e\xcb\x8a\x84\x60\xf6\x50\x96\x51\x1b\xd9\xa2\xb5\x3d\x70\x9a\x0c\xdb\x41\x47\xdb\xe9\xda\x99\x1a\xb9\x49\x34\x6a\x48\xaa\x4d\x44\x23\x09\x46\xb9\x48\x4f\x2e\x89\x46\x0f\xc8\xb7\xa1\xed\xfa\x2d\x4c\x19\x40\x0f\xbb\x86\xeb\x9b\xf7\xfb\xcf\xcb\xef\x3d\x5d\xeb\xb9\xc8\xf2\x10\x78\xd6\xd2\x0a\x02\x10\x22\xc6\x01\x91\x04\x0f\x69\x12\x31\x21\x58\xcc\xc1\xd0\x35\x8d\x05\xb0\xce\xee\x31\x2e\x8d\x4e\xdb\x04\xe0\xb1\x04\x9e\x86\x21\x58\xbe\xb7\x8a\x6d\xa7\xe7\xa2\x3e\x72\x3c\x18\xc4\x51\x44\xb9\x84\x16\x0b\x5a\x30\xa3\x6b\x9a\xd8\x14\x78\x1a\xfb\xb7\xbb\xfe\x76\x54\x1f\x7d\xaa\x7f\xbd\x2f\xdf\xfc\x54\x1c\x0b\x9a\x87\xc7\x24\x19\xdc\x27\x89\x71\xa5\xf3\x3f\xae\xfa\xfc\x72\x3c\x7a\xa1\x38\x4e\x22\x7a\x81\xcd\x76\x16\x26\x73\x65\x9e\xd5\x07\x87\x8a\x90\x64\x3d\xa4\xf8\x72\xdc\x78\xbf\x28\xf3\x4c\x71\xf1\x90\x26\x44\x52\x4c\x06\xb2\x91\x74\xce\xce\xcf\x4d\x46\xab\xb7\xd9\xc9\xf1\x5e\xb9\x5d\x9c\x1c\xef\xa9\x80\x24\x0d\xa9\x00\x49\x9f\x48\x80\x7f\xff\x38\x78\x5e\xbe\xda\x7d\x20\x62\x7e\xd6\x46\x23\x92\x3c\xbc\xc8\xef\x5c\x9d\x9f\x32\xdb\xe8\x7b\xf5\xe1\xf5\x99\x75\x49\x64\x2a\x40\x32\xbe\xd9\x68\x9f\x33\x21\xa0\x1b\x24\x0d\x25\xb4\x27\x6a\xdc\xfe\x51\x3d\xdd\x82\xf6\x22\x94\xf9\xa8\xde\x29\x60\x76\x11\xea\x2f\x5b\xf5\x4e\xa1\xe2\x6e\xbb\x76\xdf\x72\xd7\xe0\x16\x5a\x03\x83\x05\x66\x53\xf3\x1d\xfb\x8e\x8f\x4e\x4b\x29\x7f\x84\xcf\xef\x66\x98\xd0\x08\x9f\x2e\x8f\x83\xa1\x96\x3f\xd3\xb8\x35\x75\x4d\x33\x01\x39\x37\x6c\x07\x75\x6d\xce\xe3\xe5\x25\x80\x65\x74\xdd\xf2\x57\x3c\xe8\xdd\xb4\xdc\xbb\xc8\xeb\xa6\x72\x63\x01\x7a\xab\xfd\xe6\x94\xba\xad\xea\xdd\x61\x95\x7d\xad\x3e\x3e\xfb\xbd\x9b\x2b\x29\xe3\xfd\xa2\x75\x4d\xff\x13\x00\x00\xff\xff\x67\x2d\xd4\xd7\xb3\x02\x00\x00")

func outMysqlSchemeSso_data_permissionSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_data_permissionSql,
		"out/mysql/scheme/sso_data_permission.sql",
	)
}

func outMysqlSchemeSso_data_permissionSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_data_permissionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_data_permission.sql", size: 691, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_operate_logSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcf\x8a\xd3\x50\x14\xc6\xd7\xc9\x53\x9c\x5d\x13\x18\x86\xce\xe8\x80\x20\x5d\x64\x32\x57\x0d\x4e\x33\x12\x33\x8b\xae\x42\x4c\x6e\x6b\x30\xb9\x29\xc9\x8d\x90\x07\x50\x2a\x52\x0a\x6a\x0d\x62\xd1\x95\xe8\xaa\xf8\x0f\x8b\x8d\xc5\x97\xe9\x4d\xea\x5b\xc8\x6d\x52\x6c\xc5\x82\xbb\xc3\xef\xdc\xef\xdc\xef\x7c\x1c\x51\x50\x0d\xa4\x98\x08\x4c\xe5\xf4\x1c\x01\xc4\x71\x68\x85\x7d\x1c\xd9\x14\x5b\x7e\xd8\x03\x49\x14\x04\xcf\x85\x7b\x5e\xcf\x23\x54\x3a\x6e\xca\x00\x24\xa4\x40\x12\xdf\x07\xe5\xd2\xbc\xb0\x34\x5d\x35\x50\x1b\xe9\x26\x38\x61\x10\x60\x42\xa1\xe1\xb9\x0d\x38\x10\x05\x81\xa6\x7d\x0c\xd4\x23\xe9\x5a\xba\xad\xfc\xf3\xb6\xfc\x94\xb3\x37\x4f\x0f\xe0\xa8\x79\x08\xe5\xab\x9c\x2d\xc6\xc5\xf3\xe1\x72\x31\x81\x63\x0e\xbe\xe4\x65\xfe\xb6\x18\x7f\x2c\x86\xd3\x1a\xc3\x95\xe6\x21\xac\xde\x3f\x5b\x3d\xf9\xbc\xcb\xaf\x72\x3e\x9a\xb0\xe1\x78\x97\x9f\xf0\x39\x2f\x3e\x14\x83\xd9\x36\xaf\x0c\xc6\x69\x6c\xed\x5b\x6e\xcb\xe2\xda\x45\xf9\xe3\x25\x1b\xcd\x2a\x5d\x12\xe3\xe8\x3f\x84\xd5\x57\xcb\xf9\x7c\x13\x88\x13\x61\x9e\x2b\xf5\x02\x0c\xae\x4d\x71\x55\xe0\xae\x9d\xf8\x14\x9c\x24\x8a\x30\xa1\xeb\x6e\x4c\xed\xa0\xff\xaf\x91\x6c\xf0\x9a\xe5\xf3\x22\xfb\xf6\x2b\xfb\x5a\xcf\x0c\x09\xe5\xad\x87\x76\xe4\xdc\xb7\x23\xe9\xe4\x68\x4f\xd2\xec\xf1\x23\x36\xfd\xae\x55\xaa\x3b\x86\xd6\x56\x8c\x0e\xdc\x46\x1d\x90\x3c\x57\xe6\x8c\xd7\x0f\x70\x6a\xfd\x75\x02\xd6\x66\x5b\xa9\x2e\x64\x51\x10\x64\x40\xfa\x4d\x4d\x47\x2d\x8d\x90\xf0\xec\x14\xe0\x0c\xdd\x50\x2e\xcf\x4d\x50\x6f\x29\xc6\x5d\x64\xb6\x12\xda\xbd\x06\xea\x45\x9b\x5f\x46\xab\x4e\xa2\xc8\xde\xb1\x9f\x59\xe3\xba\xf8\x3b\x00\x00\xff\xff\x8f\x5b\x64\x4f\x75\x02\x00\x00")

func outMysqlSchemeSso_operate_logSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_operate_logSql,
		"out/mysql/scheme/sso_operate_log.sql",
	)
}

func outMysqlSchemeSso_operate_logSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_operate_logSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_operate_log.sql", size: 629, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_role_datapermissionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x4f\x8b\xd3\x40\x18\xc6\xcf\xc9\xa7\x78\x6f\x4d\x60\x0f\x5d\x4f\xb2\xd2\x43\x36\x3b\x6a\x70\x9b\x4a\x4c\x0f\x3d\x0d\x31\x99\x96\x81\x64\x52\x32\x13\xb0\x37\x83\x0a\xb5\x07\x05\x29\x16\xf1\x50\x8a\xa0\xbd\x88\x60\xa5\x16\xab\xf8\x65\x9a\x3f\x7e\x0b\x49\x13\x68\x0f\x15\xf7\xf6\x32\xc3\xef\x37\xef\x3c\x8f\x2c\xe9\x16\xd2\x6c\x04\xb6\x76\x79\x8d\x00\x38\x0f\x71\x14\xfa\x04\x7b\x8e\x70\x86\x24\x0a\x28\xe7\x34\x64\xa0\xc8\x92\x44\x3d\x78\x4c\x07\x94\x09\xe5\x56\x53\x05\x60\xa1\x00\x16\xfb\x3e\x68\x5d\xbb\x83\x0d\x53\xb7\x50\x1b\x99\x36\xb8\x61\x10\x10\x26\xa0\x91\x4e\xe6\xc5\xb3\x5f\xf9\xcf\xb7\xe9\xeb\xef\x0d\x38\x93\x25\x89\x8f\x38\xfe\x97\xe5\xc0\xe5\xab\x6d\xbe\x9d\x1f\x73\xfb\x8d\xfe\x0f\x16\x9f\xde\x14\x2f\xbf\x1e\x83\x87\x1f\x60\x37\x64\x7d\x3a\xb8\x91\xe5\x79\x3a\x7e\x47\xbd\x7a\x65\xe1\x88\x98\x83\xa0\x6c\x54\x42\xe7\x2a\x78\xa4\xef\xc4\xbe\x80\xe6\xc9\xdd\x27\xeb\xec\x69\x02\xe7\x17\x90\x7f\x4c\xf2\xe9\x12\x9a\x17\x90\x7d\xfe\x90\x6e\x36\x95\xce\x8d\x88\x23\x08\x16\x34\x20\xe0\x39\x82\x54\x43\x6d\x74\xe3\x28\x22\x4c\xec\x6f\xb9\x70\x82\xe1\xa9\x17\xd2\xf1\xfb\x74\xfb\x23\x9b\xad\xff\xcc\xbe\x55\xce\x87\x96\xd1\xd6\xac\x1e\x3c\x40\x3d\x50\xa8\xa7\x96\x67\xe5\x4c\xbd\x27\xb8\xec\x73\x5f\x65\x44\x02\xcc\x47\xbc\x4c\x12\x94\xaa\x87\xb3\x3a\x56\x55\x96\x24\x15\x90\x79\xcf\x30\x51\xcb\x60\x2c\xbc\xba\x04\xb8\x42\x77\xb5\xee\xb5\x0d\xfa\x7d\xcd\x7a\x84\xec\x56\x2c\xfa\xb7\x41\xef\xb4\xcb\x8a\x5b\x75\xd2\xbb\xcd\xab\x2a\xac\xf4\xc5\xaa\x48\xa6\xbb\xdf\x8b\x2c\xf9\x52\x2c\x96\x8d\x3b\xf2\xdf\x00\x00\x00\xff\xff\x92\x79\x91\xba\x58\x02\x00\x00")

func outMysqlSchemeSso_role_datapermissionSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_role_datapermissionSql,
		"out/mysql/scheme/sso_role_datapermission.sql",
	)
}

func outMysqlSchemeSso_role_datapermissionSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_role_datapermissionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_role_datapermission.sql", size: 600, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_role_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x4b\xe3\x40\x00\x85\xcf\xc9\xaf\x78\xb7\x26\xb0\x87\x76\x59\x96\xa5\x4b\x0f\x69\x3a\xbb\x1b\xb6\x4d\x35\x26\x87\x9e\xc2\x98\x4c\x35\x90\x4c\x34\x99\x08\xde\xec\x4d\x05\x41\x44\xe8\xc1\x93\x08\xb6\x17\x41\x44\x4a\xd1\xdf\x93\xa6\xfe\x0b\x49\xd3\x83\x4a\x2f\xc3\xcc\x3c\xde\xf7\xe0\x93\x25\xdd\x22\x9a\x4d\x60\x6b\xed\x2e\x01\xd2\x34\x76\x93\x38\x64\x6e\xc0\x87\x31\x14\x59\x92\xaa\x97\x8f\xdd\x60\x2f\xe0\x42\xf9\x5e\x57\x01\x1e\x0b\xf0\x2c\x0c\xa1\x39\x76\xdf\x35\x4c\xdd\x22\x3d\x62\xda\xf0\xe2\x28\x62\x5c\xa0\xb6\x9c\x5c\x2d\xcf\x9e\x02\xbf\x86\x6f\xb2\x24\x71\x1a\x31\x1c\xd1\xc4\xdb\xa7\x89\xf2\xf3\xc7\x47\xc0\xd7\x4a\x7e\x79\x51\x4c\x1e\xab\x5a\x2a\xa8\xc8\x52\x88\x80\x1f\x97\xcb\x0d\x15\x3e\x1b\xd2\x2c\x14\xa8\x6f\x02\x14\xe7\xb3\xc5\xc9\x08\x8d\x26\x8a\xfb\x51\x71\x3d\x45\xbd\x89\xc5\xc3\x5d\x3e\x9f\x57\x38\x2f\x61\x54\x30\x57\x04\x11\x83\x4f\x05\xab\x2e\x6b\xa2\x97\x25\x09\xe3\x62\x95\xa6\x82\x46\x07\x9b\x16\xf2\xd3\x9b\xfc\xf5\x65\x31\x9e\xbd\x8d\x9f\x2b\xe6\x96\x65\xf4\x34\x6b\x80\xff\x64\x00\x65\x6d\x4a\x2d\x03\xc7\x34\xb6\x1d\xb2\xfa\xcf\xf8\xa1\xfb\x49\xab\xbb\xf2\xa1\x94\xa7\x2a\x4b\x92\x0a\x62\xfe\x35\x4c\xd2\x32\x38\x8f\x3b\x6d\xa0\x43\xfe\x68\x4e\xd7\x86\xfe\x4f\xb3\x76\x88\xdd\xca\xc4\xf0\x17\xf4\x7e\xaf\x54\xdc\x5a\x7b\x5a\xde\x4e\x6b\xbf\xe5\xf7\x00\x00\x00\xff\xff\xad\x01\xeb\x69\xbe\x01\x00\x00")

func outMysqlSchemeSso_role_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_role_infoSql,
		"out/mysql/scheme/sso_role_info.sql",
	)
}

func outMysqlSchemeSso_role_infoSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_role_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_role_info.sql", size: 446, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_role_menuSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xc1\x4e\xea\x40\x14\x86\xd7\xed\x53\x9c\x1d\x34\x61\x01\x77\x75\xc3\x0d\x8b\x52\xe6\x6a\x23\x14\xad\xed\x82\x55\x53\xe8\x60\x1a\xdb\xa9\xb6\xd3\x05\x3b\xd9\xa9\x89\x92\x10\x95\x85\x26\xc6\x98\x28\x1b\x37\x6a\x90\x80\xaf\xd3\x16\xdf\xc2\x94\xb6\x31\x26\x18\x75\x37\x39\x33\xdf\x7f\xe6\x9c\xff\x67\x19\x41\x46\xbc\x82\x40\xe1\xab\x75\x04\xe0\x79\x8e\xe6\x3a\x16\xd6\x6c\x4c\x7c\xc8\xb3\x0c\x63\x1a\xd0\x36\x77\x4c\x42\xf3\x7f\x8a\x1c\x00\x71\x28\x10\xdf\xb2\x80\x57\x95\xa6\x26\x4a\x82\x8c\x1a\x48\x52\xa0\xe3\xd8\x36\x26\x14\x72\xa6\x91\x83\x02\xcb\x30\x5e\xcf\xd3\xbe\x62\x3f\x5e\x47\x4f\xf3\x68\x7e\x9d\x31\xcb\xce\x9f\x21\x03\x77\x75\xdf\xa2\x50\x5c\x45\x2f\xee\x87\x8b\xa3\xc7\x8c\x8e\xbf\xfc\x2b\x7a\x70\x15\x9c\x9c\x67\x34\x26\x7a\xdb\xc2\x40\x4d\xd2\x8b\xe9\xd2\x37\x70\x74\x3c\x09\x0f\xfa\x50\x2a\x43\x74\xd7\x8f\xce\xc6\x50\x2c\x43\xf8\x70\x1b\x4c\xa7\x89\x5c\xc7\xc5\x3a\xc5\x1a\x35\x6d\x0c\x86\x4e\x71\x72\x48\x15\x3b\xbe\xeb\x62\x42\x97\xb7\x1e\xd5\xed\xbd\x55\x1d\x82\xc3\xcb\x60\x3e\x0b\x47\x93\xb7\xd1\x73\xba\x52\xc7\xa5\xae\x4e\x76\x7f\x3c\x61\x78\x3a\x0c\x66\x83\xe8\xf5\x22\x18\xbc\x24\x12\x9b\xb2\xd8\xe0\xe5\x16\x6c\xa0\x16\xe4\x4d\x83\x8b\x6b\xaa\x24\x6e\xa9\x68\x59\xf2\xc9\xbe\x96\x45\xc0\xd6\x3c\x37\x49\x41\xba\xd8\x42\x6a\x4f\x21\xb1\x96\x63\x19\x86\x03\x24\xad\x89\x12\xaa\x88\x84\x38\xb5\x2a\x40\x0d\xfd\xe7\xd5\xba\x02\xc2\x3a\x2f\x6f\x23\xa5\xe2\xd3\xee\x5f\x10\x9a\x8d\x38\x23\x95\xd4\xaf\xc5\xcd\x38\xf7\x8f\x7d\x0f\x00\x00\xff\xff\xb7\xf4\xce\x69\x7a\x02\x00\x00")

func outMysqlSchemeSso_role_menuSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_role_menuSql,
		"out/mysql/scheme/sso_role_menu.sql",
	)
}

func outMysqlSchemeSso_role_menuSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_role_menuSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_role_menu.sql", size: 634, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_system_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\x5f\x4b\xdb\x50\x14\x7f\x4e\x3e\xc5\x79\x6b\x03\x0a\x6d\x1d\x43\x1c\x3e\xc4\x9a\x6d\x65\x5a\xb7\x2e\x7d\xf0\x29\x5c\x9b\x5b\x7b\x21\xb9\x71\xc9\xcd\x66\xdf\x2a\xa2\x03\x41\x56\x86\xce\x6d\x4c\xb4\xc2\x9c\x2f\x2e\x0e\xb1\x13\xea\x9f\x2f\xd3\x9b\x98\xa7\x7d\x85\x91\x44\xbb\x28\x32\x5f\x4a\x73\xee\xef\xcf\x39\xe7\xc7\x11\x85\x62\x45\x91\x55\x05\x54\x79\x62\x4a\x01\x70\x1c\x4b\x73\x9a\x0e\xc3\xa6\x46\x68\xdd\x82\xac\x28\x08\x44\x87\x39\x32\x4f\x28\xcb\x16\x72\x12\x00\xb5\x18\x50\xd7\x30\x40\xae\xaa\x33\x5a\xa9\x5c\xac\x28\xd3\x4a\x59\x85\x9a\x65\x9a\x98\x32\xc8\x10\x3d\x03\x43\xa2\x20\x50\x64\x62\x78\x8b\xec\x5a\x03\xd9\xd9\x91\x42\x9a\xfa\x0f\x1c\x1c\xf7\x82\xde\x0e\x6f\xaf\x07\x3f\x8e\x12\x1a\xa1\x3a\x5e\xd4\x5c\xdb\x18\x70\x1f\x3f\x92\x00\x52\x9c\x70\xff\x53\xd8\x39\xe1\xdf\x8e\xf8\x76\x2b\xe1\x60\x8a\xe6\x0c\x0c\x8c\xd0\x66\xd4\x67\x5e\x02\x1d\xd7\x91\x6b\x30\xc8\xdf\x6b\xba\xd6\xf5\x5b\x4b\x90\xff\x73\xf6\x95\xb7\xbd\x60\xe3\x00\x72\x63\x10\xec\x2f\x05\x1b\x07\x89\x9e\x61\xcd\x13\xaa\x31\x62\x62\xcb\x65\x10\x4b\xa6\x34\x47\x72\xb9\xfb\x54\xaf\xba\x2b\xfe\x56\xd7\xdf\xea\x86\x9b\x97\x03\x19\x6b\x30\x45\xbe\x30\x7a\x7b\x8c\xe8\x35\xc1\xb1\x06\x4e\xad\x2a\x06\xde\x78\x65\xe6\xe6\x87\x17\x90\x4d\x4c\x64\x37\x33\x69\x76\xff\xb4\x17\xee\x7d\xf6\x77\x7f\xf3\xb3\x0f\xd7\x6e\xa8\x19\x75\x7b\xbf\x0c\x5a\x58\x18\x6e\x60\xa4\x63\x7b\xb8\x4e\x16\xb1\x0e\x51\x01\x39\x44\xc7\xc9\xf7\x2d\xed\xb0\x73\x12\x6e\xef\xf1\xd3\x65\xfe\xab\x95\x76\x20\x7a\xf4\xfc\x50\xa4\x7c\xc3\xeb\x9f\xb6\xfc\xdd\xf7\x57\xde\x6a\x7a\x9d\xff\x8b\x34\xf8\xd2\xe3\xe7\x9b\xe9\x48\xdf\xe1\x5a\x03\x31\xcd\x61\x88\xb9\xce\x20\xd9\xc2\x03\xc9\xf2\x8b\x9f\xfd\xcb\x0e\x5f\xdb\xb9\x5a\x3e\xbf\x49\x79\x0c\xf8\x59\x8b\xb7\xbd\xa1\x28\x65\xbe\x72\x1c\x6e\x1d\x26\x16\x0e\xae\xd9\xf8\xce\x3c\xe9\x9e\x0e\x2f\x78\x7b\x9d\x7b\xab\xe1\xc7\xef\x09\xe1\x65\xa5\x34\x2d\x57\x66\xe1\x85\x32\x0b\x59\xa2\x4b\x51\xad\x5a\x2e\xbd\xaa\x2a\x71\xc9\xa5\x6f\xb4\x3b\x07\xa4\xc5\x47\x90\x8d\x7e\x63\xb4\x28\x08\x11\x92\xe8\x8b\x37\xc8\x18\xe5\xda\x46\xb2\xdb\x18\x39\x14\xff\x97\x44\x41\x90\x40\x29\x3f\x2b\x95\x95\xf1\x12\xa5\xd6\xe4\x04\xc0\xa4\xf2\x54\xae\x4e\xa9\x50\x7c\x2e\x57\x5e\x2b\xea\xb8\xcb\xea\xa3\x50\x9c\x99\x8e\x2e\x70\xfc\xfa\x98\xfa\x97\x1d\x7f\xc9\xcb\x3c\x11\xff\x06\x00\x00\xff\xff\x48\x53\x23\x97\xdd\x03\x00\x00")

func outMysqlSchemeSso_system_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_system_infoSql,
		"out/mysql/scheme/sso_system_info.sql",
	)
}

func outMysqlSchemeSso_system_infoSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_system_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_system_info.sql", size: 989, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_system_menuSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\xc1\x6b\x13\x41\x14\xc6\xcf\x9b\xbf\xe2\xdd\xb2\x0b\x1e\xd2\xa2\x45\x2a\x39\x6c\xd3\x51\x83\x4d\x2a\xeb\xf6\xd0\xd3\x32\xdd\x4c\xed\xe0\xee\x6c\xd8\x99\x2d\xe4\xd6\xe2\xc1\x5a\x88\x29\x62\xac\xd0\x4a\x15\xa1\xe9\xa5\x1a\x51\x62\x49\xaa\xff\xcd\xec\xc4\xff\x42\xb6\xbb\xd6\x55\x36\x78\x1b\x78\xf3\xfb\xde\xf7\x7d\xbc\x92\x56\xb3\x90\x69\x23\xb0\xcd\xa5\x15\x04\xc0\x79\xe0\xf0\x0e\x17\xc4\x77\x7c\xc2\x22\xd0\x4b\x9a\x46\x5b\xb0\x41\x1f\x53\x26\xf4\xf9\x8a\x01\xc0\x02\x01\x2c\xf2\x3c\x30\xd7\xec\x55\xa7\xde\xac\x59\xa8\x81\x9a\x36\xb8\x81\xef\x13\x26\xa0\x2c\xf7\x4f\xa6\x4f\xbf\xab\xcb\xd7\xb2\xf7\xad\x0c\x37\x4a\x9a\xc6\xb0\x4f\x60\x1b\x87\xee\x16\x0e\xf5\x85\x9b\x79\x91\x7f\x31\x79\xd0\x55\x83\x61\x8a\xb5\x71\x98\x4c\x0a\x97\xff\xe1\xd4\xde\x48\x8d\x07\xf9\x75\xbc\xc3\x9d\x59\xa6\x73\xdc\x97\x89\x9a\x9c\xe4\x39\x8f\x6c\x13\x2f\x21\x05\x65\x9d\x2b\x74\x06\x79\xfe\x5c\x8d\x07\x29\x43\xdd\x80\xfd\x1d\x2d\x1f\xe9\xe8\x47\xfc\xee\xd9\xef\x30\x62\xeb\xfa\xe3\xfc\xad\x85\x19\x25\x1c\x0f\xe5\xdb\x9d\x94\x20\x0c\x6f\x78\xe4\xda\xcc\x9c\x01\x2d\xb2\x89\x23\x4f\x40\xa5\xd0\xd5\xfe\x28\xde\xd9\x85\xb9\x45\x50\xa7\xbb\xea\xd5\x19\x54\x16\x21\x3e\xff\x20\x2f\x2e\x52\x39\x37\x24\x58\x10\x47\x50\x9f\x40\x0b\x0b\x92\x3e\x32\x45\x37\x0a\x93\xae\xaf\xa6\x5c\x60\xbf\x5d\x68\x6e\xef\x48\x4e\xc6\xf1\xe1\xe8\xe7\xe1\xd7\xac\xe9\x20\x14\x21\x66\x4f\xfe\xd7\x75\xfc\xe2\xa5\x1c\xf7\xf2\x5d\x53\xee\x04\x6d\xc2\x8a\xd3\xe5\xc9\x37\x9f\xe4\xc1\xa9\xfc\xdc\x97\x97\x59\x2d\x0f\xad\x7a\xc3\xb4\xd6\xe1\x01\x5a\x07\x9d\xb6\x0c\x03\x50\xf3\x5e\xbd\x89\xaa\x75\xc6\x82\xe5\x25\x80\x65\x74\xd7\x5c\x5b\xb1\xa1\x76\xdf\xb4\x1e\x21\xbb\x1a\x89\xcd\xdb\x50\x5b\x6d\x24\x57\x5a\xcd\xce\x6c\xfa\xfe\x4c\x9f\xf6\x8e\x65\xb7\x1f\xf7\x87\x71\xf7\xa3\x51\xbe\x53\xfa\x15\x00\x00\xff\xff\xb0\x46\x95\xa3\x0c\x03\x00\x00")

func outMysqlSchemeSso_system_menuSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_system_menuSql,
		"out/mysql/scheme/sso_system_menu.sql",
	)
}

func outMysqlSchemeSso_system_menuSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_system_menuSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_system_menu.sql", size: 780, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_user_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\xcf\x6e\xd3\x40\x10\xc6\xcf\xf6\x53\xcc\xcd\x8e\xd4\x43\x12\x50\x55\x15\xf5\xe0\xa6\x06\x22\x9a\x14\x82\x73\xe8\xc9\xda\xc6\x9b\xd4\x92\xbd\x0e\xf6\x9a\x84\x5b\x83\x04\x2a\x48\xd0\x1c\x12\x55\xfc\xa9\x44\x84\x40\x41\x82\x14\x24\x68\x4b\x2b\xe8\xcb\x74\x9d\xf6\x2d\x50\xbc\x4e\xec\x54\x69\xb9\x8d\x3d\xf3\x9b\x9d\xef\xdb\x1d\x51\xc8\x95\x54\x45\x53\x41\x53\x96\x57\x55\x00\xcf\x73\x74\xdf\xc3\xae\x6e\x92\xaa\x03\xb2\x28\x08\xfc\xcb\x80\x0d\xb3\x66\x12\x2a\x67\xd3\x29\x00\xe2\x50\x20\xbe\x65\x81\x52\xd6\xd6\xf4\x7c\x31\x57\x52\x0b\x6a\x51\x83\x8a\x63\xdb\x98\x50\x90\x4c\x43\x82\x39\x51\x10\xaa\xbe\x65\xe9\x04\xd9\x18\x1e\x23\xb7\xb2\x89\x5c\xf9\x46\x36\xc9\xc7\xc4\xb0\xd3\x0f\xb6\x0f\xd9\xb3\x3e\x6b\xbf\xe2\x6c\x78\xf0\x14\x3b\x7f\xf3\x5a\x76\x0c\xd6\x91\xe7\x35\x1c\xd7\xf8\xef\x99\x6c\xff\xf9\xf0\x43\x8b\x43\xd8\x46\xa6\x35\x4d\x24\x2a\xc3\x2c\x2f\xf4\x28\xa2\xbe\x07\xd4\x24\x4f\x46\x7e\x64\x52\x60\xe0\x2a\xf2\x2d\x0a\x99\x99\xb3\xbd\x3c\x08\xb6\x5a\x90\x5e\x84\xe0\xdb\x47\x76\x74\x04\x99\x45\xb8\xe8\xb4\xd8\xe0\x2d\x64\x17\x61\xf8\xb9\x35\xec\xf4\x79\x63\xdb\xd9\x30\xad\x58\x6c\xe6\x4a\xa3\x7e\x9d\xef\xef\xb1\x9d\xc3\xc9\xe8\x8d\xa6\xee\xd4\x31\x31\x8d\x69\xa3\x92\x42\xff\x0e\xce\x4e\x7b\xbc\x88\x33\x15\x17\x23\x8a\x75\x6a\xda\x18\x0c\x44\x31\x0f\x22\x21\x15\xdf\x75\x31\xa1\x61\xd6\xa3\xc8\xae\xcf\x34\x6f\xfb\x1d\x3b\x39\x0e\x76\x0f\x2e\x76\x7f\x46\x3d\x37\x11\xa9\xe1\x7a\xc3\xe0\x60\xf2\xc5\x8c\x3b\xa7\xaf\xbe\x86\xb3\xd3\x41\xd0\xf9\x1d\x7c\xed\x05\xdd\xef\xd1\x95\x34\xa9\x5e\x47\x2e\xb2\xbd\xd8\x94\x74\xf6\x92\xb4\xe0\xc5\x17\xf6\xa3\xcb\x76\x9e\x4e\x30\x0b\x79\x54\xb7\x9c\x9a\x49\x2e\xc9\x9b\xc2\xde\x6f\xb1\xf6\xeb\xe1\x9b\x13\xf6\xa7\x9b\xd4\xe0\x39\xbe\x5b\xc1\x7a\xc2\xcb\x4c\x76\x21\x9e\x5f\x4a\x4b\xb3\x14\x04\x7b\x9f\x82\xe3\xf6\xd8\x5b\xde\x23\x7e\x4b\xf3\x09\x5e\x92\xae\x69\xc0\xf1\xfb\xa5\x7c\x41\x29\xad\xc3\x3d\x75\x1d\xe4\x68\xfb\x52\xa3\x44\xb9\x98\x7f\x50\x56\xc3\xff\x3e\x79\xa4\x4f\xad\xaa\x1e\x6f\x8b\x3c\x09\x43\x4a\x14\x84\x11\x61\x1a\xcd\x09\x11\x02\xd1\x94\xf2\x44\xf1\x1c\x8f\x52\xa2\x20\xa4\x40\x2d\xde\xc9\x17\xd5\xa5\x3c\x21\xce\xca\x32\xc0\x8a\x7a\x5b\x29\xaf\x6a\x90\xbb\xab\x94\x1e\xaa\xda\x92\x4f\xab\x0b\x90\x5b\x2b\x8c\xb6\x7e\x29\xda\xbf\xf3\x5e\x5f\xba\x25\xfe\x0b\x00\x00\xff\xff\x46\xb4\xfa\x89\x51\x04\x00\x00")

func outMysqlSchemeSso_user_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_user_infoSql,
		"out/mysql/scheme/sso_user_info.sql",
	)
}

func outMysqlSchemeSso_user_infoSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_user_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_user_info.sql", size: 1105, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSchemeSso_user_roleSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xb1\x4e\xf2\x50\x1c\xc5\xe7\xdb\xa7\x38\x1b\x6d\xc2\x00\xdf\xf4\x05\xc3\x50\xca\x55\x1b\xa1\x68\x6d\x07\xa6\x06\x6c\x31\x4d\xca\x6d\xa4\xbd\x03\x9b\x0c\x26\xea\xa2\x31\x8d\xab\x71\xd1\x6e\x26\x92\xc8\x80\xaf\x03\xb7\x8f\x61\x4a\x6b\x64\xd0\x84\xf5\xe4\xfe\xce\x3d\xe7\xfc\x25\xa2\x99\x54\xb5\x28\x2c\xb5\xd5\xa1\x40\x14\x85\x0e\x8f\xbc\x89\x33\x09\x03\x0f\xb2\x44\x88\xef\x62\xe8\x9f\xfb\x2c\x96\xff\xd5\x14\x80\x85\x31\x18\x0f\x02\xa8\xb6\xd5\x73\x74\x43\x33\x69\x97\x1a\x16\xce\xc2\xf1\xd8\x63\x31\x2a\xbe\x5b\x41\x55\x22\x64\xe3\xf2\x17\xfc\xf3\x5c\x24\xe9\xfa\x7a\x21\x3e\x1f\x57\x77\x8b\x02\x8c\xa6\xd1\x2e\xdc\x7c\x29\x96\x4f\xdb\x5c\x9e\x78\x07\x30\x7b\x7d\xc8\x6e\xde\xb7\x41\x8f\x0d\x86\x81\x87\xd8\x67\xd3\x1c\xac\x2b\x70\xbd\xd1\x80\x07\x31\xea\xbf\xfe\x7c\xfb\xb1\xbe\x9c\xa1\xd6\xc0\xea\xfe\x4d\x24\x29\xea\x0d\x88\x97\x99\x48\xd2\xc2\xee\xd8\xd4\xbb\xaa\xd9\xc7\x11\xed\x43\xf6\x5d\x25\xd7\x6c\x43\x3f\xb1\xe9\x46\xe2\xec\xc2\xf9\x5e\x79\x13\x99\x47\x13\xc8\xe5\x5a\xd5\xa2\x7c\xb5\xec\xa2\x48\x84\x28\xa0\xc6\x81\x6e\xd0\xa6\xce\x58\xd8\x6e\x01\x6d\xba\xaf\xda\x1d\x0b\xda\xa1\x6a\x9e\x52\xab\xc9\xe3\xd1\x7f\x68\xbd\x6e\x7e\x86\x66\xb9\x67\x51\x72\x75\x35\xcf\x66\x49\xf6\x9c\x56\xf6\xa4\xaf\x00\x00\x00\xff\xff\xb9\x6e\x8e\x16\xe9\x01\x00\x00")

func outMysqlSchemeSso_user_roleSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSchemeSso_user_roleSql,
		"out/mysql/scheme/sso_user_role.sql",
	)
}

func outMysqlSchemeSso_user_roleSql() (*asset, error) {
	bytes, err := outMysqlSchemeSso_user_roleSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/scheme/sso_user_role.sql", size: 489, mode: os.FileMode(509), modTime: time.Unix(1609321853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"out/mysql/scheme/dds_area_info.sql":           outMysqlSchemeDds_area_infoSql,
	"out/mysql/scheme/dds_dictionary_info.sql":     outMysqlSchemeDds_dictionary_infoSql,
	"out/mysql/scheme/sso_data_permission.sql":     outMysqlSchemeSso_data_permissionSql,
	"out/mysql/scheme/sso_operate_log.sql":         outMysqlSchemeSso_operate_logSql,
	"out/mysql/scheme/sso_role_datapermission.sql": outMysqlSchemeSso_role_datapermissionSql,
	"out/mysql/scheme/sso_role_info.sql":           outMysqlSchemeSso_role_infoSql,
	"out/mysql/scheme/sso_role_menu.sql":           outMysqlSchemeSso_role_menuSql,
	"out/mysql/scheme/sso_system_info.sql":         outMysqlSchemeSso_system_infoSql,
	"out/mysql/scheme/sso_system_menu.sql":         outMysqlSchemeSso_system_menuSql,
	"out/mysql/scheme/sso_user_info.sql":           outMysqlSchemeSso_user_infoSql,
	"out/mysql/scheme/sso_user_role.sql":           outMysqlSchemeSso_user_roleSql,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"out": &bintree{nil, map[string]*bintree{
		"mysql": &bintree{nil, map[string]*bintree{
			"scheme": &bintree{nil, map[string]*bintree{
				"dds_area_info.sql":           &bintree{outMysqlSchemeDds_area_infoSql, map[string]*bintree{}},
				"dds_dictionary_info.sql":     &bintree{outMysqlSchemeDds_dictionary_infoSql, map[string]*bintree{}},
				"sso_data_permission.sql":     &bintree{outMysqlSchemeSso_data_permissionSql, map[string]*bintree{}},
				"sso_operate_log.sql":         &bintree{outMysqlSchemeSso_operate_logSql, map[string]*bintree{}},
				"sso_role_datapermission.sql": &bintree{outMysqlSchemeSso_role_datapermissionSql, map[string]*bintree{}},
				"sso_role_info.sql":           &bintree{outMysqlSchemeSso_role_infoSql, map[string]*bintree{}},
				"sso_role_menu.sql":           &bintree{outMysqlSchemeSso_role_menuSql, map[string]*bintree{}},
				"sso_system_info.sql":         &bintree{outMysqlSchemeSso_system_infoSql, map[string]*bintree{}},
				"sso_system_menu.sql":         &bintree{outMysqlSchemeSso_system_menuSql, map[string]*bintree{}},
				"sso_user_info.sql":           &bintree{outMysqlSchemeSso_user_infoSql, map[string]*bintree{}},
				"sso_user_role.sql":           &bintree{outMysqlSchemeSso_user_roleSql, map[string]*bintree{}},
			}},
		}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}