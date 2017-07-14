package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"shell", "console"}, `{"aliases":["console"],"contains":[{"className":"meta","begin":"^\\s{0,3}[\\w\\d\\[\\]()@-]*[>%$#]","starts":{"end":"$","subLanguage":["bash"]}}]}`)
}