package pip_filter

import (
	"github.com/pkg/errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (sf *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}
