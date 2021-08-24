/**
defer延迟调用，先进后出(stack)
 */

package main

import (
	"errors"
	"fmt"
)

type cc func() error

func c1(e error) cc {
	return func() (err error) {
		// 此处defer中的err为实时计算，所以err一直为空
		defer fmt.Println("c1 -->", err)
		return e
	}
}

func c2(e error) cc {
	return func() (err error) {
		// 此处defer为闭包函数，内部的err值会重新计算，最后等于返回的e
		defer func() {
			fmt.Println("c2 -->", err)
		}()
		return e
	}
}

func c3(e error) cc {
	return func() (err error) {
		// 此处defer函数的传参值为实时计算，err为nil，效果同c1
		defer func(err error) {
			fmt.Println("c3 -->", err)
		}(err)
		return e
	}
}

type ccp func() *error

func c2p(e error) ccp {
	return func() (err *error) {
		// 此处defer为闭包函数，内部的err值会重新计算，最后等于返回的e
		defer func() {
			fmt.Println("c2p -->", err, *err)
		}()
		return &e
	}
}

func main() {
	// c1 --> <nil>
	// c1 --> <nil>
	// c2 --> c2-error
	// c2 --> <nil>

	c1(errors.New("c1-error"))()
	c1(nil)()

	c2(errors.New("c2-error"))()
	c2(nil)()

	c3(errors.New("c3-error"))()
	c3(nil)()

	c2p(errors.New("c2p-error"))()
	c2p(nil)()
}
