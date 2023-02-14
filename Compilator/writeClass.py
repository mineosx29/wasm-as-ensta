class WriteClass:
    def __init__(self, liste = [],liste_etat_deb = [],liste_etat_dest = [],liste_condition = []):
        self.liste = liste
        self.liste_etat_deb = liste_etat_deb
        self.liste_etat_dest = liste_etat_dest
        self.liste_condition = liste_condition


    def WriteFunction(self):
        fichier_gv = open("code.gv", "r")
        fichier = open("cond.txt", "w")
        for i in fichier_gv:
            i = i.strip()
            if "label" in i:
                self.liste.append(i)

        for p in self.liste:
            p = p.strip()
            #self.liste_etat_deb.append(p[:p.index("-")])
            #self.liste_etat_dest.append(p[p.index(">")+1:p.index("[")])
            self.liste_condition.append(p)

        for k in self.liste_condition:
            fichier.write(k + '\n')
            

        return self.liste



