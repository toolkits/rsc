// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package keychain implements access to the passwords and other keys
// stored in the system-provided keychain.
package keychain

// BUG(rsc): Package keychain is only implemented on OS X.

// UserPasswd returns the user name and password for authenticating
// to the named server.  If the user argument is non-empty, UserPasswd
// restricts its search to passwords for the named user.
func UserPasswd(server, preferredUser string) (user, passwd string, err error) {
	return userPasswd(server, preferredUser)
}

