# Restr

A simple rest scripting language for http requests and stress testing

### Why?

Learning how to create an interpreted language using a book https://craftinginterpreters.com

### language

Go

### Draft syntax

```
var cookie = 'hello world'

var response = req POST 'http://google.com' {
    headers {
        x-my-header: 'hello'
        Cookie: cookie
    }

    json {
        "example": "json"
    }
}


stress 100 workers 10 {
    var response = req POST 'http://google.com' {
        headers {
            x-my-header: 'hello'
            Cookie: $cookie
        }

        metrics {
            capture latency
            capture status
        }

        json {
            "example": "json"
        }
    }
    print response.body
}
```

### AI

This project is NOT built with AI.
