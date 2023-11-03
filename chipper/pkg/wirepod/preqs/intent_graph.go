package processreqs

import (
	"strconv"

	pb "github.com/digital-dream-labs/api/go/chipperpb"
	"github.com/kercre123/chipper/pkg/logger"
	"github.com/kercre123/chipper/pkg/vars"
	"github.com/kercre123/chipper/pkg/vtt"
	sr "github.com/kercre123/chipper/pkg/wirepod/speechrequest"
	ttr "github.com/kercre123/chipper/pkg/wirepod/ttr"

	//luowei add
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	//sdk_wrapper "github.com/fforchino/vector-go-sdk/pkg/sdk-wrapper" //拥有很多API的vector SDK包，可以用他发送语音文件播放
	"path/filepath"

	"github.com/fforchino/vector-go-sdk/pkg/vectorpb"
	"github.com/fforchino/vector-go-sdk/pkg/vector"
	sdkWeb "github.com/kercre123/chipper/pkg/wirepod/sdkapp"
	"errors"
	"time"
	"encoding/json"
)



func PlaySound(filename string,master_volume int,robot *vector.Vector,ctx context.Context) string {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		println("File not found!")
		return "failure"
	}

	var pcmFile []byte
	tmpFileName := strings.TrimSuffix(filename, filename[len(filename)-4:]) + ".pcm" //GetTemporaryFilename("sound", "pcm", true) //ffmpeg转换出来的文件名，结尾是pcm
	if strings.Contains(filename, ".pcm") || strings.Contains(filename, ".raw") {
		//fmt.Println("Assuming already pcm")
		pcmFile, _ = os.ReadFile(filename)
	} else {
		_, conError := exec.Command("ffmpeg", "-y", "-i", filename, "-f", "s16le", "-acodec", "pcm_s16le", "-ar", "16000", "-ac", "1", tmpFileName).Output()
		if conError != nil {
			fmt.Println(conError)
		}
		//fmt.Println("FFMPEG output: " + string(conOutput))
		pcmFile, _ = os.ReadFile(tmpFileName)
	}
	var audioChunks [][]byte
	for len(pcmFile) >= 1024 {
		audioChunks = append(audioChunks, pcmFile[:1024])
		pcmFile = pcmFile[1024:]
	}
	var audioClient vectorpb.ExternalInterface_ExternalAudioStreamPlaybackClient
	audioClient, _ = robot.Conn.ExternalAudioStreamPlayback(
		ctx,
	)
	audioClient.SendMsg(&vectorpb.ExternalAudioStreamRequest{
		AudioRequestType: &vectorpb.ExternalAudioStreamRequest_AudioStreamPrepare{
			AudioStreamPrepare: &vectorpb.ExternalAudioStreamPrepare{
				AudioFrameRate: 16000,
				AudioVolume:   20*uint32(master_volume), //0~5 -> 0~100
			},
		},
	})
	//fmt.Println(len(audioChunks))
	for _, chunk := range audioChunks {
		audioClient.SendMsg(&vectorpb.ExternalAudioStreamRequest{
			AudioRequestType: &vectorpb.ExternalAudioStreamRequest_AudioStreamChunk{
				AudioStreamChunk: &vectorpb.ExternalAudioStreamChunk{
					AudioChunkSizeBytes: 1024,
					AudioChunkSamples:   chunk,
				},
			},
		})
		time.Sleep(time.Millisecond * 30)
	}
	audioClient.SendMsg(&vectorpb.ExternalAudioStreamRequest{
		AudioRequestType: &vectorpb.ExternalAudioStreamRequest_AudioStreamComplete{
			AudioStreamComplete: &vectorpb.ExternalAudioStreamComplete{},
		},
	})
	os.Remove(tmpFileName)

	return "success"
}


func getSDKSettings(robot *vector.Vector,ctx context.Context) ([]byte, error) {
	resp, err := robot.Conn.PullJdocs(ctx, &vectorpb.PullJdocsRequest{
		JdocTypes: []vectorpb.JdocType{vectorpb.JdocType_ROBOT_SETTINGS},
	})
	if err != nil {
		return nil, err
	}
	json := resp.NamedJdocs[0].Doc.JsonDoc

	// json内容: {
	// 	"button_wakeword" : 0,
	// 	"clock_24_hour" : true,
	// 	"custom_eye_color" : {
	// 	   "enabled" : false,
	// 	   "hue" : 0,
	// 	   "saturation" : 0
	// 	},
	// 	"default_location" : "San Francisco, California, United States",
	// 	"dist_is_metric" : true,
	// 	"eye_color" : 3,
	// 	"locale" : "en-US",
	// 	"master_volume" : 3,
	// 	"temp_is_fahrenheit" : false,
	// 	"time_zone" : "Asia/Hong_Kong"
	//  }


	return []byte(json), nil
}


func RefreshSDKSettings(robot *vector.Vector,ctx context.Context) map[string]interface{} {

	var settings map[string]interface{}

	settingsJSON, err := getSDKSettings(robot,ctx)
	if err != nil {
		logger.Println("ERROR: Could not load Vector settings from JDOCS")
		return settings
	}

	//println(string(settingsJSON))

	json.Unmarshal([]byte(settingsJSON), &settings)
	return settings
}

func play_sound(sound_file_name string,botSerial string) string {

	robotObj, _, _ := sdkWeb.GetRobot(botSerial)
	robot := robotObj.Vector
	ctx := robotObj.Ctx


	settings := RefreshSDKSettings(robot,ctx)
	master_volume := int(settings["master_volume"].(float64))
	println("当前音量:",master_volume)

	start := make(chan bool)
	stop := make(chan bool)
	go func() {
		err := robot.BehaviorControl(ctx, start, stop)//控制机器人不要乱动，否则看起来会不像它在说话
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case <-start:
			//sdk_wrapper.SayText(transcribedText)

			ret := PlaySound(sound_file_name,master_volume,robot,ctx)
			logger.Println(ret)

			stop <- true
			return "intent_imperative_praise"
		}
	}
}





// func sdk_wrapper_play_sound(sound_file_name string,botSerial string) string {

// 	sdk_wrapper.InitSDKForWirepod(botSerial)
// 	ctx := context.Background()
// 	start := make(chan bool)
// 	stop := make(chan bool)
// 	go func() {
// 		err := sdk_wrapper.Robot.BehaviorControl(ctx, start, stop)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	for {
// 		select {
// 		case <-start:
// 			//sdk_wrapper.SayText(transcribedText)

// 			println("当前音量=", sdk_wrapper.GetMasterVolume() )

// 			ret := sdk_wrapper.PlaySound(sound_file_name)
// 			println(ret)


// 			stop <- true
// 			return "intent_imperative_praise"
// 		}
// 	}
// }


func (s *Server) ProcessIntentGraph(req *vtt.IntentGraphRequest) (*vtt.IntentGraphResponse, error) {
	sr.BotNum = sr.BotNum + 1
	var successMatched bool
	speechReq := sr.ReqToSpeechRequest(req)
	var transcribedText string
	if !isSti {//STT语音转文本
		var err error
		transcribedText, err = sttHandler(speechReq)
		if err != nil {
			sr.BotNum = sr.BotNum - 1
			ttr.IntentPass(req, "intent_system_noaudio", "voice processing error", map[string]string{"error": err.Error()}, true, speechReq.BotNum)
			return nil, nil
		}
		//和json文件做对比，看是否和技能匹配，匹配成功则执行技能
		successMatched = ttr.ProcessTextAll(req, transcribedText, vars.MatchListList, vars.IntentsList, speechReq.IsOpus, speechReq.BotNum)
	} else { //STI语音转意图
		intent, slots, err := stiHandler(speechReq)
		if err != nil {
			if err.Error() == "inference not understood" {
				logger.Println("No intent was matched")
				sr.BotNum = sr.BotNum - 1
				ttr.IntentPass(req, "intent_system_noaudio", "voice processing error", map[string]string{"error": err.Error()}, true, speechReq.BotNum)
				return nil, nil
			}
			logger.Println(err)
			sr.BotNum = sr.BotNum - 1
			ttr.IntentPass(req, "intent_system_noaudio", "voice processing error", map[string]string{"error": err.Error()}, true, speechReq.BotNum)
			return nil, nil
		}
		ttr.ParamCheckerSlotsEnUS(req, intent, slots, speechReq.IsOpus, speechReq.BotNum, speechReq.Device)
		sr.BotNum = sr.BotNum - 1
		return nil, nil
	}
	if !successMatched {//如果和技能没有匹配成功，则进入大模型对话环节
		logger.Println("No intent was matched.")
		if vars.APIConfig.Knowledge.Enable && vars.APIConfig.Knowledge.Provider == "openai" && len([]rune(transcribedText)) >= 8 {
			apiResponse := openaiRequest(transcribedText)
			sr.BotNum = sr.BotNum - 1
			response := &pb.IntentGraphResponse{
				Session:      req.Session,
				DeviceId:     req.Device,
				ResponseType: pb.IntentGraphMode_KNOWLEDGE_GRAPH,
				SpokenText:   apiResponse,
				QueryText:    transcribedText,
				IsFinal:      true,
			}
			req.Stream.Send(response)//发送脑脑电波类型的消息，以及回答的文本
			return nil, nil//结束函数
			
		} else if transcribedText != "" {
			
			sr.BotNum = sr.BotNum - 1
			response := &pb.IntentGraphResponse{
				Session:      req.Session,
				DeviceId:     req.Device,
				ResponseType: pb.IntentGraphMode_KNOWLEDGE_GRAPH,
				SpokenText:   "",
				QueryText:    "",
				IsFinal:      true,
			}
			req.Stream.Send(response)//发送脑脑电波类型的消息，用于接下来迅飞大模型运算过程的动画展示


			//执行迅飞语音大模型
			python_cmd := os.Getenv("PYTHON_CMD")
			python_chdir := os.Getenv("PYTHON_CHDIR")
			cmd := exec.Command(python_cmd, "-c",
			"import sys; import os; os.chdir('"+python_chdir+"'); from main import chat; chat(sys.argv[1])",
			transcribedText)
			//让python继承到go代码的环境变量
            cmd.Env = os.Environ()
			
			cmd_text := python_cmd + " -c " + "\"import sys; import os; os.chdir('"+python_chdir+"'); from main import chat; chat(sys.argv[1])\" " + transcribedText

			// 调用print(chat(sys.argv[1]))可以把返回值写入output
			// 执行命令并获取输出，注意python内部打印的东西也会写入output，所以output不光是print(chat(sys.argv[1]))打印的东西
			// 所以如果python代码里有打印内容，返回值就无法使用
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err,cmd_text," python运行出错 ",string(output))
				return nil, nil//python运行错误
			}

			print(string(output)) //[]byte -> string

			file_name := filepath.Join(python_chdir, "tts.wav")
			logger.Println("sound name: ",file_name)


			play_sound(file_name,req.Device)//req.Device就是SN序列号

			//sdk_wrapper_play_sound(file_name,req.Device)
			
			return nil, nil
		}
		
		//如果大模型没有开启，或者识别文本为空，则默认处理
		sr.BotNum = sr.BotNum - 1
		ttr.IntentPass(req, "intent_system_noaudio", transcribedText, map[string]string{"": ""}, false, speechReq.BotNum)//执行默认技能
		return nil, nil//结束函数
	}
	sr.BotNum = sr.BotNum - 1
	logger.Println("Bot " + strconv.Itoa(speechReq.BotNum) + " request served.")
	return nil, nil//上面已经执行过技能，结束函数
}
