// 加载http模块
const http = require('http');
// 设置端口
const port = 80;
// 创建Web服务器
const server = http.createServer((req, res) => {
    // 设置响应的状态码
    res.statusCode = 200;
    // 设置响应的请求头
    res.setHeader('Content-Type', 'text/plain');
    // 设置响应输出文本
    res.end('Hello World !');
});
// 设置Web服务器监听端口
server.listen(port);