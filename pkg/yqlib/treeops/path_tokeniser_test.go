package treeops

import (
	"testing"

	"github.com/mikefarah/yq/v3/test"
)

var tokeniserTests = []struct {
	path           string
	expectedTokens []interface{}
}{ // TODO: Ensure ALL documented examples have tests! sheesh
	// {"len(.)", append(make([]interface{}, 0), "LENGTH", "(", "SELF", ")")},
	// {"\"len\"(.)", append(make([]interface{}, 0), "len", "TRAVERSE", "(", "SELF", ")")},
	// {".a OR (.b OR .c)", append(make([]interface{}, 0), "a", "OR", "(", "b", "OR", "c", ")")},
	// {"a OR (b OR c)", append(make([]interface{}, 0), "a", "OR", "(", "b", "OR", "c", ")")},
	// {"a .- (b OR c)", append(make([]interface{}, 0), "a", " .- ", "(", "b", "OR", "c", ")")},
	// {"(animal==3)", append(make([]interface{}, 0), "(", "animal", "==", int64(3), ")")},
	// {"(animal==f3)", append(make([]interface{}, 0), "(", "animal", "==", "f3", ")")},
	// {"apples.BANANAS", append(make([]interface{}, 0), "apples", "TRAVERSE", "BANANAS")},
	// {"appl*.BANA*", append(make([]interface{}, 0), "appl*", "TRAVERSE", "BANA*")},
	// {"a.b.**", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", "**")},
	// {"a.\"=\".frog", append(make([]interface{}, 0), "a", "TRAVERSE", "=", "TRAVERSE", "frog")},
	// {"a.b.*", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", "*")},
	// {"a.b.thin*", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", "thin*")},
	// {".a.b.[0]", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", int64(0))},
	// {".a.b.[]", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", "[]")},
	// {".a.b.[+]", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", "[+]")},
	// {".a.b.[-12]", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", int64(-12))},
	// {".a.b.0", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", "0")},
	// {".a", append(make([]interface{}, 0), "a")},
	// {".\"a.b\".c", append(make([]interface{}, 0), "a.b", "TRAVERSE", "c")},
	// {`.b."foo.bar"`, append(make([]interface{}, 0), "b", "TRAVERSE", "foo.bar")},
	// {`f | . == *og | length`, append(make([]interface{}, 0), "f", "TRAVERSE", "SELF", "EQUALS", "*og", "TRAVERSE", "LENGTH")},
	// {`.a, .b`, append(make([]interface{}, 0), "a", "OR", "b")},
	// {`[.a, .b]`, append(make([]interface{}, 0), "[", "a", "OR", "b", "]")},
	// {`."[a", ."b]"`, append(make([]interface{}, 0), "[a", "OR", "b]")},
	// {`.a[]`, append(make([]interface{}, 0), "a", "PIPE", "[]")},
	// {`.[].a`, append(make([]interface{}, 0), "[]", "PIPE", "a")},
	{`.a | (.[].b == "apple")`, append(make([]interface{}, 0), "a", "PIPE", "(", "[]", "PIPE", "b", "EQUALS", "apple", ")")},

	// {".animals | .==cat", append(make([]interface{}, 0), "animals", "TRAVERSE", "SELF", "EQUALS", "cat")},
	// {".animals | (. == cat)", append(make([]interface{}, 0), "animals", "TRAVERSE", "(", "SELF", "EQUALS", "cat", ")")},
	// {".animals | (.==c*)", append(make([]interface{}, 0), "animals", "TRAVERSE", "(", "SELF", "EQUALS", "c*", ")")},
	// {"animals(a.b==c*)", append(make([]interface{}, 0), "animals", "TRAVERSE", "(", "a", "TRAVERSE", "b", "==", "c*", ")")},
	// {"animals.(a.b==c*)", append(make([]interface{}, 0), "animals", "TRAVERSE", "(", "a", "TRAVERSE", "b", "==", "c*", ")")},
	// {"(a.b==c*).animals", append(make([]interface{}, 0), "(", "a", "TRAVERSE", "b", "==", "c*", ")", "TRAVERSE", "animals")},
	// {"(a.b==c*)animals", append(make([]interface{}, 0), "(", "a", "TRAVERSE", "b", "==", "c*", ")", "TRAVERSE", "animals")},
	// {"[1].a.d", append(make([]interface{}, 0), int64(1), "TRAVERSE", "a", "TRAVERSE", "d")},
	// {"[1]a.d", append(make([]interface{}, 0), int64(1), "TRAVERSE", "a", "TRAVERSE", "d")},
	// {"a[0]c", append(make([]interface{}, 0), "a", "TRAVERSE", int64(0), "TRAVERSE", "c")},
	// {"a.[0].c", append(make([]interface{}, 0), "a", "TRAVERSE", int64(0), "TRAVERSE", "c")},
	// {"[0]", append(make([]interface{}, 0), int64(0))},
	// {"0", append(make([]interface{}, 0), int64(0))},
	// {"a.b[+]c", append(make([]interface{}, 0), "a", "TRAVERSE", "b", "TRAVERSE", "[+]", "TRAVERSE", "c")},
	// {"a.cool(s.d.f == cool)", append(make([]interface{}, 0), "a", "TRAVERSE", "cool", "TRAVERSE", "(", "s", "TRAVERSE", "d", "TRAVERSE", "f", " == ", "cool", ")")},
	// {"a.cool.(s.d.f==cool OR t.b.h==frog).caterpillar", append(make([]interface{}, 0), "a", "TRAVERSE", "cool", "TRAVERSE", "(", "s", "TRAVERSE", "d", "TRAVERSE", "f", "==", "cool", "OR", "t", "TRAVERSE", "b", "TRAVERSE", "h", "==", "frog", ")", "TRAVERSE", "caterpillar")},
	// {"a.cool(s.d.f==cool and t.b.h==frog)*", append(make([]interface{}, 0), "a", "TRAVERSE", "cool", "TRAVERSE", "(", "s", "TRAVERSE", "d", "TRAVERSE", "f", "==", "cool", "and", "t", "TRAVERSE", "b", "TRAVERSE", "h", "==", "frog", ")", "TRAVERSE", "*")},
	// {"a.cool(s.d.f==cool and t.b.h==frog).th*", append(make([]interface{}, 0), "a", "TRAVERSE", "cool", "TRAVERSE", "(", "s", "TRAVERSE", "d", "TRAVERSE", "f", "==", "cool", "and", "t", "TRAVERSE", "b", "TRAVERSE", "h", "==", "frog", ")", "TRAVERSE", "th*")},
}

var tokeniser = NewPathTokeniser()

func TestTokeniser(t *testing.T) {
	for _, tt := range tokeniserTests {
		tokens, err := tokeniser.Tokenise(tt.path)
		if err != nil {
			t.Error(tt.path, err)
		}
		var tokenValues []interface{}
		for _, token := range tokens {
			tokenValues = append(tokenValues, token.Value)
		}
		test.AssertResultComplexWithContext(t, tt.expectedTokens, tokenValues, tt.path)
	}
}
