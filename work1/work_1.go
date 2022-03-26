package main

/*
1. 我们在数据库操作的时候，比如 dao
层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

应该抛给上层，日志应该只处理一次，
如果不用Wrap在各个地方打日志，会导致日志的不连续与重复性工作

*/

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func dao() error {
	return errors.Wrap(sql.ErrNoRows, "dao failed")
}

func method() error {
	return errors.WithMessage(dao(), "method")
}

func main() {
	err := method()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n %+v\n", err)
	}
}
