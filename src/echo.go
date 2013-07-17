/**
 * Created with IntelliJ IDEA.
 * User: shinohara_wataru
 * Date: 13/07/17
 * Time: 18:44
 * To change this template use File | Settings | File Templates.
 */
package main

import (
"os"
"flag" // コマンドラインオプションのパーサー
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse() // パラメータリストを調べてflagに設定
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if !*omitNewline {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
