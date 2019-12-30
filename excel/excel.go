package excel

import (
    `fmt`
    `io/ioutil`
    `os`
    `strings`
    `sync`
    `time`

    `github.com/tealeg/xlsx`
)

var mFile = "afterMerged.xlsx"
var lock sync.Mutex
var group sync.WaitGroup
// 合并逻辑
func mergeExcels()  {
    // 删除 afterMerged.xlsx 文件
    err := os.Remove(mFile)
    if err != nil {
        // TODO 没想好怎么处理
    }
    // 获取当前目录下所有 xlsx 文件
    dir, err := os.Getwd()
    fileInfos, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Println(err)
        return
    }
    files := make([]string, 0)
    // 不处理非 xlsx 文件
    for _, fileInfo := range fileInfos {
        file := fileInfo.Name()
        if !strings.HasSuffix(file, "xlsx"){
            continue
        }
        files = append(files, file)
    }
    // 如果没有xlsx文件，提示并退出
    if len(files) == 0 {
        fmt.Println("当前目录下没有xlsx文件！")
        return
    }
    // 挨个读取内容并合并至新文件
    newFile := xlsx.NewFile()
    sheet1, err := newFile.AddSheet("Sheet1")
    if err != nil {
        fmt.Println(err)
        return
    }
    sheet2, err := newFile.AddSheet("Sheet2")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 读取文件内容
    for _, file := range files {
        group.Add(1)
        ff := file
        go func ()  {
            fmt.Println("开始读取", ff, "文件...")
            f, err := xlsx.OpenFile(ff)
            if err != nil {
                fmt.Println(err)
                return
            }
            sheet := f.Sheets[0]
            ffRows := sheet.Rows
            lock.Lock()
            defer lock.Unlock()
            defer group.Done()
            // 追加内容
            for _, row := range ffRows {
                if len(sheet1.Rows) < (1<<20) {
                    newRow := sheet1.AddRow()
                    for _, cell := range row.Cells {
                        newCell := newRow.AddCell()
                        newCell.SetValue(cell)
                    }
                } else {
                    // 如果超出 2^20 次方行，则往 Sheet2 追加
                    newRow := sheet2.AddRow()
                    for _, cell := range row.Cells {
                        newCell := newRow.AddCell()
                        newCell.SetValue(cell)
                    }
                }
            }
        }()
    }
    group.Wait()
    fmt.Println("开始合并...")
    err = newFile.Save("afterMerged.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
}
var quit = make(chan bool)

// 合并当前目录下的 excel 文件
// 生成新文件，名为 afterMerged.xlsx
func Merge() {
    begin := time.Now()
    go func() {
    	signs := []string{"/", "|", "\\", "-"}
    	i := 0
        for {
            select {
            case <-quit:
                return
            default:
                fmt.Printf("%s\r", signs[i])
                i = (i+1) % 4
                time.Sleep(time.Millisecond * 50)
            }
        }
    }()
    mergeExcels()
    duration := time.Since(begin)
    quit <- true
    fmt.Printf("******处理完成******\n" +
        "合并后文件为 [%s] 一共耗时%4.2f秒\n", mFile, duration.Seconds())
    fmt.Println("按回车键退出程序...")
    var end string
    fmt.Scanf("%v", end)
}