package main

import (
	"flag"
	"fmt"
	"passgengo/internal/colors"
	gen "passgengo/internal/generator"
	"passgengo/internal/hashing"
	"passgengo/internal/lib"
	"passgengo/internal/logging"
	"passgengo/internal/utils"
	"runtime"
)

var (
	length         int
	hashingMethod  int
	noUpperChars   bool
	noSpecialChars bool
)

func main() {
	colors.CheckOS()

	bannerColor := ""
	if runtime.GOOS != "windows" {
		randInt, err := utils.RandomIntWithMax(len(colors.BannerColors))
		if err != nil {
			panic(err)
		}
		bannerColor = colors.BannerColors[randInt.Int64()]
	}
	fmt.Println(bannerColor + lib.Banner + colors.Reset)

	parseFlags()

	generator, err := buildGenerator()
	if err != nil {
		printError(err)
		return
	}

	logging.Info("Generating the password based on this generator:")
	generatorJson, err := generator.Json()
	if err != nil {
		logging.Fatal(err)
	}
	fmt.Println(generatorJson)

	password, err := generator.Generate()
	if err != nil {
		logging.Fatal(err)
	}

	logging.Success(fmt.Sprintf("Generated password: %s", password))
}

func parseFlags() {
	flag.IntVar(
		&length,
		"l",
		12,
		"The password length. Must be between 6 and 128 characters",
	)

	// Hashing
	hashingMethods := "\n"
	for i, method := range hashing.HashMethods {
		hashingMethods += fmt.Sprintf("\t%d -> %s", i, method)
		if i != len(hashing.HashMethods)-1 {
			hashingMethods += "\n"
		}
	}
	flag.IntVar(
		&hashingMethod,
		"hm",
		int(hashing.None),
		fmt.Sprintf("Hashing method to use. Available methods: (default 0)%s", hashingMethods),
	)

	// Char selection
	flag.BoolVar(
		&noUpperChars,
		"noupper",
		false,
		"Remove upper-case characters",
	)

	flag.BoolVar(
		&noSpecialChars,
		"nospecial",
		false,
		"Remove special characters",
	)

	flag.Parse()
}

func buildGenerator() (generator gen.Generator, err error) {
	builder := gen.NewGeneratorBuilder()

	if err = builder.Length(length); err != nil {
		return
	}

	if err = builder.HashingMethod(hashingMethod); err != nil {
		return
	}
	// TODO: Write clear text password eventhough it went through hashing ....

	if noUpperChars {
		builder.NoUpperChars()
	}

	if noSpecialChars {
		builder.NoSpecialChars()
	}

	generator = builder.Build()
	return
}

func printError(err error) {
	logging.Error(err)
	fmt.Println()
	flag.Usage()
}
