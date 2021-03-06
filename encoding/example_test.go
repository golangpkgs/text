// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encoding_test

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golangpkgs/text/encoding"
	"github.com/golangpkgs/text/encoding/charmap"
	"github.com/golangpkgs/text/transform"
)

func ExampleDecodeWindows1252() {
	sr := strings.NewReader("Gar\xe7on !")
	tr := charmap.Windows1252.NewDecoder().Reader(sr)
	io.Copy(os.Stdout, tr)
	// Output: Garçon !
}

func ExampleUTF8Validator() {
	for i := 0; i < 2; i++ {
		var transformer transform.Transformer
		transformer = charmap.Windows1252.NewEncoder()
		if i == 1 {
			transformer = transform.Chain(encoding.UTF8Validator, transformer)
		}
		dst := make([]byte, 256)
		src := []byte("abc\xffxyz") // src is invalid UTF-8.
		nDst, nSrc, err := transformer.Transform(dst, src, true)
		fmt.Printf("i=%d: produced %q, consumed %q, error %v\n",
			i, dst[:nDst], src[:nSrc], err)
	}
	// Output:
	// i=0: produced "abc\x1axyz", consumed "abc\xffxyz", error <nil>
	// i=1: produced "abc", consumed "abc", error encoding: invalid UTF-8
}
