package logCompressor

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"runtime"

	"github.com/zeromicro/go-zero/core/threading"
)

type Printf func(format string, args ...interface{})

func GetIdentify() string {
	_, file, line, _ := runtime.Caller(1)

	return fmt.Sprintf("%v-%v", file, line)
}
func getIdentify() string {
	_, file, line, _ := runtime.Caller(2)

	return fmt.Sprintf("%v-%v", file, line)
}

type Compressor struct {
	Message  string
	Format   string
	Args     []interface{}
	count    int64
	printf   Printf
	maxCount int64
	wight    int
}

func NewCompressor(maxCount int64, format string, args ...interface{}) *Compressor {
	return &Compressor{
		Format:   format,
		Args:     args,
		count:    0,
		maxCount: maxCount,
		wight:    int(math.Log10(float64(maxCount)) + 1),
	}
}

func (c *Compressor) addMessage(printf Printf, format string, args ...interface{}) {
	if fmt.Sprint(printf) != fmt.Sprint(c.printf) || c.Format != format || len(c.Args) != len(args) || !reflect.DeepEqual(c.Args, args) {
		c.formatPrint()
		c.setPrintf(printf)
		c.setArgs(format, args...)
		c.count = 0
		c.formatPrint()
	}
	c.count++
}

func (c *Compressor) setPrintf(printf Printf) {
	c.printf = printf
}

func (c *Compressor) setArgs(format string, args ...interface{}) {
	c.Format, c.Args = format, args
}

func (c *Compressor) formatPrint() {
	format := "[zip*%." + fmt.Sprint(c.wight) + "v] " + c.Format
	args := append([]interface{}{c.count}, c.Args...)
	c.printf(format, args...)
}

func (c *Compressor) tryPrint() {
	if c.count >= c.maxCount && c.maxCount != 0 {
		if c.printf != nil {
			c.formatPrint()
			c.count = 0
		}
	}
}

func (c *Compressor) setMaxCount(maxCount ...int64) {
	if len(maxCount) > 0 {
		c.maxCount = maxCount[0]
		c.wight = int(math.Log10(float64(c.maxCount)) + 1)
	}
}

type CompressorGroup struct {
	ctx            context.Context
	IdentifyMap    map[string]*Compressor
	GlobalMaxCount int64
}

func (cg *CompressorGroup) listeningStop() {
	<-cg.ctx.Done()
	for _, compressor := range cg.IdentifyMap {
		compressor.formatPrint()
	}
	cg.IdentifyMap = map[string]*Compressor{}
	cg.GlobalMaxCount = 0
}

func NewCompressorGroup(ctx context.Context, maxCount int64) *CompressorGroup {
	cg := &CompressorGroup{
		ctx:            ctx,
		IdentifyMap:    map[string]*Compressor{},
		GlobalMaxCount: maxCount,
	}

	threading.GoSafe(cg.listeningStop)

	return cg
}

func (cg *CompressorGroup) hasCompressor(identify string) bool {
	_, ok := cg.IdentifyMap[identify]

	return ok
}

func (cg *CompressorGroup) getCompressor(identify string) *Compressor {
	return cg.IdentifyMap[identify]
}

func (cg *CompressorGroup) addCompressor(identify string, format string, args ...interface{}) {
	_, ok := cg.IdentifyMap[identify]
	if !ok {
		cg.IdentifyMap[identify] = NewCompressor(cg.GlobalMaxCount, format, args...)
	}
}

func (cg *CompressorGroup) SetCompressorMaxCount(identify string, maxCount int64) {
	if cg.hasCompressor(identify) {
		cg.getCompressor(identify).setMaxCount(maxCount)
	}
}

func (cg *CompressorGroup) Loader(commonIdentify ...string) Loader {
	identify := ""
	if len(commonIdentify) > 0 {
		identify = commonIdentify[0]
	} else {
		identify = getIdentify()
	}

	return Loader{
		compressorGroup: cg,
		identify:        identify,
	}
}

func (cg *CompressorGroup) AddMessage(printf Printf, format string, args ...interface{}) {
	identify := getIdentify()
	if !cg.hasCompressor(identify) {
		cg.addCompressor(identify, format, args...)
		cg.IdentifyMap[identify].setPrintf(printf)

		cg.IdentifyMap[identify].formatPrint()
	}

	cg.IdentifyMap[identify].addMessage(printf, format, args...)
	cg.IdentifyMap[identify].tryPrint()
}

type Loader struct {
	compressorGroup *CompressorGroup
	identify        string
}

func (l Loader) AddMessage(printf Printf, format string, args ...interface{}) {
	cg, identify := l.compressorGroup, l.identify
	if !cg.hasCompressor(identify) {
		cg.addCompressor(identify, format, args...)
		cg.IdentifyMap[identify].setPrintf(printf)

		cg.IdentifyMap[identify].formatPrint()
	}

	cg.IdentifyMap[identify].addMessage(printf, format, args...)
	cg.IdentifyMap[identify].tryPrint()
}
