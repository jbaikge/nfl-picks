// Copyright 2013 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCompressHandler(t *testing.T) {
	w := httptest.NewRecorder()
	CompressHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 1024; i++ {
			io.WriteString(w, "Gorilla!\n")
		}
	})).ServeHTTP(w, &http.Request{
		Method: "GET",
		Header: http.Header{
			"Accept-Encoding": []string{"gzip"},
		},
	})

	if w.HeaderMap.Get("Content-Encoding") != "gzip" {
		t.Fatalf("wrong content encoding, got %d want %d", w.HeaderMap.Get("Content-Encoding"), "gzip")
	}
	if w.HeaderMap.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Fatalf("wrong content type, got %d want %d", w.HeaderMap.Get("Content-Type"), "text/plain; charset=utf-8")
	}
	if w.Body.Len() != 72 {
		t.Fatalf("wrong len, got %d want %d", w.Body.Len(), 72)
	}
}
