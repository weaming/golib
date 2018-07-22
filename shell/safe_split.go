package shell

// copy from https://gist.github.com/jmervine/d88c75329f98e09f5c87
// safeSplit handles quoting well for commands for use with github.com/jmervine/exec/v2
//
// Examples:
// > safeSplit("/bin/bash bash -l -c 'echo \"foo bar bah bin\"'")
// => []string{"/bin/bash", "-l", "-c", "echo \"foo bar bah bin\""}
// > safeSplit("docker run --rm -it some/image bash -c \"npm test\"")
// => []string{"docker", "run", "--rm", "-it", "some/image", "bash", "-c", "npm test"}
//----
// package main
// import "github.com/jmervine/exec/v2"
// import "fmt"
// action := safeSplit("docker run --rm -it some/image bash -c \"npm test\"")
// out, err := exec.Exec(action[0], action[1:]...)
// if err != nil { panic(err) }
// fmt.Println(out)
//----

import "strings"

func safeSplit(s string) []string {
	split := strings.Split(s, " ")

	var result []string
	var inquote string
	var block string
	for _, i := range split {
		if inquote == "" {
			if strings.HasPrefix(i, "'") || strings.HasPrefix(i, "\"") {
				inquote = string(i[0])
				block = strings.TrimPrefix(i, inquote) + " "
			} else {
				result = append(result, i)
			}
		} else {
			if !strings.HasSuffix(i, inquote) {
				block += i + " "
			} else {
				block += strings.TrimSuffix(i, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}

	return result
}
