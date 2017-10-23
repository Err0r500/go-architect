package main

import (
	"go/build"
	"log"
	"path/filepath"
	"strings"

	"github.com/err0r500/go-architect/domain"
	AstM "github.com/err0r500/go-architect/interfaces/astManager"
	FM "github.com/err0r500/go-architect/interfaces/fileManager"
	"github.com/err0r500/go-architect/interfaces/json"
	TE "github.com/err0r500/go-architect/interfaces/treeExplorer"
)

type ImportsFinderInteractor struct {
	tE    TreeExplorer
	fM    FileManager
	astM  AstManager
	jsonW JSONwriter
}

func main() {

	i := ImportsFinderInteractor{
		tE:    TE.TreeExplorer{},
		fM:    FM.FileManager{},
		astM:  AstM.AstManager{},
		jsonW: json.D3formatter{},
	}

	graph := i.GetAllImports()
	jsonPayload, _ := i.jsonW.ToJSON(graph)
	log.Print(jsonPayload)
	err := i.fM.Write(domain.File{Path: "/tmp/testGraph2.json", Content: []byte(jsonPayload)})
	log.Print(err)
}

// juste un gros bloc pour montrer l'idée initiale, surement naîve
func (i ImportsFinderInteractor) GetAllImports() *domain.Graph {
	dirs, _ := i.tE.GetDirsInTree(".")
	graph := &domain.Graph{}

	for _, dir := range *dirs {
		currVerticeID := appendVertice(graph, dir)

		fPathes, _ := i.tE.GetFilesInDir(dir)

		for _, fPath := range *fPathes {
			f := domain.File{Path: dir + "/" + fPath}
			imports, _ := i.astM.GetImportsFromFile(f.GetPath())

			for _, importPath := range *imports {
				graph.BuildGraph(currVerticeID, domain.NewPackFromPath(importPath))
			}
		}
	}
	return graph
}

func appendVertice(graph *domain.Graph, srcFolder string) int {
	src, _ := filepath.Abs(srcFolder)
	src = strings.Replace(src, build.Default.GOPATH+"/src/", "", -1)
	return graph.AppendRootNode(domain.NewPackFromPath(src))
}
