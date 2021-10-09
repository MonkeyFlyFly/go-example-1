package export

import "github.com/EDDYCJY/go-example-1/pkg/setting"

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
