package detect

import (
	"fmt"
	"regexp"
	"strings"
)

// HashPattern represents a hash type and its identifying properties
type HashPattern struct {
	Name     string
	Length   int
	Regex    *regexp.Regexp
	CrackCmd []string // Suggested commands for cracking
}

// DetectHash tries to identify a hash type based on length and pattern
func DetectHash(input string, silent bool) bool {
	input = strings.TrimSpace(input)

	hashes := []HashPattern{
		{
			Name:   "MD5",
			Length: 32,
			Regex:  regexp.MustCompile(`^[a-fA-F0-9]{32}$`),
			CrackCmd: []string{
				"john --format=raw-md5 --wordlist=rockyou.txt hash.txt",
				"hashcat -m 0 -a 0 hash.txt rockyou.txt",
				"https://hashes.com/en/decrypt/hash",
				"https://crackstation.net/",
			},
		},
		{
			Name:   "SHA-1",
			Length: 40,
			Regex:  regexp.MustCompile(`^[a-fA-F0-9]{40}$`),
			CrackCmd: []string{
				"john --format=raw-sha1 --wordlist=rockyou.txt hash.txt",
				"hashcat -m 100 -a 0 hash.txt rockyou.txt",
				"https://hashes.com/en/decrypt/hash",
				"https://crackstation.net/",
			},
		},
		{
			Name:   "SHA-256",
			Length: 64,
			Regex:  regexp.MustCompile(`^[a-fA-F0-9]{64}$`),
			CrackCmd: []string{
				"john --format=raw-sha256 --wordlist=rockyou.txt hash.txt",
				"hashcat -m 1400 -a 0 hash.txt rockyou.txt",
				"https://hashes.com/en/decrypt/hash",
				"https://crackstation.net/",
			},
		},
		{
			Name:   "SHA-512",
			Length: 128,
			Regex:  regexp.MustCompile(`^[a-fA-F0-9]{128}$`),
			CrackCmd: []string{
				"john --format=raw-sha512 --wordlist=rockyou.txt hash.txt",
				"hashcat -m 1700 -a 0 hash.txt rockyou.txt",
				"https://hashes.com/en/decrypt/hash",
				"https://crackstation.net/",
			},
		},
		{
			Name:   "NTLM",
			Length: 32,
			Regex:  regexp.MustCompile(`^[a-fA-F0-9]{32}$`),
			CrackCmd: []string{
				"john --format=nt --wordlist=rockyou.txt hash.txt",
				"hashcat -m 1000 -a 0 hash.txt rockyou.txt",
				"https://hashes.com/en/decrypt/hash",
				"https://crackstation.net/",
			},
		},
	}

	matched := false
	for _, hash := range hashes {
		if len(input) == hash.Length && hash.Regex.MatchString(input) {
			fmt.Printf("[+] Possible hash type: %s\n", hash.Name)
			fmt.Println("    Suggested commands:")
			for _, cmd := range hash.CrackCmd {
				fmt.Printf("    %s\n", cmd)
			}
			matched = true
		}
	}

	if !matched && !silent {
		fmt.Println("[!] Unknown hash format or not supported yet.")
	}

	return matched
}
