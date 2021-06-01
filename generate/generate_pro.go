package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/viper"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	LineSplit   = ":"
	ColumnIndex = 0
	TypeIndex   = 1
	ValueIndex  = 2
	STAND       = "2006-01-02 15-04-05" // yyyy-mm-dd HH:mm:ss
	TimeFormat  = "2006-01-02 15:04:05" // yyyy-mm-dd HH:mm:ss
)

var logger = log.Default()
var container = make(map[string][]interface{}, 100)
var fields = make(map[string]int, 0)
var data = make(chan string, 1000)
var concurrentTime = make([]string, 0)
var faker = gofakeit.NewCrypto()
var wg sync.WaitGroup
var config *Config
var kafkaCount int32 = 0

type Config struct {
	FieldConfigName string
	dataFormat      string
	outType         string
	outputFileName  string
	history         int
	total           int
}

type SyncWriter struct {
	L      sync.Mutex
	Writer *bufio.Writer
}

func (w *SyncWriter) writer(value string) {
	w.L.Lock()
	defer w.L.Unlock()
	w.Writer.WriteString(value)
	w.Writer.WriteString("\n")
}

func init() {
	// 读取config配置文件
	viper.SetConfigName("config")                                              // name of config file (without extension)
	viper.SetConfigType("ini")                                                 // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("C:\\Users\\think\\go\\src\\simulationData\\generate") // optionally look for config in the working directory
	err := viper.ReadInConfig()                                                // Find and read the config file
	if err != nil {                                                            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config = &Config{
		FieldConfigName: viper.GetString("input.field_config_file"),
		dataFormat:      strings.ToLower(viper.GetString("format.format")),
		outType:         strings.ToLower(viper.GetString("output.outType")),
		outputFileName:  viper.GetString("output.output_file"),
		history:         viper.GetInt("output.history"),
		total:           viper.GetInt("input.total"),
	}
	// 校验配置 TODO

	fmt.Println(config)
	// 初始化数据池
	fieldInit(config.FieldConfigName)
}

func fieldInit(filename string) {
	logger.Println("init value begnning")
	// 读取配置文件
	inputFile, err := os.Open(filename)
	if err != nil {
		logger.Panicln("read config file error : %s", err.Error())
		return
	}
	defer inputFile.Close()
	bufferIo := bufio.NewReader(inputFile)
	// 解析配置文件初始化数据
	for {
		v, err := bufferIo.ReadString('\n')
		values := strings.Split(v, LineSplit)
		value := strings.TrimSpace(values[ValueIndex])
		column := values[ColumnIndex]
		cloumnType := strings.ToLower(strings.TrimSpace(values[TypeIndex]))

		if cloumnType == "struct" {
			slice := convert2Slice(str2Slice(values[2]))
			container[column] = slice
			fields[column] = len(slice)
		} else if cloumnType == "datetime" {
			if value == "now()" {
				concurrentTime = append(concurrentTime, column)
			} else {
				d := strings.Split(value, ",")
				startTime, err := time.Parse(STAND, strings.TrimSpace(d[0]))
				endTime, err1 := time.Parse(STAND, strings.TrimSpace(d[1]))
				if err != nil || err1 != nil {
					logger.Println("error time column ", column)
				} else {
					len, err := strconv.ParseInt(strings.TrimSpace(value), 10, 32)
					if err != nil {
						len = 100
						logger.Println(" WARNING can not parse int value")
					}
					valueChannel := make([]interface{}, len)
					for i := 0; i < int(len); i++ {
						valueChannel[i] = faker.DateRange(startTime, endTime).Format(TimeFormat)
					}
					container[column] = valueChannel
					fields[column] = int(len)
					//fmt.Println(fields, container)
				}

			}
		} else {
			len, err := strconv.ParseInt(strings.TrimSpace(value), 10, 32)
			if err != nil {
				len = 100
				logger.Println(" WARNING can not parse int value")
			}
			valueChannel := make([]interface{}, len)
			switch cloumnType {
			case "username":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Username()
				}
			case "email":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Email()
				}
			case "ip4":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.IPv4Address()
				}
			case "ip6":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.IPv6Address()
				}
			case "url":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.URL()
				}
			case "phone":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Phone()
				}
			case "uuid":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.UUID()
				}
			case "domain":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.DomainName()
				}
			case "macaddress":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.MacAddress()
				}
			case "address":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Address()
				}
			case "country":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Country()
				}
			case "city":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.City()
				}
			case "street":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Street()
				}
			case "game":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Gamertag()
				}
			case "fruit":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Fruit()
				}
			case "httpcode":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.HTTPStatusCode()
				}
			case "httpmethod":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.HTTPMethod()
				}
			case "useragent":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.UserAgent()
				}
			case "company":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.Company()
				}
			case "joblevel":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.JobLevel()
				}
			case "jobtitle":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.JobTitle()
				}
			case "jobdescriptor":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.JobDescriptor()
				}
			case "appname":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.AppName()
				}
			case "appversion":
				for i := 0; i < int(len); i++ {
					valueChannel[i] = faker.AppVersion()
				}
			default:
				logger.Println("error type ", cloumnType)
			}
			container[column] = valueChannel
			fields[column] = int(len)
		}

		if err == io.EOF && strings.ContainsAny(v, ":") {
			logger.Println("READ OVER")
			break
		}
	}
	logger.Println("init value over start faker data")
}

func main() {
	for i := 0; i <= 5; i++ {
		go produceData()
	}
	//count := 0
	//timeTricker := time.NewTicker(1 * time.Second)
	//for {
	//	select {
	//	case <- data:
	//		count ++
	//	case b:= <- timeTricker.C:
	//		count ++
	//		fmt.Println(count, b)
	//	}
	//}

	// 存储数据池下次使用 TODO
	go generateFieldConfig()

	if config.total > 0 && config.outType == "file" {
		logger.Println("start file out")
		writer, err := os.OpenFile(config.outputFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logger.Println("file write fail. . .")
		}
		defer writer.Close()

		syncWriter := &SyncWriter{sync.Mutex{}, bufio.NewWriter(writer)}
		wg.Add(1)
		go consumer2File(syncWriter, config.total)
		wg.Wait()
		logger.Println("file out over")
	} else if config.outType == "kafka" {
		wg.Add(1)
		logger.Println(" start kafka out type streaming ... ")
		config1 := sarama.NewConfig()
		config1.Producer.RequiredAcks = sarama.WaitForLocal
		config1.Producer.Partitioner = sarama.NewRandomPartitioner
		config1.Producer.Return.Successes = true
		//config1.Producer.Retry.Max = 3

		client, err := sarama.NewAsyncProducer([]string{"10.50.2.56:19092", "10.50.2.56:19091", "10.50.2.56:19093"}, config1)
		if err != nil {
			fmt.Println(" Producer err ", err.Error())
			return
		}
		logger.Println("start ...")
		for i := 0; i <= 10; i++ {
			go consumer2Kafka(client, config.total)
		}
		go consumer2Kafka(client, config.total)
		wg.Wait()
	} else {
		logger.Panicln("can not support out type")
	}
}

func produceData() {
	for {
		value := make(map[string]interface{})
		for k, v := range fields {
			value[k] = container[k][rand.Intn(v)]
		}
		for _, v := range concurrentTime {
			value[v] = time.Now().Format(TimeFormat)
		}
		jsonString, err := json.Marshal(value)
		//fmt.Println(string(jsonString))
		if err != nil {
			break
		}
		data <- string(jsonString)
	}
}

func consumer2File(syncWriter *SyncWriter, total int) {
	count := 0
	defer wg.Done()
	for {
		select {
		case d := <-data:
			syncWriter.writer(d)
			count++
			if count >= total {
				syncWriter.Writer.Flush()
				return
			}
		}
	}
}

func consumer2Kafka(client sarama.AsyncProducer, total int) {
	defer wg.Done()
	defer client.Close()
	trick := time.NewTicker(1 * time.Second)
	count := 0
	for {
		select {
		case d := <-data:
			msg := &sarama.ProducerMessage{
				Topic: "gotest",
				Value: sarama.StringEncoder(d),
			}
			client.Input() <- msg
			//client.SendMessage(&sarama.ProducerMessage{
			//	Topic: "gotest",
			//	Value: sarama.StringEncoder(d),
			//})
			atomic.AddInt32(&kafkaCount, 1)
		case <-trick.C:
			count++
			logger.Println(kafkaCount, kafkaCount/int32(count))
		}
	}
}

func generateFieldConfig() {
	logger.Println("start generate field")
	logger.Println("generate field over")
}

func str2Slice(a string) []interface{} {
	slice := strings.Split(strings.TrimSpace(a), ",")
	value := make([]interface{}, len(slice))
	for ix, v := range slice {
		value[ix] = v
	}
	return value
}

// 判断是否为切片
func isSlince(arg interface{}) (reflect.Value, bool) {
	val := reflect.ValueOf(arg)
	ok := false
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return val, ok
}

// 转化为[]interface{}（装载任意元素的切片与[]int这种不兼容） 切片
func convert2Slice(arg interface{}) []interface{} {
	val, ok := isSlince(arg)
	if !ok {
		return nil
	}
	len := val.Len()
	out := make([]interface{}, len)
	for i := 0; i < len; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out
}
