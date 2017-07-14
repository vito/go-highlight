package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"awk", }, `{"keywords":{"keyword":"BEGIN END if else while do for in break continue delete next nextfile function func exit|10"},"contains":[{"className":"variable","variants":[{"begin":"\\$[\\w\\d#@][\\w\\d_]*"},{"begin":"\\$\\{(.*?)}"}]},{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}],"variants":[{"begin":"(u|b)?r?'''","end":"'''","relevance":10},{"begin":"(u|b)?r?\"\"\"","end":"\"\"\"","relevance":10},{"begin":"(u|r|ur)'","end":"'","relevance":10},{"begin":"(u|r|ur)\"","end":"\"","relevance":10},{"begin":"(b|br)'","end":"'"},{"begin":"(b|br)\"","end":"\""},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0}]}`)
}