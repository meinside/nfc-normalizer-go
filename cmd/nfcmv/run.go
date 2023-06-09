package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/meinside/version-go"
)

const (
	paramShortHelp = "-h"
	paramLongHelp  = "--help"

	paramShortVersion = "-v"
	paramLongVersion  = "--version"

	paramShortRecursive = "-r"
	paramLongRecursive  = "--recursive"

	paramShortDryrun = "-d"
	paramLongDryrun  = "--dryrun"

	paramShortInteractive = "-i"
	paramLongInteractive  = "--interactive"

	paramShortForceReplace = "-f"
	paramLongForceReplace  = "--forcereplace"

	paramShortVerbose = "-V"
	paramLongVerbose  = "--verbose"

	description = "Convert files/directories' names to NFC-normalized ones."
)

// for checking supported params
var _supportedParams = []string{
	paramShortHelp, paramLongHelp,
	paramShortVersion, paramLongVersion,
	paramShortRecursive, paramLongRecursive,
	paramShortDryrun, paramLongDryrun,
	paramShortInteractive, paramLongInteractive,
	paramShortForceReplace, paramLongForceReplace,
	paramShortVerbose, paramLongVerbose,
}

// print verbose messages or not
var _verbose = false

// check if given string has prefix of '-' or '--'
func isParam(param string) bool {
	return strings.HasPrefix(param, "-")
}

// check if given param is supported
func isSupportedParam(param string) bool {
	for _, supported := range _supportedParams {
		if supported == param {
			return true
		}
	}

	return false
}

// check if given short/long params exist in given params
func paramExists(params []string, short, long string) bool {
	for _, param := range params {
		if !isParam(param) { // skip non-param string
			continue
		}

		if !isSupportedParam(param) {
			showHelp(fmt.Errorf("not a supported parameter: %s", param))
		}

		if param == short || param == long {
			return true
		}
	}

	return false
}

// filter file/dir paths from given parameters
func filterFilepaths(params []string) (result []string) {
	result = []string{}

	for _, param := range params {
		if !isParam(param) {
			result = append(result, param)
		}
	}

	return result
}

// check if given path exists or not
func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if _verbose {
			log.Printf("error statting '%s': %s", path, err)
		}

		return !os.IsNotExist(err)
	}
	return true
}

// check if given filepaths are all valid
func checkFilepaths(paths []string) {
	for _, path := range paths {
		if !fileExists(path) {
			showHelp(fmt.Errorf("not a valid path: %s", path))
		}
	}
}

// show help message
func showHelp(err error) {
	message := description
	if err != nil {
		message = fmt.Sprintf("Error: %s", err.Error())
	}

	fmt.Printf(`%[1]s

Usage:

	$ %[2]s [parameters ...] [FILES_OR_DIRS ...]

* parameters:
	%[3]s, %[4]s: show this help message.
	%[5]s, %[6]s: show the version string of this application.
	%[7]s, %[8]s: convert all files and directories recursively.
	%[9]s, %[10]s: dry run, do not actually convert anything.
	%[11]s, %[12]s: convert interactively, asking [y/N] on each file/directory.
	%[13]s, %[14]s: force replace if there are files/directories with the same name.
	%[15]s, %[16]s: print verbose messages.
`,
		message,
		filepath.Base(os.Args[0]),
		paramShortHelp, paramLongHelp,
		paramShortVersion, paramLongVersion,
		paramShortRecursive, paramLongRecursive,
		paramShortDryrun, paramLongDryrun,
		paramShortInteractive, paramLongInteractive,
		paramShortForceReplace, paramLongForceReplace,
		paramShortVerbose, paramLongVerbose)

	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// show version string
func showVersion() {
	fmt.Printf(`%s
`, version.Minimum())

	os.Exit(0)
}

// print message to stdout
func _p(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}

	fmt.Printf(format, args...)
}

// print verbose message to stdout
func _v(format string, args ...interface{}) {
	if _verbose {
		log.Printf(format, args...)
	}
}

// run program with parameters
func runWithParams(params []string) {
	if paramExists(params, paramShortHelp, paramLongHelp) {
		showHelp(nil)
	} else if paramExists(params, paramShortVersion, paramLongVersion) {
		showVersion()
	}
	_verbose = paramExists(params, paramShortVerbose, paramLongVerbose)

	filepaths := filterFilepaths(params)
	if len(filepaths) > 0 {
		checkFilepaths(filepaths)

		recursive := paramExists(params, paramShortRecursive, paramLongRecursive)
		dryrun := paramExists(params, paramShortDryrun, paramLongDryrun)
		interactive := paramExists(params, paramShortInteractive, paramLongInteractive)
		forceReplace := paramExists(params, paramShortForceReplace, paramLongForceReplace)

		// convert all items with given params
		converted := convertAll(filepaths, recursive, dryrun, interactive, forceReplace)

		if dryrun {
			_p("[dryrun] %d item(s) will be converted.", converted)
		} else {
			_p("> converted %d item(s).", converted)
		}
	} else {
		showHelp(fmt.Errorf("no file/directory path was given."))
	}
}
