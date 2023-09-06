package main

import (
	"context"
	"golang.org/x/net/webdav"
	"os"
	"path/filepath"
)

type twdsFS struct {
	base_path string
}

func (fs *twdsFS) init(base_path string) {
	fs.base_path = base_path
}

func (fs *twdsFS) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	real_file_name := filepath.Join(fs.base_path, name)
	return os.Mkdir(real_file_name, perm)
}

func (fs *twdsFS) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	real_file_name := filepath.Join(fs.base_path, name)
	return os.OpenFile(real_file_name, flag, perm)
}

func (fs *twdsFS) RemoveAll(ctx context.Context, name string) error {
	real_file_name := filepath.Join(fs.base_path, name)
	return os.RemoveAll(real_file_name)
}

func (fs *twdsFS) Rename(ctx context.Context, old_name, new_name string) error {
	real_old_name := filepath.Join(fs.base_path, old_name)
	real_new_name := filepath.Join(fs.base_path, new_name)
	return os.Rename(real_old_name, real_new_name)
}

func (fs *twdsFS) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	real_file_name := filepath.Join(fs.base_path, name)
	return os.Stat(real_file_name)
}
