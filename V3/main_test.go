package V3

import (
	"AlgebraCalculator/log"
	"io/ioutil"
	"testing"
)

func setUpForTest() {

	var rules []string
	rules = append(rules, readFile("rules/simpRulesExpand.txt"))
	rules = append(rules, readFile("rules/simpRulesSumUp.txt"))

	Init(rules)
}
func readFile(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	content := string(buf)
	return content
}

var parsetestTerms = []struct {
	input string
}{
	{"a = 1"},
	{"a = 1 + 1"},
	{"a = 1 * 1"},
	{"a = 2 * 2"},
	{"a = 1864385 * 843"},
	{"a<t> = 3 * t"},
	{"a<t> = 4 * ( 7 + 5 )"},
	{"a<t> = t + t + t + t"},
	{"a<t> = t + ( t + t ) + t"},
	{"a<t> = t + ( ( t + t ) + t )"},
	{"a<t> = ( t + t + t ) + t"},
}

func TestParseing(t *testing.T) {
	setUpForTest()

	for i, testTerm := range parsetestTerms {
		term, err := parseTerm(testTerm.input, nil)
		if err != nil {
			t.Error(err)
		}
		log.PrintLog()

		term.printTerm()
		out := log.GetLog()
		if out != testTerm.input {
			t.Errorf("Term %d: expected %s actual %s", i, testTerm.input, out)
		}
	}

}

var runtestTerms = []struct {
	input    string
	expected string
}{
	{"a = 1", "a = 1"},
	{"a = 1 + 1", "a = 2"},
	{"a = 1 * 1", "a = 1"},
	{"a = 2 * 2", "a = 4"},
	{"a = 1864385 * 843", "a = 1571676555"},
	{"a<t> = 3 * t", "a<t> = 3 * t"},
	{"a<t> = 4 * ( 7 + 5 )", "a<t> = 48"},
	{"a<t> = t + t + t + t", "a<t> = 4 * t"},
	{"a<t> = t + ( t + t ) + t", "a<t> = 4 * t"},
	{"a<t> = t + ( ( t + t ) + t )", "a<t> = 4 * t"},
	{"a<t> = ( t + t + t + t )", "a<t> = 4 * t"},
	{"a<t> = ( t * t + t + t )", "a<t> = 2 * t + t pow 2"},
	{"g<a b c> = ( a + b ) * ( a + b )", "g<a b c> = 1 * a pow 2 + 2 * a * b + 1 * b pow 2"},
	{"a t = sin t", "a<t> = sin<t>"},
}

func TestRunning(t *testing.T) {
	setUpForTest()

	for i, testTerm := range runtestTerms {
		var termNode iNode
		termNode, err := parseTerm(testTerm.input, nil)
		if err != nil {
			t.Error(err)
		}
		termNode = solveTerm(termNode, rules)
		logged := log.GetLog()

		termNode.(*term).printTerm()
		out := log.GetLog()
		if out != testTerm.expected {
			t.Errorf("Term %d: expected %s actual %s", i, testTerm.expected, out)
			t.Error(logged)
		}
	}

}
