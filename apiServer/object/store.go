package object

import (
	"io"
	"net/http"
)

func storeObject(r io.Reader, object string) (int, error) {
	//putStream相当于创建了一个读写管道，读端交给http client,写端返回给客户，等待客户写入数据，没有数据时http client会阻塞住
	stream, err := putStream(object)
	if err != nil {
		return http.StatusServiceUnavailable, err
	}
	io.Copy(stream, r)
	//只有调用clsoe方法后才会向管道写入io.EOF，HTTP client才会返回，不然会一直阻塞在读数据部分
	err = stream.Close()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}
