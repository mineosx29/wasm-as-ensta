# Main du compilateur
#Execution : python3 main.py file.aut file.dot

from lexer import Lexer_Compilator
import sys
import parser
import generator
import generatorGo


class Compile:

    def compile(self):
        fileGo = input(str("Entrez le nom du fichier go que vous voulez Générer : \n"))
        file = sys.argv[1]
        print("Lexer Module Starting ....")
        lexer_file = Lexer_Compilator()
        tok = lexer_file.lexer(file)
        print("Parser Module Starting...")
        parser2 = parser.Parser()
        ast = parser2.parse(tok)
        print("Generating DOT Code....")
        generatorCode = generator.GeneratorDot()
        gene = generatorCode.visitAuto(ast, "code.gv" )
        generatorCodeGo = generatorGo.GeneratorGo()
        gen = generatorCodeGo.visitAuto(ast, fileGo)
        print("Process End")
        

compile = Compile()
compile.compile()
