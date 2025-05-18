package core

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

func Scan(path string) error {
	var gitRepositories []string
	var mu sync.Mutex     // Protects gitRepositories
	var wg sync.WaitGroup // Waits for all goroutines to finish
	numWorkers := runtime.NumCPU() * 4
	c := make(chan string, numWorkers*100)

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Seed initial path
	wg.Add(1)
	c <- absPath
	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for dir := range c {
				searchDirectory(dir, &gitRepositories, c, &mu, &wg)
			}
		}()
	}

	// Wait for all work to finish
	wg.Wait()
	close(c)

	fmt.Println("Found .git repositories:")
	for _, repo := range gitRepositories {
		fmt.Println(repo)
	}
	fmt.Println(numWorkers)
	return nil
}

func searchDirectory(path string, gitRepositories *[]string, c chan string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Currently processing: %s\n", path)

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error happened Hilfe")
		return // silently ignore errors like permission denied
	}

	for _, e := range entries {
		if e.IsDir() {
			subdir := filepath.Join(path, e.Name())
			if e.Name() == ".git" {
				mu.Lock()
				*gitRepositories = append(*gitRepositories, path) // not subdir, since .git is a subfolder of repo root
				mu.Unlock()
			} else if !strings.Contains(e.Name(), ".") {
				wg.Add(1)
				c <- subdir
			}
		}
	}
}
