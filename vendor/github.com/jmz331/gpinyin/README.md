HELLO WORLD

实现一个golang版本的汉字转拼音类

支持功能：
 + 》简繁体转换
 + 》带音标与不带音标的转换

使用方式：
go get github.com/jmz331/gpinyin
```go
import "github.com/jmz331/gpinyin"

const s = "台我要1234!#$翻译成繁体的汉字堡垒asdf"
r1 := ConvertToPinyinString(s, "-", PINYIN_WITHOUT_TONE)
//tai-wo-yao-1234!#$-fan-yi-cheng-fan-ti-de-han-zi-bao-lei-asdf
r2 := ConvertToPinyinString(s, "-", PINYIN_WITH_TONE_MARK)
//tái-wǒ-yào-1234!#$-fān-yì-chéng-fán-tǐ-de-hàn-zì-bǎo-lěi-asdf
r3 := ConvertToTraditionalChinese(s)
//臺我要1234!#$翻譯成繁體的漢字堡壘asdf

//特殊情况，无法转换的字符，比如日本平假名，空格之类会替换成空字符串
const s2 = "日本語にほんご1234!#$翻译成繁体的汉字堡垒 asdf"
r4 := ConvertToPinyinString(s2, "-", PINYIN_WITHOUT_TONE)
//ri-ben-yu-1234!#$-fan-yi-cheng-fan-ti-de-han-zi-bao-lei-asdf
```

###参考
https://github.com/stuxuhai/jpinyin
