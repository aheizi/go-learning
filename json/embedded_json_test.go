package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
  "basic_info": {
    "name": "aheizi",
    "age": 27
  },
  "job_info": {
    "skills": [
      "Java",
      "Go",
      "C"
    ]
  }
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(*e)

	if v, err := json.Marshal(e); err == nil {
		fmt.Print(string(v))
	} else {
		t.Error(err)
	}
}
