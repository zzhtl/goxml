// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/zzhtl/goxml"
)

type CT_TrPrChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr int64
	TrPr   *CT_TrPrBase
}

func NewCT_TrPrChange() *CT_TrPrChange {
	ret := &CT_TrPrChange{}
	ret.TrPr = NewCT_TrPrBase()
	return ret
}

func (m *CT_TrPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	setrPr := xml.StartElement{Name: xml.Name{Local: "w:trPr"}}
	e.EncodeElement(m.TrPr, setrPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TrPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TrPr = NewCT_TrPrBase()
	for _, attr := range start.Attr {
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
			continue
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_TrPrChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "trPr"}:
				if err := d.DecodeElement(m.TrPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TrPrChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TrPrChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TrPrChange and its children
func (m *CT_TrPrChange) Validate() error {
	return m.ValidateWithPath("CT_TrPrChange")
}

// ValidateWithPath validates the CT_TrPrChange and its children, prefixing error messages with path
func (m *CT_TrPrChange) ValidateWithPath(path string) error {
	if err := m.TrPr.ValidateWithPath(path + "/TrPr"); err != nil {
		return err
	}
	return nil
}
