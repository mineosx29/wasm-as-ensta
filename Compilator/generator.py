class GeneratorDot:
    def __init__(self):
        self.ecriture_etat=0
        self.etat_condition=0
        self.assignmement_etat=0
        self.etat_actuel_backup=0
        self.etat_initial=1

    def visitAuto(self,auto,file):
        self.file=open(file,"w")
        self.file.write("digraph "+ auto.name+"{\n")
        self.file.write("rankdir=LR;\n")
        self.file.write("size=\"8,5\"\n")
        for declaration in auto.declarations:
            print(declaration)
            declaration.accept(self)
        for etat in auto.etats:
            etat.accept(self)
        self.file.write("}")
        self.file.close()
        print("generator: the DOT File is generated")
    
    def visitVariable(self,variable):
        variable.type.accept(self)
        variable.ident.accept(self)

    def visitIdent(self,ident):
        print("IDENT")
        if self.ecriture_etat == 1:
            if self.etat_initial:
                self.file.write("node [shape = doublecircle];"+ ident.tok+";\n")
                self.file.write("node [shape = point ]; qi\n")
                self.file.write("node [shape = circle];\n");
                self.file.write("qi -> "+ident.tok+";\n");
                self.etat_initial=0
            else:
                self.file.write("node [shape = circle];"+ ident.tok+";\n")
            self.ecriture_etat=0
        if self.etat_condition == 1:
            self.file.write(ident.tok)
        if self.assignmement_etat == 1:
            self.file.write(ident.tok)
            self.assignmement_etat=0
        if self.etat_actuel_backup == 1:
            self.actual_state=ident.tok
            self.etat_actuel_backup=0

    def visitType(self,type):
        if type.tok=="state":
            self.ecriture_etat=1

    def visitEtats(self,etats):
        self.etat_actuel_backup=1
        etats.ident.accept(self)
        for condition in etats.condition:
            condition.accept(self)

    def visitIF(self,if_):
        print("IF")
        self.file.write(self.actual_state)
        self.assignmement_etat=1
        self.file.write(" -> ")
        if_.assign.accept(self)  
        self.etat_condition=1
        self.file.write(" [ label = \"")
        if_.condition.accept(self)
        self.etat_condition=0
        self.file.write("\"];\n")

    def visitBinary(self,binary):
        binary.lhs.accept(self)
        if binary.op and binary.rhs:
            self.file.write(' '+binary.op+' ')
            binary.rhs.accept(self) 

    def visitExpression(self,expression):
        self.file.write("(")
        expression.expression.accept(self)
        self.file.write(")")

    def visitExpression_in_Not(self,expression_in_Not_):
        self.file.write("not(")
        expression_in_Not_.expression.accept(self)
        self.file.write(")")

    def visitAssignement_Etats(self,assignement_Etats):
        assignement_Etats.part_right.accept(self)