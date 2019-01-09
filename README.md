# sample-go-backoff

## [cenkalti](https://github.com/cenkalti/backoff)

Default values for ExponentialBackOff.

|key|description|default|
|:--|:--|:--|
|InitialInterval|ベース待機時間|500 * time.Millisecond|
|RandomizationFactor|ゆらぎ比率|0.5|
|Multiplier|乗数|1.5|
|MaxInterval|最大ベース待機時間|60 * time.Second|
|MaxElapsedTime|最大経過時間|15 * time.Minute|

* InitialInterval/MaxInterval共にその設定値で処理されるわけではなく、その値をベースとしてrandom jitterの計算がなされる。

## [lestrrat-go](https://github.com/lestrrat-go/backoff)

|key|description|default|
|:--|:--|:--|
|interval|初期待機時間|500 * time.Millisecond|
|jitterFactor|ゆらぎ比率|0.5|
|maxElapsedTime|最大経過時間|ゼロ値|
|maxInterval|最大待機時間|2 * time.Minute|
|maxRetries|最大リトライ回数|10|

* 乗数(factor)はパッケージ内定数  

```golang
// lestrrat-go/backoff/exponential.go:17
factor := float64(2)
```

## [gokit](github.com/go-kit/kit/util/conn)

backoff/jitterのdurationのみ生成するモジュールを提供。  
詳細は詳細は[こちら](https://github.com/go-kit/kit/blob/master/util/conn/manager.go#L139)を参照。


## [hashicorp/go-retryablehttp](https://github.com/hashicorp/go-retryablehttp)

<!-- now ver0.5.0... -->
