package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.


func (reader MyReader) Read(data []byte) (int, error){
	for i:=0; i < len(data) ; i ++{
		data[i] = 'A'
	}
	return len(data), nil
	
}

func main() {
	reader.Validate(MyReader{})
}
