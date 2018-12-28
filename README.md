# sample-go-backoff-

## [cenkalti](https://github.com/cenkalti/backoff)

Default values for ExponentialBackOff.

|key|description|default|
|:--|:--|:--|
|InitialInterval|初期処理待機時間|500 * time.Millisecond|
|RandomizationFactor|ゆらぎ比率|0.5|
|Multiplier|乗数|1.5|
|MaxInterval|最大待機時間|60 * time.Second|
|MaxElapsedTime|最大経過時間|15 * time.Minute|

## [lestrrat-go](https://github.com/lestrrat-go/backoff)

|key|description|default|
|:--|:--|:--|
|interval|初期待機時間|500 * time.Millisecond|
|jitterFactor|ゆらぎ比率|0.5|
|maxElapsedTime|最大経過時間|ゼロ値|
|maxInterval|最大待機時間|2 * time.Minute|
|maxRetries|最大リトライ回数|10|

* 乗数(factor)はパッケージ内定数  

```golang
// lestrrat-go/backoff/exponential.go:17
factor := float64(2)
```
