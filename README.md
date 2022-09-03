# go-tcp-server



## Intro

a basic multi TCP server implemented in Go.





## Code

```shell
├── README.md
├── config
│   └── config.go
├── etc
│   └── config.yaml
├── go.mod
├── go.sum
├── interface
│   └── tcp
│       └── handler.go
├── lib
│   ├── file
│   │   └── file.go
│   ├── logger
│   │   └── logger.go
│   └── sync
│       ├── atomic
│       │   └── bool.go
│       └── wait
│           └── wait.go
├── logs
│   └── tcp-server-2022-09-03.log
├── main.go
└── tcp
    ├── echo.go
    └── server.go
```

- `config`: resolves config.yaml to config.Config struct
- `etc`: config file
- `interface`: defines tcp handler interface
- `lib`: utils which offer basic functions, like file, logger, atomic.Boolean and WaitGourpWithTimeout
- `logs`: log files created by logger
- `tcp`: tcp server handles implementation
- `main.go`: tcp server enter



## Start

1. download code

   ```sh
   git clone https://github.com/hedon954/go-tcp-server.git
   ```

2. enter go-tcp-server dir

   ```sh
   cd go-tcp-server
   ```

3. download libraries

   ```shell
   go mod tidy
   ```

4. modify `etc/config.yaml`

5. run program

   ```sh
   go run main.go
   ```



## Test

| name           | view                                                         |
| -------------- | ------------------------------------------------------------ |
| tcp-client-one | <img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h5tgsv2cdfj21360u03z2.jpg" alt="image-20220903154751842" style="zoom: 33%;" /> |
| tcp-client-two | <img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h5tgt5batrj212y0u0abq.jpg" alt="image-20220903154805898" style="zoom: 33%;" /> |
| tcp-server     | <img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h5tgtgava7j221y0b0ju1.jpg" alt="image-20220903154932564" style="zoom: 20%;" /> |









