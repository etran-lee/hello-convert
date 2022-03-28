package languages

import (
	"strings"
)

type LanguageOutput struct {
	Language string
	Output   string
}

func KeyCheck(key string) LanguageOutput {
	for i := 0; i < 10; i++ {
		if caseInsensitiveHasPrefix(LanguageList[i].Language, key) {
			return LanguageList[i]
		}
	}
	return LanguageOutput{Output: "No language found"}
}

func caseInsensitiveHasPrefix(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.HasPrefix(s, substr)
}

// Language Data
var LanguageList = []LanguageOutput{
	{Language: "Chinese", Output: "你好"},
	{Language: "Spanish", Output: "Hola"},
	{Language: "English", Output: "Hello"},
	{Language: "Hindi", Output: "नमस्ते"},
	{Language: "Bengali", Output: "হ্যালো"},
	{Language: "Portugese", Output: "Olá"},
	{Language: "Russian", Output: "Привет"},
	{Language: "Japanese", Output: "こんにちは"},
	{Language: "Punjabi", Output: "ਸਤ ਸ੍ਰੀ ਅਕਾਲ"},
	{Language: "Marathi", Output: "नमस्कार"},
	{Language: "Telugu", Output: "హలో"},
	{Language: "Turkish", Output: "Merhaba"},
	{Language: "Korean", Output: "안녕하세요"},
	{Language: "French", Output: "Bonjour"},
	{Language: "German", Output: "Hallo"},
	{Language: "Vietnamese", Output: "Xin chào"},
	{Language: "Tamil", Output: "வணக்கம்"},
	{Language: "Urdu", Output: "ہیلو"},
	{Language: "Javanese", Output: "Halo"},
	{Language: "Italian", Output: "Ciao"},
	{Language: "Egyptian", Output: "ألسّلام عليكم"},
	{Language: "Gujarati", Output: "નમસ્તે"},
	{Language: "Persian", Output: "سلام"},
	{Language: "Bhojpuri", Output: "प्रणाम"},
	{Language: "Hausa", Output: "Sannu"},
	{Language: "Kannada", Output: "ನಮಸ್ಕಾರ"},
	{Language: "Indonesian", Output: "Halo"},
	{Language: "Polish", Output: "Cześć"},
	{Language: "Yoruba", Output: "Pẹlẹ o"},
	{Language: "Malayalam", Output: "ഹലോ"},
	{Language: "Odia", Output: "Namaskar"},
	{Language: "Maithili", Output: "प्रनाम"},
	{Language: "Burmese", Output: "မင်္ဂလာပါ"},
	{Language: "Sunda", Output: "Halo"},
	{Language: "Arabic", Output: "أهلا"},
	{Language: "Ukrainian", Output: "Здравствуйте"},
	{Language: "Igbo", Output: "Nnọọ"},
	{Language: "Uzbek", Output: "Salom"},
	{Language: "Sindhi", Output: "هيلو"},
	{Language: "Romanian", Output: "Buna ziua"},
	{Language: "Tagalog", Output: "Kamusta"},
	{Language: "Dutch", Output: "Hallo"},
	{Language: "Amharic", Output: "እው ሰላም ነው"},
	{Language: "Pashto", Output: "سلام"},
	{Language: "Magahi", Output: "मगही"},
	{Language: "Thai", Output: "สวัสดี"},
	{Language: "Khmer", Output: "ជំរាបសួរ"},
	{Language: "Somali", Output: "Salaam alaykum"},
	{Language: "Cebuano", Output: "Hello"},
	{Language: "Nepali", Output: "नमस्ते"},
	{Language: "Assamese", Output: "নমস্কাৰ"},
	{Language: "Sinhalese", Output: "ආයුබෝවන්"},
	{Language: "Kurdish", Output: "Slav"},
	{Language: "Bavarian", Output: "Seavus"},
	{Language: "Greek", Output: "Χαίρετε"},
	{Language: "Kazakh", Output: "Сәлеметсіз бе"},
	{Language: "Hungarian", Output: "Helló"},
	{Language: "Kinyarwanda", Output: "Muraho"},
	{Language: "Zulu", Output: "Sawubona"},
	{Language: "Kirundi", Output: "Amahoro"},
	{Language: "Czech", Output: "Ahoj"},
	{Language: "Uyghur", Output: "Ässalamu läykum"},
	{Language: "Sylheti", Output: "Assalamu alaikum"},
}
