package main

import (
	"strings"
	"os"
	"io"
	"fmt"
	"compress/gzip"
	"crypto/sha1"
)

func main() {
	SectionReader()
}

func Copy() {
	data := strings.NewReader("Write me down")
	file, _ := os.Create("./iocopy.data")
	io.Copy(file, data)
}

func PipeReaderAndPipeWriter() {
	file, _ := os.Create("./iopipe.data")
	pr, pw :=io.Pipe()

	go func() {
		fmt.Fprint(pw,"pipe streaming")
		pw.Close()
	}()

	wait := make(chan struct{})

	go func() {
		io.Copy(file,pr)
		pr.Close()
		close(wait)
	}()

	<-wait
}

func TeeReader()  {
	fin, _ := os.Open("./reader.go")
	defer fin.Close()

	fout, _:= os.Create("./teereader.gz")
	defer fout.Close()

	zip := gzip.NewWriter(fout)
	zip.Close()
	sha := sha1.New()
	data := io.TeeReader(fin, sha)

	io.Copy(zip,data)
	fmt.Printf("SHA1 hash %x\n",sha.Sum(nil))

}

func WriteString()  {
	fout, err := os.Create("./iowritestr.data")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()
	io.WriteString(fout,"Hello there")
}

func LimitReader()  {
	str := strings.NewReader("The quick brown fox jumps over tje lazy")
	limited := &io.LimitedReader{R:str,N:19}
	io.Copy(os.Stdout,limited)
}

func SectionReader()  {
	str := strings.NewReader("The quick brown fox jumps over tje lazy")
	section := io.NewSectionReader(str,19, 23)
	io.Copy(os.Stdout,section)
}