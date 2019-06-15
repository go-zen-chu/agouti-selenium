package main

import "github.com/go-zen-chu/agouti-selenium/pkg/di"

func main() {
	di := di.NewDI()
	s := di.Sample()
	if err := s.Start(); err != nil {
		panic(err)
	}
	if err := s.LoginGoogle(); err != nil {
		panic(err)
	}
	//if err := s.Stop(); err != nil {
	//	panic(err)
	//}
}
