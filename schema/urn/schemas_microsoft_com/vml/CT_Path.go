// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/zzhtl/goxml/schema/soo/ofc/sharedTypes"
)

type CT_Path struct {
	VAttr               *string
	LimoAttr            *string
	TextboxrectAttr     *string
	FillokAttr          sharedTypes.ST_TrueFalse
	StrokeokAttr        sharedTypes.ST_TrueFalse
	ShadowokAttr        sharedTypes.ST_TrueFalse
	ArrowokAttr         sharedTypes.ST_TrueFalse
	GradientshapeokAttr sharedTypes.ST_TrueFalse
	TextpathokAttr      sharedTypes.ST_TrueFalse
	InsetpenokAttr      sharedTypes.ST_TrueFalse
	ConnecttypeAttr     OfcST_ConnectType
	ConnectlocsAttr     *string
	ConnectanglesAttr   *string
	ExtrusionokAttr     sharedTypes.ST_TrueFalse
	IdAttr              *string
}

func NewCT_Path() *CT_Path {
	ret := &CT_Path{}
	return ret
}

func (m *CT_Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.VAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"},
			Value: fmt.Sprintf("%v", *m.VAttr)})
	}
	if m.LimoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "limo"},
			Value: fmt.Sprintf("%v", *m.LimoAttr)})
	}
	if m.TextboxrectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "textboxrect"},
			Value: fmt.Sprintf("%v", *m.TextboxrectAttr)})
	}
	if m.FillokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.FillokAttr.MarshalXMLAttr(xml.Name{Local: "fillok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StrokeokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.StrokeokAttr.MarshalXMLAttr(xml.Name{Local: "strokeok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShadowokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ShadowokAttr.MarshalXMLAttr(xml.Name{Local: "shadowok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ArrowokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ArrowokAttr.MarshalXMLAttr(xml.Name{Local: "arrowok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.GradientshapeokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.GradientshapeokAttr.MarshalXMLAttr(xml.Name{Local: "gradientshapeok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TextpathokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.TextpathokAttr.MarshalXMLAttr(xml.Name{Local: "textpathok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.InsetpenokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.InsetpenokAttr.MarshalXMLAttr(xml.Name{Local: "insetpenok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ConnecttypeAttr != OfcST_ConnectTypeUnset {
		attr, err := m.ConnecttypeAttr.MarshalXMLAttr(xml.Name{Local: "connecttype"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ConnectlocsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:connectlocs"},
			Value: fmt.Sprintf("%v", *m.ConnectlocsAttr)})
	}
	if m.ConnectanglesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:connectangles"},
			Value: fmt.Sprintf("%v", *m.ConnectanglesAttr)})
	}
	if m.ExtrusionokAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ExtrusionokAttr.MarshalXMLAttr(xml.Name{Local: "extrusionok"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Path) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "connecttype" {
			m.ConnecttypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "extrusionok" {
			m.ExtrusionokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "connectangles" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ConnectanglesAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "connectlocs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ConnectlocsAttr = &parsed
			continue
		}
		if attr.Name.Local == "gradientshapeok" {
			m.GradientshapeokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "shadowok" {
			m.ShadowokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "arrowok" {
			m.ArrowokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "v" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VAttr = &parsed
			continue
		}
		if attr.Name.Local == "textpathok" {
			m.TextpathokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "insetpenok" {
			m.InsetpenokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "strokeok" {
			m.StrokeokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fillok" {
			m.FillokAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "textboxrect" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TextboxrectAttr = &parsed
			continue
		}
		if attr.Name.Local == "limo" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LimoAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Path: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Path and its children
func (m *CT_Path) Validate() error {
	return m.ValidateWithPath("CT_Path")
}

// ValidateWithPath validates the CT_Path and its children, prefixing error messages with path
func (m *CT_Path) ValidateWithPath(path string) error {
	if err := m.FillokAttr.ValidateWithPath(path + "/FillokAttr"); err != nil {
		return err
	}
	if err := m.StrokeokAttr.ValidateWithPath(path + "/StrokeokAttr"); err != nil {
		return err
	}
	if err := m.ShadowokAttr.ValidateWithPath(path + "/ShadowokAttr"); err != nil {
		return err
	}
	if err := m.ArrowokAttr.ValidateWithPath(path + "/ArrowokAttr"); err != nil {
		return err
	}
	if err := m.GradientshapeokAttr.ValidateWithPath(path + "/GradientshapeokAttr"); err != nil {
		return err
	}
	if err := m.TextpathokAttr.ValidateWithPath(path + "/TextpathokAttr"); err != nil {
		return err
	}
	if err := m.InsetpenokAttr.ValidateWithPath(path + "/InsetpenokAttr"); err != nil {
		return err
	}
	if err := m.ConnecttypeAttr.ValidateWithPath(path + "/ConnecttypeAttr"); err != nil {
		return err
	}
	if err := m.ExtrusionokAttr.ValidateWithPath(path + "/ExtrusionokAttr"); err != nil {
		return err
	}
	return nil
}
