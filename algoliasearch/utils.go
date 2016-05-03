package algoliasearch

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

// randDuration generates a pseudo-random `time.Duration` between 1 and `max`.
func randDuration(max time.Duration) time.Duration {
	rand.Seed(time.Now().Unix())
	nbNanoseconds := 1 + int64(rand.Int63n(max.Nanoseconds()))
	return time.Duration(nbNanoseconds) * time.Nanosecond
}

func invalidParameter(p string) error {
	return fmt.Errorf("`%s` doesn't exist or doesn't have the right type", p)
}

func duplicateMap(m map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})

	for k, v := range m {
		copy[k] = v
	}

	return copy
}

// encodeParams transforms `params` to a URL-safe string.
func encodeParams(params map[string]interface{}) string {
	values := url.Values{}

	if params != nil {
		for k, v := range params {
			switch v := v.(type) {
			case string:
				values.Add(k, v)
			case float64:
				values.Add(k, strconv.FormatFloat(v, 'f', -1, 64))
			case int:
				values.Add(k, strconv.Itoa(v))
			default:
				jsonValue, _ := json.Marshal(v)
				values.Add(k, string(jsonValue[:]))
			}
		}
	}

	return values.Encode()
}
