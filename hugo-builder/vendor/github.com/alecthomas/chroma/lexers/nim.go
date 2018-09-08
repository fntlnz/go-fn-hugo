package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Nim lexer.
var Nim = Register(MustNewLexer(
	&Config{
		Name:            "Nim",
		Aliases:         []string{"nim", "nimrod"},
		Filenames:       []string{"*.nim", "*.nimrod"},
		MimeTypes:       []string{"text/x-nim"},
		CaseInsensitive: true,
	},
	Rules{
		"root": {
			{`##.*$`, LiteralStringDoc, nil},
			{`#.*$`, Comment, nil},
			{`[*=><+\-/@$~&%!?|\\\[\]]`, Operator, nil},
			{"\\.\\.|\\.|,|\\[\\.|\\.\\]|\\{\\.|\\.\\}|\\(\\.|\\.\\)|\\{|\\}|\\(|\\)|:|\\^|`|;", Punctuation, nil},
			{`(?:[\w]+)"`, LiteralString, Push("rdqs")},
			{`"""`, LiteralString, Push("tdqs")},
			{`"`, LiteralString, Push("dqs")},
			{`'`, LiteralStringChar, Push("chars")},
			{`(a_?n_?d_?|o_?r_?|n_?o_?t_?|x_?o_?r_?|s_?h_?l_?|s_?h_?r_?|d_?i_?v_?|m_?o_?d_?|i_?n_?|n_?o_?t_?i_?n_?|i_?s_?|i_?s_?n_?o_?t_?)\b`, OperatorWord, nil},
			{`(p_?r_?o_?c_?\s)(?![(\[\]])`, Keyword, Push("funcname")},
			{`(a_?d_?d_?r_?|a_?n_?d_?|a_?s_?|a_?s_?m_?|a_?t_?o_?m_?i_?c_?|b_?i_?n_?d_?|b_?l_?o_?c_?k_?|b_?r_?e_?a_?k_?|c_?a_?s_?e_?|c_?a_?s_?t_?|c_?o_?n_?c_?e_?p_?t_?|c_?o_?n_?s_?t_?|c_?o_?n_?t_?i_?n_?u_?e_?|c_?o_?n_?v_?e_?r_?t_?e_?r_?|d_?e_?f_?e_?r_?|d_?i_?s_?c_?a_?r_?d_?|d_?i_?s_?t_?i_?n_?c_?t_?|d_?i_?v_?|d_?o_?|e_?l_?i_?f_?|e_?l_?s_?e_?|e_?n_?d_?|e_?n_?u_?m_?|e_?x_?c_?e_?p_?t_?|e_?x_?p_?o_?r_?t_?|f_?i_?n_?a_?l_?l_?y_?|f_?o_?r_?|f_?u_?n_?c_?|i_?f_?|i_?n_?|y_?i_?e_?l_?d_?|i_?n_?t_?e_?r_?f_?a_?c_?e_?|i_?s_?|i_?s_?n_?o_?t_?|i_?t_?e_?r_?a_?t_?o_?r_?|l_?e_?t_?|m_?a_?c_?r_?o_?|m_?e_?t_?h_?o_?d_?|m_?i_?x_?i_?n_?|m_?o_?d_?|n_?o_?t_?|n_?o_?t_?i_?n_?|o_?b_?j_?e_?c_?t_?|o_?f_?|o_?r_?|o_?u_?t_?|p_?r_?o_?c_?|p_?t_?r_?|r_?a_?i_?s_?e_?|r_?e_?f_?|r_?e_?t_?u_?r_?n_?|s_?h_?a_?r_?e_?d_?|s_?h_?l_?|s_?h_?r_?|s_?t_?a_?t_?i_?c_?|t_?e_?m_?p_?l_?a_?t_?e_?|t_?r_?y_?|t_?u_?p_?l_?e_?|t_?y_?p_?e_?|w_?h_?e_?n_?|w_?h_?i_?l_?e_?|w_?i_?t_?h_?|w_?i_?t_?h_?o_?u_?t_?|x_?o_?r_?)\b`, Keyword, nil},
			{`(f_?r_?o_?m_?|i_?m_?p_?o_?r_?t_?|i_?n_?c_?l_?u_?d_?e_?)\b`, KeywordNamespace, nil},
			{`(v_?a_?r)\b`, KeywordDeclaration, nil},
			{`(i_?n_?t_?|i_?n_?t_?8_?|i_?n_?t_?1_?6_?|i_?n_?t_?3_?2_?|i_?n_?t_?6_?4_?|f_?l_?o_?a_?t_?|f_?l_?o_?a_?t_?3_?2_?|f_?l_?o_?a_?t_?6_?4_?|b_?o_?o_?l_?|c_?h_?a_?r_?|r_?a_?n_?g_?e_?|a_?r_?r_?a_?y_?|s_?e_?q_?|s_?e_?t_?|s_?t_?r_?i_?n_?g_?)\b`, KeywordType, nil},
			{`(n_?i_?l_?|t_?r_?u_?e_?|f_?a_?l_?s_?e_?)\b`, KeywordPseudo, nil},
			{`\b((?![_\d])\w)(((?!_)\w)|(_(?!_)\w))*`, Name, nil},
			{`[0-9][0-9_]*(?=([e.]|\'f(32|64)))`, LiteralNumberFloat, Push("float-suffix", "float-number")},
			{`0x[a-f0-9][a-f0-9_]*`, LiteralNumberHex, Push("int-suffix")},
			{`0b[01][01_]*`, LiteralNumberBin, Push("int-suffix")},
			{`0o[0-7][0-7_]*`, LiteralNumberOct, Push("int-suffix")},
			{`[0-9][0-9_]*`, LiteralNumberInteger, Push("int-suffix")},
			{`\s+`, Text, nil},
			{`.+$`, Error, nil},
		},
		"chars": {
			{`\\([\\abcefnrtvl"\']|x[a-f0-9]{2}|[0-9]{1,3})`, LiteralStringEscape, nil},
			{`'`, LiteralStringChar, Pop(1)},
			{`.`, LiteralStringChar, nil},
		},
		"strings": {
			{`(?<!\$)\$(\d+|#|\w+)+`, LiteralStringInterpol, nil},
			{`[^\\\'"$\n]+`, LiteralString, nil},
			{`[\'"\\]`, LiteralString, nil},
			{`\$`, LiteralString, nil},
		},
		"dqs": {
			{`\\([\\abcefnrtvl"\']|\n|x[a-f0-9]{2}|[0-9]{1,3})`, LiteralStringEscape, nil},
			{`"`, LiteralString, Pop(1)},
			Include("strings"),
		},
		"rdqs": {
			{`"(?!")`, LiteralString, Pop(1)},
			{`""`, LiteralStringEscape, nil},
			Include("strings"),
		},
		"tdqs": {
			{`"""(?!")`, LiteralString, Pop(1)},
			Include("strings"),
			Include("nl"),
		},
		"funcname": {
			{`((?![\d_])\w)(((?!_)\w)|(_(?!_)\w))*`, NameFunction, Pop(1)},
			{"`.+`", NameFunction, Pop(1)},
		},
		"nl": {
			{`\n`, LiteralString, nil},
		},
		"float-number": {
			{`\.(?!\.)[0-9_]*`, LiteralNumberFloat, nil},
			{`e[+-]?[0-9][0-9_]*`, LiteralNumberFloat, nil},
			Default(Pop(1)),
		},
		"float-suffix": {
			{`\'f(32|64)`, LiteralNumberFloat, nil},
			Default(Pop(1)),
		},
		"int-suffix": {
			{`\'i(32|64)`, LiteralNumberIntegerLong, nil},
			{`\'i(8|16)`, LiteralNumberInteger, nil},
			Default(Pop(1)),
		},
	},
))
