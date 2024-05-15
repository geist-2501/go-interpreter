package token

type TokenType string

type Token struct {
    TokenType TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers
    IDENT = "IDENT"
    INT = "INT"

    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
)
