package common

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
)

func GetProgramPath() string {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exReal)
}

func PutJsonByData(ctx context.Context, path string, data interface{}) error {
	jsonCount, e := gjson.EncodeString(data)
	if e != nil {
		return e
	}
	return gfile.PutContents(path, jsonCount)
}

type DataFunc func(context.Context) (interface{}, error)

func PutJsonByFunc(ctx context.Context, path string, funcData DataFunc) error {
	data, e := funcData(ctx)
	if e != nil {
		return e
	}
	jsonCount, e := gjson.EncodeString(data)
	if e != nil {
		return e
	}
	return gfile.PutContents(path, jsonCount)
}