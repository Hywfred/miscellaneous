package util

import (
	`testing`
)

func TestFindDir(t *testing.T) {
	dir, err := FindDir("hello")
	CheckErr(err)
	if dir != "D:\\Git\\miscellaneous\\util\\test\\hello" {
		t.Errorf("dir is : %v\n", dir)
	}
}
