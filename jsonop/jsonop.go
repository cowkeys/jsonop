package jsonop

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Version() string {
	return "0.1.0"
}
//New new a Config
func GetJfromFile(file string) Json {
	return Json{file: file}
}

func NewJson(jsonstr string) (Json,error) {
	return Json{jsonstr: jsonstr}
}

type Json struct {
	file    string
	jsonstr string
	maps    map[string]interface{}
}

func (j *Json) readfromstring() err {
	var mapres map[string]interface{}
    err := json.Unmarshal([]byte(j.jsonstr), mapres)
    if err != nil {
        return err
    }else{
		j.maps = mapres
	}
    return nil
}

//Get name pattern key or key.key.key
//
func (j *Json) Get(name string) interface{} {

	if j.maps == nil {
		if file =="" {
			j.readfromfile()
		}else if jsonstr == ""{
			j.readfromstring()
		}else{
			return nil
		}
		
	}

	if j.maps == nil {
		return nil
	}

	// app.view.path
	keys := strings.Split(name, ".")
	l := len(keys)
	if l == 1 {
		return c.maps[name]
	}

	var ret interface{}
	for i := 0; i < l; i++ {
		if i == 0 {
			ret = c.maps[keys[i]]
			if ret == nil {
				return nil
			}
		} else {
			if m, ok := ret.(map[string]interface{}); ok {
				ret = m[keys[i]]
			} else {
				if l == i-1 {
					return ret
				}
				return nil
			}
		}
	}
	return ret
}

/*
func (j *Json) read() {
	if !filepath.IsAbs(c.file) {
		file, err := filepath.Abs(c.file)
		if err != nil {
			panic(err)
		}
		c.file = file
	}

	bts, err := ioutil.ReadFile(c.file)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bts, &c.maps)

	if err != nil {
		panic(err)
	}
}
