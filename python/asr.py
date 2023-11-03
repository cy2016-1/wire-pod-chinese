from sdk import XunfeiSpeech
from aip import AipSpeech

#迅飞
appid = ""     #填写控制台中获取的 APPID 信息
api_secret = ""   #填写控制台中获取的 APISecret 信息
api_key =""    #填写控制台中获取的 APIKey 信息

#百度
APP_ID = ''
API_KEY = ''
SECRET_KEY = ''

# 读取文件
def get_file_content(filePath):
    with open(filePath, 'rb') as fp:
        return fp.read()

class XunfeiASR():
    def __init__(self, appid, api_key, api_secret, **args):
        super(self.__class__, self).__init__()
        self.appid = appid
        self.api_key = api_key
        self.api_secret = api_secret

    def audio_to_text(self, fp):
        return XunfeiSpeech.transcribe(fp, self.appid, self.api_key, self.api_secret)


class BaiduASR():
    def __init__(self, appid, api_key, api_secret, **args):
        super(self.__class__, self).__init__()
        self.appid = appid
        self.api_key = api_key
        self.api_secret = api_secret
        self.client = AipSpeech(appid, api_key, api_secret)

    def audio_to_text(self, fp):
        res = self.client.asr(get_file_content(fp), 'wav', 16000, {'dev_pid': 1537, })  # 直接读文件
        #print("res=",res)
        if 'result' in res:
            res = res['result'][0]
        else:
            res = ""

        return res

def create_asr(type):
    if type == "xunfei":
        asr_client = XunfeiASR(appid, api_key, api_secret)
    elif type == "baidu":
        asr_client = BaiduASR(APP_ID, API_KEY, SECRET_KEY)
    else:
        print("asr type error!")
        exit(-1)

    return asr_client


