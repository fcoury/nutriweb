package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os/exec"
	"sort"
	"strings"
	"sync"
)

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func SortMap(params map[string]string) map[string]string {
	sortedMap := make(map[string]string)
	var keys []string

	for k, _ := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		sortedMap[k] = params[k]
	}

	return sortedMap
}

func Sign(base string, secret string) string {
	hasher := hmac.New(sha1.New, []byte(secret))
	io.WriteString(hasher, base)

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
