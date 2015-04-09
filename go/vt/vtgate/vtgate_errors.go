// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vtgate

// TODO(aaijazi): find the right home for this, and the right fields. Tighten it up!
import "fmt"

type vtGateError struct {
	message string
	err     error
}

func (e *vtGateError) Error() string {
	return fmt.Sprintf("%s", e.err)
}

func asVTGateError(in error) error {
	if in == nil {
		return nil
	}

	// If the input error is already a VTGateError, don't touch it. This makes
	// asVTGateError idempotent, and means we dn't have to deal with deep recursion
	// to extract the root cause error.
	vtgErr, ok := in.(*vtGateError)
	if ok {
		return in
	}

	return &vtGateError{err: in}
}

func isErrorCausedByVTGate(in error) bool {
	vtgErr, ok := e.(*vtGateError)
	if !ok {
		// We don't have any structured information about this error, so we assume
		// the worst (that this is caused by VTGate)
		return true
	}
	innerErr = vtgErr.err
	shardConnErr, ok := innerErr.(*ShardConnError)
	if ok {
		return handleError(shardConnErr.Err)
	}
	return handleError(innerErr)
}

func handleError(e error) bool {
	switch innerErr := innerErr.(type) {
	case ShardConnError:
		// TODO
		return false
	default:
		// We assume that VTGate is guilty unless proven innocent (of having produced this error)
		return true
	}
}

// ShardConnError is the shard conn specific error.
type ShardConnError struct {
	Code            int
	ShardIdentifier string
	InTransaction   bool
	// Preserve the original error
	Err error
}

func (e *ShardConnError) Error() string {
	if e.ShardIdentifier == "" {
		return e.Err.Error()
	}
	return fmt.Sprintf("%v on (shard, host): %s", e.Err, e.ShardIdentifier)
}
