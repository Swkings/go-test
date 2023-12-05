package test

// import (
// 	"fmt"

// 	"github.com/zeromicro/go-zero/core/breaker"
// )

// //‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// //  common methods for request
// //_______________________________________________________________________

// type MethodType string

// const (
// 	GET  MethodType = "GET"
// 	POST MethodType = "POST"
// )

// type RequestInput[RespType any, BodyType any] struct {
// 	Method            MethodType // GET, POST
// 	HttpClient        *resty.Client
// 	Body              func() BodyType
// 	Url               string
// 	ReplyInfo         func(resp RespType) string
// 	SetError          func(err error, resp RespType) error
// 	UseOverloadPolicy bool // if UseOverloadPolicy, ServiceName must not be empty
// 	ServiceName       string
// }

// // DoRequest do GET/POST request
// func DoRequest[RespType any, BodyType any](svcCtx *svc.ServiceContext, in RequestInput[RespType, BodyType], identification ...string) (resp RespType, err error) {
// 	fmtFuncName := util.GetFmtFuncName(identification...)
// 	svcCtx.Log.Infof("%v request(%v) %v", fmtFuncName, in.Method, in.Url)

// 	if in.HttpClient == nil {
// 		return resp, ErrorEmptyHttpClient
// 	}

// 	request := in.HttpClient.R()
// 	if in.Body != nil {
// 		request.SetBody(in.Body())
// 	}
// 	request.SetResult(&resp).SetError(&resp)

// 	switch in.Method {
// 	case GET:
// 		err = doGet[RespType](request, in)
// 	case POST:
// 		err = doPost[RespType](request, in)
// 	default:
// 		err = fmt.Errorf("%v is invalid method", in.Method)
// 	}

// 	if err != nil {
// 		svcCtx.Log.Warnf("%v request err: %v", fmtFuncName, err)
// 	} else {
// 		if in.ReplyInfo != nil {
// 			in.ReplyInfo(resp)
// 		}
// 	}

// 	if in.SetError != nil {
// 		err = in.SetError(err, resp)
// 	}

// 	return resp, err
// }

// func doGet[RespType any, BodyType any](request *resty.Request, in RequestInput[RespType, BodyType]) (err error) {
// 	if in.UseOverloadPolicy {
// 		err = breaker.Do(
// 			fmt.Sprintf("%v:%v", in.ServiceName, in.Url),
// 			func() error {
// 				_, err := request.Get(in.Url)
// 				return err
// 			},
// 		)
// 	} else {
// 		_, err = request.Get(in.Url)
// 	}

// 	return err
// }

// func doPost[RespType any, BodyType any](request *resty.Request, in RequestInput[RespType, BodyType]) (err error) {
// 	if in.UseOverloadPolicy {
// 		err = breaker.Do(
// 			fmt.Sprintf("%v:%v", in.ServiceName, in.Url),
// 			func() error {
// 				_, err := request.Post(in.Url)
// 				return err
// 			},
// 		)
// 	} else {
// 		_, err = request.Post(in.Url)
// 	}

// 	return err
// }
