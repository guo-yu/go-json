package JSON

// import "fmt"
import "testing"

func TestJSONParse(t *testing.T) {
	me, err := Parse(`{
    "age": "23",
    "name": "Guo",
    "location": "Peking, Mainland China",
    "meta": {
      "love": ["coding", "xbox"]
    }
  }`)

	// fmt.Println(me["meta"].(map[string]interface{})["love"])

	if err != nil {
		t.Error(err)
	}

	if me["age"] != "23" {
		t.Error("Failed")
	}

	t.Log("Passed")
}

func TestJSONStringify(t *testing.T) {
	dogge := map[string]interface{}{
		"age":  8,
		"name": "dogge",
	}

	doggeString, err := Stringify(dogge)

	if err != nil {
		t.Error(err)
	}

	if doggeString != `{"age":8,"name":"dogge"}` {
		t.Error("Failed")
	}

	t.Log("Passed")
}
