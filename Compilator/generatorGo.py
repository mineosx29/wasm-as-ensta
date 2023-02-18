class GeneratorGo:
    def __init__(self):
        self.write_state_ident=False
        self.write_state_cond=False
        self.write_state_assignment=False
        self.save_actual_state=False
        self.initial_state=True

    def visitAuto(self,auto,file):
        self.file=open(file,"w")
        self.file.write("package main\n\n")
        self.file.write("import (\n\n")
        self.file.write("\t\"fmt\"\n\n")
        self.file.write("\t\"syscall/js\"\n\n")
        self.file.write(")\n\n")
        self.file.write("type State uint32\n\n")
        self.file.write("var commande string\n\n")
        self.file.write("var state State\n\n")
        self.file.write("func main() {\n")
        for declaration in auto.declarations:
            print(declaration)
            declaration.accept(self)
        self.file.write(")\n")
        self.file.write("\tswitch state {\n")
        for etat in auto.etats:
            etat.accept(self)
        self.file.write("}\n")
        self.file.write("}")
        self.file.close()
        print("generatorFSM: The Go File is Generated")
    
    def visitVariable(self,variable):
        variable.type.accept(self)
        variable.ident.accept(self)


    def visitIdent(self,ident):
        print("IDENT")
        if self.write_state_ident:
            if self.initial_state:
                self.file.write("\tstate = " + ident.tok + "\n\n")
                self.file.write("\tconst (\n\t"+ ident.tok+" State = iota\n")
                self.initial_state=False
            else:
                self.file.write("\t"+ident.tok+ "\n")
            self.write_state_ident=False
        #self.file.write(")\n")
        if self.write_state_cond:
            self.file.write(ident.tok)
        if self.write_state_assignment:
            self.initial_state = True
            self.file.write("state = "+ident.tok+"\n")
            self.write_state_assignment=False
        if self.save_actual_state:
            self.actual_state=ident.tok+":"
            self.save_actual_state=False


    def visitType(self,type):
        if type.tok=="state":
            self.write_state_ident=True


    def visitEtats(self,etats):
        self.save_actual_state=True
        etats.ident.accept(self)
        for condition in etats.condition:
            condition.accept(self)

    def visitIF(self,if_):
        print("IF")
        self.file.write("\t\tcase ")
        self.file.write(self.actual_state)
        self.file.write("\n")
        self.file.write("\t\t\t")
        self.file.write("if commande == \"")
        self.write_state_cond=True
        if_.condition.accept(self)
        self.write_state_cond=False
        self.file.write("\"")
        self.file.write(" {\n\t\t\t\t")
        self.write_state_assignment=True
        if_.assign.accept(self)
        self.file.write("}\n")
        #self.file.write(" [ label = \"")
        #self.file.write("\"];\n")
        

    def visitBinary(self,binary):
        binary.lhs.accept(self)
        if binary.op is None and binary.rhs is None:
            pass
        else:
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