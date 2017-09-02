package differs

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/container-diff/utils"
	"github.com/golang/glog"
)

type AptAnalyzer struct {
}

func (a AptAnalyzer) Name() string {
	return "AptAnalyzer"
}

// AptDiff compares the packages installed by apt-get.
func (a AptAnalyzer) Diff(image1, image2 utils.Image) (utils.Result, error) {
	diff, err := singleVersionDiff(image1, image2, a)
	return diff, err
}

func (a AptAnalyzer) Analyze(image utils.Image) (utils.Result, error) {
	analysis, err := singleVersionAnalysis(image, a)
	return analysis, err
}

func (a AptAnalyzer) getPackages(image utils.Image) (map[string]utils.PackageInfo, error) {
	path := image.FSPath
	packages := make(map[string]utils.PackageInfo)
	if _, err := os.Stat(path); err != nil {
		// invalid image directory path
		return packages, err
	}
	statusFile := filepath.Join(path, "var/lib/dpkg/status")
	if _, err := os.Stat(statusFile); err != nil {
		// status file does not exist in this layer
		return packages, nil
	}
	if file, err := os.Open(statusFile); err == nil {
		// make sure it gets closed
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		var currPackage string
		for scanner.Scan() {
			currPackage = parseLine(scanner.Text(), currPackage, packages)
		}
	} else {
		return packages, err
	}

	return packages, nil
}

func parseLine(text string, currPackage string, packages map[string]utils.PackageInfo) string {
	line := strings.Split(text, ": ")
	if len(line) == 2 {
		key := line[0]
		value := line[1]

		switch key {
		case "Package":
			return value
		case "Version":
			if packages[currPackage].Version != "" {
				glog.Warningln("Multiple versions of same package detected.  Diffing such multi-versioning not yet supported.")
				return currPackage
			}
			modifiedValue := strings.Replace(value, "+", " ", 1)
			currPackageInfo, ok := packages[currPackage]
			if !ok {
				currPackageInfo = utils.PackageInfo{}
			}
			currPackageInfo.Version = modifiedValue
			packages[currPackage] = currPackageInfo
			return currPackage

		case "Installed-Size":
			currPackageInfo, ok := packages[currPackage]
			if !ok {
				currPackageInfo = utils.PackageInfo{}
			}
			var size int64
			var err error
			size, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				glog.Errorf("Could not get size for %s: %s", currPackage, err)
				size = -1
			}
			// Installed-Size is in KB, so we convert it to bytes to keep consistent with the tool's size units
			currPackageInfo.Size = size * 1024
			packages[currPackage] = currPackageInfo
			return currPackage
		default:
			return currPackage
		}
	}
	return currPackage
}
