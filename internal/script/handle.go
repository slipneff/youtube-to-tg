package script

import (
	"os"
	"os/exec"

	"github.com/google/uuid"
)

func (s *Script) Run(youtubeURL string) error {
	id := uuid.NewString()
	defer os.Remove(id)
	cmd := exec.Command("./script/download_and_convert.sh", youtubeURL, id)
	err := cmd.Run()

	if err != nil {
		os.Exit(1)
		return err
	}
	return nil
}
