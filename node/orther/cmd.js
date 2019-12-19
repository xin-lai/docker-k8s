
var exec = require('child_process').exec;
var cmdList = ['hexo server -p 8000']

var index = 0;
var execCmd = function () {
    console.info(cmdList[index]);
    exec(cmdList[index], outFunc);
}
var outFunc = function (error, stdout, stderr) {
    if (error) {
        console.error(`执行出错: ${error}`);
        return;
    }
    console.log(`stdout: ${stdout}`);
    console.log(`stderr: ${stderr}`);
    index++;
    if (index + 1 > cmdList.length) { return; }
    else {
        execCmd();
    }
};

execCmd();

// var exec = require('child_process').exec;
// var outFunc = function (error, stdout, stderr) {
//     if (error) {
//         console.error(`执行出错: ${error}`);
//         return;
//     }
//     console.log(`stdout: ${stdout}`);
//     console.log(`stderr: ${stderr}`);
// };
// var execCmd = function (cmd,func) {
//     console.info(cmd);
//     exec(cmd, outFunc);
// }
// exec('cd ../src', function (error){
//     if (error) {
//         console.error(`执行出错: ${error}`);
//         return;
//     }
//     exec('hexo generate -w -d', function (error, stdout, stderr) {
//         if (error) {
//             console.error(`执行出错: ${error}`);
//             return;
//         }
//         console.log(`stdout: ${stdout}`);
//         console.log(`stderr: ${stderr}`);
//         execCmd('hexo server -p 8000 -s -l');
//     })
// })

