# mass2dir
别问 问就是不知道用来干啥的

用法:

```
▶ cat mass.out | grep -v 443 | grep -v 80 | grep open | mass2dir | httpx | tee https_mass.txt
```

安装:

```
▶ go get -u -v github.com/flag007/mass2dir
```
