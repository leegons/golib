package golib

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func ReadAllLineCB(filename string, callback func(string) error) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	var (
		linebuf bytes.Buffer
		buf     []byte
		full    bool
	)
	for {
		buf, full, err = reader.ReadLine()
		if err != nil {
			break
		}
		if full {
			linebuf.Write(buf)
			continue
		}
		if linebuf.Len() != 0 {
			linebuf.Write(buf)
			if err := callback(linebuf.String()); err != nil {
				return err
			}
			linebuf.Reset()
			continue
		}
		if err := callback(string(buf)); err != nil {
			return err
		}
	}
	if err != io.EOF {
		return err
	}
	return nil
}

/* 读取文件所有的行并直接返回 */
func ReadAllLine(filename string) ([]string, error) {
	results := []string{}
	err := ReadAllLineCB(filename, func(line string) error {
		results = append(results, line)
		return nil
	})
	return results, err
}

func ReadAllLineSplitCB(filename string, sep string, expectCol int, callback func([]string) error) error {
	return ReadAllLineCB(filename, func(line string) error {
		col := strings.Split(line, sep)
		if len(col) != expectCol {
			return ErrUnexpectColumnCounts
		}
		return callback(col)
	})
}

func ReadAllLineSplit(filename string, sep string, expectCol int) ([][]string, error) {
	results := [][]string{}
	err := ReadAllLineSplitCB(filename, sep, expectCol, func(col []string) error {
		results = append(results, col)
		return nil
	})
	return results, err
}
