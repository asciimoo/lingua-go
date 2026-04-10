package es
import (
	"embed"
	"github.com/asciimoo/lingua-go"
)
//go:embed *.zip
var model embed.FS
func init() {
	lingua.Register("es", model)
}
