package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config - configuration of algorithn
type Config struct {
	// NumOfIteration - maximal number of generations
	NumOfIterations uint `json:"NumOfIterations"`

	// SizeOfGenerations - number of speciments in generations
	SizeOfGeneration uint `json:"SizeOfGeneration"`

	// FromConfig - if set, config should be read from file
	FromFile string

	// PathToImage - path to original image
	PathToImage string `json:"PathToImage"`

	// NumOfBest - number of best speciments used to cross
	NumOfBest uint `json:"NumOfBest"`

	// MutationChance - chance of mutation
	MutationChance float64 `json:"MutationChance"`

	// GrayScale - colors scale
	GrayScale bool `json:"GrayScale"`

	// WithCrossing - specifies if crossing operator should be used
	WithCrossing bool `json:"WithCrossing"`
}

// parseFlags - parses arguments passed to program in console
func (c *Config) parseFlags() {
	flag.StringVar(&c.FromFile, "from-file", "", "read config from given file")
	flag.UintVar(&c.NumOfIterations, "generations", 1000, "maximal number of generations")
	flag.UintVar(&c.SizeOfGeneration, "generation-size", 200, "size of generation")
	flag.UintVar(&c.NumOfBest, "best", 10, "number of best specimens")
	flag.StringVar(&c.PathToImage, "image-dir", "./mona.jpg", "path to original image")
	flag.BoolVar(&c.GrayScale, "gray-scale", false, "defines if images should be represented in gray scale")
	flag.BoolVar(&c.WithCrossing, "with-crossing", false, "specifies if crossing operator should be used")
	flag.Float64Var(&c.MutationChance, "mutation-chance", 0.2, "chance of mutation")
}

// Init - initializes config
func (c *Config) Init(fromFile bool) {
	if fromFile {
		fmt.Println("Read config from file...")
	}

	c.parseFlags()
	flag.Parse()

	// read from file if path is not empty
	if c.FromFile != "" {
		c.readFromFile()
	}
}

// readFromFile - reads config from JSON file
func (c *Config) readFromFile() {
	jsonFile, err := os.Open(c.FromFile)

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	bytesValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytesValue, &c)
}
