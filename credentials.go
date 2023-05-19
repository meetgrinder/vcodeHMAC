package vcodeHMAC

import (
	"bufio"
	"os"
	"strings"
)

func getCredentials(fileString string, profile string) ([2]string, error) {
	var credentials [2]string
    profileLine := "[" + profile + "]"
    foundProfile := false
    foundKey := false
    foundSecret := false

	file, err := os.Open(fileString)
	if err != nil {
		return credentials, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        line := scanner.Text()
        line = strings.TrimSpace(line)
        if len(line) == 0 || line[0] == '#' {
            // skip empty lines or lines starting with '#'
            continue
        }

        if line == profileLine {
            foundProfile = true
            continue
        }   

        // dont even bother with anything until youve found the profile
        if foundProfile == false {
            continue
        }
        

		//We remove spaces to account for discrepancies in user configuration of creds file
		if strings.Contains(scanner.Text(), "veracode_api_key_id") {
			removeSpaces := strings.Replace(scanner.Text(), " ", "", -1)
			credentials[0] = strings.Replace(removeSpaces, "veracode_api_key_id=", "", -1)
            foundKey = true
		} else if strings.Contains(scanner.Text(), "veracode_api_key_secret") {
			removeSpaces := strings.Replace(scanner.Text(), " ", "", -1)
			credentials[1] = strings.Replace(removeSpaces, "veracode_api_key_secret=", "", -1)
            foundSecret = true
		}

        if foundKey == true && foundSecret == true {
            break
        }
	}

	if err := scanner.Err(); err != nil {
		return credentials, err
	}

	return credentials, nil
}
