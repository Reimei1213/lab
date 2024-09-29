# Linux のしくみ
[Amazon](https://www.amazon.co.jp/dp/429713148X)

![img.png](https://m.media-amazon.com/images/I/51CKe00bsfL._SX342_SY445_.jpg)

# Usage

環境は Mac OS を想定しています。
1 行目のコマンドで docker image を作成し、2 行目のコマンドでコンテナ内に入ります。

https://github.com/Msksgm/linux-in-practice-2nd-docker

```bash
docker image build -t linux-in-practice-2nd-docker .
docker run --rm -it linux-in-practice-2nd-docker /bin/bash
```