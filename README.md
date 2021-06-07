# player-stats 简介
player stats（球员统计）应用程序，用于在《Helm 入门与实战》中演示开发 helm chart。可以在我的微信公众号（**sretech**）中查看相关内容。

# player-stats 架构
![image](https://user-images.githubusercontent.com/14903623/120955782-c20a5680-c784-11eb-9393-a6f448b3dbba.png)

player stats 是一个简单的应用程序，使用 Go 做的 Web 开发，MongoDB 作为后端存储数据库。用户输入的数据会被存储到 MongoDB 中，并展示到页面。
