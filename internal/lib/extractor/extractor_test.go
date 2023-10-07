package extractor

import (
	"testing"
)

type out struct {
	key   string
	value any
	err   error
}

type caseItem struct {
	input  string
	ftype  fileType
	output *out
}

func TestExtractLine(t *testing.T) {
	cases := []*caseItem{
		{
			input:  `"resolved": "https://registry.npmjs.org/pend/-/pend-1.2.0.tgz",`,
			ftype:  Json,
			output: &out{key: "resolved", value: "https://registry.npmjs.org/pend/-/pend-1.2.0.tgz", err: nil},
		},
		{
			input:  `"pend": {`,
			ftype:  Json,
			output: &out{key: "", value: "", err: errMatch},
		},
		{
			input:  `},`,
			ftype:  Json,
			output: &out{key: "", value: "", err: errMatch},
		},
		{
			input:  "",
			ftype:  Json,
			output: &out{key: "", value: "", err: errMatch},
		},
		{
			input:  "zod@^3.21.4:",
			ftype:  Json,
			output: &out{key: "", value: "", err: errMatch},
		},
		{
			input:  `resolved "https://registry.yarnpkg.com/zod/-/zod-3.21.4.tgz#10882231d992519f0a10b5dd58a38c9dabbb64db"`,
			ftype:  Yml,
			output: &out{key: "resolved", value: "https://registry.yarnpkg.com/zod/-/zod-3.21.4.tgz#10882231d992519f0a10b5dd58a38c9dabbb64db", err: nil},
		},
		{
			input:  `readable-stream "^3.6.0"`,
			ftype:  Yml,
			output: &out{key: "readable-stream", value: "^3.6.0", err: nil},
		},
	}

	for _, c := range cases {
		extr := New(c.ftype)
		key, value, err := extr.ExtractLine(c.input)

		if key != c.output.key || value != c.output.value || err != c.output.err {
			t.Errorf("Expect: %v ---- Got: %v, %v, %v", c.output, key, value, err)
		}
	}
}
