// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build !ent

package sqlx

import (
	"github.com/anthony-bible/atlas/sql/schema"
)

// sortViewChanges is an optional function to sort to views by their dependencies.
func sortViewChanges(changes []schema.Change) ([]schema.Change, error) {
	return changes, nil // unimplemented.
}

func (*Diff) triggerDiff(_, _ interface {
	Trigger(string) (*schema.Trigger, bool)
}, _, _ []*schema.Trigger, _ *schema.DiffOptions) ([]schema.Change, error) {
	return nil, nil // unimplemented.
}

// funcDep returns true if f1 depends on f2.
func funcDep(_, _ *schema.Func) bool {
	return false // unimplemented.
}

// procDep returns true if p1 depends on p2.
func procDep(_, _ *schema.Proc) bool {
	return false // unimplemented.
}
