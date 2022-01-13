package dictionary

import "errors"

//使用自定义dictionary类型

type Dictionary map[string]string

//使用自定义错误类型

var ErrorNotFound = errors.New("could not found a value")

/* func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
} */

// 实现自定义类的方法

func (d Dictionary) Search(key string) (string, error) {
	//return d[key], nil
	// 重构抛出错误
	definition, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
