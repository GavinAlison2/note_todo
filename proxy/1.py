import requests
import time,os

def get_proxy():
    return requests.get("http://127.0.0.1:5010/get/").json()

def delete_proxy(proxy):
    requests.get("http://127.0.0.1:5010/delete/?proxy={}".format(proxy))

# your spider code
setData = set()

def getHtml():
    retry_count = 20
    html = None
    while retry_count > 0:
        try:
            proxy = get_proxy().get("proxy")
            print(proxy)
            index = proxy.find(":")
            ip=proxy[0:index]
            # proxy = "111.3.102.207:30001"
            # proxy = "36.138.56.214:3128"
            # print({"http": "http://{}".format(proxy)})
            headers = {
                'Referer': 'http://127.0.0.1/',
                # 'Referer': 'http://117.149.243.239:9099/app/pages/login/index',
                # 'Cookie': 'PHPSESSID=a83cc506c37e72b9e25570cb61f520bb; MEIQIA_TRACK_ID=2STIFkvc9UwkcvHgOAVEHFesm8g; MEIQIA_VISIT_ID=2SVPrBrAO5NVqCcLCJamBAm3oow',
                'User-Agent': 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148',
                # 'Connection': 'keep-alive',
                # 'Host': '117.149.243.239:9099',
                # 'Accept': '*/*',
                # 'Accept-Encoding': 'gzip, deflate',
                # 'Accept-Language': 'zh-CN,zh-Hans;q=0.9',
                # 'Vary': 'Access-Control-Request-Method',
                # 'Vary': 'Access-Control-Request-Headers',
                # 'Vary': 'Origin',
                # 'Content-Type': 'application/json',
                'Server': 'nginx',
                'Proxy-Tunnel':'2033',
                'Host':'127.0.0.1:8080',
                'Remote-Addr': ip,
                'Via':'127.0.0.1',
                'X-Forwarded-For': ip,
                # 'X-Content-Type-Options': 'nosniff',
                # 'X-XSS-Protection': '1; mode=block',
                # 'Transfer-Encoding': 'chunked',
                # 'Strict-Transport-Security': 'max-age=31536000'
            }
            html = requests.get('http://175.178.234.145:9099/apiApp', 
                                proxies={"http": "http://{}".format(proxy)},
                                headers=headers, 
                                timeout=10,
                                allow_redirects=False)

            # html = requests.get('http://117.149.243.239:9099/apiApp/app/site/info', 
                                # proxies={"http": "http://{}".format(proxy)}, headers=headers,
                                # timeout=10)
            # 使用代理访问
            print(html.reason)
            if(html.status_code == 404):
                raise Exception('404')
            setData.add(proxy+"\n")
            # if(retry_count == 0):
            # return html
        except Exception as e:
            time.sleep(0.5)
            print('error: %s'%e)
            print('-'*10+str(retry_count))
            retry_count -= 1
    # 删除代理池中代理
    delete_proxy(proxy)
    # return None
    return html


def writeFile(setData: set):
    writeFilename = 'proxy/2.txt'
    if not os.path.exists(writeFilename):
        os.system(r'touch {}'.format(writeFilename))
    with open(writeFilename, 'a', encoding="utf-8") as f:
        f.writelines(setData)
    print('write data ok')

if __name__=='__main__':
    html = getHtml()
    if(html):
        print(html.text)
    else:
        print(html)
    print(setData)
    writeFile(setData=setData)