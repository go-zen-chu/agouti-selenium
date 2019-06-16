# agouti selenium

running simple selenium in golang.

## install

```
brew cask install chromedriver
brew install selenium-server-standalone
```

or may be phantomjs driver is suitable for CI/CD

```
brew install phantomjs
```

## run

Define environment variable.

```
source smaple.envrc
```

then run,

```
go run main.go
```

## package format 

[GitHub - golang-standards/project-layout: Standard Go Project Layout](https://github.com/golang-standards/project-layout)


## known issue

- when navigating invalid session id while using ChromeDriver 75.0.3770.90 

```
failed to navigate: request unsuccessful: invalid session id
```