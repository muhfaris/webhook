package pkg

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func ExecuteCommand(command string, workdir string) error {
	// Execute a shell command using 'bash -c':
	// - 'bash': Launches the Bash shell to interpret the command.
	// - '-c': Tells Bash to read and execute the command that follows.
	logrus.Infof("execute command: %s", command)
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = workdir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func HashSha1(msg string) string {
	// Create a new SHA-1 hash object
	h := sha1.New()
	h.Write([]byte(msg))

	// Compute the hash
	sha1Hash := h.Sum(nil)
	hexString := hex.EncodeToString(sha1Hash)
	return hexString
}
