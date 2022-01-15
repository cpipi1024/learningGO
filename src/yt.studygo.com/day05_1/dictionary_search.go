package dictionary

//使用自定义dictionary类型

type Dictionary map[string]string

// 创建错误类型
//var ErrorNotFound = errors.New("could not found a value")

/*var (
	ErrorNotFound = errors.New("could not found a value")
	ErrorExists   = errors.New("key already exists")
)*/
//
/* func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
} */

// 创建自定义错误类型
type DictionaryErr string

var (
	ErrorNotFound     = DictionaryErr("could not found a value")
	ErrorExists       = DictionaryErr("key already exists")
	ErrorKeyNotExists = DictionaryErr("could not update key not exists")
)

// 实现自定义类型返回接口
func (de DictionaryErr) Error() string {
	return string(de)
}

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

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorNotFound:
		return ErrorKeyNotExists
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}
