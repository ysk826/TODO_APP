package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// OpenFileは指定されたファイルを開く
	// 第一引数: ファイル名
	// 第二引数: 読み書と作成と追記の権限を付与、なければ作成
	// 第三引数: ファイルのパーミッション
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	//　MultiWriterで標準出力とログファイルの両方に書き込むライターを作成
	// multiLogFileはio.Writer型で標準出力とファイル出力を同時に行う
	// SetFlagsはログのプレフィックスに日付、時刻、ファイル名を付与
	// SetOutputはログの出力先を作成したライターに設定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
