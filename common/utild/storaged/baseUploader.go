package storaged

import (
	"errors"
	"go-zero-dandan/common/utild"
	"mime/multipart"
	"net/http"
)

// baseUploader 基础上传类
type baseUploader struct {
	MaxSize        int64
	Request        *http.Request
	File           multipart.File
	FileHeader     *multipart.FileHeader
	Type           FileType
	Hash           string
	FileMimeAccept []string
}

// defaultConfigStruct 默认的上传配置用
type defaultConfigStruct struct {
	MaxSize     int64
	AcceptMimes []string
}

var defaultConfig = map[FileType]defaultConfigStruct{
	FileTypeImage: {MaxSize: 5 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
	FileTypeFile: {MaxSize: 100 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
}

// processFileGet 处理文件获取
func (t *baseUploader) processFileGet(r *http.Request) (err error) {
	switch t.Type {
	case FileTypeImage:
		_ = t.Request.ParseMultipartForm(20 << 20) //20M,控制表单数据在内存中的存储大小，超过该值，则会自动将表单数据写入磁盘临时文件
		t.File, t.FileHeader, err = r.FormFile("img")
	case FileTypeFile:
		_ = t.Request.ParseMultipartForm(200 << 20) //200M
		t.File, t.FileHeader, err = r.FormFile("file")
	default:
		err = errors.New("unsupported file in file get")
	}

	return err
}

// processFileSize 处理文件大小和校验
func (t *baseUploader) processFileSize() (err error) {
	if t.FileHeader.Size > t.MaxSize {
		return errors.New("file size limited")
	}
	return nil
}

// processFileHash 处理文件哈希(sha1)
func (t *baseUploader) processFileHash() (err error) {
	//获取文件hash，用的sha1
	t.Hash, err = utild.GetFileHashHex(t.File)
	if err != nil {
		return err
	}
	//重新指向文件头，让后正常读取
	_, err = t.File.Seek(0, 0)
	return err
}

// processFileType 处理文件类型和校验
func (t *baseUploader) processFileType() (err error) {
	var validImageTypes map[string]bool
	for _, v := range t.FileMimeAccept {
		validImageTypes[v] = true
	}

	// 读取文件前 512 字节
	buffer := make([]byte, 512)
	if _, err = t.File.Read(buffer); err != nil {
		return errors.New("unsupport image type")
	}
	// 判断文件 MIME 类型是否为图片类型
	mime := http.DetectContentType(buffer)
	if !validImageTypes[mime] {
		return errors.New("invalid img type")
	}
	//重新指向文件头，避免后续操作问题
	if _, err = t.File.Seek(0, 0); err != nil {
		return err
	}
	return nil
}

/*
// 保存上传进度信息
func saveUploadProgress(fileName string, uploadID string, parts []Part) error {
	progressFileName := fileName + ".progress"
	progressFile, err := os.Create(progressFileName)
	if err != nil {
		return err
	}
	defer progressFile.Close()

	fmt.Fprintln(progressFile, uploadID)
	for _, part := range parts {
		fmt.Fprintf(progressFile, "%d %d %d %s\n", part.Number, part.Offset, part.Size, part.Etag)
	}

	return nil
}

// 删除上传进度信息
func deleteUploadProgress(fileName string) error {
	progressFileName := fileName + ".progress"
	return os.Remove(progressFileName)
}

// 上传分片
func uploadParts(filePath string, uploadID string, parts []Part) ([]Part, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	partCount := fileInfo.Size()/ChunkSize + 1
	if fileInfo.Size()%ChunkSize == 0 {
		partCount--
	}

	for i := range parts {
		partNumber := parts[i].Number
		partOffset := parts[i].Offset
		partLength := parts[i].Size

		if parts[i].Etag != "" {
			fmt.Printf("Part %d already uploaded, skipped\n", partNumber)
			continue
		}

		partData := make([]byte, partLength)
		_, err := file.ReadAt(partData, partOffset)
		if err != nil && err != io.EOF {
			return nil, err
		}

		partEtag, err := uploadPart(uploadID, partNumber, partData)
		if err != nil {
			return nil, err
		}

		parts[i].Etag = partEtag

		fmt.Printf("Part %d uploaded, ETag: %s\n", partNumber, partEtag)
	}

	return parts, nil
}

// 上传单个分片
func uploadPart(uploadID string, partNumber int, partData []byte) (string, error) {
	req, err := http.NewRequest(http.MethodPost, UploadURL, bytes.NewReader(partData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Upload-ID", uploadID)
	req.Header.Set("Upload-Part-Number", strconv.Itoa(partNumber))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to upload part %d, status code: %d", partNumber, resp.StatusCode)
	}

	partEtag := resp.Header.Get("ETag")
	if partEtag == "" {
		return "", fmt.Errorf("failed to get ETag for part %d", partNumber)
	}

	return partEtag, nil
}

// 完成分片上传
func completeMultipartUpload(fileName string, uploadID string, parts []Part) error {
	// 按照分片编号排序
	sort.Slice(parts, func(i, j int) bool {
		return parts[i].Number < parts[j].Number
	})

	req, err := http.NewRequest(http.MethodPost, UploadURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Upload-ID", uploadID)

	var etags []string
	for _, part := range parts {
		etags = append(etags, part.Etag)
	}

	// 构造完成分片上传请求的 JSON 数据
	bodyData := struct {
		Parts []struct {
			PartNumber int    `json:"part_number"`
			ETag       string `json:"etag"`
		} `json:"parts"`
	}{}
	for i, etag := range etags {
		bodyData.Parts = append(bodyData.Parts, struct {
			PartNumber int    `json:"part_number"`
			ETag       string `json:"etag"`
		}{
			PartNumber: i + 1,
			ETag:       etag,
		})
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to complete multipart upload, status code: %d", resp.StatusCode)
	}

	return nil
}
*/
// Part 分片上传中一个分片的信息
type Part struct {
	Number int    // 分片编号
	Offset int64  // 分片在文件中的偏移量
	Size   int64  // 分片的大小
	Etag   string // 分片上传成功后返回的 ETag
}
