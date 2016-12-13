package local

import (
	"fmt"
	"github.com/elsejj/dzhyunsdk/dzhyun"
)

func Table2Map(table *dzhyun.Table) []map[string]interface{} {

	headers := table.GetInfo()
	datas := table.GetData()

	rows := 0
	switch headers[0].Type {
	case int32(dzhyun.InfoType_Type_Int):
		rows = len(datas[0].GetIValues())
	case int32(dzhyun.InfoType_Type_SInt):
		rows = len(datas[0].GetXValues())
	case int32(dzhyun.InfoType_Type_String):
		rows = len(datas[0].GetSValues())
	case int32(dzhyun.InfoType_Type_Table):
		rows = len(datas[0].GetTValues())

	}
	mm := make([]map[string]interface{}, rows)
	sValues := make([]float64, len(headers))
	rValues := make([]int, len(headers))
	if datas != nil {

		for row := 0; row < rows; row++ {
			rec := make(map[string]interface{})
			for col, header := range headers {
				data := datas[col]
				switch header.Type {
				case int32(dzhyun.InfoType_Type_Int):
					rec[header.GetName()] = data.GetIValues()[row]
				case int32(dzhyun.InfoType_Type_SInt):
					if row == 0 {
						var err error
						sValues[col], rValues[col], err = dzhyun.UnmakeValue(data.GetXValues()[row])
						if err != nil {
							fmt.Println("unpack", data.GetXValues()[row], "failed", err)
						}
					} else {
						sValues[col] = sValues[col] + float64(data.GetXValues()[row]*int64(header.Ratio))
					}
					rec[header.GetName()] = dzhyun.YFloat(dzhyun.MakeValue(sValues[col], rValues[col]))
				case int32(dzhyun.InfoType_Type_String):
					rec[header.GetName()] = data.GetSValues()[row]
				case int32(dzhyun.InfoType_Type_Table):
					rec[header.GetName()] = Table2Map(data.GetTValues()[row])
				}
			}
			mm[row] = rec
		}
	}
	return mm
}
