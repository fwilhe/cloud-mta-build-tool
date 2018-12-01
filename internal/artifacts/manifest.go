package artifacts

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"cloud-mta-build-tool/mta"

	"github.com/pkg/errors"
)

// The deployment descriptor should be located within the META-INF folder of the JAR.
// The MANIFEST.MF file should contain at least a name section for each MTA module contained in the archive.
// Following the JAR specification, the value of a name must be a relative path to a file or directory,
// or an absolute URL referencing data outside the archive.
// It is required to add a row MTA-module: <modulename> to each name section that corresponds to an MTA module, and
// to bind archive file locations to module names as used in the deployment descriptor.
// The name sections with the MTA module attribute indicate the path to the file or directory which represents a module within the archive
// This is used by the deploy service to track the build project.

const (
	newLine         = "\n"
	contentType     = "Content-Type: "
	mtaModule       = "MTA-Module: "
	moduleName      = "Name: "
	applicationZip  = "application/zip"
	manifestVersion = "manifest-Version: 1.0"
	pathSep         = string(os.PathSeparator)
	dataZip         = pathSep + "data.zip"
)

// setManifetDesc - Set the MANIFEST.MF file
func setManifetDesc(file io.Writer, mtaStr []*mta.Module, modules []string) error {
	// TODO create dynamically
	_, err := fmt.Fprint(file, manifestVersion+newLine)
	if err != nil {
		return errors.Wrap(err, "META INFO generation failed")
	}
	// TODO set the version from external config for automatic version bump during release
	_, err = fmt.Fprint(file, "Created-By: SAP Application Archive Builder 0.0.1")
	if err != nil {
		return errors.Wrap(err, "META INFO generation failed")
	}
	for _, mod := range mtaStr {
		// Print only the required module to support the partial build
		if len(modules) > 0 && mod.Name == modules[0] {
			err := printToFile(file, mod)
			if err != nil {
				return errors.Wrap(err, "Error while printing values to mtad file")
			}
			break
		} else if len(modules) == 0 {
			// Print all the modules
			err := printToFile(file, mod)
			if err != nil {
				return errors.Wrap(err, "Error while printing values to mtad file")
			}
		}
	}
	return nil
}

// printToFile - Print to manifest.mf file
func printToFile(file io.Writer, mtaStr *mta.Module) error {
	_, err := fmt.Fprint(file, newLine+newLine, filepath.ToSlash(moduleName+mtaStr.Name+dataZip),
		newLine, mtaModule+mtaStr.Name, newLine, contentType+applicationZip)
	return err
}
