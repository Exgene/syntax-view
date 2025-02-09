package generate

import (
	"image"
	"image/png"
	"os"

	"github.com/jung-kurt/gofpdf"
)

type MDGenerator struct {
	md string
}

type PDFGenerator struct {
	pdf *gofpdf.Fpdf
}

func NewPDFGenerator() *PDFGenerator {
	pdf := gofpdf.New("P", "mm", "A4", "")
	return &PDFGenerator{pdf: pdf}
}

func NewMDGenerator(outputFile string) *MDGenerator {
	md := outputFile
	return &MDGenerator{md: md}
}

func (g *PDFGenerator) AddPage(filepath string, img image.Image) error {
	g.pdf.AddPage()

	g.pdf.SetFont("Arial", "B", 10)
	g.pdf.Cell(0, 8, filepath)
	g.pdf.Ln(10)

	tmpFile, err := saveImageToTemp(img)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	pageWidth, pageHeight := g.pdf.GetPageSize()
	marginX := float64(20)
	marginY := float64(20)
	availableWidth := pageWidth - (2 * marginX)
	availableHeight := pageHeight - (2 * marginY) - 10 // subtract header height

	imgWidth := availableWidth
	ratio := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	imgHeight := imgWidth * ratio

	// If image is too tall, scale it down
	if imgHeight > availableHeight {
		imgHeight = availableHeight
		imgWidth = imgHeight / ratio
	}

	// Center the image horizontally
	x := marginX + (availableWidth-imgWidth)/2

	g.pdf.Image(tmpFile, x, g.pdf.GetY(), imgWidth, imgHeight, false, "", 0, "")

	return nil
}

func (g *PDFGenerator) Save(outputPath string) error {
	return g.pdf.OutputFileAndClose(outputPath)
}

func saveImageToTemp(img image.Image) (string, error) {
	tmpFile, err := os.CreateTemp("", "img-*.png")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	if err := png.Encode(tmpFile, img); err != nil {
		os.Remove(tmpFile.Name())
		return "", err
	}

	return tmpFile.Name(), nil
}
