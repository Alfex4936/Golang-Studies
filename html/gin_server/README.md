<div align="center">
<p>
    <img width="190" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
</p>
<h1>Golang HTTP server using gin</h1>
    <h5>go v1.15.3</h5>

[Go gin](https://github.com/gin-gonic/gin)

</div>

## 카카오 챗봇 서버 사용법

:eye_speech_bubble: Run the server first to see the results. (localhost:8000)

```console

WIN10@DESKTOP:~$ git clone https://github.com/Alfex4936/Golang-Studies
WIN10@DESKTOP:~$ cd ./html/gin_server

WIN10@DESKTOP:~$ go get -u github.com/gin-gonic/gin

WIN10@DESKTOP:~$ go run KakaoChatBotAPI.go

[GIN-debug] GET    /                         --> main.main.func1 (4 handlers)
[GIN-debug] POST   /json                     --> main.main.func2 (4 handlers)
[GIN-debug] Listening and serving HTTP on :8000
```

## 카카오 챗봇 JSON

카카오 챗봇 JSON 예제 데이터 
(utterance = "2021 검색해줘"
Entity:sys_text = "2021")
```json
{
    "action": {
        "clientExtra": {},
        "detailParams": {
            "sys_text": {
                "groupName": "",
                "origin": "2021",
                "value": "2021"
            }
        },
        "id": "idaction",
        "name": "스킬 이름",
        "params": {
            "sys_text": "2021"
        }
    },
    "bot": {
        "id": "id",
        "name": "AjouNotice"
    },
    "contexts": [],
    "intent": {
        "extra": {
            "reason": {
                "code": 1,
                "message": "OK"
            }
        },
        "id": "idintent",
        "name": "공지 키워드 검색"
    },
    "userRequest": {
        "block": {
            "id": "iduserRe",
            "name": "공지 키워드 검색"
        },
        "lang": "kr",
        "params": {
            "ignoreMe": "true",
            "surface": "BuilderBotTest"
        },
        "timezone": "Asia/Seoul",
        "user": {
            "id": "idUser",
            "properties": {
                "botUserKey": "keybot",
                "bot_user_key": "keybotuser"
            },
            "type": "botUserKeyType"
        },
        "utterance": "2021 검색\n"
    }
}
```

## Golang Nested structs JSON Marshal
To see how to parse nested struct as Kakao's ListCard JSON, check [ListCardJSON.go](https://github.com/Alfex4936/Golang-Studies/tree/master/html/gin_server/ListCardJSON.go)

!Used [clarketm's json](https://pkg.go.dev/github.com/clarketm/json?utm_source=gopls) package!

JSON result
```json
{
  "template": {
    "outputs": [
      {
        "listCard": {
          "buttons": [
            {
              "label": "hey1",
              "action": "share"
            },
            {
              "label": "hey2",
              "action": "share"
            }
          ],
          "header": {
            "title": "hi"
          },
          "items": [
            {
              "description": "desc1",
              "imageUrl": "img",
              "link": {
                "web": "web"
              },
              "title": "title1"
            },
            {
              "description": "desc2",
              "title": "title2"
            }
          ]
        }
      }
    ],
    "quickReplies": [
      {
        "action": "message",
        "label": "어제",
        "messageText": "어제 보여줘"
      }
    ]
  },
  "version": "2.0"
}
```


## cURL POST 테스트
```console
WIN10@DESKTOP:~$ curl http://localhost:8000/json -d "@data.json"

{"message":"Reason:1 | Params['sys_text']:2021 | Utterance:2021 검색\n | UserID:idUser"}
```