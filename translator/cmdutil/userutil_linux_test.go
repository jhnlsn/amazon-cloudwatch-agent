// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

// +build linux

package cmdutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetGroupIds(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "group-file-test")
	require.Nil(t, err, "Failed to create temp file")

	defer os.Remove(tmpfile.Name())
	fmt.Fprintln(tmpfile, "root:x:0:")
	fmt.Fprintln(tmpfile, "bin:x:1:")
	fmt.Fprintln(tmpfile, "adm:x:4:root,other,test-user")
	fmt.Fprintln(tmpfile, "wheel:x:10:test-user")
	fmt.Fprintln(tmpfile, "tst:x:50:root,test-user,test")
	fmt.Fprintln(tmpfile, "test:x:100:other,test-user-2")

	gids, err := getGroupIds("test-user", tmpfile.Name())
	require.Nil(t, err, "Failed to retrieve group IDs for user: test-user")
	assert.Equal(t, []int{4, 10, 50}, gids)
	gids, err = getGroupIds("not-in-file", tmpfile.Name())
	require.Nil(t, err, "Failed to retrieve group IDs for user: not-in-file")
	assert.Len(t, gids, 0)
}
