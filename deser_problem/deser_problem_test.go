package deser_problem

import (
	"bytes"
	"encoding/json"
	assert "github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func encodeResource(t *testing.T, obj interface{}) (string, error) {
	var bodyBuffer io.Reader
	bodyBuffer = new(bytes.Buffer)
	err := json.NewEncoder(bodyBuffer.(io.Writer)).Encode(obj)
	if err != nil {
		t.Log(">>> Encoding error: ", err)
		return "", err
	}
	return readStream(bodyBuffer), nil
}

func readStream(body io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}

func decodeResource(t *testing.T, jsonString string) (*Resource, error) {
	result := &Resource{}
	responseBody := strings.NewReader(jsonString)
	err := json.NewDecoder(responseBody).Decode(result)
	if err != nil {
		t.Log(">>> Decoding error: ", err)
		return nil, err
	}
	return result, nil
}

func testEncodeDecode(t *testing.T, resourceObj *Resource, expectedJSON string) {
	t.Log("Expected JSON:                 ", expectedJSON)

	// Simulate the request path by encoding the Resource instance.
	actualJSON, err := encodeResource(t, resourceObj)
	assert.Nil(t, err)
	assert.NotNil(t, actualJSON)
	t.Log("Encoded request Resource:     ", actualJSON)
	assert.Equal(t, expectedJSON, actualJSON)

	// Simulate the response path by decoding "actualJSON" into a Resource instance.
	responseResource, err := decodeResource(t, actualJSON)
	assert.Nil(t, err)
	assert.NotNil(t, responseResource)

	// Now encode the response object so we can compare it to the encoded request object.
	responseJSON, err := encodeResource(t, responseResource)
	assert.Nil(t, err)
	assert.NotNil(t, responseJSON)
	t.Log("Decoded response Resource: ", responseJSON)
	assert.Equal(t, actualJSON, responseJSON)
}

func TestResourceWithFooInfo(t *testing.T) {
	expectedJSON := "{\"id\":\"id-1\",\"info\":{\"foo\":\"This is a Foo Info object\"}}\n"

	fooInfo := &Foo{
		Foo: "This is a Foo Info object",
	}

	resourceObj := &Resource{
		ID:   "id-1",
		Info: fooInfo,
	}

	testEncodeDecode(t, resourceObj, expectedJSON)
}

func TestResourceWithBarInfo(t *testing.T) {
	expectedJSON := "{\"id\":\"id-1\",\"info\":{\"bar\":\"This is a Bar Info object\"}}\n"

	barInfo := &Bar{
		Bar: "This is a Bar Info object",
	}

	resourceObj := &Resource{
		ID:   "id-1",
		Info: barInfo,
	}

	testEncodeDecode(t, resourceObj, expectedJSON)
}
