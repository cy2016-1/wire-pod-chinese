package wirepod_whisper

//luowei add the file

import (
	"bytes"
	// "encoding/binary"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	//"strings"

	// "github.com/go-audio/audio"
	// "github.com/go-audio/wav"
	"github.com/kercre123/chipper/pkg/logger"
	sr "github.com/kercre123/chipper/pkg/wirepod/speechrequest"
	"github.com/orcaman/writerseeker"
	
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
)

var Name string = "baidu"

type openAiResp struct {
	Text string `json:"text"`
}

func Init() error {
	if os.Getenv("BAIDU_APIKEY") == "" || os.Getenv("BADIU_APISECRET") == "" {
		logger.Println("You must export the BAIDU_APIKEY and BADIU_APISECRET env var!")
	}
	return nil
}


func baidu_asr(f io.Reader) string {
	client := voice.NewVoiceClient(os.Getenv("BAIDU_APIKEY"), os.Getenv("BADIU_APISECRET"))
	if err := client.Auth(); err != nil {
		logger.Println(err)
	}

	// f, err := os.OpenFile("16k.pcm", os.O_RDONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	//logger.Println("百度开始识别...")

	rs, err := client.SpeechToText(
		f,
		voice.Format("pcm"),
		voice.Channel(1),
		voice.Rate(16000),
	)
	if err != nil {
		logger.Println(err)
	}

	//logger.Println("百度识别完成",rs)
	return rs[0]
	
}



func makeOpenAIReq(in []byte) string {
	url := "https://api.openai.com/v1/audio/transcriptions"

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	w.WriteField("model", "whisper-1")
	sendFile, _ := w.CreateFormFile("file", "audio.mp3")
	sendFile.Write(in)//写入mp3文件
	w.Close()

	httpReq, _ := http.NewRequest("POST", url, buf)
	httpReq.Header.Set("Content-Type", w.FormDataContentType())
	httpReq.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_KEY"))

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		logger.Println(err)
		return "There was an error."
	}

	defer resp.Body.Close()

	response, _ := io.ReadAll(resp.Body)

	var aiResponse openAiResp
	json.Unmarshal(response, &aiResponse)

	return aiResponse.Text
}

func STT(req sr.SpeechRequest) (string, error) {
	logger.Println("(Bot " + strconv.Itoa(req.BotNum) + ", Baidu) Processing...")
	speechIsDone := false
	var err error
	for {
		req, _, err = sr.GetNextStreamChunk(req)
		if err != nil {
			return "", err
		}
		if err != nil {
			return "", err
		}
		// has to be split into 320 []byte chunks for VAD
		req, speechIsDone = sr.DetectEndOfSpeech(req)
		if speechIsDone {
			break
		}
	}

	pcmBufTo := &writerseeker.WriterSeeker{}
	pcmBufTo.Write(req.DecodedMicData)//麦克风数据转pcm数据

	transcribedText := baidu_asr(pcmBufTo.BytesReader())

	println("Bot " + strconv.Itoa(req.BotNum) + " 语音识别结果: " + transcribedText)
	return transcribedText, nil
}