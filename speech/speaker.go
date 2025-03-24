package speech

import (
	"context"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

var speaker *Speaker

const (
	// ref https://cloud.google.com/text-to-speech/docs/voices
	VoiceStandardA = "ja-JP-Standard-A"
	VoiceStandardB = "ja-JP-Standard-B"
	VoiceStandardC = "ja-JP-Standard-C"
	VoiceStandardD = "ja-JP-Standard-D"
	VoiceWavenetA  = "ja-JP-Wavenet-A"
	VoiceWavenetB  = "ja-JP-Wavenet-B"
	VoiceWavenetC  = "ja-JP-Wavenet-C"
	VoiceWavenetD  = "ja-JP-Wavenet-D"

	AudioEncoding_LINEAR16 = texttospeechpb.AudioEncoding_LINEAR16
	AudioEncoding_MP3      = texttospeechpb.AudioEncoding_MP3
	AudioEncoding_OGG_OPUS = texttospeechpb.AudioEncoding_OGG_OPUS
)

type SpeechOption struct {
	LanguageCode      string
	VoiceName         string
	AudioEncoding     texttospeechpb.AudioEncoding
	AudioSpeakingRate float64
	AudioPitch        float64
}

type AudioEncoding texttospeechpb.AudioEncoding

// Speaker ...
type Speaker struct {
	client *texttospeech.Client
}

// NewSpeechClient return a Speaker.
func NewSpeechClient(ctx context.Context) (*Speaker, error) {
	if speaker != nil {
		return speaker, nil
	}

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	speaker = &Speaker{client: client}
	return speaker, nil
}

// NewRequest generate perform the text-to-speech request on the text input
// with the selected voice parameters and audio file type.
func NewRequest(text string, opt *SpeechOption) *texttospeechpb.SynthesizeSpeechRequest {
	return &texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: opt.LanguageCode,
			Name:         opt.VoiceName,
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: opt.AudioEncoding,
			SpeakingRate:  opt.AudioSpeakingRate,
			Pitch:         opt.AudioPitch,
		},
	}
}

// Run ...
func (s *Speaker) Run(ctx context.Context, req *texttospeechpb.SynthesizeSpeechRequest) ([]byte, error) {
	resp, err := s.client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.AudioContent, nil
}
