# URL_to_Png_OR_PDF_DownloadService
中文描述：启动web服务，输入URL链接，将网页生成PDF文件或PNG格式图片，并返回文件实体。
English description：After starting the web service, input the URL link, and you can convert the webpage into a PDF file or a PNG format image, and return the corresponding file entity.

# Documentation    说明

本程序启动web服务后，输入URL链接，即可将网页转换成PDF文件或PNG格式图片，并返回相应的文件实体。
## 如何使用
### 方法一：自行编译。您可以运行命令 go run ./main.go启动 或者使用 go build 进行编译打包。
### 方法二：对于不熟悉编译过程的用户，您可以下载我提供的编译包。下载后，直接运行使用。
## 接口参数
程序启动后，将在端口52001上启动HTTP服务。您仅需向此服务接口发送GET请求即可开始使用。接口地址为http://localhost:52001/render，您需要携带以下参数：
url（需进行URL编码）
format（可选为png或pdf），您将获得类似以下格式的链接：http://localhost:52001/render?format=png&url=https%3a%2f%2fbujia.tyxxtb.com%2ffdd%2flsign%3ftravel_code%3d5202085520250116153732%26contact_phone%3d85cae7b39d1c66faa5f6d47e030b50a3
的链接。



After this program starts the web service, by inputting a URL link, you can convert the webpage into a PDF file or a PNG format image and return the corresponding file entity.
Project Description
## How to Use
### Method One: Compile it yourself. You can run the command go run ./main.go to start or use go build to compile and package.
### Method Two: For users unfamiliar with the compilation process, you can download the compiled package I provide. After downloading, you can run and use it directly.
## Interface Parameters
After the program starts, it will launch an HTTP service on port 52001. You only need to send a GET request to this service interface to begin using it. The interface address is http://localhost:52001/render, and you need to carry the following parameters:
url (needs to be URL encoded)
format (optional as png or pdf), and you will get a link similar to: http://localhost:52001/render?format=png&url=https//www.baidu.com
the link.
