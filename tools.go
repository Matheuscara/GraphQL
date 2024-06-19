//go:build tools

package tools

// Cria um arquivo de ferramentas para o Go,
// que nao iram ser utilizadas na hora do build para binario

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
