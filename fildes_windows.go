// Copyright (C) 2018-2021 Andrew Colin Kissa <andrew@datopdog.io>
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package clamd Golang Clamd client
Clamd - Golang clamd client
*/
package goclamd

import (
	"errors"
	"net"
	"net/textproto"
)

const fildesUnsupportErr = "Fildes is not supported"

func (c *Client) fildesScan(tc *textproto.Conn, conn net.Conn, p string) (err error) {
	return errors.New(fildesUnsupportErr)
}
