// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs01000104 struct {
	Xmlns     string                             `xml:"xmlns,attr"`
	FIDrctDbt FinancialInstitutionDirectDebitV04 `xml:"FIDrctDbt"`
}

func (doc DocumentPacs01000104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs01000104) NameSpace() string {
	return utils.DocumentPacs01000104NameSpace
}

func (doc DocumentPacs01000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIDrctDbt FinancialInstitutionDirectDebitV04 `xml:"FIDrctDbt"`
	}
	output.FIDrctDbt = doc.FIDrctDbt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPacs02800104 struct {
	Xmlns           string                        `xml:"xmlns,attr"`
	FIToFIPmtStsReq FIToFIPaymentStatusRequestV04 `xml:"FIToFIPmtStsReq"`
}

func (doc DocumentPacs02800104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs02800104) NameSpace() string {
	return utils.DocumentPacs02800104NameSpace
}

func (doc DocumentPacs02800104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtStsReq FIToFIPaymentStatusRequestV04 `xml:"FIToFIPmtStsReq"`
	}
	output.FIToFIPmtStsReq = doc.FIToFIPmtStsReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
