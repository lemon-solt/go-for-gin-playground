package main

import routerlog "app/router-log"

func main() {
	// パスに含まれるパラメータ
	// r := paraminpath.ParaminpathSample()

	// ファイルアップロード
	// r := fileupload.FileUploadSingleSample()
	// r := fileupload.FileUploadMultiSample()

	// struct bind
	// r := requestbind.RequestBindSample()

	// query-post
	// r := queryandpostform.QueryAndPostSample()

	// using-goroutine
	// r := usingmiddleware.GoroutineApiSample()

	// validation-bind
	// r := validationbind.ValidationBindSample()

	// redirects
	// r := redirectrouter.RedicretSample()

	// ルーターグルーピング
	// r := groupingrouter.GroupingRouterSample()

	// ルーターログ
	r := routerlog.RouterLogSample()
	r.Run(":8082")
	// server run.
	// fmt.Println("ports: ", config.Config.Port)
	// r.Run(config.Config.Port)
}
