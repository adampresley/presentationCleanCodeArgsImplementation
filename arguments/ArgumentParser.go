package arguments

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type ArgumentParser struct {
	arguments  []string
	marshalers map[string]ArgumentMarshaler
}

func NewArgumentParser(schema string, arguments []string) (*ArgumentParser, error) {
	var err error

	result := &ArgumentParser{
		arguments:  arguments,
		marshalers: make(map[string]ArgumentMarshaler),
	}

	if err = result.parseSchema(schema); err != nil {
		return result, err
	}

	return result, nil
}

func (argumentParser *ArgumentParser) parseSchema(schema string) error {
	var err error
	var argumentValue string

	schemaElements := strings.Split(schema, ",")

	for index, element := range schemaElements {
		if index >= len(argumentParser.arguments) {
			return ArgumentCountMismatchException()
		}

		argumentValue = argumentParser.arguments[index]
		if len(element) > 0 {
			if err = argumentParser.parseSchemaElement(strings.TrimSpace(element), argumentValue); err != nil {
				return err
			}

		}
	}

	return nil
}

func (argumentParser *ArgumentParser) parseSchemaElement(element string, argumentValue string) error {
	var err error
	var argumentID string
	var argumentType string

	argumentID = string(element[0])

	if len(element) > 1 {
		argumentType = string(element[1:])
	}

	if err = argumentParser.validateSchemaArgumentID(argumentID); err != nil {
		return err
	}

	if len(argumentType) == 0 {
		argumentParser.marshalers[argumentID] = NewBooleanArgumentMarshaler(argumentID, argumentValue)
		return nil
	}

	if argumentType == "*" {
		argumentParser.marshalers[argumentID] = NewStringArgumentMarshaler(argumentID, argumentValue)
		return nil
	}

	if argumentType == "#" {
		argumentParser.marshalers[argumentID] = NewIntegerArgumentMarshaler(argumentID, argumentValue)
		return nil
	}

	if argumentType == "##" {
		argumentParser.marshalers[argumentID] = NewFloatArgumentMarshaler(argumentID, argumentValue)
		return nil
	}

	return NewArgumentException(INVALID_ARGUMENT_FORMAT, argumentID, argumentType)
}

func (argumentParser *ArgumentParser) validateSchemaArgumentID(argumentID string) error {
	var ch rune

	ch, _ = utf8.DecodeRuneInString(argumentID)

	if !unicode.IsLetter(ch) {
		return InvalidArgumentNameException(argumentID)
	}

	return nil
}

func (argumentParser *ArgumentParser) Has(arg string) bool {
	if _, ok := argumentParser.marshalers[arg]; ok {
		return true
	}

	return false
}

func (argumentParser *ArgumentParser) GetBoolean(arg string) (bool, error) {
	var err error
	marshaler := argumentParser.marshalers[arg]

	if err = marshaler.Marshal(); err != nil {
		return false, err
	}

	return marshaler.Value().(bool), nil
}

func (argumentParser *ArgumentParser) GetString(arg string) (string, error) {
	var err error
	marshaler := argumentParser.marshalers[arg]

	if err = marshaler.Marshal(); err != nil {
		return "", err
	}

	return marshaler.Value().(string), nil
}

func (argumentParser *ArgumentParser) GetInteger(arg string) (int, error) {
	var err error
	marshaler := argumentParser.marshalers[arg]

	if err = marshaler.Marshal(); err != nil {
		return 0, err
	}

	return marshaler.Value().(int), nil
}

func (argumentParser *ArgumentParser) GetFloat(arg string) (float64, error) {
	var err error
	marshaler := argumentParser.marshalers[arg]

	if err = marshaler.Marshal(); err != nil {
		return 0.0, err
	}

	return marshaler.Value().(float64), nil
}
