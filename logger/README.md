### Logging in ImageKit SDK
The default logger in ImageKit Go SDK is `go log`.

You can use any log library by overwriting the standard SDK logging functions.

#### Using logrus with the SDK
```go
package main

import (
	"github.com/dhaval070/imagekit-go"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	var imgkit, err = imagekit.New()
	if err != nil {
		log.Fatalf("Failed to intialize ImageKit, %v", err)
	}

	// Initialize your logger somewhere in your code.
	// Set imgkit.Logger.Writer with logrus instance
	var logger = logrus.New()
	imgkit.Logger.Writer = logger.WithField("source", "ImageKit")
}
```

#### Using Zap with the SDK
```go
package main

import (
	"github.com/dhaval070/imagekit-go"
	"go.uber.org/zap"
	"log"
)

func main() {
	var imgkit, err = ImageKit.New()
	if err != nil {
		log.Fatalf("Failed to intialize ImageKit, %v", err)
	}

	// Initialize your logger somewhere in your code.
	// Set ImageKit.Logger.Writer with zap.SugaredLogger instance
	var zapLogger, _ = zap.NewDevelopment()
	imgkit.Logger.Writer = zapLogger.Sugar().With("source", "ImageKit")
}
```

#### Logging level

You can change logging level with the `Logger.SetLevel()` function.

Possible values:
- `logger.NONE`  - disabling logging from the SDK
- `logger.ERROR` - enable logging only for error messages
- `logger.DEBUG` - enable debug logs
