# hello-convert

This is my first CLI tool written in GoLang. It's a simple utility that will output the phrase 'Hello' in most common languages.

## Screenshots
The '-lang=[language]' flag is used to output a greeting in the given language
```
$ hello-convert -lang=jap
こんにちは
```
The 'list' subcommand will display available languages
```
$ hello-convert list
Chinese
Spanish
English
Hindi
Bengali
Portugese
Russian
Japanese
Punjabi
Marathi
< Page 1 of 7 >
```
there are two flags for the 'list' subcommand: <br />
'list -all' lists all available languages in a single page <br />
'list -page=[pageNum]' allows you to navigate to different pages of the list <br />
