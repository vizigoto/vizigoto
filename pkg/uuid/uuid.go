// BSD 3-Clause License
//
// Copyright (c) 2017, vizigoto
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of the copyright holder nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package uuid provides implementation of Universally Unique Identifier (UUID).
//
// This package only supports the type 4 (as specified in RFC 4122).
package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

// New creates a new uuid version 4 based on random numbers (see RFC 4122)
func New() string {
	bytes := make([]byte, 16)
	safeRandom(bytes)
	bytes[6] = (4 << 4) | (bytes[6] & 0xf)
	bytes[8] = bytes[8] & 0x3f
	bytes[8] = bytes[8] | 0x80
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], bytes[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], bytes[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], bytes[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], bytes[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], bytes[10:])
	return string(buf)
}

func safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}
