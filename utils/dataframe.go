package utils

import (
	"os"
	"github.com/kniren/gota/series"
	"github.com/kniren/gota/dataframe"
	"github.com/klauspost/compress/gzip"

	"github.com/datochan/gcom/logger"
)

/**
 * 为df重新设置index,如果没有则添加index列,如果有则重新设置index的值
 */
const (
	IndexColName = "__dato_idx__"
)

// 为df增加idx列
func ReIndex(df *dataframe.DataFrame) dataframe.DataFrame {
	idxs := GenerateIndex(0, 1, df.Nrow())

	result := df.Mutate(series.New(idxs, series.Int, IndexColName))

	return result
}

/**
 * 从指定记录集中获取某个元素
 */
func Element(data dataframe.DataFrame, row int, colname string) series.Element {
	idx := FindInStringSlice(colname, data.Names())

	if idx < 0 {
		return nil
	}

	return data.Elem(row, idx)
}

/**
 * 从CSV中加载df数据
 */
func ReadCSV(filename string, options ...dataframe.LoadOption) dataframe.DataFrame {
	inFile, err := os.OpenFile(filename, os.O_RDONLY | os.O_RDWR, 0666)
	if err != nil { return dataframe.DataFrame{Err: err} }

	defer inFile.Close()

	gzipReader, err := gzip.NewReader(inFile)
	defer gzipReader.Close()
	if nil != err { return dataframe.DataFrame{Err: err} }

	df := dataframe.ReadCSV(gzipReader, options...)

	if nil != df.Err { logger.Error(df.Err.Error()) }

	return df
}

/**
 * 保存股票数据
 * param int mode: os.O_CREATE | os.O_APPEND | os.O_RDONLY | os.O_WRONLY | os.O_RDWR
 */
func WriteCSV(filename string, mode int, df *dataframe.DataFrame, option ...dataframe.WriteOption) error {
	outFile, err := os.OpenFile(filename, mode, 0666)
	if err != nil { return err }
	defer outFile.Close()

	gWriter := gzip.NewWriter(outFile)
	defer gWriter.Close()

	return df.WriteCSV(gWriter, option...)
}

