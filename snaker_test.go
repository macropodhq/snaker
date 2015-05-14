package snaker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var content = [][3]string{
	{"one", "One", "one"},
	{"o_n_e", "ONE", "oNE"},
	{"id", "ID", "id"},
	{"i", "I", "i"},
	{"this_has_to_be_converted_correctly_id", "ThisHasToBeConvertedCorrectlyID", "thisHasToBeConvertedCorrectlyID"},
	{"this_id_is_fine", "ThisIDIsFine", "thisIDIsFine"},
	{"this_https_connection", "ThisHTTPSConnection", "thisHTTPSConnection"},
	{"hello_https_connection_id", "HelloHTTPSConnectionID", "helloHTTPSConnectionID"},
	{"https_id", "HTTPSID", "httpsID"},
	{"project_id", "ProjectID", "projectID"},
}

var functions = []struct {
	n string
	f func(string) string
}{
	{"Snake", ToSnake},
	{"UpperCamel", ToUpperCamel},
	{"LowerCamel", ToLowerCamel},
}

func TestToSnake(t *testing.T) {
	a := assert.New(t)

	a.Equal("one", ToSnake("One"), "should work with one word")
	a.Equal("o_n_e", ToSnake("ONE"), "should return an uppercase string as seperate words")
	a.Equal("id", ToSnake("ID"), "should return ID as lowercase")
	a.Equal("i", ToSnake("i"), "should work with a single lowercase character")
	a.Equal("i", ToSnake("I"), "should work with a single uppcase character")
	a.Equal("this_has_to_be_converted_correctly_id", ToSnake("ThisHasToBeConvertedCorrectlyID"), "should return a long text as expected")
	a.Equal("this_id_is_fine", ToSnake("ThisIDIsFine"), "should return the text as expected if the initialism is in the middle")
	a.Equal("this_https_connection", ToSnake("ThisHTTPSConnection"), "should work with long initialism")
	a.Equal("hello_https_connection_id", ToSnake("HelloHTTPSConnectionID"), "should work with multi initialisms")
	a.Equal("https_id", ToSnake("HTTPSID"), "sould work with concat initialisms")
}

func TestToUpperCamel(t *testing.T) {
	a := assert.New(t)

	a.Equal("ThisHasToBeUppercased", ToUpperCamel("this_has_to_be_uppercased"), "should return a snaked text as camel case")
	a.Equal("ThisIsAnID", ToUpperCamel("this_is_an_id"), "should return a snaked text as camel case, except the word ID")
	a.Equal("ThisIsAnIdentifier", ToUpperCamel("this_is_an_identifier"), "should return 'id' not as uppercase")
	a.Equal("ID", ToUpperCamel("id"), "should simply work with id")
}

func TestToLowerCamel(t *testing.T) {
	a := assert.New(t)

	a.Equal("thisHasToBeUppercased", ToLowerCamel("this_has_to_be_uppercased"), "should return a snaked text as camel case")
	a.Equal("thisIsAnID", ToLowerCamel("this_is_an_id"), "should return a snaked text as camel case, except the word ID")
	a.Equal("thisIsAnIdentifier", ToLowerCamel("this_is_an_identifier"), "should return 'id' not as uppercase")
	a.Equal("id", ToLowerCamel("id"), "should simply work with id")
}

func TestRoundTrips(t *testing.T) {
	a := assert.New(t)

	for _, s := range content {
		for i := range functions {
			f0, f1, f2 := functions[i], functions[(i+1)%3], functions[(i+2)%3]

			a.Equal(s[i], f0.f(s[i]), f0.n+" -> "+f0.n)
			a.Equal(s[i], f0.f(s[(i+2)%3]), f1.n+" -> "+f0.n)
			a.Equal(s[i], f0.f(s[(i+3)%3]), f2.n+" -> "+f0.n)

			a.Equal(s[i], f0.f(f1.f(s[i])), f0.n+" -> "+f1.n+" -> "+f0.n)
			a.Equal(s[i], f0.f(f1.f(s[(i+2)%3])), f1.n+" -> "+f1.n+" -> "+f0.n)
			a.Equal(s[i], f0.f(f1.f(s[(i+3)%3])), f2.n+" -> "+f1.n+" -> "+f0.n)

			a.Equal(s[i], f0.f(f1.f(f2.f(s[i]))), f0.n+" -> "+f2.n+" -> "+f1.n+" -> "+f0.n)
			a.Equal(s[i], f0.f(f1.f(f2.f(s[(i+2)%3]))), f1.n+" -> "+f2.n+" -> "+f1.n+" -> "+f0.n)
			a.Equal(s[i], f0.f(f1.f(f2.f(s[(i+3)%3]))), f2.n+" -> "+f2.n+" -> "+f1.n+" -> "+f0.n)
		}
	}
}
