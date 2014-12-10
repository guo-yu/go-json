package JSON

import "encoding/json"

func Parse(jsonString string) (map[string]interface{}, error) {
	var ret map[string]interface{}
	var f interface{}

	err := json.Unmarshal([]byte(jsonString), &f)

	if err != nil {
		return ret, err
	}

	return f.(map[string]interface{}), nil
}

func Stringify(jsonContent interface{}) (string, error) {
	jsonByte, err := json.Marshal(jsonContent)

	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}

func Byteify(jsonContent interface{}) ([]byte, error) {
	return json.Marshal(jsonContent)
}
