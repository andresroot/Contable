package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"path/filepath"
)

// ESTRUCTURA PDF
const (
	cnGofpdfDir = "."
	cnFontDir   = cnGofpdfDir + "/font"
	cnImgDir    = cnGofpdfDir + "/image"
	cnTextDir   = cnGofpdfDir + "/text"
)
type SomeStruct struct {
	SomeValue bool `json:"result"`
}
type nullWriter struct {
}

func (nw *nullWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	return
}

func (nw *nullWriter) Close() (err error) {
	return
}

type pdfWriter struct {
	pdf *gofpdf.Fpdf
	fl  *os.File
	idx int
}

func (pw *pdfWriter) Write(p []byte) (n int, err error) {
	if pw.pdf.Ok() {
		return pw.fl.Write(p)
	}
	return
}

func (pw *pdfWriter) Close() (err error) {
	if pw.fl != nil {
		pw.fl.Close()
		pw.fl = nil
	}
	if pw.pdf.Ok() {
		fmt.Printf("Successfully generated pdf/tutorial%02d.pdf\n", pw.idx)
	} else {
		fmt.Printf("%s\n", pw.pdf.Error())
	}
	return
}

func docWriter(pdf *gofpdf.Fpdf, idx int) *pdfWriter {
	pw := new(pdfWriter)
	pw.pdf = pdf
	pw.idx = idx
	if pdf.Ok() {
		var err error
		fileStr := fmt.Sprintf("%s/pdf/tutorial%02d.pdf", cnGofpdfDir, idx)
		pw.fl, err = os.Create(fileStr)
		if err != nil {
			pdf.SetErrorf("Error opening output file %s", fileStr)
		}
	}
	return pw
}

func imageFile(fileStr string) string {
	return filepath.Join(cnImgDir, fileStr)
}

func fontFile(fileStr string) string {
	return filepath.Join(cnFontDir, fileStr)
}

func textFile(fileStr string) string {
	return filepath.Join(cnTextDir, fileStr)
}

