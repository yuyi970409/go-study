/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co.,Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package general

import (
	"ShopApi/general/errcode"
)

type ErrorResp struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

type Resp struct {
	Code int `json:"status"`
}

type DataResp struct {
	Code int         `json:"status"`
	Data interface{} `json:"data"`
}

type ProductListResp struct {
	Code        int `json:"status"`
	ProductList `json:"data"`
}

type ProductList struct {
	Header interface{} `json:"adPics"`
	Image  interface{} `json:"wares"`
}

func NewErrorWithMessage(code int, msg string) *ErrorResp {
	if code == errcode.ErrSucceed {
		msg = ""
	}

	return &ErrorResp{
		Code:    code,
		Message: msg,
	}
}

func (this *ErrorResp) Error() string {
	return this.Message
}

func NewMessage(code int) *Resp {
	return &Resp{
		Code: code,
	}
}

func NewMessageWithData(code int, data interface{}) *DataResp {
	return &DataResp{
		Code: code,
		Data: data,
	}
}

func NewMessageForProductList(code int, header, image interface{}) *ProductListResp {
	return &ProductListResp{
		Code: code,
		ProductList: ProductList{
			Header: header,
			Image:  image,
		},
	}
}
