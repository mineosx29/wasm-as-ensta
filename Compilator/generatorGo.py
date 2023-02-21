class GeneratorGo:
    def __init__(self):
        self.ecriture_etat=0
        self.etat_condition=0
        self.assignement_etat=0
        self.etat_actuel_backup=0
        self.etat_initial=1
        self.assignement_etat_ini = 0

    def visitAuto(self,auto,file):
        self.file=open(file,"w")
        self.file.write("package main\n\n")
        self.file.write("import (\n\n")
        self.file.write("\t\"fmt\"\n\n")
        self.file.write("\t\"syscall/js\"\n\n")
        self.file.write(")\n\n")
        self.file.write("type State uint32\n\n")
        self.file.write("var commande string\n\n")
        self.file.write("var document js.Value\n\n")
        #self.file.write("var state State\n\n")
        self.file.write("func getElementById(elem string) js.Value {\n\n\t")
        self.file.write("document = js.Global().Get(\"document\")\n\t")
        self.file.write("return document.Call(\"getElementById\", elem)\n}\n")
        self.file.write("func getElementValue(elem string, value string) js.Value {\n\n\t")
        self.file.write("return getElementById(elem).Get(value)\n}\n\n")
        self.file.write("func GetString(elem string, value string) string {\n\n\t")
        self.file.write("return getElementValue(elem, value).String()\n}\n\n")
        self.file.write("func main() {\n")
        self.file.write("\tquit := make(chan struct{}, 0)\n")
        for declaration in auto.declarations:
            print(declaration)
            declaration.accept(self)
        self.file.write(")\n")
        self.file.write("\tbutton := js.Global().Get(\"document\").Call(\"getElementById\", \"ok\")\n")       
        self.file.write("\tvar cb js.Func\n")
        self.file.write("\tcb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {\n")
        self.file.write("\tcommande = GetString(\"in\", \"value\")\n\n")
        self.file.write("\tswitch state {\n")
        for etat in auto.etats:
            etat.accept(self)
        self.file.write("\t}\n")
        self.file.write("\t\t\treturn nil\n\n")
        self.file.write("\t\t})\n")
        self.file.write("\tbutton.Call(\"addEventListener\", \"click\", cb)\n")
        self.file.write("\t<-quit\n")
        self.file.write("}")
        self.file.close()
        print("generatorFSM: The Go File is Generated")
    
    def visitVariable(self,variable):
        variable.type.accept(self)
        variable.ident.accept(self)


    def visitIdent(self,ident):
        print("IDENT")
        if self.ecriture_etat == 1:
            if self.etat_initial == 1:
                self.file.write("\tconst (\n\t"+ ident.tok+" State = iota\n")
                self.file.write("\tstate := " + ident.tok + "\n\n")
                self.etat_initial=0
            else:
                self.file.write("\t"+ident.tok+ "\n")
            self.ecriture_etat=0
        if self.etat_condition == 1:
            self.file.write(ident.tok)
        # if self.assignement_etat_ini == 1:
        #     if self.etat_initial == 1:
        #         self.file.write("\tstate = " + ident.tok)
        #         self.etat_initial = 0
        #     self.assignement_etat_ini = 0
        if self.assignement_etat == 1:
            self.etat_initial = 1
            self.file.write("fmt.Println(\" Passage à Etat : \")\n");
            self.file.write("\t\t\t\tfmt.Println(" + ident.tok+")\n")
            self.file.write("\t\t\t\tstate = "+ident.tok+"\n")
            self.assignement_etat=0
        if self.etat_actuel_backup == 1:
            self.actual_state=ident.tok+":"
            self.act_state = ident.tok
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
        self.file.write("\t\tcase ")
        self.file.write(self.actual_state)
        self.file.write("\n")
        self.file.write("\t\t\t")
        self.file.write("if commande == \"")
        self.etat_condition=1
        if_.condition.accept(self)
        self.etat_condition=0
        self.file.write("\"")
        self.file.write(" {\n\t\t\t\t")
        self.assignement_etat=1
        if_.assign.accept(self)
        self.file.write("\t\t\t}\n")
        
       
        

    def visitBinary(self,binary):
        binary.lhs.accept(self)
        # if binary.op is None and binary.rhs is None:
        #     pass
        # else:
        #     self.file.write(' '+binary.op+' ')
        #     binary.rhs.accept(self) 
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