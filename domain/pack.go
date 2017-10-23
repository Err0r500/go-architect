package domain

import (
	"go/build"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var currPackage string

func init() {
	setCurrPackageImportPath()
}

type Pack struct {
	packagePath
	packageClass
}

func NewPackFromPath(p string) *Pack {
	pP := packagePath(p)
	return &Pack{
		packagePath:  pP,
		packageClass: pP.getPackageClass(),
	}
}

func (p Pack) String() string {
	return string(p.packagePath) + " (" + string(p.packageClass) + ")"
}

func (p Pack) GetPath() string {
	return string(p.packagePath)
}
func (p Pack) GetClass() string {
	return string(p.packageClass)
}

type packageClass string

const (
	corePackage       packageClass = "corePackage"
	internalpackage                = "projectPackage"
	thirdPartyPackage              = "thirdPartyPackage"
)

type packagePath string

func (pP packagePath) getPackageClass() packageClass {
	var internal = regexp.MustCompile(currPackage + `.*`)
	if internal.MatchString(string(pP)) {
		return internalpackage
	}

	if isStandardImportPath(string(pP)) {
		return corePackage
	}

	return thirdPartyPackage
}

// copy-pasted from go source code ! :)

// isStandardImportPath reports whether $GOROOT/src/path should be considered
// part of the standard distribution. For historical reasons we allow people to add
// their own code to $GOROOT instead of using $GOPATH, but we assume that
// code will start with a domain name (dot in the first element).
func isStandardImportPath(path string) bool {
	i := strings.Index(path, "/")
	if i < 0 {
		i = len(path)
	}
	elem := path[:i]
	return !strings.Contains(elem, ".")
}

func setCurrPackageImportPath() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	currPackage = strings.Replace(filepath.Dir(ex), build.Default.GOPATH+"/src/", "", -1)
}
