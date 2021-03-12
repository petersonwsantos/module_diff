package main
import (
  "io/ioutil"
  "bytes"
  "fmt"
  "os"
)

func diff(f1 string, f2 string) bool {
  one, err := ioutil.ReadFile(f1)
  if err != nil{
    panic(err)
  }
  two, err2 := ioutil.ReadFile(f2)
  if err2 != nil{
    panic(err2)
  }
  same := bytes.Equal(one, two)
  //fmt.Println(same)
  //fmt.Println("rolou")
  return same
}

func main(){
  file1 := os.Args[1]
  file2 := os.Args[2]
  resultDiff := diff(file1, file2)
  fmt.Println(resultDiff)
}