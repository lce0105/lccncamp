本次作业采用istio的官方bookinfo示例完成

1.将bookinfo作为http服务发布出去
![image](https://github.com/lce0105/lccncamp/assets/163086376/e3273fd0-a5b6-4b17-82f9-47b56c51dc00)

2.发布为https服务, 并开启全链路的mTLS
![image](https://github.com/lce0105/lccncamp/assets/163086376/ecd537f7-7179-49d3-b510-a568d761ac95)
![image](https://github.com/lce0105/lccncamp/assets/163086376/0638917a-acd5-4424-9807-b3b94fce18aa)

3.灰度发布,登录用户为jason,转到reviews服务的v2版本
![image](https://github.com/lce0105/lccncamp/assets/163086376/6a68ea3f-4aaa-4ec0-932f-c85bb3951f2c)
登录用户为非jason或者不登录,转到reviews服务的v1版本
![image](https://github.com/lce0105/lccncamp/assets/163086376/d4bec649-e4d1-46ae-8a55-cdc84c6c6328)


