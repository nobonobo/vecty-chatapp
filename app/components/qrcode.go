package components

import (
	"bytes"
	"log"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"

	"github.com/aaronarduino/goqrsvg"
	"github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
)

// QRCode ...
type QRCode struct {
	vecty.Core
	CellSize int    `vecty:"prop"`
	Text     string `vecty:"prop"`
}

// Render ...
func (c *QRCode) Render() vecty.ComponentOrHTML {
	sz := 8
	if c.CellSize > 0 {
		sz = c.CellSize
	}
	code, err := qr.Encode(c.Text, qr.M, qr.Auto)
	if err != nil {
		log.Println(err)
		return nil
	}
	buff := bytes.NewBuffer(nil)
	g := svg.New(buff)
	qs := goqrsvg.NewQrSVG(code, sz)
	qs.StartQrSVG(g)
	qs.WriteQrSVG(g)
	g.End()
	return elem.Div(
		vecty.Markup(vecty.UnsafeHTML(buff.String())),
	)
}
