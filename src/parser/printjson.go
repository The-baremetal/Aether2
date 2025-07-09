package parser

import (
  "encoding/json"
)

func MarshalAST(ast interface{}) ([]byte, error) {
  return json.MarshalIndent(ast, "", "  ")
} 