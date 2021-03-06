// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"log"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebookincubator/ent-contrib/entgql"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Header: `
// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.`,
		Templates: entgql.AllTemplates,
		Features: []gen.Feature{
			gen.FeaturePrivacy,
			gen.FeatureSnapshot,
		},
	}, entc.TemplateDir("./template"))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
