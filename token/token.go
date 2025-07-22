package token

type TokenType string

//type TokenType basically creates a new type which is a string but is treated
//as a distinct type.eg:Categorization of keywords,operators

type Token struct {
	Type    TokenType //type tells what is the type of keyword or operator {like LET,+,-}
	Literal string    //This is the literal value of the string {x,5,=} etc
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers+literals
	IDENT = "IDENT" //add,foobar,x,y
	INT   = "INT"

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	//more Operators
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"

	LT       = "<"
	GT       = ">"
	EQ       = "=="
	N_EQ     = "!="
	STRING   = "STRING"
	LBRACKET = "["
	RBRACKET = "]"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}

//tok will hold the TokenType if ident is found in the map.
//ok is a boolean that indicates whether the key ident exists in the map.
