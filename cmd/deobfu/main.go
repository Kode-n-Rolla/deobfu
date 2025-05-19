package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"deobfu/internal/cli"
	"deobfu/internal/detect"

	"github.com/manifoldco/promptui"
)

func main() {
	// Flags
	decodeFlag := flag.Bool("decode", false, "Decode common encoding like bases or hex")
	flag.BoolVar(decodeFlag, "dcd", false, "Alias for --decode")

	hashFlag := flag.Bool("hash-identify", false, "Identify possible hash type")
	flag.BoolVar(hashFlag, "hi", false, "Alias for --hash-identify")

	cipherFlag := flag.Bool("decypher", false, "Try to decypher cyphers")
	flag.BoolVar(cipherFlag, "dcyph", false, "Alias for --decypher")

	autoFlag := flag.Bool("auto", false, "Auto analyze the input")

	recurseFlag := flag.Bool("recurse-decode", false, "Recursively decode chain of encodings")
	flag.BoolVar(recurseFlag, "rdc", false, "Alias for --recurse-decode")

	inputString := flag.String("string", "", "Input string to analyze")
	flag.StringVar(inputString, "str", "", "Alias for --string")

	flag.Usage = func() {
		fmt.Println(`
DeObFU 🥋 — decoding, deobfuscation and decypher tool
		
Usage:
	deobfu
	deobfu [flags] [string]
		
Flags:
	-dcd, --decode              Try to decode common encodings (base64, base32, etc.)
	-hi,  --hash-identify       Try to identify hash type (MD5, SHA1, etc.)
	-dcyph, --decypher          Try to brute-force simple ciphers (Caesar, ROT13)
	-auto                       Automatically detect encodings and hashes
	-rdc, --recurse-decode      Recursively decode encoding chains
	-string                     Input string to analyze
	-h, --help                  Show this help message
		
	Supported Encodings: Base64, Base32, Base85, Hex, JWT (header + payload decoding)
	
	Supported hashes: MD5, SHA-1, SHA-256, SHA-512, NTLM
	    NOTE! MD5 and NTLM look identical. Use context to distinguish them.

	Supported ciphers: Caesar, ROT13, ROT47, Atbash
	`)
	}

	flag.Parse()

	// Nothing to pass
	if len(os.Args) == 1 {
		runInteractive()
		return
	}

	// Flag exist, string not
	if *inputString == "" && flag.NArg() > 0 {
		*inputString = flag.Arg(0)
	}

	if *inputString == "" {
		prompt := promptui.Prompt{
			Label: "Enter the string",
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}
		*inputString = result
	}

	// Choose logic
	switch {
	case *decodeFlag:
		fmt.Println("[*] Decode selected")
		fmt.Println("===================")
		detect.Decode(*inputString, false)
	case *hashFlag:
		fmt.Println("[*] Hash identification selected")
		fmt.Println("================================")
		detect.DetectHash(*inputString, false)
	case *cipherFlag:
		fmt.Println("[*] Decypher mode selected")
		fmt.Println("==========================")
		cli.RunDecypherMode(*inputString)
	case *autoFlag:
		fmt.Println("[*] Auto-analysis selected")
		fmt.Println("==========================")

		var wg sync.WaitGroup
		var mu sync.Mutex
		found := false

		wg.Add(2)

		go func() {
			defer wg.Done()
			if detect.Decode(*inputString, true) {
				mu.Lock()
				found = true
				mu.Unlock()
			}
		}()

		go func() {
			defer wg.Done()
			if detect.DetectHash(*inputString, true) {
				mu.Lock()
				found = true
				mu.Unlock()
			}
		}()

		wg.Wait()

		if !found {
			fmt.Println("[!] Nothing detected: no known encoding or hash matched.")
		}
	case *recurseFlag:
		fmt.Println("[*] Recursive decode selected")
		fmt.Println("=============================")
		detect.RecurseDecode(*inputString, 10, true)
	default:
		// If only string without flag
		runModeMenu(*inputString)
	}
}

// Request mode from the user
func runModeMenu(input string) {
	prompt := promptui.Select{
		Label: "Select mode",
		Items: []string{"Decode", "Hash Identify", "Decypher", "Auto", "Recursive Decode"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
	}

	switch result {
	case "Decode":
		fmt.Println("[*] Decode selected")
		fmt.Println("===================")
		detect.Decode(input, false)
	case "Hash Identify":
		fmt.Println("[*] Hash identification selected")
		fmt.Println("================================")
		detect.DetectHash(input, false)
	case "Decypher":
		fmt.Println("[*] Decypher mode selected")
		fmt.Println("==========================")
		cli.RunDecypherMode(input)
	case "Auto":
		fmt.Println("[*] Auto-analysis selected")
		fmt.Println("==========================")

		var wg sync.WaitGroup
		var mu sync.Mutex
		found := false

		wg.Add(2)

		go func() {
			defer wg.Done()
			if detect.Decode(input, true) {
				mu.Lock()
				found = true
				mu.Unlock()
			}
		}()

		go func() {
			defer wg.Done()
			if detect.DetectHash(input, true) {
				mu.Lock()
				found = true
				mu.Unlock()
			}
		}()

		wg.Wait()

		if !found {
			fmt.Println("[!] Nothing detected: no known encoding or hash matched.")
		}
	case "Recursive Decode":
		fmt.Println("[*] Recursive decode selected")
		fmt.Println("==============================")
		detect.RecurseDecode(input, 10, true)
	}
}

// Full interactive: ask mode and string
func runInteractive() {
	prompt := promptui.Prompt{
		Label: "Enter the string",
	}

	input, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	runModeMenu(input)
}
