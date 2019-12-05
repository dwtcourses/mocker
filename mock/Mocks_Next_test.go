package mock

import "testing"

func TestNextWorkSuccess(t *testing.T) {
	// Arrange

	path := "/test"
	method := "GET"

	group := RequestModelGroup{
		URL:             path,
		Method:          method,
		iteratorIndexes: map[string]int{},
		models: []RequestModel{
			RequestModel{
				FilePath: path,
			},
		},
	}

	// Act

	result := group.Next(path)

	// Assert

	if result == nil {
		t.Fail()
	}
}

func TestNextWorkSuccessWithWrongPath(t *testing.T) {
	// Arrange

	path := "/test"
	method := "GET"

	group := RequestModelGroup{
		URL:             path,
		Method:          method,
		iteratorIndexes: map[string]int{},
		models: []RequestModel{
			RequestModel{
				FilePath: "tmp" + path,
			},
		},
	}

	// Act

	result := group.Next(path)

	// Assert

	if result != nil {
		t.Fail()
	}
}

func TestNextUpdateCounter(t *testing.T) {
	// Arrange

	path := "/test"
	method := "GET"

	group := RequestModelGroup{
		URL:             path,
		Method:          method,
		iteratorIndexes: map[string]int{},
		models: []RequestModel{
			RequestModel{
				FilePath: path,
			},
		},
	}

	// Act

	group.Next(path)

	// Assert

	if group.iteratorIndexes[path] != 1 {
		t.Fail()
	}
}

func TestNextReturnsIsOnlyMock(t *testing.T) {
	// Arrange

	path := "/test"
	method := "GET"

	group := RequestModelGroup{
		URL:             path,
		Method:          method,
		iteratorIndexes: map[string]int{},
		models: []RequestModel{
			RequestModel{
				FilePath: path,
			},
			RequestModel{
				FilePath: path,
				IsOnly:   newTrue(),
			},
		},
	}

	// Act

	result := group.Next(path)

	// Assert

	if result == nil {
		t.Fail()
	}

	if result.IsOnly == nil {
		t.Fail()
	}
}

func TestNextReturnsIsOnlyMockNotMutateIterator(t *testing.T) {
	// Arrange

	path := "/test"
	method := "GET"

	group := RequestModelGroup{
		URL:             path,
		Method:          method,
		iteratorIndexes: map[string]int{},
		models: []RequestModel{
			RequestModel{
				FilePath: path,
			},
			RequestModel{
				FilePath: path,
				IsOnly:   newTrue(),
			},
		},
	}

	index := group.iteratorIndexes[path]

	// Act

	_ = group.Next(path)

	// Assert

	if index != group.iteratorIndexes[path] {
		t.Fail()
	}
}
