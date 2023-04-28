package fsx

import (
	"archive/zip"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type FS struct {
	path     string
	fileInfo os.FileInfo
}

func New(path string) (*FS, error) {
	var err error
	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}
	}
	fs := &FS{path: path}
	fs.fileInfo, _ = fs.State()

	return fs, nil
}

func (fs *FS) State() (os.FileInfo, error) {
	if fs.fileInfo == nil {
		var err error
		fs.fileInfo, err = os.Stat(fs.path)
		if err != nil {
			return nil, err
		}
	}
	return fs.fileInfo, nil
}

// exists
func (fs *FS) Exists() bool {
	if _, err := os.Stat(fs.path); err != nil {
		return false
	}
	return true
}

func (fs *FS) Path() string {
	return fs.path
}

// path base
func (fs *FS) Base() string {
	return filepath.Base(fs.path)
}

func (fs *FS) Ext() string {
	return filepath.Ext(fs.path)
}

// mkdir
func (fs *FS) Mkdir(perm fs.FileMode) error {
	return os.MkdirAll(fs.path, perm)
}

// open file
func (fs *FS) Open() (*os.File, error) {
	return os.Open(fs.path)
}

// open file with flag
func (fs *FS) CreateFile(perm fs.FileMode) (*os.File, error) {
	return os.Create(fs.path)
}

// open file with flag
func (fs *FS) OpenFile(flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(fs.path, flag, perm)
}

// remove file
func (fs *FS) Remove() error {
	return os.Remove(fs.path)
}

// move file
func (fs *FS) Move(newPath string) error {
	return os.Rename(fs.path, newPath)
}

// remove all
func (fs *FS) RemoveAll() error {
	return os.RemoveAll(fs.path)
}

// rename file
func (fs *FS) Rename(newPath string) error {
	return os.Rename(fs.path, newPath)
}

// truncate file
func (fs *FS) Truncate(size int64) error {
	return os.Truncate(fs.path, size)
}

// change permission
func (fs *FS) Chmod(perm fs.FileMode) error {
	return os.Chmod(fs.path, perm)
}

// change owner
func (fs *FS) Chown(uid, gid int) error {
	return os.Chown(fs.path, uid, gid)
}

// change owner and group
func (fs *FS) ChownUid(uid int) error {
	return os.Chown(fs.path, uid, -1)
}

// cal file md5 hash
func (fs *FS) Md5() (string, error) {
	file, err := fs.Open()
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// cal file sha256 hash
func (fs *FS) Sha256() (string, error) {
	file, err := fs.Open()
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// is file or directory
func (fs *FS) IsDir() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.IsDir()
}

// is file
func (fs *FS) IsFile() bool {
	if !fs.Exists() {
		return false
	}
	return !fs.fileInfo.IsDir()
}

// is symlink
func (fs *FS) IsSymlink() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeSymlink != 0
}

// is socket
func (fs *FS) IsSocket() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeSocket != 0
}

// is named pipe
func (fs *FS) IsNamedPipe() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeNamedPipe != 0
}

// is character device
func (fs *FS) IsCharDevice() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeCharDevice != 0
}

// is block device
func (fs *FS) IsBlockDevice() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeDevice != 0
}

// is setuid
func (fs *FS) IsSetuid() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeSetuid != 0
}

// is setgid
func (fs *FS) IsSetgid() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeSetgid != 0
}

// is sticky
func (fs *FS) IsSticky() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeSticky != 0
}

// is regular
func (fs *FS) IsRegular() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode().IsRegular()
}

// is append-only
func (fs *FS) IsAppend() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeAppend != 0
}

// is exclusive use
func (fs *FS) IsExclusive() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeExclusive != 0
}

// is temporary
func (fs *FS) IsTemporary() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeTemporary != 0
}

// is device
func (fs *FS) IsDevice() bool {
	if !fs.Exists() {
		return false
	}
	return fs.fileInfo.Mode()&os.ModeDevice != 0
}

// zip dir
func (fs *FS) Zip(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		return errors.New("file already exists")
	}

	// zip a dir to a file
	zipFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	w := zip.NewWriter(zipFile)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		f, err := w.Create(strings.ReplaceAll(path, fs.path, ""))
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}
		return nil
	}

	return filepath.Walk(fs.path, walker)
}

// unzip to dst path
func (fs *FS) Unzip(dst string) error {
	r, err := zip.OpenReader(fs.path)
	if err != nil {
		return err
	}
	defer r.Close()

	dstFs, err := New(dst)
	if err != nil {
		return err
	}
	if !dstFs.Exists() {
		if err := dstFs.Mkdir(os.ModePerm); err != nil {
			return err
		}
	}

	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}
		if err := fs.unzipFile(f, dstFs); err != nil {
			return err
		}
	}

	return nil
}

// unzip file
func (fs *FS) unzipFile(f *zip.File, dstFs *FS) error {
	// nolint gosec // ignore G305: File traversal when extracting zip archive
	dstPath := filepath.Join(dstFs.path, f.Name)
	if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
		return err
	}

	file, err := f.Open()
	if err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// nolint G110 // ignore G110: Potential DoS vulnerability via decompression bomb
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		return err
	}
	return nil
}
