package cli

import (
	"fmt"

	"deobfu/internal/detect"

	"github.com/manifoldco/promptui"
)

func RunDecypherMode(input string) {
	// Prompt: choose cipher
	prompt := promptui.Select{
		Label: "Select cipher to try",
		Items: []string{"Caesar (ROT-N)", "ROT13", "ROT47", "Atbash"},
	}

	_, choice, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	switch choice {
	case "Caesar (ROT-N)":
		detect.BruteCaesar(input)
	case "ROT13":
		result := detect.ROT13(input)
		fmt.Println("[+] ROT13:")
		fmt.Println(result)
	case "ROT47":
		result := detect.ROT47(input)
		fmt.Println("[+] ROT47:")
		fmt.Println(result)
	case "Atbash":
		result := detect.Atbash(input)
		fmt.Println("[+] Atbash:")
		fmt.Println(result)
	}
}
