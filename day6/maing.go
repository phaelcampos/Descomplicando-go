package main

type DataBucket struct {
	Data []byte
}

type DataLoader interface {
	LoadData(filename string, bucket DataBucket) error
}

type XLSDataLoader struct{}

func (xlsdataloader *XLSDataLoader) LoadData(filename string, bucket DataBucket) error {
	//do something
	//Carrega dados de um arquivo de nome filename dentro desse databucket, fazendo as tranfosrma~çoes de dados necessários para um xls
	return nil
}

type CSVDataLoader struct{}

func (csvdataloader *CSVDataLoader) LoadData(filename string, bucket DataBucket) error {
	//do something
	//Carrega dados de um arquivo de nome filename dentro desse databucket, fazendo as tranfosrma~çoes de dados necessários para um xls
	return nil
}

func LoadData(filename string, dataloader DataLoader) DataBucket {
	bucket := DataBucket{}
	dataloader.LoadData(filename, bucket)
	return bucket
}

func main() {
	LoadData("arquivo.csv", &CSVDataLoader{})
	LoadData("arquivo.xls", &XLSDataLoader{})

}
