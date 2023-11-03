package tool

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// SentModelSSEResponse 把模型的SSE响应数据使用sse格式发送给前端
func sentModelSSEResponse(w http.ResponseWriter, sseResp *http.Response) (answer string, err error) {
	reader := bufio.NewReader(sseResp.Body)
	headerData := []byte("data: ")
	var buffer bytes.Buffer
	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("not http.Flusher")
		err = errors.New("not http.Flusher")
		return
	}
	for {
		var line []byte
		if line, err = reader.ReadBytes('\n'); err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
				break
			}
			err = errors.w
			//err = errors.Wrap(err, "streamChatHandler:读取响应数据出错")
			return "", err
		}
		line = bytes.TrimSpace(line)
		line = bytes.TrimPrefix(line, headerData)
		//发送给前端
		strData := string(line)
		if strData == "" || strData == "\n" || strData == "\r" || strData == "data:" {
			continue
		}
		_, err = fmt.Fprintf(w, "data: %v\n\n", strData)
		if err != nil {
			err = errors.Wrap(err, "streamChatHandler:发送响应数据出错")
			return
		}

		flusher.Flush()
		//把回答缓存到answer
		//answer = strings.Join([]string{answer, strData}, "")
		_, err = buffer.WriteString(strData)
		if err != nil {
			return
		}
	}
	answer = buffer.String()
	return
}
