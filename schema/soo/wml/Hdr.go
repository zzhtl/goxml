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

	"github.com/zzhtl/goxml"
	"github.com/zzhtl/goxml/schema/soo/ofc/math"
)

type Hdr struct {
	CT_HdrFtr
}

func NewHdr() *Hdr {
	ret := &Hdr{}
	ret.CT_HdrFtr = *NewCT_HdrFtr()
	return ret
}

func (m *Hdr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:hdr"
	return m.CT_HdrFtr.MarshalXML(e, start)
}

func (m *Hdr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_HdrFtr = *NewCT_HdrFtr()
lHdr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "altChunk"}:
				tmp := NewCT_AltChunk()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AltChunk = append(m.AltChunk, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXml"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.CustomXml = NewCT_CustomXmlBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdt"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmpcontentblockcontent.Sdt = NewCT_SdtBlock()
				if err := d.DecodeElement(tmpcontentblockcontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "p"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_P()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.P = append(tmpcontentblockcontent.P, tmp)
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tbl"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmp := NewCT_Tbl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentblockcontent.Tbl = append(tmpcontentblockcontent.Tbl, tmp)
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofErr"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "del"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFrom"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveTo"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeStart"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeEnd"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMathPara"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMath"}:
				tmpcontentblockcontent := NewEG_ContentBlockContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_ContentBlockContent = append(m.EG_ContentBlockContent, tmpcontentblockcontent)
				tmpcontentblockcontent.EG_RunLevelElts = append(tmpcontentblockcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				gooxml.Log("skipping unsupported element on Hdr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lHdr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Hdr and its children
func (m *Hdr) Validate() error {
	return m.ValidateWithPath("Hdr")
}

// ValidateWithPath validates the Hdr and its children, prefixing error messages with path
func (m *Hdr) ValidateWithPath(path string) error {
	if err := m.CT_HdrFtr.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
