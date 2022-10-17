package helpers

import (
	"html"
	"log"
	"os/exec"
	"strings"
)

func GenerateUUID() (uuid string) {

	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	strUid := string(out)

	return html.EscapeString(strings.TrimSpace(strUid))
}
