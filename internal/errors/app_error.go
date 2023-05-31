package errors

func NewSpeechBotError(message string) *SpeechBotError {
	return &SpeechBotError{
		Message: message,
	}
}

type SpeechBotError struct {
	Message string
}

func (sbe *SpeechBotError) Error() string {
	return sbe.Message
}
