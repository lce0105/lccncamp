本次作业采用istio的官方bookinfo示例完成

1.将bookinfo作为http服务发布出去
![image](https://github.com/lce0105/lccncamp/assets/163086376/e3273fd0-a5b6-4b17-82f9-47b56c51dc00)

2.发布为https服务, 并开启全链路的mTLS
![image](https://github.com/lce0105/lccncamp/assets/163086376/ecd537f7-7179-49d3-b510-a568d761ac95)
![image](https://github.com/lce0105/lccncamp/assets/163086376/0638917a-acd5-4424-9807-b3b94fce18aa)

3.灰度发布,登录用户为jason,转到reviews服务的v2版本
![image](https://github.com/lce0105/lccncamp/assets/163086376/6a68ea3f-4aaa-4ec0-932f-c85bb3951f2c)
登录用户为非jason或者不登录,转到reviews服务的v3版本
![image](https://github.com/lce0105/lccncamp/assets/163086376/3e63f0bd-018a-46f1-8ffe-12aee9e30f63)

4.接入jaeger-tracing
![322ee15fdfd0e6c46e7ee038d2d628c](https://github.com/lce0105/lccncamp/assets/163086376/497c34da-8e69-4078-a27f-5554ed314cc1)




