package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"leaf", }, `{"contains":[{"className":"function","begin":"#+[A-Za-z_0-9]*\\(","end":" {","returnBegin":true,"excludeEnd":true,"contains":[{"className":"keyword","begin":"#+"},{"className":"title","begin":"[A-Za-z_][A-Za-z_0-9]*"},{"className":"params","begin":"\\(","end":"\\)","endsParent":true,"contains":[{"className":"string","begin":"\"","end":"\""},{"className":"variable","begin":"[A-Za-z_][A-Za-z_0-9]*"}]}]}]}`)
}