package di

import (
	"go.uber.org/dig"
)

func MustProvide(container *dig.Container, constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor, opts...)
	if err != nil {
		panic(err)
	}
}

func MustInvoke(container *dig.Container, function interface{}, opts ...dig.InvokeOption) {
	err := container.Invoke(function, opts...)
	if err != nil {
		panic(err)
	}
}
