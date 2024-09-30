package script

import (
	"os/exec"

	"github.com/google/uuid"
)

func (s *Script) Run(youtubeURL string) (string, error) {
	id := uuid.NewString() + ".mp3"
	// defer os.Remove(id)
	cmd := exec.Command("./script/download_and_convert.sh", youtubeURL, id)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return id, nil
}
