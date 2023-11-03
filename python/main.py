import asr,llm,tts
import os


#asr_client = asr.create_asr("baidu")
llm_client = llm.create_llm("xunfei")
tts_client = tts.create_tts("edge")

def chat(text):
    #text_input = "你是我的宠物猫，你的名字叫“歪克特”，你要用一只软萌猫咪的语气和我聊天，并且你每次聊天都要以“喵呜”结尾，回答字数必须小于30字，现在我聊天的内容是：" + text
    text_input = "你是我的宠物机器人，你的名字叫vector，外形酷似一个小推土机，你要用一只软萌机器人的语气和我聊天，并且你每次聊天都要以“喵呜”结尾，回答字数必须小于30字，现在我聊天的内容是：" + text
    answer = llm_client.chat(text_input)
    print("AI:",answer)
    if answer == "":
        print("聊天回复为空") #可能问了不合法的问题，不合法的问题有可能返回的是空，也有可能会有回复
        answer = "喵呜,喵呜～"

    mp3_name = tts_client.text_to_speech(answer)

    out_name =  "tts.wav"

    #os.system("ffmpeg -i %s -f s16le -acodec pcm_s16le %s"%(mp3_name,pcm_name))
    
    #cmd = "sox -v 10.0  %s -r 16000 -c 1 %s  >/dev/null 2>&1" % (fname, wave_fname) #会失真，但音量可以放大
    #cmd = "sox --norm=-1 %s -r 16000 -c 1 %s  >/dev/null 2>&1" % (fname, wave_fname) #归一化放大
    
    os.system("sox -v 10.0 %s -r 16000 -c 1 %s  >/dev/null 2>&1" % (mp3_name,out_name)) #会输出tts.wav文件，在go语言中需要读取tts.wav文件
   
    return out_name

if __name__ == "__main__":
    chat("你好")
