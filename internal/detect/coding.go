package detect

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"unicode"
)

func isMostlyPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}

	printable := 0
	for _, r := range s {
		if unicode.IsPrint(r) && r != '\uFFFD' { // exclude replacement character
			printable++
		}
	}

	ratio := float64(printable) / float64(len(s))
	return ratio >= 0.85 // allow ~15% garbage
}

// Main func. All detect run in parallel
func Decode(input string, silent bool) bool {
	var wg sync.WaitGroup
	results := make(chan string, 10)

	found := false

	type decoderFunc func(string, bool) (bool, string, string)

	detectors := []decoderFunc{
		DetectBase64URL,
		DetectBase64,
		DetectBase32,
		DetectBase85,
		DetectHex,
		DetectJWT,
	}

	for _, detector := range detectors {
		wg.Add(1)

		go func(det decoderFunc) {
			defer wg.Done()
			ok, decoded, label := det(input, silent)
			if ok {
				results <- fmt.Sprintf("[+] %s detected\n%s", label, decoded)
			}
		}(detector)
	}

	// Close channel only after when go routines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
		found = true
	}

	if !found && !silent {
		fmt.Println("[!] No known encoding matched.")
	}

	return found
}

// === Base64 ===
func DetectBase64(input string, silent bool) (bool, string, string) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return false, "", ""
	}

	result := string(decoded)

	if !isMostlyPrintable(result) {
		return false, "", ""
	}

	return true, result, "Base64"
}

// === Base64URL ===
func DetectBase64URL(input string, silent bool) (bool, string, string) {
	if strings.Contains(input, "=") { // Avoid Base64URL if '=' is present
		return false, "", ""
	}

	decoded, err := base64.RawURLEncoding.DecodeString(input)
	if err != nil {
		//fmt.Println("Decode error:", err)
		return false, "", ""
	}

	result := string(decoded)

	if !isMostlyPrintable(result) {
		return false, "", ""
	}

	return true, result, "Base64URL"
}

// === Base32 ===
func DetectBase32(input string, silent bool) (bool, string, string) {
	decoded, err := base32.StdEncoding.DecodeString(input)
	if err != nil {
		return false, "", ""
	}

	result := string(decoded)

	if !isMostlyPrintable(result) {
		return false, "", ""
	}

	return true, result, "Base32"
}

// === Base85 ===
func DetectBase85(input string, silent bool) (bool, string, string) {
	decoded := make([]byte, len(input)*4/5)
	n, _, err := ascii85.Decode(decoded, []byte(input), true)
	if err != nil {
		return false, "", ""
	}

	result := string(decoded[:n])

	if !isMostlyPrintable(result) {
		return false, "", ""
	}

	return true, result, "Base85"
}

// === Hex ===
func DetectHex(input string, silent bool) (bool, string, string) {
	if len(input)%2 != 0 {
		return false, "", ""
	}

	decoded, err := hex.DecodeString(input)
	if err != nil {
		return false, "", ""
	}

	result := string(decoded)

	if !isMostlyPrintable(result) {
		return false, "", ""
	}

	return true, result, "Hex"
}

// === JWT ===
func DetectJWT(input string, silent bool) (bool, string, string) {
	parts := strings.Split(input, ".")
	if len(parts) != 3 {
		return false, "", ""
	}

	decoded := ""
	for i, part := range parts[:2] {
		padded := padBase64(part)
		data, err := base64.RawURLEncoding.DecodeString(padded)
		if err != nil {
			return false, "", ""
		}

		var out map[string]interface{}
		if err := json.Unmarshal(data, &out); err != nil {
			return false, "", ""
		}

		partJSON, _ := json.MarshalIndent(out, "", "  ")
		label := "Header"
		if i == 1 {
			label = "Payload"
		}
		decoded += fmt.Sprintf("[%s]\n%s\n", label, string(partJSON))
	}

	return true, decoded, "JWT"
}

// Add padding for base64url
func padBase64(s string) string {
	switch len(s) % 4 {
	case 2:
		return s + "=="
	case 3:
		return s + "="
	}
	return s
}

// RecurseDecode tries to decode the input multiple times using known encodings
func RecurseDecode(input string, maxDepth int, silent bool) {
	current := input
	steps := []string{input}
	depth := 0
	changed := true

	for changed && depth < maxDepth {
		changed = false
		type decoderFunc func(string, bool) (bool, string, string)

		detectors := []decoderFunc{
			DetectBase64URL,
			DetectBase64,
			DetectBase32,
			DetectBase85,
			DetectHex,
			DetectJWT,
		}

		for _, det := range detectors {
			ok, decoded, label := det(current, silent)
			if ok && decoded != current {
				fmt.Printf("[+] Step %d: %s detected\n", depth+1, label)
				fmt.Printf("%s\n\n", decoded)
				current = decoded
				steps = append(steps, decoded)
				changed = true
				break // only one step during iteration
			}
		}
		depth++
	}

	if len(steps) == 1 {
		fmt.Println("[!] Nothing could be recursively decoded.")
	}
}
