/*
 * @Author: Go不浪队
 * @Date: 2023-02-07 00:51:13
 * @LastEditTime: 2023-02-07 01:42:32
 * @Description:
 */
package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Context struct {
	Transaction      *gorm.DB
	context context.Context
	cancel  context.CancelFunc

	commit bool
}

func GetContext() *Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	c := &Context{
		Transaction:      beginTransactionWithCtx(ctx),
		context: ctx,
		cancel:  cancel,
		commit:   false,
	}

	return c
}