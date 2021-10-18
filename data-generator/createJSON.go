package data_generator

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	jsonDir   string
	structure map[string]interface{}
)

func newStruct(mod string) interface{} {
	if structure[mod] != nil {
		return structure[mod]
	} else {
		return nil
	}
}

func ToArray(fileName string, arr [][]string) map[int32]interface{} {
	var Employees = make(map[int32]interface{})

	for _, row := range arr[4:] {
		if len(row) == 0 {
			break
		}
		if row[0] == "" {
			break
		}

		employee := newStruct(fileName)
		//klog.Info("employee: ", employee)
		//klog.Info("type employee: ", reflect.TypeOf(employee))
		getType := reflect.TypeOf(employee)
		getValue := reflect.ValueOf(&employee).Elem()
		//klog.Info("type getValue: ", reflect.TypeOf(getValue))
		newValue := reflect.New(getValue.Elem().Type()).Elem()
		newValue.Set(getValue.Elem())
		//klog.Info("type newValue: ", reflect.TypeOf(newValue))

		var id int32

		j := 0
		for i := 0; i < len(row); i++ {

			if j == getType.NumField() {
				break
			}

			if ok := IsIgnored(arr[0][i]); ok {
				continue
			}

			field := getType.Field(j)
			var value interface{}
			var err error = nil
			rowValue := row[i]
			switch field.Type.String() {
			case "int32":
				value, err = ToInt(rowValue)
				if i == 0 {
					id = value.(int32)
				}
			case "string":
				value = row[i]
			case "float32":
				value, err = ToFloat(rowValue)
			case "[]int32":
				value, err = ToIntSlice(ToStringSlice(rowValue))
			case "[]float32":
				value, err = ToFloatSlice(ToStringSlice(rowValue))
			case "[]string":
				value = ToStringSlice(rowValue)
			case "map[int32][]int32":
				value, err = ToListIntMap(rowValue)
			case "map[int32][]string":
				value, err = ToListStringMap(rowValue)
			case "map[int32]int32":
				value, err = ToIntMap(rowValue)
			case "map[int32]string":
				value, err = ToStringMap(rowValue)
			case "map[int32][]float32":
				value, err = ToListFloatMap(rowValue)
			case "map[int32]float32":
				value, err = ToFloatMap(rowValue)
			}

			if err != nil {
				fmt.Println(err)
				fmt.Println("当前行: 表: %s\tid: %d\t列: %s\t值: %s", fileName, id, field.Name, rowValue)
				fmt.Println("↑↑↑↑数值生成发生错误")
				os.Exit(1)
			}
			newValue.FieldByName(field.Name).Set(reflect.ValueOf(value))
			j++
		}
		getValue.Set(newValue)

		Employees[id] = employee

	}

	return Employees
}

func CreateJSON(jsDir, souDir string, s map[string]interface{}) {
	structure = s
	jsonDir = jsDir

	fmt.Println("Create JSON START")
	fileMap := GetFileMap(souDir)
	for sheetName, f := range fileMap {
		fmt.Println(sheetName, "    start")
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			fmt.Println(err)
			fmt.Println("表%s发生错误, 请确认存在Sheet1标签", sheetName)
			fmt.Println()
		}
		//structure.InitStructMap()
		// TODO 生成数据[]
		temp := ToArray(sheetName, rows)

		// TODO 转换成JSON文件并存储
		jsons, err := json.MarshalIndent(temp, "", "  ")

		ff := CreateFile(sheetName, "json", jsonDir)
		_, err = io.WriteString(ff, string(jsons))
		if err != nil {
			fmt.Println("file: ", sheetName, "err: ", err.Error())
		}
		ff.Close()
		fmt.Println(sheetName, "    end")
	}
	fmt.Println("Create JSON END")
}

func ToInt(str string) (int32, error) {
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func ToFloat(str string) (float32, error) {
	if str == "" {
		return 0, nil
	}
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), err
}

// 格式化成[]string
func ToStringSlice(str string) []string {
	if str == "" {
		return make([]string, 0)
	}
	reg := regexp.MustCompile(`[{}]`)
	tmp := reg.ReplaceAllString(str, ``)

	return strings.Split(tmp, ",")
}

// 将[]string 转化成 []int32
func ToIntSlice(str []string) ([]int32, error) {
	var ret []int32
	if len(str) == 0 {
		return make([]int32, 0), nil
	}

	for _, i := range str {
		j, err := strconv.ParseInt(i, 10, 32)
		if err != nil {
			return nil, err
		}
		ret = append(ret, int32(j))
	}
	return ret, nil
}

// 将[]string 转化成 []float32
func ToFloatSlice(str []string) ([]float32, error) {
	var ret []float32
	if len(str) == 0 {
		return make([]float32, 0), nil
	}

	for _, i := range str {
		j, err := strconv.ParseFloat(i, 32)
		if err != nil {
			return nil, err
		}
		ret = append(ret, float32(j))
	}
	return ret, nil
}

// 转化成map[int32]int32
func ToStringMap(str string) (map[int32]string, error) {
	var ret = make(map[int32]string)
	if str == "" {
		return ret, nil
	}

	tmp := strings.Split(str, ";")

	for _, val := range tmp {
		tmp1 := strings.Split(val, ":")

		index, err := strconv.ParseInt(tmp1[0], 10, 32)
		if err != nil {
			return nil, err
		}

		ret[int32(index)] = tmp1[1]

	}

	return ret, nil
}

// 转化成map[int32]int32
func ToIntMap(str string) (map[int32]int32, error) {
	var ret = make(map[int32]int32)
	if str == "" {
		return ret, nil
	}

	tmp := strings.Split(str, ";")

	for _, val := range tmp {
		if val == "" {
			continue
		}
		tmp1 := strings.Split(val, ":")

		j, err := strconv.ParseInt(tmp1[1], 10, 32)
		if err != nil {
			return nil, err
		}

		index, err := strconv.ParseInt(tmp1[0], 10, 32)
		if err != nil {
			return nil, err
		}

		ret[int32(index)] = int32(j)

	}

	return ret, nil
}

// 转化成map[int32][]string
func ToListStringMap(str string) (map[int32][]string, error) {
	var ret = make(map[int32][]string)
	if str == "" {
		return ret, nil
	}

	tmp := strings.Split(str, ";")

	for _, val := range tmp {
		tmp1 := strings.Split(val, ":")

		value := strings.Split(tmp1[1], ",")

		index, err := strconv.ParseInt(tmp1[0], 10, 32)
		if err != nil {
			return nil, err
		}

		ret[int32(index)] = value

	}

	return ret, nil
}

// 转化成map[int32]float32
func ToFloatMap(str string) (map[int32]float32, error) {
	var ret = make(map[int32]float32)
	if str == "" {
		return ret, nil
	}

	tmp := strings.Split(str, ";")

	for _, val := range tmp {
		if val == "" {
			continue
		}
		tmp1 := strings.Split(val, ":")

		j, err := strconv.ParseFloat(tmp1[1], 10)
		if err != nil {
			return nil, err
		}

		index, err := strconv.ParseInt(tmp1[0], 10, 32)
		if err != nil {
			return nil, err
		}

		ret[int32(index)] = float32(j)

	}

	return ret, nil
}

// 转化成map[int32][]int32
func ToListIntMap(str string) (map[int32][]int32, error) {
	var ret = make(map[int32][]int32)
	if str == "" {
		return ret, nil
	}

	tmp := strings.Split(str, ";")

	for _, val := range tmp {
		if val == "" {
			continue
		}
		tmp1 := strings.Split(val, ":")

		value := strings.Split(tmp1[1], ",")
		var arr []int32
		for _, val := range value {
			j, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return nil, err
			}
			arr = append(arr, int32(j))
		}

		index, err := strconv.ParseInt(tmp1[0], 10, 32)
		if err != nil {
			return nil, err
		}

		ret[int32(index)] = arr

	}

	return ret, nil
}

// 转化成map[int32][]float32
func ToListFloatMap(str string) (map[int32][]float32, error) {
	var ret = make(map[int32][]float32)
	if str == "" {
		return ret, nil
	}

	tmp := strings.Split(str, ";")

	for _, val := range tmp {
		tmp1 := strings.Split(val, ":")

		value := strings.Split(tmp1[1], ",")
		var arr []float32
		for _, val := range value {
			j, err := strconv.ParseFloat(val, 10)
			if err != nil {
				return nil, err
			}
			arr = append(arr, float32(j))
		}

		index, err := strconv.ParseInt(tmp1[0], 10, 32)
		if err != nil {
			return nil, err
		}

		ret[int32(index)] = arr

	}

	return ret, nil
}
