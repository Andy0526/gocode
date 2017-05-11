package eval

import (
	"fmt"
	"strings"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next() { lex.token = lex.scan.Scan() }

func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token))
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func parsePrimary(lex *lexer) Expr{
    switch lex.token {
    case scanner.Ident:
        id:=lex.text()
        lex.next()
        if lex.token!='('{
            for {
                args=append(args,parseExpr(lex))
                if lex.token!=','{
                    break
                }
                lex.next()
            }
            if lex.token !=')'{
                msg:=fmt.Sprintf("got %q want ')'", lex.token)
                panic(lexPanic(msg))
            }
        }
        lex.next()
        return call{id,args}
    }
    case scanner.Int,scanner.Float:
}

func parseBinary(lex *lexer,prec1 int) Expr{
    lhs:=ParseUnary(lex)
}

func parseExpr(lex *lexer) Expr{ return }

func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
		case lexPanic:
			err := fmt.Errorf("%s", x)
		default:
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
    e:=
}
