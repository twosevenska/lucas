package crawl

import (
	"fmt"
	"testing"
)

type testEU struct {
	line    string
	address string
	ok      bool
}

func (expected testEU) isExpected(address string, ok bool) error {
	result := testEU{
		line:    expected.line,
		address: address,
		ok:      ok,
	}

	if expected != result {
		return fmt.Errorf("value mismatch. Expected VS Result\n    %+v\n    %+v\n", expected, result)
	}

	return nil
}

func TestExtractURL(t *testing.T) {
	cases := []testEU{
		testEU{
			line:    "<li><a href=\"http://info.cern.ch/hypertext/WWW/TheProject.html\">Browse the first website</a></li>",
			address: "http://info.cern.ch/hypertext/WWW/TheProject.html",
			ok:      true,
		},
		testEU{
			line:    "<p>From here you can:</p>",
			address: "",
			ok:      false,
		},
		testEU{
			line:    "",
			address: "",
			ok:      false,
		},
		testEU{
			line:    "<li><a href=\"\">Browse the first website</a></li>",
			address: "",
			ok:      false,
		},
	}

	for _, c := range cases {
		a, ok := ExtractURL(c.line)
		err := c.isExpected(a, ok)
		if err != nil {
			t.Errorf("Failed assertion: \n %s", err.Error())
		}
	}
}
