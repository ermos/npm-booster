package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

type PackageJSON struct {
	Name			string 		`json:"name"`
	Version			string 		`json:"version"`
	Description		string 		`json:"description"`
	Main			string 		`json:"main"`
	Scripts			Scripts		`json:"scripts"`
	Author			string 		`json:"author"`
	License			string 		`json:"license"`
}

type Scripts struct  {
	Test 			string			`json:"test"`
}

var RandomName = []string{"MARTIN", "BERNARD", "THOMAS", "PETIT", "ROBERT", "RICHARD", "DURAND", "DUBOIS", "MOREAU", "LAURENT", "SIMON", "MICHEL", "LEFEBVRE", "LEROY", "ROUX", "DAVID", "BERTRAND", "MOREL", "FOURNIER", "GIRARD", "BONNET", "DUPONT", "LAMBERT", "FONTAINE", "ROUSSEAU", "VINCENT", "MULLER", "LEFEVRE", "FAURE", "ANDRE", "MERCIER", "BLANC", "GUERIN", "BOYER", "GARNIER", "CHEVALIER", "FRANCOIS", "LEGRAND", "GAUTHIER", "GARCIA", "PERRIN", "ROBIN", "CLEMENT", "MORIN", "NICOLAS", "HENRY", "ROUSSEL", "MATHIEU", "GAUTIER", "MASSON", "MARCHAND", "DUVAL", "DENIS", "DUMONT", "MARIE", "LEMAIRE", "NOEL", "MEYER", "DUFOUR", "MEUNIER", "BRUN", "BLANCHARD", "GIRAUD", "JOLY", "RIVIERE", "LUCAS", "BRUNET", "GAILLARD", "BARBIER", "ARNAUD", "MARTINEZ", "GERARD", "ROCHE", "RENARD", "SCHMITT", "ROY", "LEROUX", "COLIN", "VIDAL", "CARON", "PICARD", "ROGER", "FABRE", "AUBERT", "LEMOINE", "RENAUD", "DUMAS", "LACROIX", "OLIVIER", "PHILIPPE", "BOURGEOIS", "PIERRE", "BENOIT", "REY", "LECLERC", "PAYET", "ROLLAND", "LECLERCQ", "GUILLAUME", "LECOMTE", "LOPEZ", "JEAN", "DUPUY", "GUILLOT", "HUBERT", "BERGER", "CARPENTIER", "SANCHEZ", "DUPUIS", "MOULIN", "LOUIS", "DESCHAMPS", "HUET", "VASSEUR", "PEREZ", "BOUCHER", "FLEURY", "ROYER", "KLEIN", "JACQUET", "ADAM", "PARIS", "POIRIER", "MARTY", "AUBRY", "GUYOT", "CARRE", "CHARLES", "RENAULT", "CHARPENTIER", "MENARD", "MAILLARD", "BARON", "BERTIN", "Nom", "BAILLY", "HERVE", "SCHNEIDER", "FERNANDEZ", "LE", "COLLET", "LEGER", "BOUVIER", "JULIEN", "PREVOST", "MILLET", "PERROT", "DANIEL", "LE", "COUSIN", "GERMAIN", "BRETON", "BESSON", "LANGLOIS", "REMY", "LE", "PELLETIER", "LEVEQUE", "PERRIER", "LEBLANC", "BARRE", "LEBRUN", "MARCHAL", "WEBER", "MALLET", "HAMON", "BOULANGER", "JACOB", "MONNIER", "MICHAUD", "RODRIGUEZ", "GUICHARD", "GILLET", "ETIENNE", "GRONDIN", "POULAIN", "TESSIER", "CHEVALLIER", "COLLIN", "CHAUVIN", "DA", "BOUCHET", "GAY", "LEMAITRE", "BENARD", "MARECHAL", "HUMBERT", "REYNAUD", "ANTOINE", "HOARAU", "PERRET", "BARTHELEMY", "CORDIER", "PICHON", "LEJEUNE", "GILBERT", "LAMY", "DELAUNAY", "PASQUIER", "CARLIER", "LAPORTE"}

func newPackageJSON () PackageJSON {
	return PackageJSON{
		Name: uuid.New().String(),
		Version: fmt.Sprintf("%d.%d.%d", rand.Intn(9), rand.Intn(9), rand.Intn(9)),
		Description: "",
		Main: "index.js",
		Scripts: Scripts{
			Test: "node index.js",
		},
		Author: fmt.Sprintf("%s %s", usrName, RandomName[rand.Intn(len(RandomName))]),
		License: "MIT",
	}
}