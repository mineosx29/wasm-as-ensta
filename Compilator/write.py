import writeClass

liste_etat = []
liste_dest = []
liste_cond = []

def main():
    w = writeClass.WriteClass()
    liste_etat, liste_dest, liste_cond = w.WriteFunction()
    print(liste_etat)
    print(liste_dest)
    print(liste_cond)

main()