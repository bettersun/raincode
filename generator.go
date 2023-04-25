package main

import (
	"github.com/bettersun/rain"
	yml "github.com/bettersun/rain/yaml"
	"log"
	"strings"
	"time"
)

func RunGenerate() {
	configFile := "config/config.yml"
	cfg, err := ReadConfig(configFile)
	if err != nil {
		log.Println(err)
	}

	GenerateCode(cfg)
}

func ReadConfig(configFile string) (Config, error) {
	var cfg Config

	err := yml.YamlFileToStruct(configFile, &cfg)
	if err != nil {
		log.Println("加载配置文件发生错误")
	}

	return cfg, err
}

func GenerateCode(cfg Config) {

	// 读取代码模板
	result := rain.SearchByType([]string{cfg.TmplPath}, []string{".dart", ".json"}, nil)

	log.Println(result)
	for _, f := range cfg.FileName {

		for _, v := range result {

			// 替换文件名路径中的模块物理名部分
			// 模板目录 -> 输出目录
			s := strings.Replace(v, cfg.TmplPath, cfg.OutputPath, -1)
			// 模块内目录 / 文件名
			s = strings.ReplaceAll(s, f.PlaceHolder, f.Fill)

			// 复制模板代码文件内容
			err := rain.CopyFile(v, s)
			if err != nil {
				log.Println(err)
			}

			// 读取目标代码文件内容
			code, err := rain.ReadFile(s)
			if err != nil {
				log.Println(err)
			}

			// 替换代码文件内容
			code = ReplaceCode(code, cfg)

			// 输入到文件
			rain.WriteFile(s, []string{code})
		}
	}
}

func ReplaceCode(codeTmpl string, cfg Config) string {

	code := codeTmpl

	for _, v := range cfg.Code {
		var dt string
		if strings.TrimSpace(cfg.DateFormat) == "" {
			dt = time.Now().Format(cfg.DateFormat)
		} else {
			dt = time.Now().Format("2016/01/02")
		}
		// 日期
		code = strings.Replace(code, cfg.Date, dt, -1)

		//
		code = strings.Replace(code, v.PlaceHolder, v.Fill, -1)
	}

	return code
}
