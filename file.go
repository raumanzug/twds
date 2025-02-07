package main

import (
	"context"
	"golang.org/x/net/webdav"
	"os"
)

type twdsFS struct {
	base_path webdav.Dir
}

func (fs *twdsFS) init(base_path string) {
	fs.base_path = webdav.Dir(base_path)
}

func (fs *twdsFS) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	return fs.base_path.Mkdir(ctx, name, perm)
}

func (fs *twdsFS) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	return fs.base_path.OpenFile(ctx, name, flag, perm)
}

func (fs *twdsFS) RemoveAll(ctx context.Context, name string) error {
	return fs.base_path.RemoveAll(ctx, name)
}

func (fs *twdsFS) Rename(ctx context.Context, old_name, new_name string) error {
	return fs.base_path.Rename(ctx, old_name, new_name)
}

func (fs *twdsFS) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	return fs.base_path.Stat(ctx, name)
}
