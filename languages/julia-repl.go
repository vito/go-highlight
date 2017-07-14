package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"julia-repl", }, `{"contains":[{"className":"meta","begin":"^julia>","relevance":10,"starts":{"end":"^(?![ ]{6})","subLanguage":["julia"]},"aliases":["jldoctest"]}]}`)
}