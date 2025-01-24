// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package core_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/zzhtl/goxml/schema/soo/pkg/metadata/core_properties"
)

func TestCT_KeywordConstructor(t *testing.T) {
	v := core_properties.NewCT_Keyword()
	if v == nil {
		t.Errorf("core_properties.NewCT_Keyword must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed core_properties.CT_Keyword should validate: %s", err)
	}
}

func TestCT_KeywordMarshalUnmarshal(t *testing.T) {
	v := core_properties.NewCT_Keyword()
	buf, _ := xml.Marshal(v)
	v2 := core_properties.NewCT_Keyword()
	xml.Unmarshal(buf, v2)
}
