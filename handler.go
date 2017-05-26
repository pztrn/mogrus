package mogrus

import (
    // stdlib
    "io"
    "strings"

    // github
    "github.com/sirupsen/logrus"
)

type LoggerHandler struct {
    // Logrus instances
    instances map[string]*logrus.Logger
}

// Adds output for logger handler.
// This actually creates new Logrus's Logger instance and configure
// it to write to given writer.
func (lh *LoggerHandler) CreateOutput(name string, writer io.Writer, colors bool) {
    // Formatter.
    logrus_formatter := new(logrus.TextFormatter)
    logrus_formatter.DisableTimestamp = false
    logrus_formatter.FullTimestamp = true
    logrus_formatter.QuoteEmptyFields = true
    logrus_formatter.TimestampFormat = "2006-01-02 15:04:05.000000000"

    if colors {
        logrus_formatter.DisableColors = false
        logrus_formatter.ForceColors = true
    } else {
        logrus_formatter.DisableColors = true
        logrus_formatter.ForceColors = false
    }

    logrus_instance := logrus.New()
    logrus_instance.Out = writer
    // ToDo: configurable level.
    logrus_instance.Level = logrus.DebugLevel
    logrus_instance.Formatter = logrus_formatter

    lh.instances[name] = logrus_instance

    for _, logger := range lh.instances {
        logger.Println("Added new logger output:", name)
    }
}

// Formats string by replacing "{{ var }}"'s with data from passed map.
func (lh *LoggerHandler) FormatString(data string, replacers map[string]string) string {
    for k, v := range replacers {
        data = strings.Replace(data, "{{ " + k + " }}", v, -1)
    }

    return data
}

// Initializes logger handler.
// It will only initializes LoggerHandler structure, see CreateOutput()
// for configuring output for this logger handler.
func (lh *LoggerHandler) Initialize() {
    lh.instances = make(map[string]*logrus.Logger)
}

// Proxy for Logrus's Logger.Debug() function.
func (lh *LoggerHandler) Debug(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Debug(args...)
    }
}

// Proxy for Logrus's Logger.Debugf() function.
func (lh *LoggerHandler) Debugf(format string, args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Debugf(format, args...)
    }
}

// Proxy for Logrus's Logger.Debugln() function.
func (lh *LoggerHandler) Debugln(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Debugln(args...)
    }
}

// Proxy for Logrus's Logger.Error() function.
func (lh *LoggerHandler) Error(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Error(args...)
    }
}

// Proxy for Logrus's Logger.Errorf() function.
func (lh *LoggerHandler) Errorf(format string, args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Errorf(format, args...)
    }
}

// Proxy for Logrus's Logger.Errorln() function.
func (lh *LoggerHandler) Errorln(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Errorln(args...)
    }
}

// Proxy for Logrus's Logger.Fatal() function.
func (lh *LoggerHandler) Fatal(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Fatal(args...)
    }
}

// Proxy for Logrus's Logger.Fatalf() function.
func (lh *LoggerHandler) Fatalf(format string, args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Fatalf(format, args...)
    }
}

// Proxy for Logrus's Logger.Fatalln() function.
func (lh *LoggerHandler) Fatalln(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Fatalln(args...)
    }
}

// Proxy for Logrus's Logger.Info() function.
func (lh *LoggerHandler) Info(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Info(args...)
    }
}

// Proxy for Logrus's Logger.Infof() function.
func (lh *LoggerHandler) Infof(format string, args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Infof(format, args...)
    }
}

// Proxy for Logrus's Logger.Infoln() function.
func (lh *LoggerHandler) Infoln(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Infoln(args...)
    }
}

// Proxy for Logrus's Logger.Print() function.
func (lh *LoggerHandler) Print(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Print(args...)
    }
}

// Proxy for Logrus's Logger.Printf() function.
func (lh *LoggerHandler) Printf(format string, args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Printf(format, args...)
    }
}

// Proxy for Logrus's Logger.Println() function.
func (lh *LoggerHandler) Println(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Println(args...)
    }
}

// Proxy for Logrus's Logger.Warn() function.
func (lh *LoggerHandler) Warn(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Warn(args...)
    }
}

// Proxy for Logrus's Logger.Warnf() function.
func (lh *LoggerHandler) Warnf(format string, args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Warnf(format, args...)
    }
}

// Proxy for Logrus's Logger.Warnln() function.
func (lh *LoggerHandler) Warnln(args ...interface{}) {
    for _, logger := range lh.instances {
        logger.Warnln(args...)
    }
}
