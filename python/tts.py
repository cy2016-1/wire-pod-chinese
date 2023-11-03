from sdk import XunfeiSpeech
import edge_tts
import asyncio
import pyttsx3
import requests

#迅飞
appid = ""     #填写控制台中获取的 APPID 信息
api_secret = ""   #填写控制台中获取的 APISecret 信息
api_key =""    #填写控制台中获取的 APIKey 信息

# def play_music(file_name):
#     pygame.mixer.init()
#     pygame.mixer.music.set_volume(1.0)  # 音量0~1
#     pygame.mixer.music.load(file_name)
#     pygame.mixer.music.play()
#
#     while pygame.mixer.music.get_busy():
#         time.sleep(0.1)
#
#     pygame.mixer.music.stop()

class XunfeiTTS():
    def __init__(self, appid, api_key, api_secret, voice="x3_xiaofang"): #xiaoyan x3_xiaofang
        self.appid, self.api_key, self.api_secret, self.voice_name = (
            appid,
            api_key,
            api_secret,
            voice,
        )

    def text_to_speech(self, phrase):
        return XunfeiSpeech.synthesize(
            phrase, self.appid, self.api_key, self.api_secret, self.voice_name
        )

class EdgeTTS():
    def __init__(self, voice="zh-CN-XiaoxiaoNeural"):
        self.voice = voice

    async def async_get_speech(self, text):
        try:
            tmpfile = "tts.mp3"
            tts = edge_tts.Communicate(text=text, voice=self.voice)
            await tts.save(tmpfile)
            return tmpfile
        except Exception as e:
            print("edge tts合成失败")
            return None

    def text_to_speech(self, text):
        event_loop = asyncio.new_event_loop()
        tmpfile = event_loop.run_until_complete(self.async_get_speech(text))
        event_loop.close()
        return tmpfile


class LocalTTS():
    def __init__(self):
        pass

    def text_to_speech(self, text):
        engine = pyttsx3.init()
        engine.setProperty('voice', 'zh')  # 开启支持中文
        engine.say(text)
        engine.runAndWait()
        return None



def create_tts(type):
    if type == "xunfei":
        tts_client = XunfeiTTS(appid, api_key, api_secret)
    elif type == "edge":
        tts_client = EdgeTTS()
    elif type == "local":
        tts_client = LocalTTS()
    else:
        print("tts type error!")
        exit(-1)

    return tts_client
