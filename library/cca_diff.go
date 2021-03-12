// This file is part of Ansible
//
// Ansible is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Ansible is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Ansible.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"bytes"
	//"strconv"
)

type ModuleArgs struct {
	Name string
    File1 string
    File2 string
}

type Response struct {
	Msg     string `json:"msg"`
	Changed bool   `json:"changed"`
	Failed  bool   `json:"failed"`
	Isdiff  bool   `json:"isdiff"`
}

func ExitJson(responseBody Response) {
	returnResponse(responseBody)
}

func FailJson(responseBody Response) {
	responseBody.Failed = true
	returnResponse(responseBody)
}

func returnResponse(responseBody Response) {
	var response []byte
	var err error
	response, err = json.Marshal(responseBody)
	if err != nil {
		response, _ = json.Marshal(Response{Msg: "Invalid response object"})
	}
	fmt.Println(string(response))
	if responseBody.Failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func diff(f1 string, f2 string) bool {

  one, err := ioutil.ReadFile(f1)
  //one, err := ioutil.ReadFile("/workspace/ansible/test/integration/targets/binary_modules/1.txt")
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


func main() {
	var response Response

	if len(os.Args) != 2 {
		response.Msg = "No argument file provided"
		FailJson(response)
	}

	argsFile := os.Args[1]

	text, err := ioutil.ReadFile(argsFile)
	if err != nil {
		response.Msg = "Could not read configuration file: " + argsFile
		FailJson(response)
	}

	var moduleArgs ModuleArgs
	err = json.Unmarshal(text, &moduleArgs)
	if err != nil {
		response.Msg = "Configuration file not valid JSON: " + argsFile
		FailJson(response)
	}

	var name string = "Go CCA"
	if moduleArgs.Name != "" {
		name = moduleArgs.Name
	}

	var file1 string = moduleArgs.File1
	var file2 string = moduleArgs.File2
    
    resultDiff := diff(file1, file2)
    //pepi := strconv.FormatBool(resultDiff)

	response.Isdiff = resultDiff
	response.Msg = name
	ExitJson(response)
}
