// Package winenv helps Go developers on windows manage their ENV files
// Without having to type of long strings in the console
package winenv

import (
	"bufio"
	"os"
	"strings"
)

// Parse takes an optional filename path to your .env file(s) and places them in the env
// of your application
func Parse(paths ...string) error {
	if len(paths) == 0 {
		// Use the default path
		defaultPath := "./.env"
		paths = append(paths, defaultPath)
	}

	for _, filePath := range paths {
		// Open a connect to the file
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// read each line
		lineReader := bufio.NewScanner(file)
		for lineReader.Scan() {
			// seperate each line by the "="
			line := lineReader.Text()

			if len(line) >= 3 {
				// set the name equal to the 0th param
				// and the value equal to the ...rest
				ss := strings.Split(line, "=")
				key := ss[0]
				val := strings.Join(ss[1:], "=")

				os.Setenv(key, val)
			}
		}

	}

	return nil
}
