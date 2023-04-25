package main

type Config struct {
	TmplPath   string     `yaml:"tmplPath"`
	OutputPath string     `yaml:"outputPath"`
	Date       string     `yaml:"date"`
	DateFormat string     `yaml:"dateFormat"`
	FileName   []FillInfo `yaml:"fileName"`
	Code       []FillInfo `yaml:"code"`
}

type FillInfo struct {
	PlaceHolder string `yaml:"placeHolder"`
	Fill        string `yaml:"fill"`
}
