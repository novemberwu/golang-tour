package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(data []byte)(int, error){
	
	for {
		
		n, err := rot.r.Read(data)
		//fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		//fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			return n, err
		}else{
			for i,_ := range data {
				if data[i] >='A' && data[i] <= 'Z'{
					data[i] = 'A' + (data[i] - 'A'+13)%26
				}
				if data[i] >= 'a' && data[i] <= 'z'{
					data[i] = 'a' + (data[i] - 'a' +13)%26
				}
				
			}
			return 8, nil
						
		}
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
