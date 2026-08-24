package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type (
	ObjectType string

	Object interface {
		Inspect() string
		Type() ObjectType
	}

	Boolean struct {
		Value bool
	}

	Error struct {
		Message string
	}

	Function struct {
		Body       *ast.BlockStatement
		Env        *Environment
		Parameters []*ast.Identifier
	}

	Integer struct {
		Value int64
	}

	Null struct{}

	ReturnValue struct {
		Value Object
	}
)

const (
	BOOLEAN_OBJ      = "BOOLEAN"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	INTEGER_OBJ      = "INTEGER"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}
func (i *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Null) Inspect() string {
	return "null"
}
func (i *Null) Type() ObjectType {
	return NULL_OBJ
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
