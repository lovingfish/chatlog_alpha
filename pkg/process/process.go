package process

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v4/process"
)

// CheckSingleInstance checks if another instance is running using a PID file.
// If another instance is found, it prompts the user to force close it.
// Returns a cleanup function to be called on exit.
func CheckSingleInstance(workDir string) (func(), error) {
	pidFile := filepath.Join(workDir, "chatlog.pid")

	// Read existing PID file
	if content, err := os.ReadFile(pidFile); err == nil {
		pidStr := strings.TrimSpace(string(content))
		if pid, err := strconv.Atoi(pidStr); err == nil {
			// Check if process exists
			if exists, _ := process.PidExists(int32(pid)); exists {
				// Process exists, check if it's really us (optional, but good practice)
				// For now, just assume if PID exists it might be us or a zombie.
				// We can check process name if needed, but pid file is strong hint.
				
				fmt.Printf("Detected another instance running (PID: %d).\n", pid)
				fmt.Print("Do you want to force close it and continue? [y/N]: ")
				
				reader := bufio.NewReader(os.Stdin)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(strings.ToLower(input))

				if input == "y" || input == "yes" {
					if p, err := process.NewProcess(int32(pid)); err == nil {
						if err := p.Kill(); err != nil {
							return nil, fmt.Errorf("failed to kill process: %w", err)
						}
						fmt.Println("Process killed.")
					} else {
						// Process might have exited in the meantime
						fmt.Println("Process not found, continuing...")
					}
				} else {
					return nil, fmt.Errorf("application already running")
				}
			}
		}
	}

	// Write current PID
	currentPID := os.Getpid()
	if err := os.WriteFile(pidFile, []byte(strconv.Itoa(currentPID)), 0644); err != nil {
		return nil, fmt.Errorf("failed to write pid file: %w", err)
	}

	// Cleanup function
	return func() {
		os.Remove(pidFile)
	},
	nil
}
