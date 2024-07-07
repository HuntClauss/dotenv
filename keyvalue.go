package dotenv

import "strings"

type KeyValue struct {
	Key   string
	Value string
}

func (kv KeyValue) String() string {
	return kv.Key + "=" + kv.Value
}

func keyValueFromString(s string) KeyValue {
	parts := strings.SplitN(s, "=", 1)
	if len(parts) != 2 {
		panic("cannot convert environ value to key value")
	}

	return KeyValue{Key: parts[0], Value: parts[1]}
}
