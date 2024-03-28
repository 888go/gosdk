
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package godebug makes the settings in the $GODEBUG environment variable
// available to other packages. These settings are often used for compatibility
// tweaks, when we need to change a default behavior but want to let users
// opt back in to the original. For example GODEBUG=http2server=0 disables
// HTTP/2 support in the net/http server.
//
// In typical usage, code should declare a Setting as a global
// and then call Value each time the current setting value is needed:
//
//	var http2server = godebug.New("http2server")
//
//	func ServeConn(c net.Conn) {
//		if http2server.Value() == "0" {
//			disallow HTTP/2
//			...
//		}
//		...
//	}
<原文结束>

# <翻译开始>
// Package godebug makes the settings in the $GODEBUG environment variable
// available to other packages. These settings are often used for compatibility
// tweaks, when we need to change a default behavior but want to let users
// opt back in to the original. For example GODEBUG=http2server=0 disables
// HTTP/2 support in the net/http server.
//
// In typical usage, code should declare a Setting as a global
// and then call Value each time the current setting value is needed:
//
//	var http2server = godebug.New("http2server")
//
//	func ServeConn(c net.Conn) {
//		if http2server.Value() == "0" {
//			disallow HTTP/2
//			...
//		}
//		...
//	}
# <翻译结束>


<原文开始>
// A Setting is a single setting in the $GODEBUG environment variable.
<原文结束>

# <翻译开始>
// A Setting is a single setting in the $GODEBUG environment variable.
# <翻译结束>


<原文开始>
// New returns a new Setting for the $GODEBUG setting with the given name.
<原文结束>

# <翻译开始>
// New returns a new Setting for the $GODEBUG setting with the given name.
# <翻译结束>


<原文开始>
// Name returns the name of the setting.
<原文结束>

# <翻译开始>
// Name returns the name of the setting.
# <翻译结束>


<原文开始>
// String returns a printable form for the setting: name=value.
<原文结束>

# <翻译开始>
// String returns a printable form for the setting: name=value.
# <翻译结束>


<原文开始>
// cache is a cache of all the GODEBUG settings,
// a locked map[string]*atomic.Pointer[string].
//
// All Settings with the same name share a single
// *atomic.Pointer[string], so that when GODEBUG
// changes only that single atomic string pointer
// needs to be updated.
//
// A name appears in the values map either if it is the
// name of a Setting for which Value has been called
// at least once, or if the name has ever appeared in
// a name=value pair in the $GODEBUG environment variable.
// Once entered into the map, the name is never removed.
<原文结束>

# <翻译开始>
// cache is a cache of all the GODEBUG settings,
// a locked map[string]*atomic.Pointer[string].
//
// All Settings with the same name share a single
// *atomic.Pointer[string], so that when GODEBUG
// changes only that single atomic string pointer
// needs to be updated.
//
// A name appears in the values map either if it is the
// name of a Setting for which Value has been called
// at least once, or if the name has ever appeared in
// a name=value pair in the $GODEBUG environment variable.
// Once entered into the map, the name is never removed.
# <翻译结束>


<原文开始>
// name string -> value *atomic.Pointer[string]
<原文结束>

# <翻译开始>
// name string -> value *atomic.Pointer[string]
# <翻译结束>


<原文开始>
// Value returns the current value for the GODEBUG setting s.
//
// Value maintains an internal cache that is synchronized
// with changes to the $GODEBUG environment variable,
// making Value efficient to call as frequently as needed.
// Clients should therefore typically not attempt their own
// caching of Value's result.
<原文结束>

# <翻译开始>
// Value returns the current value for the GODEBUG setting s.
//
// Value maintains an internal cache that is synchronized
// with changes to the $GODEBUG environment variable,
// making Value efficient to call as frequently as needed.
// Clients should therefore typically not attempt their own
// caching of Value's result.
# <翻译结束>


<原文开始>
// update records an updated GODEBUG setting.
// def is the default GODEBUG setting for the running binary,
// and env is the current value of the $GODEBUG environment variable.
<原文结束>

# <翻译开始>
// update records an updated GODEBUG setting.
// def is the default GODEBUG setting for the running binary,
// and env is the current value of the $GODEBUG environment variable.
# <翻译结束>


<原文开始>
	// Update all the cached values, creating new ones as needed.
	// We parse the environment variable first, so that any settings it has
	// are already locked in place (did[name] = true) before we consider
	// the defaults.
<原文结束>

# <翻译开始>
	// Update all the cached values, creating new ones as needed.
	// We parse the environment variable first, so that any settings it has
	// are already locked in place (did[name] = true) before we consider
	// the defaults.
# <翻译结束>


<原文开始>
// Clear any cached values that are no longer present.
<原文结束>

# <翻译开始>
// Clear any cached values that are no longer present.
# <翻译结束>


<原文开始>
// parse parses the GODEBUG setting string s,
// which has the form k=v,k2=v2,k3=v3.
// Later settings override earlier ones.
// Parse only updates settings k=v for which did[k] = false.
// It also sets did[k] = true for settings that it updates.
<原文结束>

# <翻译开始>
// parse parses the GODEBUG setting string s,
// which has the form k=v,k2=v2,k3=v3.
// Later settings override earlier ones.
// Parse only updates settings k=v for which did[k] = false.
// It also sets did[k] = true for settings that it updates.
# <翻译结束>


<原文开始>
	// Scan the string backward so that later settings are used
	// and earlier settings are ignored.
	// Note that a forward scan would cause cached values
	// to temporarily use the ignored value before being
	// updated to the "correct" one.
<原文结束>

# <翻译开始>
	// Scan the string backward so that later settings are used
	// and earlier settings are ignored.
	// Note that a forward scan would cause cached values
	// to temporarily use the ignored value before being
	// updated to the "correct" one.
# <翻译结束>

