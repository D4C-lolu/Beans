package parser

import (
	"Beans/beans/pkg/lexer"
	"testing"
)


func TestLetStatements (t *testing.T){
	input:=`
	let x=5;
	let y=10;
	let foobar =8383838;
	`


	l:lexer.New(input)
}