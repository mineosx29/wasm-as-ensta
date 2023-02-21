# Main du compilateur

from lexer import Lexer_Compilator
import sys
import parser
import generator
import generatorGo
import argparse


class Compile:

    def compile(self, inputFile = None, outputFile = None, outputFileDot = None):
        if outputFile != "fsm.go":
            print("Error, the file must be called fsm.go")
            exit()
        print("Lexer Module Starting ....")
        lexer_file = Lexer_Compilator()
        tok = lexer_file.lexer(inputFile)
        print("Parser Module Starting...")
        parser2 = parser.Parser()
        ast = parser2.parse(tok)
        print("Generating DOT Code....")
        generatorCode = generator.GeneratorDot()
        gene = generatorCode.visitAuto(ast, outputFileDot)
        generatorCodeGo = generatorGo.GeneratorGo()
        gen = generatorCodeGo.visitAuto(ast, outputFile)
        print("Process End")

if __name__ == "__main__":
    parse = argparse.ArgumentParser(description="This Compilator compiles a program and Generates DOT FSM and a FSM in GO Langages", epilog="This Compilator realized in AS Courses")
    parse.add_argument("-i","--inputFile", help="input Auto Code", required=True)
    parse.add_argument("-oDot","--outputFileDot", help="gv DOT Code output File", required=True)
    parse.add_argument("-oGO","--outputFileGo", help="Go Code output File", required=True)

    args = parse.parse_args()


    compile = Compile()
    compile.compile(args.inputFile, args.outputFileGo, args.outputFileDot)
