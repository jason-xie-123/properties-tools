package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	packageVersion "properties-cli/version"

	"github.com/urfave/cli/v2"
)

// propRead：读取 key 的值（只取第一次出现的）
func propRead(keyName, propPath string) string {
	data, err := os.ReadFile(propPath)
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	prefix := keyName + "="

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		lineNoLeft := strings.TrimLeft(line, " \t")
		if strings.HasPrefix(lineNoLeft, prefix) {
			return strings.TrimSpace(strings.TrimPrefix(lineNoLeft, prefix))
		}
	}
	return ""
}

// propWrite：只替换第一次出现的 key；不存在则追加
func propWrite(keyName, keyValue, propPath string) {
	data, _ := os.ReadFile(propPath)

	scanner := bufio.NewScanner(bytes.NewReader(data))
	var lines []string

	prefix := keyName + "="
	newLine := keyName + "=" + keyValue
	found := false

	for scanner.Scan() {
		raw := scanner.Text()
		trim := strings.TrimSpace(raw)

		if trim == "" || strings.HasPrefix(trim, "#") || strings.HasPrefix(trim, ";") {
			lines = append(lines, raw)
			continue
		}

		lineNoLeft := strings.TrimLeft(raw, " \t")
		if !found && strings.HasPrefix(lineNoLeft, prefix) {
			lines = append(lines, newLine)
			found = true
		} else {
			lines = append(lines, raw)
		}
	}

	if !found {
		lines = append(lines, newLine)
	}

	// 按 Go / 系统常规方式写回（统一 \n）
	_ = os.WriteFile(propPath, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func main() {
	AppName := "properties-cli"

	app := &cli.App{
		Name:    AppName,
		Usage:   "CLI Tool to read and write properties files",
		Version: packageVersion.Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "read",
				Usage: "read flag",
			},
			&cli.BoolFlag{
				Name:  "write",
				Usage: "write flag",
			},
			&cli.StringFlag{
				Name:  "key",
				Usage: "property key name",
			},
			&cli.StringFlag{
				Name:  "value",
				Usage: "property key value",
			},
			&cli.StringFlag{
				Name:  "path",
				Usage: "path to properties file",
			},
		},
		Action: func(c *cli.Context) error {
			readFlag := c.Bool("read")
			writeFlag := c.Bool("write")
			keyName := c.String("key")
			keyValue := c.String("value")
			filePath := c.String("path")

			if !readFlag && !writeFlag {
				return fmt.Errorf("either read or write flag must be set")
			}

			if keyName == "" {
				return fmt.Errorf("key name is required")
			}

			if filePath == "" {
				return fmt.Errorf("property file path is required")
			}

			if !fileExists(filePath) {
				return fmt.Errorf("property file does not exist at path: %s", filePath)
			}

			if readFlag {
				value := propRead(keyName, filePath)
				fmt.Print(value)
			} else if writeFlag {
				propWrite(keyName, keyValue, filePath)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
