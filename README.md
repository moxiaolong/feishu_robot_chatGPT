融合了以下项目：

[飞书机器人](https://github.com/waro163/feishu_robot.git)

[chatgpt-dingtalk](https://github.com/eryajf/chatgpt-dingtalk)

配置：

APP_ID: "飞书APP ID"

APP_SECRET: "飞书App Secret"

APP_VERIFICATION_TOKEN: "飞书事件订阅Verification Token"

GTP_API_KEY: "OPEN AI api-keys"

使用：

启动后将外网地址（如http://domain:8081/api/event/call_back）配置到飞书应用事件订阅地址：

订阅事件：接收消息

权限添加：im:message:send_as_bot

发布应用后在群聊中添加机器人。