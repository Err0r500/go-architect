package AstManager

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/err0r500/go-architect/testHelpers"
)

var src = `
package thePackageName
import (
	"flag"
	"fmt"
	"path/filepath"

	mA "github.com/err0r500/codeAnalyzer/analyzer"
)
const c = 1.0
var X = f(3.14)*2 + c
func myfunc(myInterface){
	return
}
type myInterface interface {
	doThis()
}
`

func TestAstManager_GetImports(t *testing.T) {
	expected := []string{
		"flag",
		"path/filepath",
		"fmt",
		"github.com/err0r500/codeAnalyzer/analyzer",
	}

	testFilePath := "./testFile.go"

	ioutil.WriteFile(testFilePath, []byte(src), 0644)
	defer os.Remove(testFilePath)
	astM := AstManager{}
	returned, err := astM.GetImportsFromFile(testFilePath)
	if err != nil {
		t.Error(err)
	}

	if err := testHelpers.CheckStringSliceEqual(expected, *returned); err != nil {
		t.Error(err)
	}
}
