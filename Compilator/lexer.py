import re
import sys

reg_expression = [
    (r'[ \n\t]+', None),
    (r'#[^\n]*', None),
    (r'case\b', 'CASE'),
    (r'if\b', 'IF'),
    (r'begin\b', 'BEGIN'),
    (r'state\b', 'STATE'),
    (r'not\b', 'NOT'),
    (r'next_state\b', 'NEXTST'),
    (r'input\b', 'INPUT'),
    (r'state\b', 'STATE'),
    (r'then\b', 'THEN'),
    (r'end\b', 'END'),
    (r'\(', 'LPAREN'),
    (r'\)', 'RPAREN'),
    (r'\{', 'LBRACE'),
    (r'\}', 'RBRACE'),
    (r'\,', 'COMMA'),
    (r'\:', 'COLON'),
    (r'\;', 'SEMICOLON'),
    (r'\=', 'ASSIGN'),
    (r'and', 'AND'),
    (r'or', 'OR'),
    (r'[a-zA-Z][a-zA-Z0-9]*', 'IDENTIFIER')
    
]


class Token_Compilator:
    def __init__ (self, kind, value, position):
        self.kind = kind
        self.value = value
        self.position = position

class Lexer_Compilator:
    def __init__(self):
        self.tokens = []
        self.file = None

    def lexer(self, file):
        print("Lexing program.....")
        self.file = open(file).readlines()
        lineNumber = 0
        for line in self.file:
            lineNumber += 1
            position = 0
            while position < len(line):
                match = None
                for tokenRegex in reg_expression:
                    pattern, tag = tokenRegex
                    regex = re.compile(pattern)
                    match = regex.match(line, position)
                    if match:
                        data = match.group(0)
                        if tag:
                            token = Token_Compilator(tag, data, [lineNumber, position])
                            self.tokens.append(token)
                        break
                if not match:
                    print(self.file[position])
                    print("no match")
                    sys.exit(1)
                else:
                    position = match.end(0)
        print("Lexer part is success ! ")
        return self.tokens
