
# slack-horoscope
API that retrieves horoscope from other APIs and sends to a Slack app

This app can consume data from 
- https://github.com/andrechavesg/babi-api-horoscopo (portuguese)
- https://github.com/tapaswenipathak/Horoscope-API (english)

This project is based on https://github.com/JonathonGore/knowledge-base

**Setup**

```
git clone https://github.com/aon1/slack-horoscope.git
```

Choose source API

Edit file main.go and change the following lines
```
horoscopeService "github.com/aon1/slack-horoscope/services/babi.hefesto.io"
``` 
```
service, err = horoscopeService.New(restClient, conf.HoroscopeServices.BabiHefestoIO)
```
or
```
horoscopeService "github.com/aon1/slack-horoscope/services/horoscope-api.herokuapp.com"
```  
```
service, err = horoscopeService.New(restClient, conf.HoroscopeServices.HoroscopeAPIHeroku)
```

Install 
```
go install
```

Run
```
go run main.go
```

App is running on http://localhost:3000/daily

**Usage**
```
curl -d "text=libra" -X POST http://localhost:3000/daily
```

Available sun signs for https://github.com/andrechavesg/babi-api-horoscopo
```
aquario
peixes
aries
touro
gemeos
cancer
leao
virgem
libra
escorpiao
sagitario
capricornio
```

Available sun signs for https://github.com/tapaswenipathak/Horoscope-API
```
aries 
taurus
gemini
cancer
leo
virgo
libra
scorpio
sagittarius
capricorn
aquarius 
pisces
```

Slack config

- Go to https://api.slack.com/apps and create a new app
- Choose Slash command
- Create a command that will be issued on chat
- Fill in the address of the app on Request URL field

**TODO**

Endpoint to retrieve weekly horoscope
