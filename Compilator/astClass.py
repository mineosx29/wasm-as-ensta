from ast import AST


class ASTNode:
    def accept(self, visitor):
        method_name=getattr(visitor,"visit"+self.__class__.__name__)
        method_name(self)
        
class Auto(ASTNode):
    def __init__(self, name):
        self.name = name
        self.declarations = []
        self.etats=[]

class Variable(ASTNode):
    def __init__(self, ident , type):
        self.ident = ident
        self.type = type

class Type(ASTNode):
    def __init__(self, tok):
        self.tok = tok

class Ident(ASTNode):
    def __init__(self, tok):
        self.tok = tok

class Etats(ASTNode):
    def __init__(self, ident, condition = []):
        self.condition=condition
        self.ident = ident

class IF(ASTNode):
    def __init__(self, condition, assign):
        self.condition = condition
        self.assign = assign

class Binary(ASTNode):
    def __init__(self, lhs = None, op = None, rhs = None):
        self.lhs = lhs
        self.op = op
        self.rhs = rhs

class Declaration(ASTNode):
    def __init__(self, type = None, name = None):
        self.type = type
        self.name = name

class Expression(ASTNode):
    def __init__(self, expression):
        self.expression = expression
        
class Expression_in_Not(ASTNode):
    def __init__(self, expression):
        self.expression = expression

class Assignement_Etats(ASTNode):
    def __init__(self, part_right = None):
        self.part_right = part_right