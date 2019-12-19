using Microsoft.Extensions.Configuration;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using System;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading.Tasks;

namespace DingTalk.NET
{
    internal class Program
    {
        /// <summary>
        /// 环境变量列表
        /// </summary>
        private static readonly string[] EnvList =
        {
            //钉钉机器人地址
            "WEBHOOK",
            //@的手机号码
            "AT_MOBILES",
            //@所有人
            "IS_AT_ALL",
            //消息内容
            "MESSAGE",
            //消息类型（仅支持文本和markdown）
            "MSG_TYPE"
        };

        private static void Main(string[] args)
        {
            var config = new ConfigurationBuilder()
                        //支持命令行参数
                        .AddCommandLine(args)
                        //支持环境变量
                        .AddEnvironmentVariables()
                        .Build();
            #region 参数检查
            foreach (var envName in EnvList)
            {
                var value = config[envName];
                if (string.IsNullOrWhiteSpace(value) && envName != "AT_MOBILES" && envName != "IS_AT_ALL")
                {
                    Console.WriteLine($"{envName} 不能为空！");
                    return;
                }
            }
            if (string.IsNullOrWhiteSpace(config["AT_MOBILES"]) && string.IsNullOrWhiteSpace(config["IS_AT_ALL"]))
            {
                Console.WriteLine("必须设置参数 AT_MOBILES 和 IS_AT_ALL 两者之一！");
                return;
            }
            #endregion
            try
            {
                //推送消息
                SetDataAndSendWebhooks(config).Wait();
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.ToString());
            }
        }

        /// <summary>
        /// 设置消息并调用Webhook
        /// </summary>
        /// <param name="config"></param>
        /// <returns></returns>
        private static async Task SetDataAndSendWebhooks(IConfigurationRoot config)
        {
            var at = new
            {
                AtMobiles = config["AT_MOBILES"]?.Split(','),
                IsAtAll = Convert.ToBoolean(config["IS_AT_ALL"] ?? "false")
            };
            switch (config["MSG_TYPE"])
            {
                case "text":
                    {
                        var data = new
                        {
                            Msgtype = "text",
                            Text = new
                            {
                                Content = config["MESSAGE"]
                            },
                            At = at
                        };
                        await SendWebhooks(config["WEBHOOK"], data);
                        break;
                    }
                case "markdown":
                    {
                        var data = new
                        {
                            Msgtype = "markdown",
                            Markdown = new
                            {
                                Title = "钉钉通知",
                                Text = config["MESSAGE"]
                            },
                            At = at
                        };
                        await SendWebhooks(config["WEBHOOK"], data);
                        break;
                    }
                default:
                    {
                        Console.WriteLine($"不支持的格式:{config["MSG_TYPE"]}");
                        break;
                    }
            }
        }

        /// <summary>
        /// 调用webhook
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="url">webhook地址</param>
        /// <param name="data">消息</param>
        /// <returns></returns>
        private static async Task SendWebhooks<T>(string url, T data) where T : class
        {
            JsonConvert.DefaultSettings = new Func<JsonSerializerSettings>(() => new JsonSerializerSettings()
            {
                NullValueHandling = NullValueHandling.Ignore,
                ContractResolver = new CamelCasePropertyNamesContractResolver()
            });
            var jsonData = JsonConvert.SerializeObject(data);
            Console.WriteLine(jsonData);
            using (var httpClient = new HttpClient())
            {
                var content = new StringContent(jsonData);
                content.Headers.ContentType = new MediaTypeHeaderValue("application/json");
                var result = await httpClient.PostAsync(url, content);
                result.EnsureSuccessStatusCode();
                Console.WriteLine($"Send webhook succeed. StatusCode:{result.StatusCode}");
            }
        }
    }
}
