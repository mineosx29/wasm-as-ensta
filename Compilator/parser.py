import sys
import astClass
from lexer import Token_Compilator

class Parser:
    def __init__(self):
        pass

    def expect(self, kind):
        next = self.showNext()
        if next.kind == kind:
            self.acceptIt()
        else:
            print("Expect Error : syntax error line " + str(self.showNext().position) )
            print("Expected : ", kind, "but have : " , next.kind)
            exit()
        return next.value

    def showNext(self):
        return self.tokens[0]

    def acceptIt(self):
        self.tokens = self.tokens[1:]
        return self.tokens

    def parse(self, tokens):
        self.tokens = tokens 
        ast2 = self.parse_program()
        return ast2



    def parse_program(self):
        self.expect("BEGIN")
        value = self.showNext().value
        self.expect("IDENTIFIER")
        self.expect("COLON")
        auto = astClass.Auto(value)
        auto.declarations = self.parse_declarations()
        auto.etats = self.parse_etats()
        print(self.showNext().value)
        self.expect("END")
        return auto

    
    
    def parse_declarations(self):
        declara = []
        while self.showNext().kind in ["INPUT", "STATE"]:
            ty = astClass.Type(self.showNext().value)
            self.acceptIt()
            self.expect("ASSIGN")
            val = self.showNext().value
            self.expect("IDENTIFIER")
            ident_id = astClass.Ident(val)
            declara.append(astClass.Variable(ident_id, ty))
            while self.showNext().kind == "SEMICOLON":
                self.acceptIt()
                value = self.showNext().value
                ident_id = astClass.Ident(value)
                self.expect("IDENTIFIER")
                declara.append(astClass.Variable(ident_id, ty))
            while self.showNext().kind == "COMMA":
                self.acceptIt()
                value = self.showNext().value
                ident_id = astClass.Ident(value)
                self.expect("IDENTIFIER")
                declara.append(astClass.Variable(ident_id, ty))
        return declara
    
           
    
    def parse_etats(self):
        etatss = []
        while self.showNext().kind == "CASE":
            self.acceptIt()
            print(self.showNext().kind)
            value = self.showNext().value
            ident_id = astClass.Ident(value)
            self.expect("IDENTIFIER")
            etats = astClass.Etats(ident_id)
            self.expect("COLON")
            etats.condition = self.parse_condition()
            etatss.append(etats)
            #print(etats)
        return etatss

    def parse_condition(self):
        conditions=[]
        self.expect("BEGIN")
        self.expect("COLON")
        while self.showNext().kind == "IF":
            conditions.append(self.parse_if())
        self.expect("END")
        self.expect("CASE")
        return conditions

    def parse_if(self):
        self.expect("IF")
        self.expect("LPAREN")
        cond = self.parse_expression()
        self.expect("RPAREN")
        self.expect("THEN")
        assignment = self.parse_expression_states()
        return astClass.IF(cond, assignment)

    def parse_expression(self):
        lhs = self.parse_pr()
        op = None
        rhs = None
        if self.showNext().kind in ["OR", "AND"]:
            op=self.showNext().value
            self.acceptIt()
            rhs=self.parse_pr()
        return astClass.Binary(lhs,op,rhs)

    def parse_pr(self):
        if self.showNext().kind == ("IDENTIFIER"):
            valeur = self.showNext().value
            self.acceptIt()
            return astClass.Ident(valeur)
        elif self.showNext().kind == ("LPAREN"):
            self.acceptIt()
            expression_cond = self.parse_expression()
            self.expect("RPAREN")
            return astClass.Expression(expression_cond)
        if self.showNext().kind == "NOT":
            self.acceptIt()
            self.expect("LPAREN")
            expression_cond_not = self.parse_expression()
            self.expect("RPAREN")
            return astClass.Expression_in_Not(expression_cond_not)

        
    def parse_expression_states(self):
        state = astClass.Assignement_Etats()
        self.expect("NEXTST")
        self.expect("ASSIGN")
        state.part_right = astClass.Ident(self.showNext().value)
        self.expect("IDENTIFIER")
        return state


    












