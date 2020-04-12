# 性能结果

本地测试，结果仅供参考

# segment
step 为2000
### 场景1
单key， 10万次请求
```
goos: darwin
goarch: amd64
pkg: github.com/jerryzhou343/leaf-go/test
BenchmarkSegment
time cost [25.071975441]BenchmarkSegment-4   	       1	25072347845 ns/op
PASa
```
### 场景2
双key，每个key10万次请求，两个客户端同时请求
```
./client segment  --key leaf-segment-test1
time cost [48.667793728]

./client segment  --key leaf-segment-test
time cost [48.626422017]

```

### 场景3
三key， 每个key 10 万次请求， 三个客户端同时请求
```



```



# snowflake 
单链接10万次请求
```
goos: darwin
goarch: amd64
pkg: github.com/jerryzhou343/leaf-go/test
BenchmarkSnowflake
time cost [26.694874262]BenchmarkSnowflake-4   	       1	26695240963 ns/op
PASS
```

