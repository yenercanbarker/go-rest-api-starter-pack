package utils

import "github.com/jinzhu/copier"

func AutoMap[S any, D any](source S, destination *D) error {
	return copier.Copy(destination, source)
}

func MapSlice[S any, D any](source []S, destination *[]D) error {
	return copier.Copy(destination, source)
}

func MapWithIgnoreEmpty[S any, D any](source S, destination *D) error {
	return copier.CopyWithOption(destination, source, copier.Option{
		IgnoreEmpty: true,
	})
}

func MapWithDeepCopy[S any, D any](source S, destination *D) error {
	return copier.CopyWithOption(destination, source, copier.Option{
		DeepCopy: true,
	})
}
