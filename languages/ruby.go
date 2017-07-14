package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"ruby", "rb", "gemspec", "podspec", "thor", "irb"}, `{"aliases":["rb","gemspec","podspec","thor","irb"],"keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"illegal":"\\/\\*","contains":[{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"begin":"^\\s*=>","starts":{"end":"$","contains":[{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"variants":[{"begin":"'","end":"'"},{"begin":"\"","end":"\""},{"begin":"`+"`"+`","end":"`+"`"+`"},{"begin":"%[qQwWx]?\\(","end":"\\)"},{"begin":"%[qQwWx]?\\[","end":"\\]"},{"begin":"%[qQwWx]?{","end":"}"},{"begin":"%[qQwWx]?<","end":">"},{"begin":"%[qQwWx]?/","end":"/"},{"begin":"%[qQwWx]?%","end":"%"},{"begin":"%[qQwWx]?-","end":"-"},{"begin":"%[qQwWx]?\\|","end":"\\|"},{"begin":"\\B\\?(\\\\\\d{1,3}|\\\\x[A-Fa-f0-9]{1,2}|\\\\u[A-Fa-f0-9]{4}|\\\\?\\S)\\b"},{"begin":"<<(-?)\\w+$","end":"^\\s*\\w+$"}]},{"begin":"#<","end":">"},{"className":"class","beginKeywords":"class module","end":"$|;","illegal":"=","contains":[{"className":"title","begin":"[A-Za-z_]\\w*(::\\w+)*(\\?|\\!)?","relevance":0},{"begin":"<\\s*","contains":[{"begin":"([a-zA-Z]\\w*::)?[a-zA-Z]\\w*"}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"className":"function","beginKeywords":"def","end":"$|;","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*[!?=]?|[-+~]\\@|<<|>>|=~|===?|<=>|[<>]=?|\\*\\*|[-/+%^&*~`+"`"+`|]|\\[\\]=?","relevance":0},{"className":"params","begin":"\\(","end":"\\)","endsParent":true,"keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"[a-zA-Z]\\w*::"},{"className":"symbol","begin":"[a-zA-Z_]\\w*(\\!|\\?)?:","relevance":0},{"className":"symbol","begin":":(?!\\s)","contains":[{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"variants":[{"begin":"'","end":"'"},{"begin":"\"","end":"\""},{"begin":"`+"`"+`","end":"`+"`"+`"},{"begin":"%[qQwWx]?\\(","end":"\\)"},{"begin":"%[qQwWx]?\\[","end":"\\]"},{"begin":"%[qQwWx]?{","end":"}"},{"begin":"%[qQwWx]?<","end":">"},{"begin":"%[qQwWx]?/","end":"/"},{"begin":"%[qQwWx]?%","end":"%"},{"begin":"%[qQwWx]?-","end":"-"},{"begin":"%[qQwWx]?\\|","end":"\\|"},{"begin":"\\B\\?(\\\\\\d{1,3}|\\\\x[A-Fa-f0-9]{1,2}|\\\\u[A-Fa-f0-9]{4}|\\\\?\\S)\\b"},{"begin":"<<(-?)\\w+$","end":"^\\s*\\w+$"}]},{"begin":"[a-zA-Z_]\\w*[!?=]?|[-+~]\\@|<<|>>|=~|===?|<=>|[<>]=?|\\*\\*|[-/+%^&*~`+"`"+`|]|\\[\\]=?"}],"relevance":0},{"className":"number","begin":"(\\b0[0-7_]+)|(\\b0x[0-9a-fA-F_]+)|(\\b[1-9][0-9_]*(\\.[0-9_]+)?)|[0_]\\b","relevance":0},{"begin":"(\\$\\W)|((\\$|\\@\\@?)(\\w+))"},{"className":"params","begin":"\\|","end":"\\|","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"}},{"begin":"(!|!=|!==|%|%=|&|&&|&=|\\*|\\*=|\\+|\\+=|,|-|-=|/=|/|:|;|<<|<<=|<=|<|===|==|=|>>>=|>>=|>=|>>>|>>|>|\\?|\\[|\\{|\\(|\\^|\\^=|\\||\\|=|\\|\\||~|unless)\\s*","keywords":"unless","contains":[{"begin":"#<","end":">"},{"className":"regexp","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"illegal":"\\n","variants":[{"begin":"/","end":"/[a-z]*"},{"begin":"%r{","end":"}[a-z]*"},{"begin":"%r\\(","end":"\\)[a-z]*"},{"begin":"%r!","end":"![a-z]*"},{"begin":"%r\\[","end":"\\][a-z]*"}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}],"relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}},{"className":"meta","begin":"^([>?]>|[\\w#]+\\(\\w+\\):\\d+:\\d+>|(\\w+-)?\\d+\\.\\d+\\.\\d(p\\d+)?[^>]+>)","starts":{"end":"$","contains":[{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"variants":[{"begin":"'","end":"'"},{"begin":"\"","end":"\""},{"begin":"`+"`"+`","end":"`+"`"+`"},{"begin":"%[qQwWx]?\\(","end":"\\)"},{"begin":"%[qQwWx]?\\[","end":"\\]"},{"begin":"%[qQwWx]?{","end":"}"},{"begin":"%[qQwWx]?<","end":">"},{"begin":"%[qQwWx]?/","end":"/"},{"begin":"%[qQwWx]?%","end":"%"},{"begin":"%[qQwWx]?-","end":"-"},{"begin":"%[qQwWx]?\\|","end":"\\|"},{"begin":"\\B\\?(\\\\\\d{1,3}|\\\\x[A-Fa-f0-9]{1,2}|\\\\u[A-Fa-f0-9]{4}|\\\\?\\S)\\b"},{"begin":"<<(-?)\\w+$","end":"^\\s*\\w+$"}]},{"begin":"#<","end":">"},{"className":"class","beginKeywords":"class module","end":"$|;","illegal":"=","contains":[{"className":"title","begin":"[A-Za-z_]\\w*(::\\w+)*(\\?|\\!)?","relevance":0},{"begin":"<\\s*","contains":[{"begin":"([a-zA-Z]\\w*::)?[a-zA-Z]\\w*"}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"className":"function","beginKeywords":"def","end":"$|;","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*[!?=]?|[-+~]\\@|<<|>>|=~|===?|<=>|[<>]=?|\\*\\*|[-/+%^&*~`+"`"+`|]|\\[\\]=?","relevance":0},{"className":"params","begin":"\\(","end":"\\)","endsParent":true,"keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"[a-zA-Z]\\w*::"},{"className":"symbol","begin":"[a-zA-Z_]\\w*(\\!|\\?)?:","relevance":0},{"className":"symbol","begin":":(?!\\s)","contains":[{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"variants":[{"begin":"'","end":"'"},{"begin":"\"","end":"\""},{"begin":"`+"`"+`","end":"`+"`"+`"},{"begin":"%[qQwWx]?\\(","end":"\\)"},{"begin":"%[qQwWx]?\\[","end":"\\]"},{"begin":"%[qQwWx]?{","end":"}"},{"begin":"%[qQwWx]?<","end":">"},{"begin":"%[qQwWx]?/","end":"/"},{"begin":"%[qQwWx]?%","end":"%"},{"begin":"%[qQwWx]?-","end":"-"},{"begin":"%[qQwWx]?\\|","end":"\\|"},{"begin":"\\B\\?(\\\\\\d{1,3}|\\\\x[A-Fa-f0-9]{1,2}|\\\\u[A-Fa-f0-9]{4}|\\\\?\\S)\\b"},{"begin":"<<(-?)\\w+$","end":"^\\s*\\w+$"}]},{"begin":"[a-zA-Z_]\\w*[!?=]?|[-+~]\\@|<<|>>|=~|===?|<=>|[<>]=?|\\*\\*|[-/+%^&*~`+"`"+`|]|\\[\\]=?"}],"relevance":0},{"className":"number","begin":"(\\b0[0-7_]+)|(\\b0x[0-9a-fA-F_]+)|(\\b[1-9][0-9_]*(\\.[0-9_]+)?)|[0_]\\b","relevance":0},{"begin":"(\\$\\W)|((\\$|\\@\\@?)(\\w+))"},{"className":"params","begin":"\\|","end":"\\|","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"}},{"begin":"(!|!=|!==|%|%=|&|&&|&=|\\*|\\*=|\\+|\\+=|,|-|-=|/=|/|:|;|<<|<<=|<=|<|===|==|=|>>>=|>>=|>=|>>>|>>|>|\\?|\\[|\\{|\\(|\\^|\\^=|\\||\\|=|\\|\\||~|unless)\\s*","keywords":"unless","contains":[{"begin":"#<","end":">"},{"className":"regexp","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"illegal":"\\n","variants":[{"begin":"/","end":"/[a-z]*"},{"begin":"%r{","end":"}[a-z]*"},{"begin":"%r\\(","end":"\\)[a-z]*"},{"begin":"%r!","end":"![a-z]*"},{"begin":"%r\\[","end":"\\][a-z]*"}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}],"relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}},{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"variants":[{"begin":"'","end":"'"},{"begin":"\"","end":"\""},{"begin":"`+"`"+`","end":"`+"`"+`"},{"begin":"%[qQwWx]?\\(","end":"\\)"},{"begin":"%[qQwWx]?\\[","end":"\\]"},{"begin":"%[qQwWx]?{","end":"}"},{"begin":"%[qQwWx]?<","end":">"},{"begin":"%[qQwWx]?/","end":"/"},{"begin":"%[qQwWx]?%","end":"%"},{"begin":"%[qQwWx]?-","end":"-"},{"begin":"%[qQwWx]?\\|","end":"\\|"},{"begin":"\\B\\?(\\\\\\d{1,3}|\\\\x[A-Fa-f0-9]{1,2}|\\\\u[A-Fa-f0-9]{4}|\\\\?\\S)\\b"},{"begin":"<<(-?)\\w+$","end":"^\\s*\\w+$"}]},{"begin":"#<","end":">"},{"className":"class","beginKeywords":"class module","end":"$|;","illegal":"=","contains":[{"className":"title","begin":"[A-Za-z_]\\w*(::\\w+)*(\\?|\\!)?","relevance":0},{"begin":"<\\s*","contains":[{"begin":"([a-zA-Z]\\w*::)?[a-zA-Z]\\w*"}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"className":"function","beginKeywords":"def","end":"$|;","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*[!?=]?|[-+~]\\@|<<|>>|=~|===?|<=>|[<>]=?|\\*\\*|[-/+%^&*~`+"`"+`|]|\\[\\]=?","relevance":0},{"className":"params","begin":"\\(","end":"\\)","endsParent":true,"keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"[a-zA-Z]\\w*::"},{"className":"symbol","begin":"[a-zA-Z_]\\w*(\\!|\\?)?:","relevance":0},{"className":"symbol","begin":":(?!\\s)","contains":[{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"variants":[{"begin":"'","end":"'"},{"begin":"\"","end":"\""},{"begin":"`+"`"+`","end":"`+"`"+`"},{"begin":"%[qQwWx]?\\(","end":"\\)"},{"begin":"%[qQwWx]?\\[","end":"\\]"},{"begin":"%[qQwWx]?{","end":"}"},{"begin":"%[qQwWx]?<","end":">"},{"begin":"%[qQwWx]?/","end":"/"},{"begin":"%[qQwWx]?%","end":"%"},{"begin":"%[qQwWx]?-","end":"-"},{"begin":"%[qQwWx]?\\|","end":"\\|"},{"begin":"\\B\\?(\\\\\\d{1,3}|\\\\x[A-Fa-f0-9]{1,2}|\\\\u[A-Fa-f0-9]{4}|\\\\?\\S)\\b"},{"begin":"<<(-?)\\w+$","end":"^\\s*\\w+$"}]},{"begin":"[a-zA-Z_]\\w*[!?=]?|[-+~]\\@|<<|>>|=~|===?|<=>|[<>]=?|\\*\\*|[-/+%^&*~`+"`"+`|]|\\[\\]=?"}],"relevance":0},{"className":"number","begin":"(\\b0[0-7_]+)|(\\b0x[0-9a-fA-F_]+)|(\\b[1-9][0-9_]*(\\.[0-9_]+)?)|[0_]\\b","relevance":0},{"begin":"(\\$\\W)|((\\$|\\@\\@?)(\\w+))"},{"className":"params","begin":"\\|","end":"\\|","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"}},{"begin":"(!|!=|!==|%|%=|&|&&|&=|\\*|\\*=|\\+|\\+=|,|-|-=|/=|/|:|;|<<|<<=|<=|<|===|==|=|>>>=|>>=|>=|>>>|>>|>|\\?|\\[|\\{|\\(|\\^|\\^=|\\||\\|=|\\|\\||~|unless)\\s*","keywords":"unless","contains":[{"begin":"#<","end":">"},{"className":"regexp","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"contains":[{"Ref":["contains","3","starts","contains"],"IsArray":true}]}],"illegal":"\\n","variants":[{"begin":"/","end":"/[a-z]*"},{"begin":"%r{","end":"}[a-z]*"},{"begin":"%r\\(","end":"\\)[a-z]*"},{"begin":"%r!","end":"![a-z]*"},{"begin":"%r\\[","end":"\\][a-z]*"}]},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}],"relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"^\\=begin","end":"^\\=end","contains":[{"className":"doctag","begin":"@[A-Za-z]+"},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^__END__","end":"\\n$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}`)
}