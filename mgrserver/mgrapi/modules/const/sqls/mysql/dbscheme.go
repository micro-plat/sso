// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package mysql generated by go-bindata.// sources:
// out/mysql/dds_area_info.sql
// out/mysql/dds_dictionary_info.sql
// out/mysql/sso_data_permission.sql
// out/mysql/sso_operate_log.sql
// out/mysql/sso_role_datapermission.sql
// out/mysql/sso_role_info.sql
// out/mysql/sso_role_menu.sql
// out/mysql/sso_system_info.sql
// out/mysql/sso_system_menu.sql
// out/mysql/sso_user_info.sql
// out/mysql/sso_user_role.sql
package mysql

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

var _outMysqlDds_area_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd1\xc1\x8a\xd3\x40\x1c\xc7\xf1\x73\xfb\x14\xff\x5b\x52\x10\x6c\xeb\x65\x59\xe9\x21\xdb\x1d\xb5\xb8\xad\x12\xe3\x61\x4f\x61\x4c\x26\xbb\x81\xc9\x4c\x49\x26\x82\xb7\xee\x41\x5a\x6c\xb5\x45\x4a\x4f\x42\xe9\xc5\x06\x44\xad\x97\x2a\xa1\xe0\xcb\x38\x99\x3e\x86\x94\xa0\x8d\x10\xdc\xfb\x7c\xe6\xfb\x1b\xa6\x5a\x69\x9b\xc8\xb0\x10\x58\xc6\xd9\x05\x02\x70\xdd\xc8\xc6\x21\xc1\xb6\xcf\x3c\x0e\x7a\xb5\x52\x71\x30\x13\x9c\xd9\x0e\x77\x09\xbc\xc4\xa1\x73\x8d\x43\xfd\x5e\xb3\x06\xc0\xb8\x00\x16\x53\x0a\xe0\xf0\x20\x20\x4c\x80\x26\x27\xa9\x5c\x2e\xd5\x6e\x21\xa7\xdf\x35\xb8\x73\xd0\xd7\x3e\x23\x11\xb1\x19\x0e\x8e\xbc\xd1\x3c\x29\xf7\xbf\x7e\x7c\xce\x16\x43\x39\x7b\xab\xd6\x9b\xdc\xf7\x71\x48\x98\x28\xa9\x17\x94\x1a\x6d\x55\xba\x2e\x56\xaf\x42\xec\x12\x10\x3e\x7b\xe5\x33\xa1\x37\xca\x63\xfb\xd5\x24\x9b\xff\x54\xe9\x5a\x8e\x3e\xe5\xcc\x8b\x29\xb5\xa3\x3e\xa1\xf4\x6f\xab\x59\xff\xb7\xb5\x1f\x7f\xcb\x16\xc3\xbb\xf2\x75\x92\x8d\x77\xb9\x8a\xfc\xa0\x4f\xc9\x2d\x4e\x7d\x19\x1c\x01\x0f\x85\xcd\x38\xbc\xf0\xaf\x0e\xf3\x0e\x47\x5d\xe2\xe1\x98\x0a\xa8\x97\x0d\xcd\xde\xbd\x97\xe9\x54\x0e\xfe\x70\x81\x45\x1c\x15\x5f\xf7\x5f\xad\xde\x6c\xb3\xc1\x0d\x34\x4e\x41\x7d\xbc\x51\xf3\x04\xea\xa7\x20\x67\x5f\xd5\x3c\xc9\xaf\x7b\x6a\x76\xba\x86\x79\x09\x8f\xd1\x25\xe8\x85\xcf\xae\xd5\x00\xf5\x1e\x76\x7a\xa8\xd5\x61\x8c\x9f\x9f\x01\x9c\xa3\x07\xc6\xf3\x0b\x0b\xda\x8f\x0c\xf3\x19\xb2\x5a\xb1\xf0\x4e\xa0\xfd\xa4\xdb\x45\x3d\xab\xa5\xc9\x0f\x1b\x39\x49\xf7\xab\x44\xbb\x5f\xfd\x1d\x00\x00\xff\xff\xe2\x97\x59\x34\x56\x02\x00\x00")

func outMysqlDds_area_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlDds_area_infoSql,
		"out/mysql/dds_area_info.sql",
	)
}

func outMysqlDds_area_infoSql() (*asset, error) {
	bytes, err := outMysqlDds_area_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/dds_area_info.sql", size: 598, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlDds_dictionary_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xce\x31\x4b\xc3\x40\x18\xc6\xf1\x39\xf9\x14\xcf\x96\x14\x1c\xda\x2a\x22\x95\x0e\x69\x7b\x6a\xb1\xad\x12\xe3\xd0\xa9\x9c\xbd\x56\x0f\xd2\x3b\x49\x2e\x85\x6c\xed\xe0\xa2\xa8\x45\x0a\x6e\x82\x8b\x66\x11\x75\x11\x51\xbf\xce\x35\x7e\x0c\x91\x80\x3a\xb4\xb8\xff\x7f\xef\xf3\x9a\x46\xd5\x25\x8e\x47\xe0\x39\x95\x06\x01\x18\x0b\x3b\x8c\x77\x15\x97\x82\x06\x71\x87\x8b\xbe\x84\x6d\x1a\x06\x67\x38\xe0\x87\x5c\x28\xbb\x98\xcf\x01\x42\x2a\x88\xc8\xf7\x81\xae\x1c\x0c\x7a\x42\xc1\xe2\xcc\xc2\x92\x69\x18\x82\x0e\x7a\x18\xd2\xa0\x7b\x44\x03\x7b\x75\x65\x7e\xac\x27\xe7\xe9\xfd\x53\x06\x86\xd4\x8f\x7e\xc5\x72\x71\x81\x18\x7d\x64\xb9\x8a\x8f\xff\xaf\xd3\xe7\x77\x7d\x73\x96\x81\x50\x06\xaa\x23\xe4\xdf\xff\x59\xaf\x4f\x23\x5f\x21\x3f\xcf\xce\x2e\xae\xf4\xdb\xe5\xcf\x5e\xa8\xa8\x8a\x42\x28\x2e\xe2\x6f\x5e\x58\x30\x78\xfa\x32\x1b\x8d\x51\x28\x21\xbd\x1b\xa7\xd3\x04\xf9\x12\xf4\xe4\x31\x9d\x26\xd9\x95\x5d\xb7\xde\x74\xdc\x36\xb6\x49\x1b\x36\x67\xb9\x1c\x48\x6b\xb3\xde\x22\xe5\xba\x10\xb2\x56\x01\x6a\x64\xc3\xd9\x6f\x78\xa8\x6e\x39\xee\x1e\xf1\xca\x91\xea\xaf\xa1\xba\xd3\x6c\x92\x96\x57\xb6\xf4\xc3\xb5\x3e\x79\xfd\xbc\x4d\xac\x75\xf3\x2b\x00\x00\xff\xff\x15\x82\x1a\x07\xb3\x01\x00\x00")

func outMysqlDds_dictionary_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlDds_dictionary_infoSql,
		"out/mysql/dds_dictionary_info.sql",
	)
}

func outMysqlDds_dictionary_infoSql() (*asset, error) {
	bytes, err := outMysqlDds_dictionary_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/dds_dictionary_info.sql", size: 435, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_data_permissionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x4d\x8f\xd2\x40\x1c\x06\xf0\x73\xfb\x29\xfe\xb7\xb6\x89\x07\xb6\xae\x9b\xcd\x1a\x0e\x5d\x18\x95\x08\x68\x6a\x3d\x70\x6a\x06\x3a\xe8\x68\x3b\x25\x9d\xa9\x91\x9b\x44\xa3\x86\xa4\xda\x44\x34\x92\x60\x94\x8b\xf4\x24\x24\x1a\x3d\x20\xdf\x86\xb6\xf8\x2d\x0c\x29\x12\x0f\xf8\x72\x9e\xe7\xf7\xcc\xcc\x23\x4b\x15\x13\x19\x16\x02\xcb\x38\xaf\x23\x00\xce\x7d\xdb\xc1\x02\xdb\x3d\x12\x78\x94\x73\xea\x33\x50\x65\x49\xa2\x0e\xb4\xe9\x1d\xca\x84\xaa\x97\x34\x00\xe6\x0b\x60\xa1\xeb\x02\x74\x7c\xcf\x23\x4c\x80\x42\x1d\x05\x2e\xc8\x92\xc4\xfb\xdc\xfe\x77\x3a\xff\xbc\xcc\x97\xef\xf3\xef\x6f\xd2\x97\xdf\x0a\x47\x9d\xed\xc1\x03\x1c\x74\xee\xe2\x40\xbd\xa8\xff\xcd\x65\x1f\x9e\x6d\xe6\x4f\x0b\xc7\xb0\x47\xf6\xec\x48\x3f\x3d\xec\xd2\x38\xca\x67\x8b\x42\x08\xdc\x76\x89\xfd\x7f\x6e\x33\x4d\xd2\x38\x2a\x9c\xdf\x23\x01\x16\xc4\xc6\x1d\xb1\x9d\xe5\x97\x3d\x39\x3e\x4c\xb3\x57\xd1\x7a\x35\x49\x87\xc9\x7a\x35\x29\x0a\x82\xd0\x25\x1c\x04\x79\x28\x00\x7e\xbf\x63\xf6\x24\x7d\x3e\xbe\xc7\x7d\xb6\x8b\x11\x0f\x07\xf7\xf7\xfd\xfa\xa5\x93\x3f\xbc\x6d\xfe\x25\x7b\xfb\x62\xb7\xba\xc0\x22\xe4\x20\x28\xeb\x6f\x67\x3f\xd6\xc0\x21\x5d\x1c\xba\x02\x4a\x07\x67\x1c\x7e\xcd\x1e\x0d\xa0\x74\x06\x69\x3c\xcf\x47\x09\x1c\x9d\x41\xfe\x71\x90\x8f\x92\xa2\xee\xa6\x59\x6b\x18\x66\x0b\xae\xa3\x16\xa8\xd4\xd1\x34\x40\xcd\xab\xb5\x26\x2a\xd7\x18\xf3\xab\xe7\x00\x55\x74\xc5\xb8\x5d\xb7\xa0\x72\xcd\x30\x6f\x21\xab\x1c\x8a\xee\x29\x54\x6e\x34\x1a\xa8\x69\x95\x95\xec\xf5\x22\x8b\x3e\x65\xef\x1e\xff\x18\xc7\xc5\xef\x36\xd3\x44\xb9\x2c\xff\x0c\x00\x00\xff\xff\x86\x6b\xef\x6b\x6e\x02\x00\x00")

func outMysqlSso_data_permissionSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_data_permissionSql,
		"out/mysql/sso_data_permission.sql",
	)
}

func outMysqlSso_data_permissionSql() (*asset, error) {
	bytes, err := outMysqlSso_data_permissionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_data_permission.sql", size: 622, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_operate_logSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x5f\x6b\xda\x50\x18\xc6\xaf\x93\x4f\xf1\xde\x19\x41\x42\x74\x13\x06\xc3\x8b\xa8\x67\x5b\x98\xba\x91\x65\x17\x5e\x49\x96\x1c\x5d\x20\x39\x91\xe4\x64\xe0\x07\xd8\x70\x0c\x11\xb6\xb9\x30\x2a\xed\x55\x69\xaf\xa4\xff\xa8\xd4\x54\xfa\x65\x3c\x89\xfd\x16\x45\x63\xa1\x82\xa5\xbd\x3b\xfc\xde\xf7\x79\xcf\xf3\x3c\x3c\x57\x51\x91\xac\x21\xd0\xe4\x72\x0d\x01\xf8\xbe\xdb\x72\xbb\xd8\xd3\x29\x6e\xd9\x6e\x07\x04\x9e\xe3\x2c\x13\xbe\x58\x1d\x8b\x50\xa1\x20\x65\x01\x88\x4b\x81\x04\xb6\x0d\x60\xb8\x8e\x83\x09\x85\x8c\x65\x66\x20\xc7\x73\x1c\xed\x75\x31\x50\x8b\xf4\xd6\xcb\xbb\x77\x93\xd3\x88\xed\xff\xca\xe5\x25\x31\xf9\x1f\xb1\xf9\x28\xfe\x33\x58\xcc\xc7\x50\x90\xc4\xe4\x3c\x4a\xa2\x83\x78\x74\x12\x0f\x26\x1b\x0a\x2f\x24\x71\x79\xf4\x7b\xf9\xf3\x6c\x1b\xbf\x94\xc4\xe5\x70\xcc\x06\xa3\x6d\x5c\x94\xc4\xe4\xef\x71\xdc\x9f\x3e\xc4\xa9\x35\xbf\xe7\xb7\x9e\x0e\x92\x5a\x48\xae\xff\xb1\xe1\x34\xd5\x05\x3e\xf6\x9e\x21\x4c\xbf\x5a\xcc\x66\xf7\x55\x18\x1e\x5e\x75\x48\x2d\x07\x83\xa9\x53\x9c\x3e\x70\x5b\x0f\x6c\x0a\x46\xe0\x79\x98\xd0\xf5\xd4\xa7\xba\xd3\xdd\x75\x92\xf5\xf7\x58\x34\x8b\xc3\xcb\xdb\xf0\x62\x73\xd3\x25\x74\x35\xfa\xa6\x7b\xc6\x57\xdd\x13\x8a\xf9\x47\x3a\x66\x3f\xbe\xb3\xc9\x95\x92\xaa\x3e\xaa\x4a\x5d\x56\x9b\xf0\x1e\x35\x41\xb0\xcc\x6c\x16\x50\xe3\xad\xd2\x40\x25\x85\x10\xb7\x5a\x06\xa8\xa2\x37\xf2\xe7\x9a\x06\x95\x77\xb2\xfa\x09\x69\xa5\x80\xb6\x5f\x41\xe5\x43\xbd\x8e\x1a\x5a\x69\x93\x2c\x0e\x0f\xd9\x4d\x98\x79\xcd\xdf\x05\x00\x00\xff\xff\xa3\x28\x42\xb4\x31\x02\x00\x00")

func outMysqlSso_operate_logSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_operate_logSql,
		"out/mysql/sso_operate_log.sql",
	)
}

func outMysqlSso_operate_logSql() (*asset, error) {
	bytes, err := outMysqlSso_operate_logSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_operate_log.sql", size: 561, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_role_datapermissionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x4f\x4b\x2a\x51\x18\x87\xd7\x33\x9f\xe2\xdd\xe9\xc0\x5d\xe8\x5d\x5d\xbc\xb8\x18\xf5\xdc\x9b\xa4\x16\xd3\xb4\x70\x35\x4c\x33\x47\x39\x30\x73\x46\xe6\x9c\x59\xb8\x6b\xa8\xc0\x5c\x14\x84\x24\xd1\x42\x24\x28\x37\x11\x64\x98\x64\xd1\x97\x71\xfe\xf4\x2d\x22\x47\xc8\x85\x60\xbb\x03\x87\xe7\xe1\x79\x7f\xa2\x50\x54\x90\xac\x22\x50\xe5\x42\x05\x01\x30\xe6\x68\xae\x63\x61\xcd\xd4\xb9\xde\xc2\xae\x4d\x18\x23\x0e\x85\xb4\x28\x08\xc4\x84\x03\xd2\x24\x94\xa7\x7f\x67\x24\x00\xea\x70\xa0\x9e\x65\x01\x18\x8e\x6d\x63\xca\x21\x15\x74\x07\xf1\xd1\x5b\xf4\x7a\x19\x9c\x3f\xa7\xe0\x97\x28\x08\xac\xcd\xb4\xcd\x5c\x34\x9e\x45\xb3\xc1\x2a\xb7\x68\xd8\x0c\xc6\x77\x17\xf1\xe9\xe3\x2a\xf8\xdd\xac\x19\x0e\x6d\x90\xe6\x8f\x2c\xc7\x41\xe7\x8a\x98\xcb\x64\xae\x73\x8f\x01\x27\xb4\xfd\x05\x65\x25\x30\x71\x43\xf7\x2c\x0e\x99\xb5\xed\xdd\x49\x78\xe8\x43\x36\x07\xd1\xad\x1f\xf5\x46\x90\xc9\x41\x78\x7f\x13\x4c\xa7\x89\xce\x70\xb1\xce\xb1\xc6\x89\x8d\xc1\xd4\x39\x4e\x1e\x4b\xa3\xe1\xb9\x2e\xa6\x7c\xf1\xcb\xb8\x6e\xb7\xd6\xae\xda\xb9\x0e\x66\x2f\x61\x7f\xf2\xd1\x7f\x4a\x9c\xbb\x4a\xb9\x2a\x2b\x75\xd8\x46\x75\x48\x13\x53\x92\x00\xd5\xfe\x97\x6b\x28\x5f\xa6\xd4\x29\x15\x00\x4a\xe8\x9f\xbc\x5f\x51\xa1\xb8\x25\x2b\x7b\x48\xcd\x7b\xbc\xf1\x07\x8a\x3b\xd5\x2a\xaa\xa9\xf9\xe5\x6a\xf3\xe9\x59\x72\x78\x70\x32\x8e\xfd\xde\xfc\x7d\x18\xfa\x0f\xf1\x70\x94\xfa\x2b\x7e\x06\x00\x00\xff\xff\xa1\x6d\x7f\x6b\x16\x02\x00\x00")

func outMysqlSso_role_datapermissionSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_role_datapermissionSql,
		"out/mysql/sso_role_datapermission.sql",
	)
}

func outMysqlSso_role_datapermissionSql() (*asset, error) {
	bytes, err := outMysqlSso_role_datapermissionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_role_datapermission.sql", size: 534, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_role_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x31\x4b\xf3\x50\x14\x86\xe7\xe4\x57\xbc\x5b\x13\xf8\x86\xf6\x43\x44\x2a\x1d\xd2\xf6\xaa\xc5\xb6\x4a\x8c\x43\xa7\x72\x4d\x6e\x35\x90\xdc\x2b\x37\x27\x82\x9b\xdd\x54\x10\x44\x84\x0e\x4e\x22\xd8\x2e\x82\x88\x94\xa2\xbf\x27\x4d\xfd\x17\xa2\xe9\x58\xdc\xce\xe1\xf0\x3e\x0f\xe7\x35\x8d\x86\xcb\x1c\x8f\xc1\x73\xea\x6d\x06\x24\x89\xea\x6b\x15\x89\x7e\x28\x07\x0a\x96\x69\x18\xc5\x16\xe0\x28\x3c\x0e\x25\x59\xff\xcb\x36\x20\x15\x41\xa6\x51\x04\xf8\x2a\x8e\x85\x24\x94\x16\xe3\xbb\xc5\xd5\x5b\x18\x94\xf0\xcf\x34\x0c\xc9\x63\x81\x33\xae\xfd\x13\xae\xad\xf5\xb5\xbf\x22\xd9\xed\x4d\x3e\x7e\x2d\x62\x09\x71\x4a\x13\x50\x28\xcf\x7f\x5c\x15\x1b\x81\x18\xf0\x34\x22\x94\x57\x01\xf2\xeb\xe9\xfc\x62\x88\x4a\x15\xf9\xf3\x30\xbf\x9f\xa0\x5c\xc5\xfc\xe5\x29\x9b\xcd\x0a\x9c\xaf\x05\x27\xd1\xa7\x30\x16\x08\x38\x89\x62\x58\x12\xfd\x54\x6b\x21\xe9\xf7\x9a\x10\x8f\x4f\x57\x19\xb2\xcb\x87\xec\xf3\x63\x3e\x9a\x7e\x8d\xde\x0b\xe6\xbe\xdb\xea\x38\x6e\x0f\xbb\xac\x07\x6b\xd9\x8d\x6d\x83\x75\xb7\x5b\x5d\x56\x6b\x49\xa9\x9a\x75\xa0\xc9\xb6\x9c\xc3\xb6\x87\xc6\x8e\xe3\x1e\x30\xaf\x96\xd2\x60\x03\x8d\xbd\x4e\x87\x75\xbd\xda\xf2\xf1\xc5\xe3\xa4\xb4\x69\x7e\x07\x00\x00\xff\xff\x7f\x8d\xf1\xbb\x81\x01\x00\x00")

func outMysqlSso_role_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_role_infoSql,
		"out/mysql/sso_role_info.sql",
	)
}

func outMysqlSso_role_infoSql() (*asset, error) {
	bytes, err := outMysqlSso_role_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_role_info.sql", size: 385, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_role_menuSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xd0\x3f\x4f\xc2\x40\x18\x06\xf0\xb9\xfd\x14\xcf\x06\x24\x0e\xe0\x64\x30\x0c\x05\x4e\x25\x02\x9a\x5a\x07\xa6\xa6\xd0\xc3\x5c\x6c\xaf\xa6\xbd\x0e\x6c\xb2\xa9\x89\x92\x10\x95\x41\x13\x63\x4c\x94\xc5\x45\x0d\x12\xf0\xeb\xb4\xc5\x6f\x61\xf8\xd3\xc1\x84\x04\xdd\x2e\xb9\xfb\x3d\xef\xbd\x8f\x2c\x15\x54\xa2\x68\x04\x9a\x92\x2f\x13\xc0\xf3\x1c\xdd\x75\x2c\xaa\xdb\x94\xfb\x48\xca\x92\xc4\x4c\xd4\xd9\x11\xe3\x22\xb9\x9e\x4e\x01\xdc\x11\xe0\xbe\x65\x01\x0d\xc7\xb6\x29\x17\x48\x30\x33\x81\x35\x59\x92\xbc\x96\xa7\xaf\x7e\x1d\xbd\x8f\xa3\xf1\x43\x6c\x66\xb3\x7e\x23\x93\x36\x0d\xdf\x12\x48\x2f\xd3\x93\x97\xee\xe4\xfc\x2d\xd6\xd3\x4f\xfe\x4b\x77\xee\x83\xcb\x9b\x58\x53\x6e\xd4\x2d\x0a\xc1\x78\x6b\xaa\x33\x2b\x70\x74\x31\x08\x4f\xdb\xc8\x64\x11\x3d\xb7\xa3\xeb\x3e\xd2\x59\x84\xaf\x4f\xc1\x70\x38\x8f\x6b\xb8\xd4\x10\x54\x17\xcc\xa6\x30\x0d\x41\xe7\x87\x45\x62\xc3\x77\x5d\xca\xc5\xec\xd6\x13\x86\x7d\xb2\x6c\x42\x70\x76\x17\x8c\x47\x61\x6f\xf0\xdd\xfb\x58\x54\xea\xb8\xc2\x35\xf8\xf1\x9f\x37\x0c\xaf\xba\xc1\xa8\x13\x7d\xdd\x06\x9d\xcf\x79\xc4\xbe\x5a\xaa\x28\x6a\x0d\xbb\xa4\x86\x24\x33\x53\x29\x90\xea\x76\xa9\x4a\x72\x25\xce\x9d\x62\x1e\x28\x92\x2d\xe5\xb0\xac\xa1\xb0\xa3\xa8\x07\x44\xcb\xf9\xa2\xb9\x81\xc2\x5e\xa5\x42\xaa\x5a\x6e\x51\xf8\xe4\xb1\x9f\xd8\x94\x7f\x02\x00\x00\xff\xff\xca\x39\xae\xef\x2d\x02\x00\x00")

func outMysqlSso_role_menuSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_role_menuSql,
		"out/mysql/sso_role_menu.sql",
	)
}

func outMysqlSso_role_menuSql() (*asset, error) {
	bytes, err := outMysqlSso_role_menuSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_role_menu.sql", size: 557, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_system_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x51\x6b\xda\x50\x14\xc7\x9f\xf5\x53\x9c\x37\x15\x5a\x50\x3b\x46\xe9\xf0\xc1\xda\x6c\x93\x55\x37\x5c\xf6\xd0\xa7\x70\x35\x57\xbd\x90\xdc\x48\x72\xb3\xd5\xb7\x48\x69\x07\x85\x32\x19\x76\x6e\x63\xa5\xb5\xb0\xae\x2f\x5d\x3a\x4a\x9d\xa0\xad\x5f\xc6\x9b\x98\xa7\x7d\x85\x91\xa6\x95\x74\xc8\xfa\x98\x9c\xf3\xff\xff\xcf\x39\xbf\x1b\x8d\xe4\x4a\x42\x56\x14\x40\xcc\xae\xae\x0b\x00\x86\xa1\x49\x46\xd3\x60\x58\x95\x08\xad\x6a\x10\x8f\x46\x22\x44\x86\x32\xa9\x11\xca\xe2\xe9\x64\x02\x80\x6a\x0c\xa8\xa9\x28\x00\x15\x4d\x55\x31\x65\x10\x23\x72\x0c\x16\xa2\x91\x08\x45\x2a\x86\xb7\x48\xaf\xd4\x91\x1e\x5f\x4a\xcf\x6f\x76\x2f\x86\xee\xf0\x90\xb7\xf7\xdc\x1f\xe7\x81\x8c\x50\x19\x6f\x4a\xa6\xae\xcc\xb4\x8f\x1f\x25\x00\x42\x1a\xef\xe4\x93\xd7\xbb\xe4\xdf\xce\xf9\x81\x15\x68\x30\x45\x65\x05\x03\x23\xb4\xe9\x4f\x96\x4a\x80\x8c\xab\xc8\x54\x18\xa4\xe6\x86\xee\xf6\x1d\xab\x05\xa9\x3f\xa3\xaf\xbc\x6d\xbb\x9d\x53\x48\xae\x80\x7b\xd2\x72\x3b\xa7\x81\x9f\xa2\xd5\x08\x95\x18\x51\xb1\x66\x32\xb8\xb1\x0c\x79\x2e\x25\x93\xf3\x5c\xa7\xfd\x6d\xa7\xdb\x77\xba\x7d\x6f\x7f\x3c\xb3\xd1\x66\x5b\xa4\xd2\xcb\xf7\xd7\xf0\xab\x41\x1f\xab\xe3\xd0\xa9\x6e\x1a\xef\xb2\x62\xe5\xda\x62\x03\xe9\x44\x45\x7a\x33\x16\x56\x4f\x06\x43\xef\xf8\xb3\x73\xf4\x9b\x8f\x3e\xdc\xa6\xa1\xa6\x3f\xed\x7c\x1b\xd4\x68\x2c\xd6\x31\x92\xb1\xbe\x58\x25\x9b\x58\x06\xff\x07\x32\x88\x8c\x83\xef\x7b\xde\x5e\xef\xd2\x3b\x38\xe6\x83\x2d\xfe\xcb\x0a\x27\x10\xd9\x2f\x3f\x84\x94\x77\xec\xc9\xc0\x72\x8e\xde\x4f\xed\x9d\xf0\x39\xff\x87\xd4\xfd\x32\xe4\x57\xfb\x61\xa4\xef\x70\xa5\x8e\x98\x64\x30\xc4\x4c\x63\x46\x36\xfd\x00\x59\x7e\xfd\x73\x32\xee\xf1\xdd\xc3\xe9\xd6\xd5\x1d\xe5\x15\xe0\x23\x8b\xb7\xed\x05\x9f\x32\xdf\xbe\xf0\xba\x67\x41\x84\x81\x2b\x3a\xfe\x67\x9f\xf0\x4c\x67\xd7\xbc\xbd\xc7\xed\x1d\xef\xe3\xf7\x40\xf0\xaa\x94\x2f\x64\x4b\x1b\xf0\x42\xd8\x80\x38\x91\x13\x09\x10\x8a\xcf\xf2\x45\x21\x93\xa7\x54\x5b\x5b\x05\x58\x13\x9e\x66\xdf\xac\x8b\x90\x7b\x9e\x2d\xbd\x16\xc4\x8c\xc9\xaa\xcb\x90\x7b\x59\x28\x08\x45\x31\x73\xfb\xd6\x27\xe3\x9e\xd3\xb2\x63\x4f\xa2\x7f\x03\x00\x00\xff\xff\x5f\x7c\x39\x88\x6e\x03\x00\x00")

func outMysqlSso_system_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_system_infoSql,
		"out/mysql/sso_system_info.sql",
	)
}

func outMysqlSso_system_infoSql() (*asset, error) {
	bytes, err := outMysqlSso_system_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_system_info.sql", size: 878, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_system_menuSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x41\x6b\x13\x41\x1c\xc5\xcf\x9b\x4f\xf1\xbf\x65\x17\x3c\xa4\x51\x8b\x54\x72\xd8\xa6\xa3\x06\x9b\x28\x71\x3d\xf4\xb4\x4c\x77\xa7\x76\x70\x77\x36\xec\xcc\x16\x72\x6b\xf1\x60\x2d\xc4\x14\x31\x56\x68\xa5\x8a\xd0\xf4\x52\x8d\x28\xb1\x6c\xaa\xdf\x66\x76\xe2\xb7\x90\x74\xd7\xba\x42\x62\xe6\x34\xf0\x9f\xdf\xfb\xbf\xf7\x98\x82\x56\x6d\x22\xd3\x42\x60\x99\xcb\xab\x08\x80\xf3\xc0\xe6\x6d\x2e\x88\x6f\xfb\x84\x45\xa0\x17\x34\x8d\xba\xb0\x4e\x9f\x50\x26\xf4\x72\xc9\x00\x60\x81\x00\x16\x79\x1e\x80\x13\xf8\x3e\x61\x02\x8a\x72\xef\x78\xfc\xec\x87\xba\x78\x23\xbb\xdf\xe1\xef\xb9\x5e\x2e\xc2\xb5\x82\xa6\x31\xec\x13\xd8\xc2\xa1\xb3\x89\x43\x7d\xf1\xc6\xff\x34\xe4\x7e\x47\xf5\x07\x29\xd6\xc2\xe1\x64\x32\x67\xb7\xda\x1d\xaa\xb8\x9f\xee\x4e\x39\xde\xe6\xf6\x7c\xcf\xea\xeb\x48\x8d\x8e\xf3\x9c\x47\xb6\x88\x37\x21\x05\x65\xed\x4b\x74\x06\x79\xf6\x42\xc5\xfd\x94\xa1\x4e\xc0\xfe\x8d\x96\x8f\x74\xf8\x33\x79\xff\xfc\x4f\x18\xb1\x79\xf5\xb0\x7c\x73\x71\x46\x09\x47\x03\xf9\x6e\x3b\x25\x08\xc3\xeb\x1e\xb9\x32\xb3\x60\x80\x4b\x36\x70\xe4\x09\x28\x4d\x75\xb5\x37\x4c\xb6\x77\x60\x61\x09\xd4\xc9\x8e\x7a\x7d\x0a\xa5\x25\x48\xce\x3e\xca\xf3\xf3\x54\xce\x09\x09\x16\xc4\x16\xd4\x27\xe0\x62\x41\xd2\x4b\xa6\xe8\x44\xe1\xa4\xeb\xcb\x29\x17\xd8\x6f\x4d\x35\xb7\x7b\x28\x47\x71\x72\x30\xfc\x75\xf0\x2d\x6b\x3a\x08\x45\x88\xd9\xd3\x79\x5d\x27\x2f\x5f\xc9\xb8\x9b\xef\x9a\x72\x3b\x68\x11\x36\x3d\x5d\x9e\x7c\xfb\x59\xee\x9f\xc8\x2f\x3d\x79\x91\xd5\xf2\xb0\x59\xab\x9b\xcd\x35\xb8\x8f\xd6\x40\xa7\xae\x61\x00\x6a\xdc\xad\x35\x50\xa5\xc6\x58\xb0\xb2\x0c\xb0\x82\xee\x98\x8f\x57\x2d\xa8\xde\x33\x9b\x8f\x90\x55\x89\xc4\xc6\x2d\xa8\x3e\xa8\xd7\x51\xc3\xaa\x64\xdf\x6c\xfc\xe1\x54\x1f\x77\x8f\x64\xa7\x97\xf4\x06\x49\xe7\x93\x51\xbc\x5d\xf8\x1d\x00\x00\xff\xff\x8b\x8d\xb0\x78\x0b\x03\x00\x00")

func outMysqlSso_system_menuSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_system_menuSql,
		"out/mysql/sso_system_menu.sql",
	)
}

func outMysqlSso_system_menuSql() (*asset, error) {
	bytes, err := outMysqlSso_system_menuSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_system_menu.sql", size: 779, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_user_infoSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\xcf\x6b\x13\x41\x14\xc7\xcf\xc9\x5f\xf1\x6e\x4d\xc0\x43\x12\xa5\x94\x48\x0e\x69\xba\x6a\xb0\x89\x12\xe3\xa1\xa7\x65\xba\x3b\x49\x07\x76\x67\x96\x9d\x59\x53\x6f\x8d\xa0\x54\x41\x9b\x43\x42\xf1\x47\xc1\x20\x4a\x04\x4d\x15\xb4\xad\x2d\xda\x7f\xa6\x33\xdb\xfe\x17\x62\x66\xd3\x6c\x4a\x5a\x6f\xbb\xbc\xf7\x79\xef\x7d\xbf\xef\x4d\x32\x51\xaa\x19\xc5\xba\x01\xf5\xe2\xe2\xb2\x01\xc0\x39\x33\x03\x8e\x7d\x93\xd0\x06\x83\x54\x32\x91\xd0\x7f\x36\xac\x92\x26\xa1\x22\x95\xcb\xa4\x01\x28\x13\x40\x03\xc7\x01\xb0\x98\xeb\x62\x2a\x60\x8e\xd8\x73\x70\x2d\x99\x48\x34\x02\xc7\x31\x29\x72\x31\x3c\x42\xbe\xb5\x86\xfc\xd4\xf5\xdc\x6c\x22\xec\x0e\xd4\xe6\xbe\x7c\x3a\x90\x9d\x97\x9a\x1d\xb5\x9a\x62\xe7\x6f\x5c\xc9\x8e\x41\x0f\x71\xde\x62\xbe\xfd\xdf\x9e\x72\xf7\x59\xf8\xbe\xad\x21\xec\x22\xe2\x4c\x13\xb1\xcc\x51\x54\x27\x72\x81\x44\xc0\x41\x10\xfa\xf8\x9f\x03\xd9\x34\xd8\xb8\x81\x02\x47\x40\x76\xe6\x6c\x2f\xf6\xd4\x46\x1b\x32\x79\x50\x5f\x3f\xc8\x83\x03\xc8\xe6\xe1\xac\xdb\x96\xc3\x37\x90\xcb\x43\xf8\xa9\x1d\x76\x07\xba\xb0\xcb\x56\x89\x33\x11\x9b\xbd\xd4\xa8\x9f\xa7\xbb\x3b\x72\x6b\xff\x7c\xf4\xd6\xba\xc9\x3c\x4c\x89\x3d\x6d\x54\x5c\xe8\x9f\xe1\xc9\x71\x5f\x27\x69\xc6\xf2\x31\x12\xd8\x14\xc4\xc5\x60\x23\x81\xf5\x47\x24\xc4\x0a\x7c\x1f\x53\x31\x8a\x72\x81\x5c\x6f\xa6\x79\x9b\x6f\xe5\xd1\xa1\xda\xde\x3b\xdb\xfe\x11\xd5\x5c\x43\xb4\x89\xbd\x96\xad\xc1\xf8\x8d\x8c\x2b\x67\x2e\x5f\xc3\xc9\xf1\x50\x75\x7f\xa9\x2f\x7d\xd5\xfb\x16\xad\x64\x5d\x98\x1e\xf2\x91\xcb\x27\xa6\x64\x72\x17\xa4\xa9\xe7\x9f\xe5\xf7\x9e\xdc\x7a\x72\x8e\x39\x88\x0b\xd3\x61\x4d\x42\x2f\xc8\x9b\xc2\xde\x6d\xc8\xce\xab\xf0\xf5\x91\xfc\xdd\x8b\x6b\xe0\x2c\xf0\x2d\x6c\xc6\xbc\xcc\xe6\x16\x66\x2f\x42\xed\x7c\x54\x87\x9d\xb1\xa1\x1a\x9c\x1c\xd0\xfc\x55\x90\x46\xee\xd7\xca\x95\x62\x6d\x05\xee\x1a\x2b\x90\x8a\x1e\x56\x3a\x0d\x46\xf5\x76\xb9\x6a\x14\xca\x94\xb2\xa5\x45\x80\x25\xe3\x56\xf1\xe1\x72\x1d\x4a\x77\x8a\xb5\x07\x46\xbd\x10\x88\xc6\x02\x94\xee\x55\x2a\x46\xb5\x5e\x88\x4e\xff\xb4\x3f\x98\xbb\x99\xfc\x1b\x00\x00\xff\xff\x52\xaf\xd2\x5f\xbe\x03\x00\x00")

func outMysqlSso_user_infoSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_user_infoSql,
		"out/mysql/sso_user_info.sql",
	)
}

func outMysqlSso_user_infoSql() (*asset, error) {
	bytes, err := outMysqlSso_user_infoSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_user_info.sql", size: 958, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _outMysqlSso_user_roleSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\xb1\x4a\xc3\x50\x14\xc6\xf1\x39\x79\x8a\x6f\x4b\x02\x0e\x8d\x93\x54\x32\xa4\xe9\x55\x83\x4d\x94\x18\x87\x4e\xa5\x35\xb7\x72\x21\xbd\x81\xe6\x66\xe8\x66\x06\x41\x5d\x14\x09\xae\xe2\xa2\xd9\x04\x0b\x76\xa8\xaf\x93\xde\x3c\x86\xd4\x38\x74\x10\xec\x7e\x7e\xe7\xfc\x39\xaa\xe2\x04\xc4\x0e\x09\x42\xbb\xd3\x23\x40\x9a\x26\x83\x2c\xa5\xd3\xc1\x34\x89\x29\x74\x55\x51\x58\x84\x11\xbb\x64\x5c\xe8\xbb\x2d\x03\xe0\x89\x00\xcf\xe2\x18\xb8\x48\x26\x13\xca\x05\x34\x16\x69\xd8\x51\x15\xe5\xc7\xfd\x3f\x2e\x8b\x72\x75\xb3\x90\x5f\x4f\xd5\xfd\xa2\x81\xe9\x2c\xdd\xc6\xcd\x97\x72\xf9\xbc\xe9\xd6\x8d\x5b\xc0\xfa\xed\xb1\xbe\xfd\xd8\x84\x94\x0f\x47\x31\x85\x60\x7c\xb6\x86\xa6\x81\x88\x8e\x87\x59\x2c\x60\xfe\x79\xf9\xee\x73\x75\x95\xa3\xd5\x46\xf5\xf0\x2e\x8b\x12\x66\x1b\xf2\x35\x97\x45\xd9\xac\x3b\x0d\x5c\xcf\x0e\xfa\x38\x26\x7d\xe8\x2c\x32\x0c\x10\xff\xd0\xf5\x89\xe5\x72\x9e\x74\x3b\x40\x97\x1c\xd8\xe7\xbd\x10\xce\x91\x1d\x9c\x91\xd0\xca\xc4\x78\x0f\xce\x89\xe7\x11\x3f\xb4\x7e\x1f\xd2\x54\x56\xd7\xf3\x3a\x2f\xea\x97\x52\xdb\x57\xbf\x03\x00\x00\xff\xff\xcb\x54\xb4\xa2\x9c\x01\x00\x00")

func outMysqlSso_user_roleSqlBytes() ([]byte, error) {
	return bindataRead(
		_outMysqlSso_user_roleSql,
		"out/mysql/sso_user_role.sql",
	)
}

func outMysqlSso_user_roleSql() (*asset, error) {
	bytes, err := outMysqlSso_user_roleSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out/mysql/sso_user_role.sql", size: 412, mode: os.FileMode(509), modTime: time.Unix(1609152482, 0)}
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
	"out/mysql/dds_area_info.sql":           outMysqlDds_area_infoSql,
	"out/mysql/dds_dictionary_info.sql":     outMysqlDds_dictionary_infoSql,
	"out/mysql/sso_data_permission.sql":     outMysqlSso_data_permissionSql,
	"out/mysql/sso_operate_log.sql":         outMysqlSso_operate_logSql,
	"out/mysql/sso_role_datapermission.sql": outMysqlSso_role_datapermissionSql,
	"out/mysql/sso_role_info.sql":           outMysqlSso_role_infoSql,
	"out/mysql/sso_role_menu.sql":           outMysqlSso_role_menuSql,
	"out/mysql/sso_system_info.sql":         outMysqlSso_system_infoSql,
	"out/mysql/sso_system_menu.sql":         outMysqlSso_system_menuSql,
	"out/mysql/sso_user_info.sql":           outMysqlSso_user_infoSql,
	"out/mysql/sso_user_role.sql":           outMysqlSso_user_roleSql,
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
			"dds_area_info.sql":           &bintree{outMysqlDds_area_infoSql, map[string]*bintree{}},
			"dds_dictionary_info.sql":     &bintree{outMysqlDds_dictionary_infoSql, map[string]*bintree{}},
			"sso_data_permission.sql":     &bintree{outMysqlSso_data_permissionSql, map[string]*bintree{}},
			"sso_operate_log.sql":         &bintree{outMysqlSso_operate_logSql, map[string]*bintree{}},
			"sso_role_datapermission.sql": &bintree{outMysqlSso_role_datapermissionSql, map[string]*bintree{}},
			"sso_role_info.sql":           &bintree{outMysqlSso_role_infoSql, map[string]*bintree{}},
			"sso_role_menu.sql":           &bintree{outMysqlSso_role_menuSql, map[string]*bintree{}},
			"sso_system_info.sql":         &bintree{outMysqlSso_system_infoSql, map[string]*bintree{}},
			"sso_system_menu.sql":         &bintree{outMysqlSso_system_menuSql, map[string]*bintree{}},
			"sso_user_info.sql":           &bintree{outMysqlSso_user_infoSql, map[string]*bintree{}},
			"sso_user_role.sql":           &bintree{outMysqlSso_user_roleSql, map[string]*bintree{}},
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