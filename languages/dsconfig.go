package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"dsconfig", }, `{"keywords":"dsconfig","contains":[{"className":"keyword","begin":"^dsconfig","end":"\\s","excludeEnd":true,"relevance":10},{"className":"built_in","begin":"(list|create|get|set|delete)-(\\w+)","end":"\\s","excludeEnd":true,"illegal":"!@#$%^&*()","relevance":10},{"className":"built_in","begin":"--(\\w+)","end":"\\s","excludeEnd":true},{"className":"string","begin":"\"","end":"\""},{"className":"string","begin":"'","end":"'"},{"className":"string","begin":"[\\w-?]+:\\w+","end":"\\W","relevance":0},{"className":"string","begin":"\\w+-?\\w+","end":"\\W","relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}`)
}