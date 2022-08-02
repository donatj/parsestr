// Package parsestr is a simple Golang library emulating the query parsing
// behaviour of the PHP parse_str.
//
// See: http://us2.php.net/manual/en/function.parse-str.php
package parsestr

import (
	"net/url"
	"regexp"
)

func ParseURLValues(val url.Values) Values {
	return Values(val)
}

func ParseQuery(query string) (Values, error) {
	val, err := url.ParseQuery(query)
	return Values(val), err
}

type Values url.Values

func (v Values) Get(key string) string {
	return url.Values(v).Get(key)
}

func (v Values) GetSub(key string) Values {
	ret := Values{}

	r := regexp.MustCompile(`^` + regexp.QuoteMeta(key) + `\[([^\]]*)]`)
	for i, e := range v {
		m := r.FindStringSubmatch(i)
		if len(m) > 0 {
			if ret[m[1]] == nil {
				ret[m[1]] = []string{}
			}

			ret[m[1]] = append(ret[m[1]], e...)
		}
	}

	return ret
}
