package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"monkey", }, `{"case_insensitive":true,"keywords":{"keyword":"public private property continue exit extern new try catch eachin not abstract final select case default const local global field end if then else elseif endif while wend repeat until forever for to step next return module inline throw import","built_in":"DebugLog DebugStop Error Print ACos ACosr ASin ASinr ATan ATan2 ATan2r ATanr Abs Abs Ceil Clamp Clamp Cos Cosr Exp Floor Log Max Max Min Min Pow Sgn Sgn Sin Sinr Sqrt Tan Tanr Seed PI HALFPI TWOPI","literal":"true false null and or shl shr mod"},"illegal":"\\/\\*","contains":[{"className":"comment","begin":"#rem","end":"#end","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"'","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"className":"function","beginKeywords":"function method","end":"[(=:]|$","illegal":"\\n","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"className":"class","beginKeywords":"class interface","end":"$","contains":[{"beginKeywords":"extends implements"},{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"className":"built_in","begin":"\\b(self|super)\\b"},{"className":"meta","begin":"\\s*#","end":"$","keywords":{"meta-keyword":"if else elseif endif end then"}},{"className":"meta","begin":"^\\s*strict\\b"},{"beginKeywords":"alias","end":"=","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","relevance":0,"variants":[{"begin":"[$][a-fA-F0-9]+"},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0}]}]}`)
}