package main

import (
	"errors"
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
	hexString      bool
	base64String   bool
	hashingMethod  int
	noUpperChars   bool
	noSpecialChars bool
	wizard         bool
)

type GeneratorBuildFunc func() (gen.Generator, error)

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

	var generatorBuildFunc GeneratorBuildFunc
	if wizard {
		generatorBuildFunc = gen.GetGeneratorFromWizard
	} else {
		generatorBuildFunc = buildGenerator
	}

	generator, err := generatorBuildFunc()
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

	flag.BoolVar(
		&hexString,
		"hex",
		false,
		"Hexadecimal encoded string. The length represents the number of bytes",
	)

	flag.BoolVar(
		&base64String,
		"b64",
		false,
		"Standard base64 encoded string. The length represents the number of bytes",
	)

	// Hashing
	hashingMethods := hashing.GetHashingMethodsString()
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

	flag.BoolVar(
		&wizard,
		"wizard",
		false,
		"Generate a password using the wizard",
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

	if hexString && base64String {
		err = errors.New("cannot use hex string and base64 string in the same command")
		return
	}

	if hexString {
		builder.HexString()
	}

	if base64String {
		builder.Base64String()
	}

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
