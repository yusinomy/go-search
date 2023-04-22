package cmd

import (
	"fmt"
	"log"
	"os"
)

func Exec() {
	fofa := Fofa()
	zone := Zone()
	hunter := Hunter()
	fofa = append(fofa, zone...)
	fofa = append(fofa, hunter...)
	array := deduplicated_array(fofa)
	for i := range array {
		fmt.Println(array[i])
		file, err := os.OpenFile("all.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			log.Println(err)
		}
		defer func() { file.Close() }()
		file.WriteString(array[i] + "\n")
	}
}

func deduplicated_array(arr []string) []string {
	// 创建一个整型key和布尔类型value的哈希表
	hash := make(map[string]bool)
	// 创建一个空的整型数组
	result := []string{}
	// 遍历原始数组
	for _, value := range arr {
		// 如果哈希表(hash)中不存在该值，则加入结果数组和哈希表(hash)
		if _, ok := hash[value]; !ok {
			result = append(result, value)
			hash[value] = true
		}
	}
	// 返回去重后的数组
	return result
}
