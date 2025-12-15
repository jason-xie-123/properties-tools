package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	packageVersion "properties-cli/version"

	"github.com/urfave/cli/v2"
)

func detectPlatformEOL() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

// propRead：读取 key 的值（只取第一次出现的）
func propRead(keyName, propPath string) string {
	data, err := os.ReadFile(propPath)
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	prefix := keyName + "="

	for scanner.Scan() {
		// Scanner 会去掉行尾的 \n，并且对 \r\n 的 \r 也会处理掉（不会出现在 Text() 里）
		raw := scanner.Text()
		line := strings.TrimSpace(raw)

		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		lineNoLeft := strings.TrimLeft(raw, " \t")
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
	lines := make([]string, 0, 64)

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

	eol := detectPlatformEOL()

	// 按平台常规方式写回：Windows 用 \r\n，其它用 \n；并确保文件末尾有换行
	_ = os.WriteFile(propPath, []byte(strings.Join(lines, eol)+eol), 0o644)
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
