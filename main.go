package protoc_gowrap

import (
	"io"
	"io/ioutil"

	"github.com/RangelReale/protoc-gowrap/generator"
	"github.com/golang/protobuf/proto"
)

func Main(g *generator.Generator, reader io.Reader, writer io.Writer) {

	var data []byte
	var err error

	data, err = ioutil.ReadAll(reader)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = writer.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}
