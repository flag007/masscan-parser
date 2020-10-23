# masscan-parser
别问 问就是不知道用来干啥的

用法:

```
▶ cat masscan.out | grep open | masscan-parser | httpx | tee https_masscan.txt
```


```
去除， 80， 443端口
▶ cat masscan.out | grep -v 443 | grep -v 80 | grep open | masscan-parser | httpx | tee https_masscan.txt
```

安装:

```
▶ go get -u -v github.com/flag007/masscan-parser
```
