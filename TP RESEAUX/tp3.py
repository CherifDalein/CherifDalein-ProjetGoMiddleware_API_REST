import re
#saisi=input("Saisir un texte")
#print(saisi)


#----------------------------------------
#saisi=input("Saisir un texte")
#minuscule=saisi.lower()
#sansCaractere=re.sub(r'[^a-zA-Z0-9]','', minuscule)
#print(sansCaractere)

#------------------------------------------
#def position_alphabet(lettre):
    #return ord(lettre.lower())-ord('a')+1


#print(position_alphabet('b'))

#------------------------------------------


maliste=['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
lemot="a"
lacle="b"
def recupIndice(lettre):
    for i in range(26):
        if(maliste[i]==lettre):
            indice=i+1
            return indice

    
def crypto(mot, cle):
    IndiceMot=recupIndice(mot)
    IndiceCle=recupIndice(cle)
    IndiceCrypto=IndiceCle+IndiceMot
    return maliste[IndiceCrypto-1], IndiceCrypto


print(crypto(lemot, lacle))

#--------------------------------------------------------------------------------------
