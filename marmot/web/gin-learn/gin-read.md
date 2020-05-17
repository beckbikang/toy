# gin


## 提供的功能

基本的http服务

基本的插件


## go的http服务

基本流程

```
func IndexHandler(w http.ResponseWriter, r *http.Request)
http.HandleFunc("/", IndexHandler)
http.ListenAndServe("127.0.0.0:8000", nil)

或者
server := &Server{Addr: addr, Handler: handler}
server.ListenAndServe()



client->req -> [multiplexer(router)] -> handler -> reponse -> clinet
```
结构

```

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}


type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry
    hosts bool 
}
 
type muxEntry struct {
    explicit bool
    h        Handler
    pattern  string
}


type Server struct {
    Addr         string        
    Handler      Handler       
    ReadTimeout  time.Duration 
    WriteTimeout time.Duration 
    TLSConfig    *tls.Config   
 
    MaxHeaderBytes int
 
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
 
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
    disableKeepAlives int32     nextProtoOnce     sync.Once 
    nextProtoErr      error     
}



var DefaultServeMux = &defaultServeMux 
var defaultServeMux ServeMux
func NewServeMux() *ServeMux {
    return new(ServeMux)
}
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}

默认的包装
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}

http的接口

type ResponseWriter interface {
    Header() Header//header投
    WriteHeader(int)//http状态码
    Write([]byte) (int, error)//写byte
}

type Flusher interface {
    Flush()//Flush将缓冲中的所有数据发送到客户端
}
type CloseNotifier interface {
    CloseNotify() <-chan bool// CloseNotify返回一个通道，该通道会在客户端连接丢失时接收到唯一的值
}

type Hijacker interface {
    // Hijack让调用者接管连接，返回连接和关联到该连接的一个缓冲读写器。,调用本方法后，HTTP服务端将不再对连接进行任何操作, 调用者有责任管理、关闭返回的连接。
    Hijack() (net.Conn, *bufio.ReadWriter, error)
}

type RoundTripper interface {
    // RoundTrip执行单次HTTP事务，接收并发挥请求req的回复。RoundTrip不应试图解析/修改得到的回复。 尤其要注意，只要RoundTrip获得了一个回复，不管该回复的HTTP状态码如何，它必须将返回值err设置为nil。非nil的返回值err应该留给获取回复失败的情况。类似的，RoundTrip不能试图管理高层次的细节，如重定向、认证、cookie。除了从请求的主体读取并关闭主体之外，RoundTrip不应修改请求，包括（请求的）错误。RoundTrip函数接收的请求的URL和Header字段可以保证是（被）初始化了的。
    RoundTrip(*Request) (*Response, error)
}

```



## 主流程分析

gin的基础流程
```
r := gin.New()
r.GET("/", func(c *gin.Context) {}
r.Run(":" + port)
```

### gin简单主流程分析


1 新建一个引擎，引擎注入了RouterGroup

```
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		FuncMap:                template.FuncMap{},
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		ForwardedByClientIP:    true,
		AppEngine:              defaultAppEngine,
		UseRawPath:             false,
		RemoveExtraSlash:       false,
		UnescapePathValues:     true,
		MaxMultipartMemory:     defaultMultipartMemory,
		trees:                  make(methodTrees, 0, 9),
		delims:                 render.Delims{Left: "{{", Right: "}}"},
		secureJSONPrefix:       "while(1);",
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}
```

2 注册方法和请求，通过Group处理

	1 通过group的handler增加addRoute方法
	2 
	
	

3 使用  http.ListenAndServe(address, engine) 运行服务

```
1 从pool获取context，
2 重置responsed
3 设置request
4 重置context
5 处理请求
6 回收context
7 

```





## 功能分析

doc 文档配置

mode.go 设置编程模式，设置binding的方式

debug.go 打印各种debug信息

util.go 
	
```
type H map[string]interface{}

	1 绑定变量
	2 http.HandlerFunc转HandlerFunc
	3 http.Handler转HandlerFunc
	4 type H map[string]interface{}
	5 Map解析成map
	6 assert检查
	7 截断，chooseData，，函数名，最后一个字符串
	8 获取地址
```

error.go，error的string化

```
两种类型的错误出出力

type Error struct {
	Err  error
	Type ErrorType
	Meta interface{}
}
type errorMsgs []*Error

1 各类错误的定义
2 errorjson化
3 合适的类型处理
4 Last()
5 Errors()错误的string


```

logger.go 输出默认的日志

```
打印日志，设置颜色
type consoleColorModeValue int
consoleColorMode
//日志配置
type LoggerConfig struct {
	Formatter LogFormatter
	Output io.Writer
	SkipPaths []string
}
//日志函数
type LogFormatter func(params LogFormatterParams) string
//日志参数
type LogFormatterParams struct {
	Request *http.RequestTimeStamp time.Time
	StatusCode int
	Latency time.Duration
	ClientIP string
	Method string
	Path string
	ErrorMessage string
	isTerm bool//是否输出到终端
	BodySize int
	Keys map[string]interface{}
}

defaultLogFormatter

func ErrorLogger() HandlerFunc()

定义通用的Logger，LoggerWithFormatter，LoggerWithWriter

SkipPaths is a url path array which logs are not written.

```

gin.go结构

```
type HandlerFunc func(*Context)//方法处理函数
type HandlersChain []HandlerFunc//处理的slice

//路由信息
type RouteInfo struct {
	Method      string
	Path        string
	Handler     string //名称
	HandlerFunc HandlerFunc
}
type RoutesInfo []RouteInfo //路由列表

type Engine struct {
	RouterGroup//分组信息
	RedirectTrailingSlash bool//是否去除/301跳转
	RedirectFixedPath bool去除//划线
	HandleMethodNotAllowed bool//是否处理没有的方法
	ForwardedByClientIP    bool
	AppEngine bool //开启特殊header
	UseRawPath bool//是否允许url.RawPath
	UnescapePathValues bool//是否需要反解值？
	MaxMultipartMemory int64//最大分片内存
	RemoveExtraSlash bool//是否移除额外的反斜杠
	
	//模板相关的
	delims           render.Delims//模板分隔符
	secureJSONPrefix string
	HTMLRender       render.HTMLRender
	FuncMap          template.FuncMap
		
	allNoRoute       HandlersChain//没有路由的处理链条
	allNoMethod      HandlersChain//无方法
	noRoute          HandlersChain//无路由
	noMethod         HandlersChain//无方法
	pool             sync.Pool//池
	trees            methodTrees//方法树
	maxParams        uint16//最大参数数
}


//pool的使用
func (engine *Engine) allocateContext() *Context {
	v := make(Params, 0, engine.maxParams)
	return &Context{engine: engine, params: &v}
}

 设置Delims， secureJSONPrefix
 设置模板，通过全局变量，通过文件路径，设置html模板
 设置模板func
 设置处理函数
 
 添加路由
 获取所有的路由信息
 运行htt和https服务器，通过socket运行，通过fd运行，listener运行
 
 ServeHTTP(w http.ResponseWriter, req *http.Request) 处理http请求
 
 
 handleHTTPRequest(c *Context)处理请求

serveError处理error
 



//对head X-Forwarded-Prefix处理，处理/
redirectTrailingSlash(c *Context)

//301跳转， 
redirectRequest(c *Context)

redirectFixedPath?


```



tree.go 路由树

```
type Param struct {
	Key   string
	Value string
}
type Params []Param

//方法数据
type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

countParams 计算*和:的个数


```



context.go的结构

```
type Context struct {
	writermem responseWriter //对Writer的包装
	Request   *http.Request//请求头
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine *Engine
	params *Params
	mu sync.RWMutex
	Keys map[string]interface{}
	Errors errorMsgs
	Accepted []string
	queryCache url.Values//Request.URL.Query()
	formCache url.Values //req.PostForm
	sameSite http.SameSite
}

重置context

func (c *Context) reset() {
	c.Writer = &c.writermem
	c.Params = c.Params[0:0]
	c.handlers = nil
	c.index = -1

	c.fullPath = ""
	c.Keys = nil
	c.Errors = c.Errors[0:0]
	c.Accepted = nil
	c.queryCache = nil
	c.formCache = nil
	*c.params = (*c.params)[0:0]
}

复制context

停止context的运行：Abort()


Error(err error)

获取参数的方法

上传文件

各种绑定

获取客户端ip



```




routergroup.go

```
type IRouter interface {
	IRoutes//单个路由接口
	Group(string, ...HandlerFunc) *RouterGroup//路由组接口
}
//路由处理接口
type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}
//路由grouper
type RouterGroup struct {
	Handlers HandlersChain//方法链
	basePath string//基本路径
	engine   *Engine//引擎
	root     bool//是否是router
}

1 支持增加多个处理方法[]HandlerFunc
2 returnObj() IRoutes
3 合并处理器，合并路径
4 根据path创建Group, 添加处理器
5 Handler，POST,GET,DELETE,各类方法
6 静态文件，静态文件处理器

```


response_writer.go 响应写的接口

```
重置，写status，写入[]byte，写入string
连接器管理，返回一个通道，该通道会在客户端连接丢失时接收到唯一的值
获取http的pusher


noWritten     = -1
defaultStatus = http.StatusOK
//响应的接口
type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier
	Status() int
	Size() int
	WriteString(string) (int, error)
	Written() bool
	WriteHeaderNow()
	Pusher() http.Pusher
}
//响应的返回
type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

```


auth.go 认证

```
const AuthUserKey = "user"
type Accounts map[string]string
type authPair struct {
	value string
	user  string
}
type authPairs []authPair


processAccounts(map[string]string) 得到pairs

BasicAuthForRealm做个认证

```


path.go

```
cleanPath(p string)  去除../和./
```


recovery.go

```

//捕获异常继续执行
RecoveryWithWriter(out io.Writer) HandlerFunc



获取函数名
function(pc uintptr)
	1 / dunno dot

返回栈信息
stack(skip int)

trim空格
func source(lines [][]byte, n int) []byte 

获取当前时间
```


fs.go http文件处理

```
type onlyfilesFS struct {
	fs http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

```



### ginS

构建一个内部的包

```
var once sync.Once
var internalEngine *gin.Engine
```


### internal包

json处理

三方包 json-iterator

不复制的类型转换

### render包

render.go

```
type Render interface {
	// Render writes data with custom ContentType.
	Render(http.ResponseWriter) error
	// WriteContentType writes custom ContentType.
	WriteContentType(w http.ResponseWriter)
}
```

data.go写入数据

```
type Data struct {
	ContentType string
	Data        []byte
}

// Render (Data) writes data with custom ContentType.
func (r Data) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	_, err = w.Write(r.Data)
	return
}

// WriteContentType (Data) writes custom ContentType.
func (r Data) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, []string{r.ContentType})
}
```

html.go go的html模板

json.go

msgpack.go

protobuf.go

redict.go

text.go

xml.go

yaml.go












