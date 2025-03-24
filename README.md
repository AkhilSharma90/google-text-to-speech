# Google Text to speech

so, i think many of you know that my sister stays in kawasaki, japan, i've found google translate to be wrong on many occasions that's because it will translate from colloquial speech to text. a more accurate way of getting around japan is to write english text and convert it to japanese in a very japanese-accent voice (japanese people will not understand japanese that's spoken in a non-japanese accent) and carry some recordings of very commonly used sentences with you so you can play these to get your point across.

So i built this simple program just to scratch my itch.

in the readme file, i've given it text "hi how are you doing, my name is akhil.", feel free to go to google translate, convert your english to japanese and then bring the japanese characters here for a more accurate translation into a voice that has a strong japanese accent.

command line tool that synthesizes speech from text using Google Cloud Text-to-Speech.
Currently configured for japanese
 
### set up

```
export GOOGLE_APPLICATION_CREDENTIALS="$PWD/service-account-key.json"
```

### Usage
 
```
go run main.go -text "こんにちは、お元気ですか。私の名前はアキルです。ソフトウェアエンジニアです。" -o test2.mp3
```

```
Usage of spokesman:
  -o string
        output audio file (support format of the audio: LINEAR16, MP3)
  -pitch float
        speaking pitch (-20.0 ~ 20.0)
  -rate float
        speech rate (0.25 ~ 4.0) (default 1)
  -text string
        text to speech
  -voice string
        speaker's voice name (default "stand-a")

```

##### speaker's voice name

| speaker's voice name                                                         | Description                                                                                                                                         |
|:-------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------|
| [stand-a](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-A, SSML Gender: FEMALE                                                                                                                                         |
| [stand-b](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-B, SSML Gender: FEMALE                                                                                                                                         |
| [stand-c](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-C, SSML Gender: MALE                                                                           |
| [stand-d](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-D, SSML Gender: MALE                                                                           |
| [wave-a](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-A, SSML Gender: FEMALE                                                                                                                           |
| [wave-b](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-B, SSML Gender: FEMALE                                                                             |
| [wave-c](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-C, SSML Gender: MALE                                                                             |
| [wave-d](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-D, SSML Gender: MALE                                                                              |
