package govdf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderRoundTrip(t *testing.T) {
	input := "\"z\" {\n    \"b\" \"one\"\n    \"a\" \"two\"\n    \"b\" \"three\"\n}\n\"a\" {\n}\n\"z\" {\n    \"x\" \"last\"\n}\n"
	var node Node
	require.NoError(t, Unmarshal([]byte(input), &node))
	require.Equal(t, []string{"z", "a", "z"}, node.Order)
	text, err := Marshal(&node)
	require.NoError(t, err)
	require.Equal(t, input, string(text))
	binary, err := MarshalBinary(&node)
	require.NoError(t, err)
	var decoded Node
	require.NoError(t, UnmarshalBinary(binary, &decoded))
	text, err = Marshal(&decoded)
	require.NoError(t, err)
	require.Equal(t, input, string(text))
	again, err := MarshalBinary(&decoded)
	require.NoError(t, err)
	require.Equal(t, binary, again)
}

func TestOrderWithModifiedChildren(t *testing.T) {
	var node Node
	require.NoError(t, Unmarshal([]byte("\"z\" \"first\"\n\"a\" \"removed\"\n\"z\" \"removed\"\n"), &node))
	delete(node.Children, "a")
	node.Children["z"] = node.Children["z"][:1]
	node.Children["b"] = []*Node{{Type: NodeTypeScalar, Value: "new"}}
	text, err := Marshal(&node)
	require.NoError(t, err)
	require.Equal(t, "\"z\" \"first\"\n\"b\" \"new\"\n", string(text))
}
