package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"rust", "rs"}, `{"aliases":["rs"],"keywords":{"keyword":"alignof as be box break const continue crate do else enum extern false fn for if impl in let loop match mod mut offsetof once priv proc pub pure ref return self Self sizeof static struct super trait true type typeof unsafe unsized use virtual while where yield move default","literal":"true false Some None Ok Err","built_in":"drop i8 i16 i32 i64 i128 isize u8 u16 u32 u64 u128 usize f32 f64 str char bool Box Option Result String Vec Copy Send Sized Sync Drop Fn FnMut FnOnce ToOwned Clone Debug PartialEq PartialOrd Eq Ord AsRef AsMut Into From Default Iterator Extend IntoIterator DoubleEndedIterator ExactSizeIterator SliceConcatExt ToString assert! assert_eq! bitflags! bytes! cfg! col! concat! concat_idents! debug_assert! debug_assert_eq! env! panic! file! format! format_args! include_bin! include_str! line! local_data_key! module_path! option_env! print! println! select! stringify! try! unimplemented! unreachable! vec! write! writeln! macro_rules! assert_ne! debug_assert_ne!"},"lexemes":"[a-zA-Z]\\w*!?","illegal":"</","contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"Ref":["contains","1"]},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|they|like|more)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"b?\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","variants":[{"begin":"r(#*)\"(.|\\n)*?\"\\1(?!#)"},{"begin":"b?'\\\\?(x\\w{2}|u\\w{4}|U\\w{8}|.)'"}]},{"className":"symbol","begin":"'[a-zA-Z_][a-zA-Z0-9_]*"},{"className":"number","variants":[{"begin":"\\b0b([01_]+)([ui](8|16|32|64|128|size)|f(32|64))?"},{"begin":"\\b0o([0-7_]+)([ui](8|16|32|64|128|size)|f(32|64))?"},{"begin":"\\b0x([A-Fa-f0-9_]+)([ui](8|16|32|64|128|size)|f(32|64))?"},{"begin":"\\b(\\d[\\d_]*(\\.[0-9_]+)?([eE][+-]?[0-9_]+)?)([ui](8|16|32|64|128|size)|f(32|64))?"}],"relevance":0},{"className":"function","beginKeywords":"fn","end":"(\\(|<)","excludeEnd":true,"contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"className":"meta","begin":"#\\!?\\[","end":"\\]","contains":[{"className":"meta-string","begin":"\"","end":"\""}]},{"className":"class","beginKeywords":"type","end":";","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0,"endsParent":true}],"illegal":"\\S"},{"className":"class","beginKeywords":"trait enum struct union","end":"{","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0,"endsParent":true}],"illegal":"[\\w\\d]"},{"begin":"[a-zA-Z]\\w*::","keywords":{"built_in":"drop i8 i16 i32 i64 i128 isize u8 u16 u32 u64 u128 usize f32 f64 str char bool Box Option Result String Vec Copy Send Sized Sync Drop Fn FnMut FnOnce ToOwned Clone Debug PartialEq PartialOrd Eq Ord AsRef AsMut Into From Default Iterator Extend IntoIterator DoubleEndedIterator ExactSizeIterator SliceConcatExt ToString assert! assert_eq! bitflags! bytes! cfg! col! concat! concat_idents! debug_assert! debug_assert_eq! env! panic! file! format! format_args! include_bin! include_str! line! local_data_key! module_path! option_env! print! println! select! stringify! try! unimplemented! unreachable! vec! write! writeln! macro_rules! assert_ne! debug_assert_ne!"}},{"begin":"->"}]}`)
}