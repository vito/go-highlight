package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"erb", }, `{"subLanguage":["xml"],"contains":[{"className":"comment","begin":"<%#","end":"%>","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"begin":"<%[%=-]?","end":"[%-]?%>","subLanguage":["ruby"],"excludeBegin":true,"excludeEnd":true}]}`)
}