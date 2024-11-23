package logconfig

import (
	"encoding/json"
	"log"
	"os"
)

// JSONファイルの構造をマッピング
type Config struct {
	Log struct {
		Output string
		Format string
		Flags  []string
		// Level  string
	}
}

// フラグを集約して、数値型に変換
func parseFlags(flagList []string) int {
	flagMap := map[string]int{
		"Ldate":         log.Ldate,
		"Ltime":         log.Ltime,
		"Lmicroseconds": log.Lmicroseconds,
		"Llongfile":     log.Llongfile,
		"Lshortfile":    log.Lshortfile,
		"LUTC":          log.LUTC,
		"Lmsgprefix":    log.Lmsgprefix,
	}
	flags := 0
	for _, flag := range flagList {
		if val, exists := flagMap[flag]; exists {
			flags |= val
		}
	}
	return flags
}

func SetupLogger(configFile string) (*os.File, error) {
	// 設定ファイルを開く
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// JSONをデコード
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	// 出力先を設定
	var logOutput *os.File
	if config.Log.Output == "stdout" {
		logOutput = os.Stdout
	} else if config.Log.Output == "stderr" {
		logOutput = os.Stderr
	} else {
		logOutput, err = os.OpenFile(config.Log.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		// defer logOutput.Close()
		// defer logOutput.Sync()
	}
	log.SetOutput(logOutput)

	// フラグを設定
	log.SetFlags(parseFlags(config.Log.Flags))

	return logOutput, nil
}
